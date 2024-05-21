<script setup lang="ts">
const { data } = await useFetch('/api/provider');
</script>

<template>
  <section class="bg-white dark:bg-gray-900">
    <div class="mx-auto max-w-screen-xl px-4 py-8 lg:px-6 lg:py-16">
      <div class="mx-auto mb-8 max-w-screen-sm text-center lg:mb-16">
        <h2
          class="mb-4 text-4xl font-extrabold tracking-tight text-gray-900 dark:text-white"
        >
          Find a Provider
        </h2>
        <p
          class="font-light text-gray-500 dark:text-gray-400 sm:text-xl lg:mb-16"
        >
          Explore the whole collection of open-source web components and
          elements built with the utility classes from Tailwind
        </p>
      </div>
      <div class="mb-6 grid gap-8 md:grid-cols-2 lg:mb-16">
        <ProviderProfile v-for="provider in data" :provider="provider" />
      </div>
    </div>
  </section>
</template>
