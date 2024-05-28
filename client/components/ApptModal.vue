<script setup lang="ts">
import { initFlowbite } from 'flowbite';

defineProps({
  provider: {
    type: Object as () => Provider,
  },
});

onMounted(() => {
  initFlowbite();
});
</script>

<template>
  <div
    :id="'apptModal' + provider?.email"
    tabindex="-1"
    aria-hidden="true"
    class="fixed left-0 right-0 top-0 z-50 hidden h-[calc(100%-1rem)] max-h-full w-full items-center justify-center overflow-y-auto overflow-x-hidden md:inset-0"
  >
    <div class="relative max-h-full w-full max-w-2xl p-4">
      <!-- Modal content -->
      <div class="relative rounded-lg bg-white shadow">
        <!-- Modal header -->
        <div class="flex justify-between rounded-t border-b p-4 md:p-5">
          <div class="flex flex-col">
            <p class="text-lg font-semibold">
              Book with {{ provider?.firstname }} {{ provider?.lastname }},
              {{ provider?.suffix }}
            </p>
            <img
              class="mt-2 w-20 rounded-full"
              src="/img/skillissue.jpg"
              alt=""
            />
          </div>
          <button
            type="button"
            class="ms-auto inline-flex h-8 w-8 items-center justify-center rounded-lg bg-transparent text-sm text-gray-400 hover:bg-gray-200 hover:text-gray-900"
            :data-modal-hide="'apptModal' + provider?.email"
          >
            <Icon name="ic:round-close" />
            <span class="sr-only">Close modal</span>
          </button>
        </div>
        <!-- Modal body -->
        <div class="space-y-4 p-4 md:p-5">
          <div>
            <label for="email"></label>
            <input type="text" name="" id="" />
          </div>
        </div>
        <!-- Modal footer -->
        <div
          class="flex items-center rounded-b border-t border-gray-200 p-4 md:p-5"
        >
          <button
            :data-modal-hide="'apptModal' + provider?.email"
            type="button"
            class="cta-btn"
          >
            Book Appointment
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
