<script setup>
import { initFlowbite } from 'flowbite';

onMounted(() => {
  initFlowbite();
});
</script>

<template>
  <header>
    <nav class="bg-white px-4 py-2.5 lg:px-6">
      <div
        class="mx-auto flex max-w-screen-xl flex-wrap items-center justify-between"
      >
        <a class="flex items-center">
          <img class="mr-3 h-6 sm:h-9" />
          <span
            class="self-center whitespace-nowrap text-xl font-semibold text-black"
            >HealthMark</span
          >
        </a>
        <div class="flex items-center lg:order-2">
          <a
            href="#"
            class="mr-2 rounded-lg px-4 py-2 text-sm font-medium text-black hover:bg-gray-50 focus:outline-none focus:ring-4 focus:ring-gray-300 lg:px-5 lg:py-2.5"
            >Log in</a
          >
          <a
            href="#"
            class="mr-2 rounded-lg bg-primary-700 px-4 py-2 text-sm font-medium text-white hover:bg-primary-800 focus:outline-none focus:ring-4 focus:ring-primary-300 lg:px-5 lg:py-2.5"
            >Book an Appointment</a
          >
          <button
            data-collapse-toggle="mobile-menu-2"
            type="button"
            class="ml-1 inline-flex items-center rounded-lg p-2 text-sm text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 lg:hidden"
            aria-controls="mobile-menu-2"
            aria-expanded="false"
          >
            <span class="sr-only">Open main menu</span>
            <svg
              class="h-6 w-6"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <svg
              class="hidden h-6 w-6"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </button>
        </div>
        <div
          class="hidden w-full items-center justify-between lg:order-1 lg:flex lg:w-auto"
          id="mobile-menu-2"
        >
          <ul
            class="mt-4 flex flex-col font-medium lg:mt-0 lg:flex-row lg:space-x-8"
          >
            <li>
              <NuxtLink to="/" class="navlink" aria-current="page"
                >Home</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="/find-a-doctor" class="navlink" aria-current="page"
                >Find a Doctor</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="/services" class="navlink" aria-current="page"
                >Services</NuxtLink
              >
            </li>
            <li>
              <NuxtLink to="/about" class="navlink" aria-current="page"
                >About</NuxtLink
              >
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>
</template>

<style scoped>
.navlink {
  @apply block rounded py-2 pl-3 pr-4 text-black lg:bg-transparent lg:p-0;
}
.activelink {
  @apply block rounded bg-[#243566] py-2 pl-3 pr-4 text-white lg:bg-transparent lg:p-0 lg:text-primary-700;
}
</style>
