<script setup lang="ts">
const { data } = await useFetch('/api/provider');
</script>

<template>
  <p>providers: {{ data }}</p>
</template>
