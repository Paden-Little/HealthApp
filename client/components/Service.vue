<script setup lang="ts">
defineProps({
  name: String,
  img: String,
  description: String,
});
</script>
<template>
  <main class="rounded-md bg-white p-4 drop-shadow-2xl flex content-center flex-col">
    <img :src="img" />
    <h1
      class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl"
    >
      {{ name }}
    </h1>
    <p>{{ description }}</p>
  </main>
</template>
