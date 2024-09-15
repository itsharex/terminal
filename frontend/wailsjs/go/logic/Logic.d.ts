// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {logic} from '../models';
import {termx} from '../models';

export function AddFoldOrHost(arg1:logic.HostEntry):Promise<void>;

export function AddKey(arg1:logic.KeyEntry):Promise<void>;

export function ClosePty(arg1:string):Promise<void>;

export function CloseSftpClient(arg1:string):Promise<void>;

export function CreateLocalPty(arg1:termx.SystemShell):Promise<void>;

export function CreateSshPty(arg1:string,arg2:number,arg3:number,arg4:number):Promise<void>;

export function CreateSshPtyWithJumper(arg1:string,arg2:number,arg3:number,arg4:number,arg5:number):Promise<void>;

export function DelFoldOrHost(arg1:number,arg2:boolean):Promise<void>;

export function DelKey(arg1:number):Promise<void>;

export function ExportData():Promise<string>;

export function GetFolds():Promise<Array<logic.HostEntry>>;

export function GetFoldsAndHosts(arg1:number):Promise<Array<logic.HostEntry>>;

export function GetHosts(arg1:number):Promise<Array<logic.HostEntry>>;

export function GetKeyList(arg1:boolean):Promise<Array<logic.KeyEntry>>;

export function GetLocalPtyList():Promise<Array<termx.SystemShell>>;

export function ImportData():Promise<void>;

export function IsRunAsAdmin():Promise<boolean>;

export function OpenLink(arg1:string):Promise<void>;

export function OpenPlayerWindow():Promise<void>;

export function OsGoos():Promise<string>;

export function ResizePty(arg1:string,arg2:number,arg3:number):Promise<void>;

export function RunAsAdmin():Promise<void>;

export function SftpDelete(arg1:string,arg2:string):Promise<void>;

export function SftpDir(arg1:string,arg2:string):Promise<Array<logic.FileInfo>>;

export function SftpDownload(arg1:string,arg2:string):Promise<void>;

export function SftpHomeDir(arg1:string):Promise<string>;

export function SftpUploadDirectory(arg1:string,arg2:string):Promise<void>;

export function SftpUploadMultipleFiles(arg1:string,arg2:string):Promise<void>;

export function StartRec(arg1:string):Promise<string>;

export function StopRec(arg1:string):Promise<void>;

export function UpdFoldOrHost(arg1:logic.HostEntry):Promise<void>;

export function UpdKey(arg1:logic.KeyEntry):Promise<void>;

export function WriteToPty(arg1:string,arg2:Array<number>):Promise<void>;
