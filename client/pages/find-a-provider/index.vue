<script setup lang="ts">
const { data } = await useFetch('/api/provider');
</script>

<template>
  <section class="bg-slate-200">
    <div class="mx-auto max-w-screen-xl px-4 py-8 lg:px-6 lg:py-16">
      <div class="mx-auto max-w-screen-sm text-center">
        <h2 class="mb-4 text-4xl font-extrabold tracking-tight text-gray-900">
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
              <svg
                class="h-4 w-4 text-gray-500"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 20 20"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                />
              </svg>
            </div>
            <input
              type="search"
              id="default-search"
              class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-4 ps-10 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500"
              placeholder="Filter by providers name"
              required
            />
          </div>
        </form>
      </div>
    </div>
  </section>
  <section class="bg-white">
    <div class="mx-auto flex max-w-screen-xl flex-col">
      <p class="self-center pb-4 text-3xl text-gray-900">Resulting Providers</p>
      <div class="mb-6 grid gap-8 md:grid-cols-2 lg:mb-16">
        <ProviderProfile v-for="provider in data" :provider="provider" />
      </div>
    </div>
  </section>
</template>
