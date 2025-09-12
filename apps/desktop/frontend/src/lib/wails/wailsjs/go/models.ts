export namespace app {
	
	export class RequestPairingResult {
	    Accepted: boolean;
	    Message: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestPairingResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Accepted = source["Accepted"];
	        this.Message = source["Message"];
	    }
	}
	export class SendFile {
	    Path: string;
	    Size: number;
	
	    static createFrom(source: any = {}) {
	        return new SendFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Size = source["Size"];
	    }
	}

}

export namespace color {
	
	export class Color {
	    Name: string;
	    Hex: string;
	    R: number;
	    G: number;
	    B: number;
	
	    static createFrom(source: any = {}) {
	        return new Color(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Hex = source["Hex"];
	        this.R = source["R"];
	        this.G = source["G"];
	        this.B = source["B"];
	    }
	}

}

export namespace config {
	
	export class FileSendSettings {
	    Enabled: boolean;
	    MaxSize: number;
	    SaveFolder: string;
	
	    static createFrom(source: any = {}) {
	        return new FileSendSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.MaxSize = source["MaxSize"];
	        this.SaveFolder = source["SaveFolder"];
	    }
	}
	export class SysInfoSettings {
	    Enabled: boolean;
	    CacheSeconds: number;
	
	    static createFrom(source: any = {}) {
	        return new SysInfoSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.CacheSeconds = source["CacheSeconds"];
	    }
	}

}

export namespace connected {
	
	export class Connected {
	    Info: device.Info;
	    PairedDevice: device.PairedDevice;
	    IP: number[];
	    Port: number;
	
	    static createFrom(source: any = {}) {
	        return new Connected(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Info = this.convertValues(source["Info"], device.Info);
	        this.PairedDevice = this.convertValues(source["PairedDevice"], device.PairedDevice);
	        this.IP = source["IP"];
	        this.Port = source["Port"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace device {
	
	export class Info {
	    ID: string;
	    Name: string;
	    Color: color.Color;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Color = this.convertValues(source["Color"], color.Color);
	        this.Type = source["Type"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PairedDevice {
	    Certificate: number[];
	    Token: number[];
	
	    static createFrom(source: any = {}) {
	        return new PairedDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Certificate = source["Certificate"];
	        this.Token = source["Token"];
	    }
	}

}

export namespace discovery {
	
	export class FoundDevice {
	    IP: number[];
	    Port: number;
	    Info: device.Info;
	
	    static createFrom(source: any = {}) {
	        return new FoundDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IP = source["IP"];
	        this.Port = source["Port"];
	        this.Info = this.convertValues(source["Info"], device.Info);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace filesend {
	
	export class FileTransfersStatusTransfer {
	    FileName: string;
	    Percent: number;
	    FileSize: number;
	    Status: string;
	
	    static createFrom(source: any = {}) {
	        return new FileTransfersStatusTransfer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.FileName = source["FileName"];
	        this.Percent = source["Percent"];
	        this.FileSize = source["FileSize"];
	        this.Status = source["Status"];
	    }
	}
	export class FileTransfersStatusDevice {
	    Device: device.Info;
	    Transfers: FileTransfersStatusTransfer[];
	    Percent: number;
	
	    static createFrom(source: any = {}) {
	        return new FileTransfersStatusDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Device = this.convertValues(source["Device"], device.Info);
	        this.Transfers = this.convertValues(source["Transfers"], FileTransfersStatusTransfer);
	        this.Percent = source["Percent"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class FileTransfersStatus {
	    PercentSending: number;
	    PercentReceiving: number;
	    Sending: FileTransfersStatusDevice[];
	    Receiving: FileTransfersStatusDevice[];
	
	    static createFrom(source: any = {}) {
	        return new FileTransfersStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.PercentSending = source["PercentSending"];
	        this.PercentReceiving = source["PercentReceiving"];
	        this.Sending = this.convertValues(source["Sending"], FileTransfersStatusDevice);
	        this.Receiving = this.convertValues(source["Receiving"], FileTransfersStatusDevice);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace node {
	
	export class Status {
	    DiscoveryRunning: boolean;
	    ServerRunning: boolean;
	    WorkerRunning: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Status(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DiscoveryRunning = source["DiscoveryRunning"];
	        this.ServerRunning = source["ServerRunning"];
	        this.WorkerRunning = source["WorkerRunning"];
	    }
	}

}

export namespace pbuf {
	
	export class SysInfo {
	    hostName?: string;
	    os?: string;
	    platform?: string;
	    memTotal?: number;
	    memUsed?: number;
	    diskTotal?: number;
	    diskUsed?: number;
	    batteryCharging?: boolean;
	    batteryPercent?: number;
	    cpuUsage?: number;
	
	    static createFrom(source: any = {}) {
	        return new SysInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostName = source["hostName"];
	        this.os = source["os"];
	        this.platform = source["platform"];
	        this.memTotal = source["memTotal"];
	        this.memUsed = source["memUsed"];
	        this.diskTotal = source["diskTotal"];
	        this.diskUsed = source["diskUsed"];
	        this.batteryCharging = source["batteryCharging"];
	        this.batteryPercent = source["batteryPercent"];
	        this.cpuUsage = source["cpuUsage"];
	    }
	}

}

