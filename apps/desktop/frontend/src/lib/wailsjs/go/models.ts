export namespace address {
	
	export class Address {
	    Port: number;
	    IP: number[];
	    Phrase: string;
	
	    static createFrom(source: any = {}) {
	        return new Address(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Port = source["Port"];
	        this.IP = source["IP"];
	        this.Phrase = source["Phrase"];
	    }
	}

}

export namespace config {
	
	export class ServicesConfig {
	    RCFile: rcfile.ServiceConfig;
	
	    static createFrom(source: any = {}) {
	        return new ServicesConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RCFile = this.convertValues(source["RCFile"], rcfile.ServiceConfig);
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
	    Address: address.Address;
	    Name: string;
	    Color: options.RGBA;
	    Code: string;
	    Type: number;
	
	    static createFrom(source: any = {}) {
	        return new DeviceInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Address = this.convertValues(source["Address"], address.Address);
	        this.Name = source["Name"];
	        this.Color = this.convertValues(source["Color"], options.RGBA);
	        this.Code = source["Code"];
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
	    Name: string;
	    Enabled: boolean;
	    Protected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Service(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Enabled = source["Enabled"];
	        this.Protected = source["Protected"];
	    }
	}

}

