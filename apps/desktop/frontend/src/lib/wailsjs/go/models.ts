export namespace address {
	
	export class Address {
	    port: number;
	    ip: number[];
	    phrase: string;
	
	    static createFrom(source: any = {}) {
	        return new Address(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.port = source["port"];
	        this.ip = source["ip"];
	        this.phrase = source["phrase"];
	    }
	}

}

export namespace config {
	
	export class ServicesConfig {
	    RCFile: rcfile.ServiceConfig;
	    SysInfo: sysinfo.ServiceConfig;
	
	    static createFrom(source: any = {}) {
	        return new ServicesConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RCFile = this.convertValues(source["RCFile"], rcfile.ServiceConfig);
	        this.SysInfo = this.convertValues(source["SysInfo"], sysinfo.ServiceConfig);
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
	export class Config {
	    DeviceName: string;
	    DeviceColor: options.RGBA;
	    DeviceCode: string;
	    Language: string;
	    AutoStart: boolean;
	    Theme: string;
	    Services: ServicesConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DeviceName = source["DeviceName"];
	        this.DeviceColor = this.convertValues(source["DeviceColor"], options.RGBA);
	        this.DeviceCode = source["DeviceCode"];
	        this.Language = source["Language"];
	        this.AutoStart = source["AutoStart"];
	        this.Theme = source["Theme"];
	        this.Services = this.convertValues(source["Services"], ServicesConfig);
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
	
	export class DeviceInfo {
	    address: address.Address;
	    name: string;
	    color: options.RGBA;
	    code: string;
	    type: number;
	    services: service.Service[];
	
	    static createFrom(source: any = {}) {
	        return new DeviceInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], address.Address);
	        this.name = source["name"];
	        this.color = this.convertValues(source["color"], options.RGBA);
	        this.code = source["code"];
	        this.type = source["type"];
	        this.services = this.convertValues(source["services"], service.Service);
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

export namespace options {
	
	export class RGBA {
	    r: number;
	    g: number;
	    b: number;
	    a: number;
	
	    static createFrom(source: any = {}) {
	        return new RGBA(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.r = source["r"];
	        this.g = source["g"];
	        this.b = source["b"];
	        this.a = source["a"];
	    }
	}

}

export namespace pair {
	
	export class Pair {
	    DeviceInfo: device.DeviceInfo;
	    Token: string;
	    OldToken: string;
	    Services: service.Service[];
	    ExposingServices: service.Service[];
	
	    static createFrom(source: any = {}) {
	        return new Pair(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DeviceInfo = this.convertValues(source["DeviceInfo"], device.DeviceInfo);
	        this.Token = source["Token"];
	        this.OldToken = source["OldToken"];
	        this.Services = this.convertValues(source["Services"], service.Service);
	        this.ExposingServices = this.convertValues(source["ExposingServices"], service.Service);
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

export namespace pbuf {
	
	export class Battery {
	    current_capacity?: number;
	    full_charged_capacity?: number;
	    state?: string;
	
	    static createFrom(source: any = {}) {
	        return new Battery(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current_capacity = source["current_capacity"];
	        this.full_charged_capacity = source["full_charged_capacity"];
	        this.state = source["state"];
	    }
	}
	export class CPUInfoStat {
	    cpu?: number;
	    vendor_id?: string;
	    family?: string;
	    model?: string;
	    stepping?: number;
	    physical_id?: string;
	    core_id?: string;
	    cores?: number;
	    model_name?: string;
	    mhz?: number;
	    cache_size?: number;
	    flags?: string[];
	
	    static createFrom(source: any = {}) {
	        return new CPUInfoStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cpu = source["cpu"];
	        this.vendor_id = source["vendor_id"];
	        this.family = source["family"];
	        this.model = source["model"];
	        this.stepping = source["stepping"];
	        this.physical_id = source["physical_id"];
	        this.core_id = source["core_id"];
	        this.cores = source["cores"];
	        this.model_name = source["model_name"];
	        this.mhz = source["mhz"];
	        this.cache_size = source["cache_size"];
	        this.flags = source["flags"];
	    }
	}
	export class DeviceColor {
	    r?: number;
	    g?: number;
	    b?: number;
	
	    static createFrom(source: any = {}) {
	        return new DeviceColor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.r = source["r"];
	        this.g = source["g"];
	        this.b = source["b"];
	    }
	}
	export class Service {
	    name?: string;
	    enabled?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Service(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.enabled = source["enabled"];
	    }
	}
	export class DeviceInfo {
	    code?: string;
	    name?: string;
	    type?: number;
	    color?: DeviceColor;
	    services?: Service[];
	    ip?: string;
	    port?: number;
	
	    static createFrom(source: any = {}) {
	        return new DeviceInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.color = this.convertValues(source["color"], DeviceColor);
	        this.services = this.convertValues(source["services"], Service);
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
	export class DiskUsageStat {
	    path?: string;
	    fstype?: string;
	    total?: number;
	    free?: number;
	    used?: number;
	    used_percent?: number;
	    inodes_total?: number;
	    inodes_used?: number;
	    inodes_free?: number;
	    inodes_used_percent?: number;
	
	    static createFrom(source: any = {}) {
	        return new DiskUsageStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.fstype = source["fstype"];
	        this.total = source["total"];
	        this.free = source["free"];
	        this.used = source["used"];
	        this.used_percent = source["used_percent"];
	        this.inodes_total = source["inodes_total"];
	        this.inodes_used = source["inodes_used"];
	        this.inodes_free = source["inodes_free"];
	        this.inodes_used_percent = source["inodes_used_percent"];
	    }
	}
	export class HostInfoStat {
	    hostname?: string;
	    uptime?: number;
	    boot_time?: number;
	    procs?: number;
	    os?: string;
	    platform?: string;
	    platform_family?: string;
	    platform_version?: string;
	    kernel_version?: string;
	    kernel_arch?: string;
	    virtualization_system?: string;
	    virtualization_role?: string;
	    host_id?: string;
	
	    static createFrom(source: any = {}) {
	        return new HostInfoStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.uptime = source["uptime"];
	        this.boot_time = source["boot_time"];
	        this.procs = source["procs"];
	        this.os = source["os"];
	        this.platform = source["platform"];
	        this.platform_family = source["platform_family"];
	        this.platform_version = source["platform_version"];
	        this.kernel_version = source["kernel_version"];
	        this.kernel_arch = source["kernel_arch"];
	        this.virtualization_system = source["virtualization_system"];
	        this.virtualization_role = source["virtualization_role"];
	        this.host_id = source["host_id"];
	    }
	}
	export class PairRequest {
	    token?: string;
	    device?: DeviceInfo;
	    services?: Service[];
	
	    static createFrom(source: any = {}) {
	        return new PairRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.token = source["token"];
	        this.device = this.convertValues(source["device"], DeviceInfo);
	        this.services = this.convertValues(source["services"], Service);
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
	export class PairResponse {
	    token?: string;
	    device?: DeviceInfo;
	    accepted?: boolean;
	    services?: Service[];
	
	    static createFrom(source: any = {}) {
	        return new PairResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.token = source["token"];
	        this.device = this.convertValues(source["device"], DeviceInfo);
	        this.accepted = source["accepted"];
	        this.services = this.convertValues(source["services"], Service);
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
	
	export class VirtualMemoryStat {
	    total?: number;
	    available?: number;
	    used?: number;
	    used_percent?: number;
	    free?: number;
	
	    static createFrom(source: any = {}) {
	        return new VirtualMemoryStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.available = source["available"];
	        this.used = source["used"];
	        this.used_percent = source["used_percent"];
	        this.free = source["free"];
	    }
	}
	export class SysInfo {
	    host?: HostInfoStat;
	    cpu?: CPUInfoStat[];
	    cpu_usage?: number[];
	    memory?: VirtualMemoryStat;
	    disk?: DiskUsageStat;
	    battery?: Battery[];
	
	    static createFrom(source: any = {}) {
	        return new SysInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.host = this.convertValues(source["host"], HostInfoStat);
	        this.cpu = this.convertValues(source["cpu"], CPUInfoStat);
	        this.cpu_usage = source["cpu_usage"];
	        this.memory = this.convertValues(source["memory"], VirtualMemoryStat);
	        this.disk = this.convertValues(source["disk"], DiskUsageStat);
	        this.battery = this.convertValues(source["battery"], Battery);
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

export namespace rcfile {
	
	export class ServiceConfig {
	    Enabled: boolean;
	    Destination: string;
	    FileMaxSize: number;
	
	    static createFrom(source: any = {}) {
	        return new ServiceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.Destination = source["Destination"];
	        this.FileMaxSize = source["FileMaxSize"];
	    }
	}

}

export namespace service {
	
	export class Service {
	    name: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Service(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.enabled = source["enabled"];
	    }
	}

}

export namespace sysinfo {
	
	export class ServiceConfig {
	    Enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ServiceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	    }
	}

}

