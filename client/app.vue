<script setup lang="ts">
let pid = useCookie("pid");
let token = useCookie("token");
let type = useCookie("type");

onMounted(() => {
  pid.value = '';
  token.value = '';
  type.value = '';
});
</script>
<template>
  <NuxtLoadingIndicator />
  <Navbar />
  <NuxtPage />
</template>
