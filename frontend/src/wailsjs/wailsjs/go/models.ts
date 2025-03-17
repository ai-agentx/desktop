export namespace main {
	
	export class Agent {
	    id: string;
	    name: string;
	    instructions: string;
	    handoff_description?: string;
	    model?: string;
	    tools: string[];
	
	    static createFrom(source: any = {}) {
	        return new Agent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.instructions = source["instructions"];
	        this.handoff_description = source["handoff_description"];
	        this.model = source["model"];
	        this.tools = source["tools"];
	    }
	}
	export class RunAgentResponse {
	    agent_id: string;
	    result: string;
	    raw_output?: Record<string, any>;
	
	    static createFrom(source: any = {}) {
	        return new RunAgentResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.agent_id = source["agent_id"];
	        this.result = source["result"];
	        this.raw_output = source["raw_output"];
	    }
	}
	export class Tool {
	    name: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new Tool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}

}

