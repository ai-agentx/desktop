import { defineStore } from 'pinia';
import { 
  ListAgents, 
  GetAgent, 
  CreateAgent, 
  DeleteAgent, 
  RunAgent, 
  CloneAgent,
  GetTools
} from '../wailsjs/wailsjs/go/main/App';

interface Agent {
  id: string;
  name: string;
  instructions: string;
  handoff_description?: string;
  model?: string;
  tools: string[];
}

interface Tool {
  name: string;
  description: string;
}

interface RunAgentResponse {
  agent_id: string;
  result: string;
  raw_output?: any;
}

export const useAgentStore = defineStore('agent', {
  state: () => ({
    agents: [] as Agent[],
    currentAgent: null as Agent | null,
    tools: [] as Tool[],
    loading: false,
    error: null as string | null,
    runResult: null as RunAgentResponse | null,
  }),
  
  getters: {
    getAgentById: (state) => (id: string) => {
      return state.agents.find(agent => agent.id === id) || null;
    },
    
    toolsOptions: (state) => {
      return state.tools.map(tool => ({
        value: tool.name,
        label: `${tool.name} - ${tool.description}`
      }));
    },
  },
  
  actions: {
    async fetchAgents() {
      this.loading = true;
      this.error = null;
      
      try {
        const agents = await ListAgents();
        this.agents = agents;
      } catch (err: any) {
        this.error = err.toString();
        console.error('Failed to fetch agents:', err);
      } finally {
        this.loading = false;
      }
    },
    
    async fetchAgent(id: string) {
      this.loading = true;
      this.error = null;
      
      try {
        const agent = await GetAgent(id);
        this.currentAgent = agent;
        return agent;
      } catch (err: any) {
        this.error = err.toString();
        console.error(`Failed to fetch agent ${id}:`, err);
        return null;
      } finally {
        this.loading = false;
      }
    },
    
    async createAgent(
      name: string, 
      instructions: string, 
      handoffDescription: string, 
      model: string,
      temperature: number,
      topP: number,
      maxTokens: number,
      tools: string[]
    ) {
      this.loading = true;
      this.error = null;
      
      try {
        const agent = await CreateAgent(
          name, 
          instructions, 
          handoffDescription, 
          model, 
          temperature, 
          topP, 
          maxTokens, 
          tools
        );
        
        this.agents.push(agent);
        return agent;
      } catch (err: any) {
        this.error = err.toString();
        console.error('Failed to create agent:', err);
        return null;
      } finally {
        this.loading = false;
      }
    },
    
    async deleteAgent(id: string) {
      this.loading = true;
      this.error = null;
      
      try {
        await DeleteAgent(id);
        this.agents = this.agents.filter(agent => agent.id !== id);
        if (this.currentAgent && this.currentAgent.id === id) {
          this.currentAgent = null;
        }
        return true;
      } catch (err: any) {
        this.error = err.toString();
        console.error(`Failed to delete agent ${id}:`, err);
        return false;
      } finally {
        this.loading = false;
      }
    },
    
    async runAgent(agentId: string, input: string, context: any = {}) {
      this.loading = true;
      this.error = null;
      this.runResult = null;
      
      try {
        const result = await RunAgent(agentId, input, context);
        this.runResult = result;
        return result;
      } catch (err: any) {
        this.error = err.toString();
        console.error(`Failed to run agent ${agentId}:`, err);
        return null;
      } finally {
        this.loading = false;
      }
    },
    
    async cloneAgent(agentId: string, updates: any) {
      this.loading = true;
      this.error = null;
      
      try {
        const agent = await CloneAgent(agentId, updates);
        this.agents.push(agent);
        return agent;
      } catch (err: any) {
        this.error = err.toString();
        console.error(`Failed to clone agent ${agentId}:`, err);
        return null;
      } finally {
        this.loading = false;
      }
    },
    
    async fetchTools() {
      this.loading = true;
      this.error = null;
      
      try {
        const tools = await GetTools();
        this.tools = tools;
        return tools;
      } catch (err: any) {
        this.error = err.toString();
        console.error('Failed to fetch tools:', err);
        return [];
      } finally {
        this.loading = false;
      }
    },
  },
});
