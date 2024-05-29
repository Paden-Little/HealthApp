<script setup lang="ts">
import { initFlowbite } from 'flowbite';
defineProps({
  name: String,
  img: String,
  description: String,
});

onMounted(() => {
  initFlowbite();
});
</script>
<template>
  <div
    class="flex flex-col items-center rounded-lg border border-gray-200 bg-white shadow hover:bg-gray-100 md:max-w-xl md:flex-row"
  >
    <img
      class="h-96 w-full rounded-t-lg object-cover md:h-auto md:w-48 md:rounded-none md:rounded-s-lg"
      :src="img"
    />
    <div class="flex flex-col justify-between p-4 leading-normal">
      <h5
        class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
      >
        {{ name }}
      </h5>
      <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">
        {{ description }}
      </p>
    </div>
  </div>
</template>
