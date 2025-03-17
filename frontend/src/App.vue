<template>
  <div id="app">
    <Toast position="top-right"></Toast>
    <div class="menubar">
      <div class="menubar-start">
        <div class="flex align-items-center">
          <img src="./assets/logo.svg" height="40" class="mr-2" />
          <h1 class="m-0">AgentX</h1>
        </div>
      </div>
      <div class="menubar-end">
        <button class="icon-button" @click="showSettings = true">
          <span class="pi pi-cog"></span>
        </button>
      </div>
    </div>
    
    <div class="content">
      <router-view />
    </div>
    
    <!-- Settings Dialog -->
    <Dialog v-model:visible="showSettings" header="Settings" :style="{ width: '50vw' }">
      <div class="p-fluid">
        <div class="field">
          <label for="apiUrl">API URL</label>
          <input id="apiUrl" v-model="apiUrl" class="p-inputtext" />
        </div>
      </div>
      <template #footer>
        <button class="p-button p-button-text" @click="showSettings = false">
          <span class="pi pi-times"></span>
          Cancel
        </button>
        <button class="p-button" @click="saveSettings">
          <span class="pi pi-check"></span>
          Save
        </button>
      </template>
    </Dialog>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import Dialog from 'primevue/dialog';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { GetAPIURL, SetAPIURL } from './wailsjs/wailsjs/go/main/App';

export default {
  components: {
    Dialog,
    Toast
  },
  setup() {
    const router = useRouter();
    const toast = useToast();

    const showSettings = ref(false);
    const apiUrl = ref('http://localhost:9091');

    onMounted(async () => {
      try {
        const currentUrl = await GetAPIURL();
        apiUrl.value = currentUrl;
      } catch (error) {
        console.error('Failed to get current API URL:', error);
      }
    });

    const saveSettings = async () => {
      try {
        const result = await SetAPIURL(apiUrl.value);
        toast.add({ severity: 'success', summary: 'Settings Saved', detail: result, life: 3000 });
        showSettings.value = false;
      } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to save settings', life: 3000 });
      }
    };

    const navigateToAgents = () => {
      router.push('/agents');
    };

    const navigateToTools = () => {
      router.push('/tools');
    };

    return {
      showSettings,
      apiUrl,
      saveSettings,
      navigateToAgents,
      navigateToTools
    };
  }
}
</script>

<style>
/* Add your global styles here */
@import 'primevue/resources/themes/lara-light-indigo/theme.css';
@import 'primevue/resources/primevue.min.css';
@import 'primeicons/primeicons.css';

html, body, #app {
  height: 100%;
  margin: 0;
  font-family: var(--font-family);
}

#app {
  display: flex;
  flex-direction: column;
}

.menubar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.5rem 1rem;
  background-color: var(--surface-0);
  border-bottom: 1px solid var(--surface-200);
}

.menubar-start, .menubar-end {
  display: flex;
  align-items: center;
}

.content {
  flex: 1;
  padding: 1rem;
  overflow-y: auto;
}

.icon-button {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 50%;
  border: none;
  background-color: transparent;
  cursor: pointer;
  transition: background-color 0.2s;
}

.icon-button:hover {
  background-color: var(--surface-200);
}

.flex {
  display: flex;
}

.align-items-center {
  align-items: center;
}

.m-0 {
  margin: 0;
}

.mr-2 {
  margin-right: 0.5rem;
}

.field {
  margin-bottom: 1rem;
}

.field label {
  display: block;
  margin-bottom: 0.5rem;
}

.p-inputtext {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid var(--surface-300);
  border-radius: 4px;
}

.p-button {
  display: inline-flex;
  align-items: center;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  border: none;
  background-color: var(--primary-color);
  color: var(--primary-color-text);
  cursor: pointer;
  transition: background-color 0.2s;
}

.p-button-text {
  background-color: transparent;
  color: var(--text-color);
}

.p-button .pi {
  margin-right: 0.5rem;
}
</style>
