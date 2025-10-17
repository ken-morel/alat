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
	    name: string;
	    hex: string;
	    r: number;
	    g: number;
	    b: number;
	
	    static createFrom(source: any = {}) {
	        return new Color(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.hex = source["hex"];
	        this.r = source["r"];
	        this.g = source["g"];
	        this.b = source["b"];
	    }
	}

}

export namespace config {
	
	export class FileSendConfig {
	    enabled: boolean;
	    maxSize: number;
	    saveFolder: string;
	
	    static createFrom(source: any = {}) {
	        return new FileSendConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.maxSize = source["maxSize"];
	        this.saveFolder = source["saveFolder"];
	    }
	}
	export class SysInfoConfig {
	    enabled: boolean;
	    cacheSeconds: number;
	
	    static createFrom(source: any = {}) {
	        return new SysInfoConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.cacheSeconds = source["cacheSeconds"];
	    }
	}

}

export namespace connected {
	
	export class Connected {
	    info: device.Info;
	    pairedDevice: device.PairedDevice;
	    ip: number[];
	    port: number;
	
	    static createFrom(source: any = {}) {
	        return new Connected(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.info = this.convertValues(source["info"], device.Info);
	        this.pairedDevice = this.convertValues(source["pairedDevice"], device.PairedDevice);
	        this.ip = source["ip"];
	        this.port = source["port"];
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
	    id: string;
	    name: string;
	    color: color.Color;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.color = this.convertValues(source["color"], color.Color);
	        this.type = source["type"];
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
	    certificate: number[];
	    token: number[];
	
	    static createFrom(source: any = {}) {
	        return new PairedDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.certificate = source["certificate"];
	        this.token = source["token"];
	    }
	}

}

export namespace discovery {
	
	export class FoundDevice {
	    ip: number[];
	    port: number;
	    info: device.Info;
	
	    static createFrom(source: any = {}) {
	        return new FoundDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ip = source["ip"];
	        this.port = source["port"];
	        this.info = this.convertValues(source["info"], device.Info);
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
	    fileName: string;
	    percent: number;
	    fileSize: number;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new FileTransfersStatusTransfer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.percent = source["percent"];
	        this.fileSize = source["fileSize"];
	        this.status = source["status"];
	    }
	}
	export class FileTransfersStatusDevice {
	    device: device.Info;
	    transfers: FileTransfersStatusTransfer[];
	    percent: number;
	
	    static createFrom(source: any = {}) {
	        return new FileTransfersStatusDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.device = this.convertValues(source["device"], device.Info);
	        this.transfers = this.convertValues(source["transfers"], FileTransfersStatusTransfer);
	        this.percent = source["percent"];
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
	    percentSending: number;
	    percentReceiving: number;
	    sending: FileTransfersStatusDevice[];
	    receiving: FileTransfersStatusDevice[];
	
	    static createFrom(source: any = {}) {
	        return new FileTransfersStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.percentSending = source["percentSending"];
	        this.percentReceiving = source["percentReceiving"];
	        this.sending = this.convertValues(source["sending"], FileTransfersStatusDevice);
	        this.receiving = this.convertValues(source["receiving"], FileTransfersStatusDevice);
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
	    discoveryRunning: boolean;
	    serverRunning: boolean;
	    workerRunning: boolean;
	    port: number;
	
	    static createFrom(source: any = {}) {
	        return new Status(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.discoveryRunning = source["discoveryRunning"];
	        this.serverRunning = source["serverRunning"];
	        this.workerRunning = source["workerRunning"];
	        this.port = source["port"];
	    }
	}

}

export namespace sysinfo {
	
	export class SysInfo {
	    hostname: string;
	    os: string;
	    platform: string;
	    memTotal: number;
	    memUsed: number;
	    diskTotal: number;
	    diskUsed: number;
	    batteryCharging: boolean;
	    batteryPercent: number;
	    cpuUsage: number;
	
	    static createFrom(source: any = {}) {
	        return new SysInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
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

