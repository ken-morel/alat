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
	
	export class FileSharingSettings {
	    DefaultDownloadLocation: string;
	    AskBeforeReceiving: boolean;
	    MaxFileSizeMB: number;
	
	    static createFrom(source: any = {}) {
	        return new FileSharingSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DefaultDownloadLocation = source["DefaultDownloadLocation"];
	        this.AskBeforeReceiving = source["AskBeforeReceiving"];
	        this.MaxFileSizeMB = source["MaxFileSizeMB"];
	    }
	}
	export class SyncedFolderEntry {
	    LocalPath: string;
	    SyncDirection: string;
	    ConflictStrategy: string;
	    IgnorePatterns: string[];
	
	    static createFrom(source: any = {}) {
	        return new SyncedFolderEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LocalPath = source["LocalPath"];
	        this.SyncDirection = source["SyncDirection"];
	        this.ConflictStrategy = source["ConflictStrategy"];
	        this.IgnorePatterns = source["IgnorePatterns"];
	    }
	}
	export class FolderSyncSettings {
	    Enabled: boolean;
	    SyncedFolders: SyncedFolderEntry[];
	
	    static createFrom(source: any = {}) {
	        return new FolderSyncSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.SyncedFolders = this.convertValues(source["SyncedFolders"], SyncedFolderEntry);
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
	export class MediaControlSettings {
	    Enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new MediaControlSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	    }
	}
	export class NotificationSyncSettings {
	    Enabled: boolean;
	    AppWhitelist: string[];
	    AppBlacklist: string[];
	    QuickReplies: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NotificationSyncSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.AppWhitelist = source["AppWhitelist"];
	        this.AppBlacklist = source["AppBlacklist"];
	        this.QuickReplies = source["QuickReplies"];
	    }
	}
	export class RemoteInputSettings {
	    Enabled: boolean;
	    MouseSensitivity: number;
	    NaturalScrolling: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RemoteInputSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.MouseSensitivity = source["MouseSensitivity"];
	        this.NaturalScrolling = source["NaturalScrolling"];
	    }
	}
	
	export class UniversalClipboardSettings {
	    Enabled: boolean;
	    SyncText: boolean;
	    SyncImages: boolean;
	    IgnorePasswordManagers: boolean;
	
	    static createFrom(source: any = {}) {
	        return new UniversalClipboardSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Enabled = source["Enabled"];
	        this.SyncText = source["SyncText"];
	        this.SyncImages = source["SyncImages"];
	        this.IgnorePasswordManagers = source["IgnorePasswordManagers"];
	    }
	}

}

export namespace connected {
	
	export class Connected {
	    Info: device.Info;
	    PairedDevice: device.PairedDevice;
	
	    static createFrom(source: any = {}) {
	        return new Connected(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Info = this.convertValues(source["Info"], device.Info);
	        this.PairedDevice = this.convertValues(source["PairedDevice"], device.PairedDevice);
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

export namespace node {
	
	export class Status {
	    DiscoveryRunning: boolean;
	    ServerRunning: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Status(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DiscoveryRunning = source["DiscoveryRunning"];
	        this.ServerRunning = source["ServerRunning"];
	    }
	}

}

