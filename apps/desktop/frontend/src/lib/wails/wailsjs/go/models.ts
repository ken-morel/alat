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

