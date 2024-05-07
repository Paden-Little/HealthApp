<script setup lang="ts">
const responses = reactive<string[]>([]);

async function getServerData() {
  const resp = await $fetch<string>('/api/record');
  responses.push(resp)
}
</script>

<template>
  <button @click="getServerData()">GET http://localhost:3000</button>
  <li v-for="response in responses">
    {{ response }}
  </li>
  <NuxtPage />
</template>
