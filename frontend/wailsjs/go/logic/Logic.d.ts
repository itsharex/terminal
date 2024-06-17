// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {logic} from '../models';
import {termx} from '../models';

export function AddFoldOrHost(arg1:logic.HostEntry):Promise<void>;

export function AddKey(arg1:logic.KeyEntry):Promise<void>;

export function ClosePty(arg1:string):Promise<void>;

export function CreateLocalPty(arg1:termx.SystemShell):Promise<void>;

export function CreateSshPty(arg1:string,arg2:number,arg3:number,arg4:number):Promise<void>;

export function DelFoldOrHost(arg1:number,arg2:boolean):Promise<void>;

export function DelKey(arg1:number):Promise<void>;

export function Dir(arg1:string,arg2:string):Promise<Array<logic.FileInfo>>;

export function Download(arg1:string,arg2:string):Promise<void>;

export function GetFolds():Promise<Array<logic.HostEntry>>;

export function GetFoldsAndHosts(arg1:number):Promise<Array<logic.HostEntry>>;

export function GetHosts(arg1:number):Promise<Array<logic.HostEntry>>;

export function GetKeyList(arg1:boolean):Promise<Array<logic.KeyEntry>>;

export function GetLocalPtyList():Promise<Array<termx.SystemShell>>;

export function ResizePty(arg1:string,arg2:number,arg3:number):Promise<void>;

export function UpdFoldOrHost(arg1:logic.HostEntry):Promise<void>;

export function UpdKey(arg1:logic.KeyEntry):Promise<void>;

export function UploadDirectory(arg1:string,arg2:string):Promise<void>;

export function UploadMultipleFiles(arg1:string,arg2:string):Promise<void>;

export function WriteToPty(arg1:string,arg2:Array<number>):Promise<void>;