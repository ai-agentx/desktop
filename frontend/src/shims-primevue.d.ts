declare module 'primevue/usetoast' {
  import { ToastServiceMethods } from 'primevue/toastservice';
  export function useToast(): ToastServiceMethods;
}

declare module 'primevue/useconfirm' {
  import { ConfirmationServiceMethods } from 'primevue/confirmationservice';
  export function useConfirm(): ConfirmationServiceMethods;
}

declare module 'primevue/usedialog' {
  import { DialogServiceMethods } from 'primevue/dialogservice';
  export function useDialog(): DialogServiceMethods;
}
