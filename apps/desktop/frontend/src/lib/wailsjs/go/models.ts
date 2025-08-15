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
	
	export class Config {
	    DeviceName: string;
	    DeviceColor: options.RGBA;
	    DeviceCode: string;
	    Language: string;
	    AutoStart: boolean;
	    Theme: string;
	
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

