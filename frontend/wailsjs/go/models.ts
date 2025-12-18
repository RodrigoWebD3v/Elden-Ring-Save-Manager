export namespace main {
	
	export class Save {
	    Id: number;
	    Name: string;
	    IsAtivo: boolean;
	    LastModified: string;
	    Size: string;
	
	    static createFrom(source: any = {}) {
	        return new Save(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Name = source["Name"];
	        this.IsAtivo = source["IsAtivo"];
	        this.LastModified = source["LastModified"];
	        this.Size = source["Size"];
	    }
	}

}

