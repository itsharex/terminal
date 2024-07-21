package api

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
	"terminal/lib/common/taskrunner"
	"terminal/lib/utils"
)

type FileSystem struct {
	MountPoint string `json:"mountPoint"`
	Used       uint64 `json:"used"`
	Free       uint64 `json:"free"`
}

type Network struct {
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
	Rx   uint64 `json:"rx"`
	Tx   uint64 `json:"tx"`
}

type cpuRaw struct {
	User    uint64 // time spent in user mode
	Nice    uint64 // time spent in user mode with low priority (nice)
	System  uint64 // time spent in system mode
	Idle    uint64 // time spent in the idle task
	Iowait  uint64 // time spent waiting for I/O to complete (since Linux 2.5.41)
	Irq     uint64 // time spent servicing  interrupts  (since  2.6.0-test4)
	SoftIrq uint64 // time spent servicing softirqs (since 2.6.0-test4)
	Steal   uint64 // time spent in other OSes when running in a virtualized environment
	Guest   uint64 // time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.
	Total   uint64 // total of all time fields
}

type CPU struct {
	User    float32 `json:"user"`
	Nice    float32 `json:"nice"`
	System  float32 `json:"system"`
	Idle    float32 `json:"idle"`
	IOWait  float32 `json:"ioWait"`
	Irq     float32 `json:"irq"`
	SoftIrq float32 `json:"softIrq"`
	Steal   float32 `json:"steal"`
	Guest   float32 `json:"guest"`
}

type Stat struct {
	Uptime         int64              `json:"uptime"`
	Hostname       string             `json:"hostname"`
	Load1          string             `json:"load1"`
	Load5          string             `json:"load5"`
	Load10         string             `json:"load10"`
	RunningProcess string             `json:"runningProcess"`
	TotalProcess   string             `json:"totalProcess"`
	MemTotal       uint64             `json:"memTotal"`
	MemAvailable   uint64             `json:"memAvailable"`
	MemFree        uint64             `json:"memFree"`
	MemBuffers     uint64             `json:"memBuffers"`
	MemCached      uint64             `json:"memCached"`
	SwapTotal      uint64             `json:"swapTotal"`
	SwapFree       uint64             `json:"swapFree"`
	FileSystems    []FileSystem       `json:"fileSystems"`
	Network        map[string]Network `json:"network"`
	CPU            CPU                `json:"cpu"`
	runner         *taskrunner.Runner `json:"-" yaml:"-" xml:"-"`
}

func NewStats() *Stat {
	return &Stat{
		runner: &taskrunner.Runner{},
	}
}

func (s *Stat) GetAllStats(client *ssh.Client) error {
	//start := time.Now()
	if s.Uptime == 0 {
		if err := s.getUptime(client); err != nil {
			return err
		}
	}

	if s.Hostname == "" {
		if err := s.getHostname(client); err != nil {
			return err
		}
	}

	s.runner.Add(func() error {
		return s.getLoad(client)
	})
	s.runner.Add(func() error {
		return s.getMem(client)
	})
	s.runner.Add(func() error {
		return s.getFileSystems(client)
	})
	s.runner.Add(func() error {
		return s.getInterfaces(client)
	})
	s.runner.Add(func() error {
		return s.getInterfaceInfo(client)
	})
	s.runner.Add(func() error {
		return s.getCPU(client)
	})
	errs := s.runner.Wait()
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	//cost := time.Since(start)
	//fmt.Printf("%s: %v\n", "GetAllStats", cost)
	return nil
}

func (s *Stat) getHostname(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getHostname")
	hostname, err := utils.RunCommand(client, "/bin/hostname -f")
	if err != nil {
		return
	}
	s.Hostname = strings.TrimSpace(hostname)
	return
}

func (s *Stat) getUptime(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getUptime")
	uptime, err := utils.RunCommand(client, "/bin/cat /proc/uptime")
	if err != nil {
		return
	}

	parts := strings.Fields(uptime)
	if len(parts) == 2 {
		var upSeconds float64
		upSeconds, err = strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return
		}
		s.Uptime = int64(upSeconds * 1000)
	}
	return
}

func (s *Stat) getLoad(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getLoad")
	line, err := utils.RunCommand(client, "/bin/cat /proc/loadavg")
	if err != nil {
		return
	}

	parts := strings.Fields(line)
	if len(parts) == 5 {
		s.Load1 = parts[0]
		s.Load5 = parts[1]
		s.Load10 = parts[2]
		if i := strings.Index(parts[3], "/"); i != -1 {
			s.RunningProcess = parts[3][0:i]
			if i+1 < len(parts[3]) {
				s.TotalProcess = parts[3][i+1:]
			}
		}
	}
	return
}

func (s *Stat) getMem(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getMem")
	lines, err := utils.RunCommand(client, "/bin/cat /proc/meminfo")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 3 {
			val, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				continue
			}
			val *= 1024
			switch parts[0] {
			case "MemTotal:":
				s.MemTotal = val
			case "MemFree:":
				s.MemFree = val
			case "MemAvailable:":
				s.MemAvailable = val
			case "Buffers:":
				s.MemBuffers = val
			case "Cached:":
				s.MemCached = val
			case "SwapTotal:":
				s.SwapTotal = val
			case "SwapFree:":
				s.SwapFree = val
			}
		}
	}
	return
}

func (s *Stat) getFileSystems(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getFileSystems")
	lines, err := utils.RunCommand(client, "/bin/df -B1")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	flag := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		n := len(parts)
		dev := n > 0 && strings.Index(parts[0], "/dev/") == 0
		if n == 1 && dev {
			flag = 1
		} else if (n == 5 && flag == 1) || (n == 6 && dev) {
			i := flag
			flag = 0
			used, err := strconv.ParseUint(parts[2-i], 10, 64)
			if err != nil {
				continue
			}
			free, err := strconv.ParseUint(parts[3-i], 10, 64)
			if err != nil {
				continue
			}
			s.FileSystems = append(s.FileSystems, FileSystem{
				parts[5-i], used, free,
			})
		}
	}

	return
}

func (s *Stat) getInterfaces(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getInterfaces")
	var lines string
	lines, err = utils.RunCommand(client, "/bin/ip -o addr")
	if err != nil {
		// try /sbin/ip
		lines, err = utils.RunCommand(client, "/sbin/ip -o addr")
		if err != nil {
			return
		}
	}

	if s.Network == nil {
		s.Network = make(map[string]Network)
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 4 && (parts[2] == "inet" || parts[2] == "inet6") {
			ipv4 := parts[2] == "inet"
			intfname := parts[1]
			if info, ok := s.Network[intfname]; ok {
				if ipv4 {
					info.IPv4 = parts[3]
				} else {
					info.IPv6 = parts[3]
				}
				s.Network[intfname] = info
			} else {
				info := Network{}
				if ipv4 {
					info.IPv4 = parts[3]
				} else {
					info.IPv6 = parts[3]
				}
				s.Network[intfname] = info
			}
		}
	}

	return
}

func (s *Stat) getInterfaceInfo(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getInterfaceInfo")

	if s.Network == nil {
		return
	} // should have been here already

	lines, err := utils.RunCommand(client, "/bin/cat /proc/net/dev")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 17 {
			intf := strings.TrimSpace(parts[0])
			intf = strings.TrimSuffix(intf, ":")
			if info, ok := s.Network[intf]; ok {
				rx, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					continue
				}
				tx, err := strconv.ParseUint(parts[9], 10, 64)
				if err != nil {
					continue
				}
				info.Rx = rx
				info.Tx = tx
				s.Network[intf] = info
			}
		}
	}

	return
}

func parseCPUFields(fields []string, stat *cpuRaw) {
	numFields := len(fields)
	for i := 1; i < numFields; i++ {
		val, err := strconv.ParseUint(fields[i], 10, 64)
		if err != nil {
			continue
		}

		stat.Total += val
		switch i {
		case 1:
			stat.User = val
		case 2:
			stat.Nice = val
		case 3:
			stat.System = val
		case 4:
			stat.Idle = val
		case 5:
			stat.Iowait = val
		case 6:
			stat.Irq = val
		case 7:
			stat.SoftIrq = val
		case 8:
			stat.Steal = val
		case 9:
			stat.Guest = val
		}
	}
}

// the CPU stats that were fetched last time round
var preCPU cpuRaw

func (s *Stat) getCPU(client *ssh.Client) (err error) {
	defer utils.TimeWatcher("getCPU")
	lines, err := utils.RunCommand(client, "/bin/cat /proc/stat")
	if err != nil {
		return
	}

	var (
		nowCPU cpuRaw
		total  float32
	)

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) > 0 && fields[0] == "cpu" { // changing here if want to get every cpu-core's stats
			parseCPUFields(fields, &nowCPU)
			break
		}
	}
	if preCPU.Total == 0 { // having no pre raw cpu data
		goto END
	}

	total = float32(nowCPU.Total - preCPU.Total)
	s.CPU.User = float32(nowCPU.User-preCPU.User) / total * 100
	s.CPU.Nice = float32(nowCPU.Nice-preCPU.Nice) / total * 100
	s.CPU.System = float32(nowCPU.System-preCPU.System) / total * 100
	s.CPU.Idle = float32(nowCPU.Idle-preCPU.Idle) / total * 100
	s.CPU.IOWait = float32(nowCPU.Iowait-preCPU.Iowait) / total * 100
	s.CPU.Irq = float32(nowCPU.Irq-preCPU.Irq) / total * 100
	s.CPU.SoftIrq = float32(nowCPU.SoftIrq-preCPU.SoftIrq) / total * 100
	s.CPU.Guest = float32(nowCPU.Guest-preCPU.Guest) / total * 100
END:
	preCPU = nowCPU
	return
}
