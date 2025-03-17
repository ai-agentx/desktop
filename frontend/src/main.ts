import { createApp } from 'vue';
import { createRouter, createWebHashHistory } from 'vue-router';
import { createPinia } from 'pinia';
import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';
import ConfirmationService from 'primevue/confirmationservice';
import DialogService from 'primevue/dialogservice';

import App from './App.vue';

// Import views using raw-loader to avoid type checking issues
const AgentsList = () => import('./views/AgentsList.vue');
const AgentCreate = () => import('./views/AgentCreate.vue');
const AgentDetail = () => import('./views/AgentDetail.vue');
const ToolsList = () => import('./views/ToolsList.vue');

// Define routes
const routes = [
  { path: '/', redirect: '/agents' },
  { path: '/agents', component: AgentsList },
  { path: '/agents/create', component: AgentCreate },
  { path: '/agents/:id', component: AgentDetail },
  { path: '/tools', component: ToolsList }
];

// Create router
const router = createRouter({
  history: createWebHashHistory(),
  routes
});

// Create pinia store
const pinia = createPinia();

// Create app
const app = createApp(App);

// Configure plugins
app.use(PrimeVue, { ripple: true });
app.use(ToastService);
app.use(ConfirmationService);
app.use(DialogService);
app.use(router);
app.use(pinia);

// Mount app
app.mount('#app');
