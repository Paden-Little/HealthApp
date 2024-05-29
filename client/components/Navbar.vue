<script setup lang="ts">
import { initFlowbite } from 'flowbite';

const type = useCookie('type');
const pid = useCookie('pid');

function isLoggedIn() {
  if (pid.value !== undefined && pid.value !== undefined) {
    return true;
  } else {
    return false;
  }
}

onMounted(() => {
  initFlowbite();
});
</script>

<template>
  <header>
    <nav class="px-4 py-2.5 shadow lg:px-6">
      <div
        class="mx-auto flex max-w-screen-xl flex-wrap items-center justify-between"
      >
        <NuxtLink to="/" class="flex items-center">
          <img class="mr-3 h-6 sm:h-9" src="/img/heart.png" />
          <span
            class="self-center whitespace-nowrap text-xl font-semibold text-blue-900"
            >HealthMark</span
          >
        </NuxtLink>
        <div class="flex items-center lg:order-2">
          <NuxtLink
            to="/provider/dashboard"
            v-if="isLoggedIn() && type === 'provider'"
            class="secondary-btn"
          >
            Dashboard
          </NuxtLink>
          <NuxtLink
            v-else-if="isLoggedIn() && type === 'patient'"
            to="/my-health/dashboard"
            class="secondary-btn"
          >
            Dashboard
          </NuxtLink>
          <NuxtLink
            to="/my-health/authentication/login"
            class="secondary-btn"
            v-else
            >Log in</NuxtLink
          >
          <NuxtLink to="/find-a-provider" class="cta-btn"
            >Book an Appointment</NuxtLink
          >
          <button
            data-collapse-toggle="mobile-menu-2"
            type="button"
            class="ml-1 inline-flex items-center rounded-lg p-2 text-sm hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 lg:hidden"
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
            class="mt-4 flex flex-col items-center font-medium lg:mt-0 lg:flex-row lg:space-x-8"
          >
            <li>
              <NuxtLink
                to="/find-a-provider"
                class="navlink"
                :class="{ activelink: $route.path === '/find-a-provider' }"
                aria-current="page"
                >Find a Provider</NuxtLink
              >
            </li>
            <li>
              <NuxtLink
                to="/services"
                class="navlink"
                :class="{ activelink: $route.path === '/services' }"
                aria-current="page"
                >Services</NuxtLink
              >
            </li>
            <li>
              <NuxtLink
                to="/about"
                class="navlink"
                :class="{ activelink: $route.path === '/about' }"
                aria-current="page"
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
  @apply font-bold tracking-tight text-gray-600 transition-all hover:text-blue-800;
}
.activelink {
  @apply text-blue-800;
}
</style>
