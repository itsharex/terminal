//go:build windows

package termx

import (
	"log"
	"strings"
	"sync/atomic"
)

type windowsPty struct {
	pty    *conpty.ConPty
	closed *atomic.Bool
}

func (t *windowsPty) Sftp() (*sftp.Client, error) {
	//TODO implement me
	return nil, errors.New("sftp pty not supported")
}

func (t *windowsPty) Resize(rows, cols int) error {
	return t.pty.Resize(cols, rows)
}

func (t *windowsPty) Read(p []byte) (n int, err error) {
	return t.pty.Read(p)
}

func (t *windowsPty) Write(p []byte) (n int, err error) {
	return t.pty.Write(p)
}

func (t *windowsPty) Close() error {
	if t.closed.CompareAndSwap(false, true) {
		log.Println("do swap")
		return t.pty.Close()
	}
	return nil
}

func NewPTY(s *SystemShell) (PtyX, error) {
	if !conpty.IsConPtyAvailable() {
		return nil, conpty.ErrConPtyUnsupported
	}
	cmdLine := s.Command + " " + strings.Join(s.Args, " ")
	var options []conpty.ConPtyOption
	options = append(options, conpty.ConPtyDimensions(1000, 700))
	if s.Cwd != "" {
		options = append(options, conpty.ConPtyWorkDir(s.Cwd))
	}
	wPty, err := conpty.Start(cmdLine, options...)
	if err != nil {
		return nil, err
	}
	return &windowsPty{pty: wPty, closed: &atomic.Bool{}}, nil
}
