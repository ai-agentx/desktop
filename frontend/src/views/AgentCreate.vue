<template>
  <div class="agent-create">
    <h2>Create New Agent</h2>
    
    <p-card class="mt-3">
      <template #content>
        <form @submit.prevent="submitForm" class="p-fluid">
          <div class="field">
            <label for="name">Agent Name*</label>
            <p-inputtext id="name" v-model="agent.name" :class="{ 'p-invalid': v$.name.$invalid && submitted }" />
            <small v-if="v$.name.$invalid && submitted" class="p-error">{{ v$.name.$errors[0].$message }}</small>
          </div>
          
          <div class="field">
            <label for="instructions">Instructions*</label>
            <p-textarea 
              id="instructions" 
              v-model="agent.instructions" 
              rows="5" 
              :class="{ 'p-invalid': v$.instructions.$invalid && submitted }"
            />
            <small v-if="v$.instructions.$invalid && submitted" class="p-error">{{ v$.instructions.$errors[0].$message }}</small>
          </div>
          
          <div class="field">
            <label for="handoffDescription">Handoff Description</label>
            <p-textarea id="handoffDescription" v-model="agent.handoffDescription" rows="3" />
          </div>
          
          <div class="field">
            <label for="model">Model</label>
            <p-inputtext id="model" v-model="agent.model" placeholder="e.g., gpt-4" />
          </div>
          
          <div class="formgrid grid">
            <div class="field col">
              <label for="temperature">Temperature</label>
              <p-inputnumber 
                id="temperature" 
                v-model="agent.temperature" 
                :min="0" 
                :max="2" 
                :step="0.1" 
                mode="decimal" 
              />
            </div>
            
            <div class="field col">
              <label for="topP">Top P</label>
              <p-inputnumber 
                id="topP" 
                v-model="agent.topP" 
                :min="0" 
                :max="1" 
                :step="0.05" 
                mode="decimal" 
              />
            </div>
            
            <div class="field col">
              <label for="maxTokens">Max Tokens</label>
              <p-inputnumber 
                id="maxTokens" 
                v-model="agent.maxTokens" 
                :min="1" 
                :step="100" 
              />
            </div>
          </div>
          
          <div class="field">
            <label for="tools">Tools</label>
            <p-multi-select 
              id="tools" 
              v-model="agent.tools" 
              :options="agentStore.toolsOptions" 
              option-label="label" 
              option-value="value"
              placeholder="Select Tools" 
              class="w-full"
              display="chip"
            />
          </div>
          
          <div class="flex justify-content-between mt-4">
            <p-button 
              type="button" 
              label="Cancel" 
              icon="pi pi-times" 
              class="p-button-secondary" 
              @click="goBack"
            />
            <p-button 
              type="submit" 
              label="Create Agent" 
              icon="pi pi-check" 
              :loading="agentStore.loading"
            />
          </div>
        </form>
      </template>
    </p-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAgentStore } from '../stores/AgentStore';
import { useToast } from 'primevue/usetoast';
import { useVuelidate } from '@vuelidate/core';
import { required, minLength } from '@vuelidate/validators';

const router = useRouter();
const agentStore = useAgentStore();
const toast = useToast();
const submitted = ref(false);

const agent = reactive({
  name: '',
  instructions: '',
  handoffDescription: '',
  model: '',
  temperature: 0.7,
  topP: 0.9,
  maxTokens: 1000,
  tools: [] as string[]
});

const rules = computed(() => {
  return {
    name: { required, minLength: minLength(3) },
    instructions: { required, minLength: minLength(10) }
  };
});

const v$ = useVuelidate(rules, agent);

onMounted(async () => {
  if (agentStore.tools.length === 0) {
    await agentStore.fetchTools();
  }
});

const submitForm = async () => {
  submitted.value = true;
  const isValid = await v$.value.$validate();
  
  if (!isValid) {
    toast.add({ 
      severity: 'error', 
      summary: 'Validation Error', 
      detail: 'Please check the form for errors', 
      life: 3000 
    });
    return;
  }
  
  try {
    const newAgent = await agentStore.createAgent(
      agent.name,
      agent.instructions,
      agent.handoffDescription,
      agent.model,
      agent.temperature,
      agent.topP,
      agent.maxTokens,
      agent.tools
    );
    
    if (newAgent) {
      toast.add({ 
        severity: 'success', 
        summary: 'Success', 
        detail: `Agent "${newAgent.name}" created successfully`, 
        life: 3000 
      });
      router.push('/agents');
    }
  } catch (error) {
    console.error('Error creating agent:', error);
  }
};

const goBack = () => {
  router.push('/agents');
};
</script>
