<template>
  <div class="tools-list">
    <h2>Available Tools</h2>
    
    <p-card class="mt-3">
      <template #content>
        <p-data-table 
          :value="agentStore.tools" 
          :loading="agentStore.loading"
          striped-rows
          responsive-layout="stack"
          class="p-datatable-sm"
          :empty-message="agentStore.error || 'No tools found'"
        >
          <template #empty>
            <div class="text-center p-4">
              <p v-if="agentStore.error" class="text-red-500">{{ agentStore.error }}</p>
              <p v-else>No tools found</p>
            </div>
          </template>
          
          <p-column field="name" header="Name" :sortable="true">
            <template #body="{ data }">
              <span class="font-bold">{{ data.name }}</span>
            </template>
          </p-column>
          
          <p-column field="description" header="Description">
            <template #body="{ data }">
              <span>{{ data.description }}</span>
            </template>
          </p-column>
        </p-data-table>
      </template>
    </p-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useAgentStore } from '../stores/AgentStore';

const agentStore = useAgentStore();

onMounted(async () => {
  await agentStore.fetchTools();
});
</script>
