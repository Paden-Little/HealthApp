<script setup lang="ts">
const { data } = await useFetch<Provider[]>('/api/provider');

const searchQuery = ref('');

// Filtered list of providers
const filteredProviders = computed(() => {
  return data.value!.filter(provider =>
    provider.firstname.toLowerCase().includes(searchQuery.value.toLowerCase()),
  );
});
</script>

<template>
  <section class="bg-slate-200">
    <div class="mx-auto max-w-screen-xl px-4 py-8 lg:px-6 lg:py-16">
      <div class="mx-auto max-w-screen-sm text-center">
        <h2 class="mb-4 text-4xl font-bold tracking-tight text-gray-900">
          Find a Provider
        </h2>
        <form class="mx-auto max-w-md">
          <label
            for="default-search"
            class="sr-only mb-2 text-sm font-medium text-gray-900"
            >Search</label
          >
          <div class="relative">
            <div
              class="pointer-events-none absolute inset-y-0 start-0 flex items-center ps-3"
            >
              <Icon name="material-symbols:search-rounded" />
            </div>
            <input
              type="search"
              id="default-search"
              class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-4 ps-10 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500"
              placeholder="Filter by providers name"
              required
              v-model="searchQuery"
            />
          </div>
        </form>
      </div>
    </div>
  </section>
  <section class="bg-white">
    <div class="mx-auto flex max-w-screen-xl flex-col items-center">
      <p
        class="mt-4 w-full border-b-[1px] border-gray-400 pb-2 text-xl tracking-tight text-gray-900"
      >
        {{ filteredProviders.length }} Resulting Providers
      </p>
      <div v-if="filteredProviders.length != 0" class="w-full">
        <ProviderListing
          v-for="provider in filteredProviders"
          :provider="provider"
        />
      </div>
      <div v-else class="flex h-48 items-center">
        <p class="text-4xl tracking-tighter">No providers found...</p>
      </div>
    </div>
  </section>
</template>
