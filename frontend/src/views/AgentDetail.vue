<template>
  <div class="agent-detail">
    <div v-if="agentStore.loading" class="loading-spinner">
      <p-progress-spinner />
    </div>
    
    <div v-else-if="agentStore.error" class="p-4 text-center">
      <p class="text-red-500">{{ agentStore.error }}</p>
      <p-button label="Go Back" icon="pi pi-arrow-left" @click="goBack" class="mt-3" />
    </div>
    
    <div v-else-if="agent">
      <div class="flex align-items-center justify-content-between mb-3">
        <div class="flex align-items-center">
          <p-button icon="pi pi-arrow-left" class="p-button-text mr-2" @click="goBack" />
          <h2 class="m-0">{{ agent.name }}</h2>
        </div>
        <div>
          <p-button 
            icon="pi pi-copy" 
            class="p-button-secondary mr-2" 
            @click="showCloneDialog = true" 
            tooltip="Clone Agent"
          />
          <p-button 
            icon="pi pi-trash" 
            class="p-button-danger" 
            @click="confirmDelete" 
            tooltip="Delete Agent"
          />
        </div>
      </div>
      
      <div class="grid">
        <div class="col-12 md:col-8">
          <!-- Agent Information Card -->
          <p-card class="mb-3">
            <template #title>Agent Information</template>
            <template #content>
              <div class="field">
                <label class="font-bold">ID:</label>
                <div>{{ agent.id }}</div>
              </div>
              
              <div class="field mt-3">
                <label class="font-bold">Instructions:</label>
                <p class="whitespace-pre-line">{{ agent.instructions }}</p>
              </div>
              
              <div v-if="agent.handoff_description" class="field mt-3">
                <label class="font-bold">Handoff Description:</label>
                <p>{{ agent.handoff_description }}</p>
              </div>
              
              <div class="field mt-3">
                <label class="font-bold">Model:</label>
                <p>{{ agent.model || 'Default Model' }}</p>
              </div>
              
              <div class="field mt-3">
                <label class="font-bold">Tools:</label>
                <div v-if="agent.tools && agent.tools.length > 0">
                  <p-chip 
                    v-for="tool in agent.tools" 
                    :key="tool" 
                    :label="tool" 
                    class="mr-2 mb-2"
                  />
                </div>
                <p v-else class="text-500">No tools</p>
              </div>
            </template>
          </p-card>
          
          <!-- Run Agent Card -->
          <p-card>
            <template #title>Run Agent</template>
            <template #content>
              <div class="field">
                <label for="agentInput">Input:</label>
                <p-textarea 
                  id="agentInput" 
                  v-model="runInput" 
                  rows="5" 
                  placeholder="Enter input for the agent..."
                  class="w-full"
                />
              </div>
              
              <div class="field">
                <label for="agentContext">Context (JSON):</label>
                <p-textarea 
                  id="agentContext" 
                  v-model="runContextText" 
                  rows="3" 
                  placeholder="{}"
                  class="w-full" 
                />
                <small v-if="contextError" class="p-error">{{ contextError }}</small>
              </div>
              
              <div class="flex justify-content-end">
                <p-button 
                  label="Run" 
                  icon="pi pi-play" 
                  @click="runAgent" 
                  :loading="agentStore.loading"
                  :disabled="!runInput"
                />
              </div>
            </template>
          </p-card>
        </div>
        
        <div class="col-12 md:col-4">
          <!-- Results Card -->
          <p-card>
            <template #title>Results</template>
            <template #content>
              <div v-if="agentStore.runResult" class="result-container">
                <h4>Response:</h4>
                <div class="p-2 border-round bg-gray-100 whitespace-pre-line mb-3">
                  {{ agentStore.runResult.result }}
                </div>
                
                <h4>Raw Output:</h4>
                <pre class="p-2 border-round bg-gray-100 overflow-auto">{{ JSON.stringify(agentStore.runResult.raw_output, null, 2) }}</pre>
              </div>
              <div v-else class="text-center p-4 text-500">
                Run the agent to see results here
              </div>
            </template>
          </p-card>
        </div>
      </div>
      
      <!-- Clone Dialog -->
      <p-dialog 
        v-model:visible="showCloneDialog" 
        header="Clone Agent" 
        :style="{ width: '450px' }" 
        :modal="true"
      >
        <div class="p-fluid">
          <div class="field">
            <label for="cloneName">New Agent Name:</label>
            <p-inputtext id="cloneName" v-model="cloneName" autofocus />
          </div>
          
          <div class="field">
            <label for="cloneInstructions">New Instructions (Optional):</label>
            <p-textarea id="cloneInstructions" v-model="cloneInstructions" rows="5" />
          </div>
        </div>
        
        <template #footer>
          <p-button 
            label="Cancel" 
            icon="pi pi-times" 
            @click="showCloneDialog = false" 
            class="p-button-text"
          />
          <p-button 
            label="Clone" 
            icon="pi pi-copy" 
            @click="cloneAgent" 
            :loading="agentStore.loading"
            :disabled="!cloneName"
          />
        </template>
      </p-dialog>
      
      <!-- Confirm Delete Dialog -->
      <p-confirm-dialog></p-confirm-dialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAgentStore } from '../stores/AgentStore';
import { useToast } from 'primevue/usetoast';
import { useConfirm } from 'primevue/useconfirm';

const route = useRoute();
const router = useRouter();
const agentStore = useAgentStore();
const toast = useToast();
const confirm = useConfirm();

const agentId = computed(() => route.params.id as string);
const agent = computed(() => agentStore.currentAgent);

const runInput = ref('');
const runContextText = ref('{}');
const contextError = ref('');

const showCloneDialog = ref(false);
const cloneName = ref('');
const cloneInstructions = ref('');

onMounted(async () => {
  await agentStore.fetchAgent(agentId.value);
});

const goBack = () => {
  router.push('/agents');
};

const runAgent = async () => {
  contextError.value = '';
  let contextObj = {};
  
  try {
    if (runContextText.value.trim() !== '' && runContextText.value !== '{}') {
      contextObj = JSON.parse(runContextText.value);
    }
  } catch (error) {
    contextError.value = 'Invalid JSON context';
    return;
  }
  
  try {
    await agentStore.runAgent(agentId.value, runInput.value, contextObj);
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to run agent',
      life: 3000
    });
  }
};

const cloneAgent = async () => {
  if (!cloneName.value) return;
  
  const updates: any = {
    name: cloneName.value
  };
  
  if (cloneInstructions.value) {
    updates.instructions = cloneInstructions.value;
  }
  
  try {
    const newAgent = await agentStore.cloneAgent(agentId.value, updates);
    
    if (newAgent) {
      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: `Agent "${newAgent.name}" cloned successfully`,
        life: 3000
      });
      
      showCloneDialog.value = false;
      router.push(`/agents/${newAgent.id}`);
    }
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to clone agent',
      life: 3000
    });
  }
};

const confirmDelete = () => {
  confirm.require({
    message: `Are you sure you want to delete agent "${agent.value?.name}"?`,
    header: 'Delete Confirmation',
    icon: 'pi pi-exclamation-triangle',
    acceptClass: 'p-button-danger',
    accept: async () => {
      try {
        const success = await agentStore.deleteAgent(agentId.value);
        
        if (success) {
          toast.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Agent deleted successfully',
            life: 3000
          });
          
          router.push('/agents');
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

<style scoped>
.result-container {
  max-height: 600px;
  overflow-y: auto;
}

pre {
  margin: 0;
  font-size: 0.9em;
}
</style>
