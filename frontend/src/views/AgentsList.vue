<template>
  <div class="agents-list">
    <div class="flex justify-content-between align-items-center mb-3">
      <h2>Agents</h2>
      <p-button 
        label="Create New Agent" 
        icon="pi pi-plus" 
        @click="navigateToCreate" 
      />
    </div>
    
    <p-card class="mb-3">
      <template #content>
        <p-data-table 
          :value="agentStore.agents" 
          :loading="agentStore.loading"
          :paginator="agentStore.agents.length > 10" 
          :rows="10"
          striped-rows
          responsive-layout="stack"
          class="p-datatable-sm"
          v-model:selection="selectedAgent"
          selection-mode="single"
          data-key="id"
          :empty-message="agentStore.error || 'No agents found'"
          @row-select="onRowSelect"
        >
          <template #empty>
            <div class="text-center p-4">
              <p v-if="agentStore.error" class="text-red-500">{{ agentStore.error }}</p>
              <p v-else>No agents found. Create one to get started!</p>
            </div>
          </template>
          
          <p-column field="id" header="ID" :sortable="true">
            <template #body="{ data }">
              <span class="text-sm text-500">{{ truncateId(data.id) }}</span>
            </template>
          </p-column>
          
          <p-column field="name" header="Name" :sortable="true">
            <template #body="{ data }">
              <div class="flex align-items-center">
                <span class="font-bold">{{ data.name }}</span>
              </div>
            </template>
          </p-column>
          
          <p-column field="instructions" header="Instructions">
            <template #body="{ data }">
              <div class="text-sm text-truncate" style="max-width: 300px;">
                {{ data.instructions }}
              </div>
            </template>
          </p-column>
          
          <p-column field="model" header="Model">
            <template #body="{ data }">
              <span v-if="data.model" class="text-sm">{{ data.model }}</span>
              <span v-else class="text-sm text-500">Default</span>
            </template>
          </p-column>
          
          <p-column field="tools" header="Tools">
            <template #body="{ data }">
              <div v-if="data.tools && data.tools.length > 0" class="flex flex-wrap gap-1">
                <p-chip 
                  v-for="tool in data.tools.slice(0, 2)" 
                  :key="tool" 
                  :label="tool" 
                  class="text-xs"
                />
                <p-chip 
                  v-if="data.tools.length > 2" 
                  :label="`+${data.tools.length - 2}`" 
                  class="text-xs bg-primary text-white"
                />
              </div>
              <span v-else class="text-sm text-500">No tools</span>
            </template>
          </p-column>
          
          <p-column style="width: 100px">
            <template #body="{ data }">
              <div class="flex justify-content-end">
                <p-button 
                  icon="pi pi-ellipsis-v" 
                  class="p-button-text p-button-rounded" 
                  @click="openMenu($event, data)" 
                />
              </div>
            </template>
          </p-column>
        </p-data-table>
      </template>
    </p-card>
    
    <p-menu id="agent-menu" ref="menu" :model="menuModel" :popup="true" />
    <p-confirm-dialog></p-confirm-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAgentStore } from '../stores/AgentStore';
import { useToast } from 'primevue/usetoast';
import { useConfirm } from 'primevue/useconfirm';

interface Agent {
  id: string;
  name: string;
  instructions?: string;
  handoff_description?: string;
  model?: string;
  model_settings?: Record<string, any>;
  tools?: string[];
}

const router = useRouter();
const agentStore = useAgentStore();
const toast = useToast();
const confirm = useConfirm();
const menu = ref();
const selectedAgent = ref<Agent | null>(null);
const currentAgent = ref<Agent | null>(null);

const menuModel = ref([
  {
    label: 'View Details',
    icon: 'pi pi-search',
    command: () => {
      if (currentAgent.value) {
        router.push(`/agents/${currentAgent.value.id}`);
      }
    }
  },
  {
    label: 'Run Agent',
    icon: 'pi pi-play',
    command: () => {
      if (currentAgent.value) {
        router.push(`/agents/${currentAgent.value.id}`);
      }
    }
  },
  { separator: true },
  {
    label: 'Delete',
    icon: 'pi pi-trash',
    command: () => {
      if (currentAgent.value) {
        confirmDelete(currentAgent.value);
      }
    }
  }
]);

onMounted(async () => {
  await agentStore.fetchAgents();
});

const truncateId = (id: string) => {
  if (!id) return '';
  return id.length > 8 ? `${id.substring(0, 8)}...` : id;
};

const navigateToCreate = () => {
  router.push('/agents/create');
};

const onRowSelect = (event: any) => {
  router.push(`/agents/${event.data.id}`);
};

const openMenu = (event: any, agent: any) => {
  currentAgent.value = agent;
  menu.value.toggle(event);
};

const confirmDelete = (agent: any) => {
  confirm.require({
    message: `Are you sure you want to delete agent "${agent.name}"?`,
    header: 'Delete Confirmation',
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: async () => {
      try {
        const success = await agentStore.deleteAgent(agent.id);
        
        if (success) {
          toast.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Agent deleted successfully',
            life: 3000
          });
        }
      } catch (error) {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: 'Failed to delete agent',
          life: 3000
        });
      }
    }
  });
};
</script>
