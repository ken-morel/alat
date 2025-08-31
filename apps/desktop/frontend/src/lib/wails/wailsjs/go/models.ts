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

}

