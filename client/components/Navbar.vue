<template>
  <nav class="flex px-20 py-5 shadow-md">
    <div class="flex-1 flex items-center m-auto">
      <p>Could you imagine if healthcare was affordable?</p>
    </div>
    <div class="flex items-center">
      <NuxtLink
        :class="['nav-link', { 'active-link': $route.path === '/' }]"
        to="/"
        >Home</NuxtLink
      >
      <NuxtLink
        :class="[
          'nav-link',
          { 'active-link': $route.path === '/find-a-doctor' },
        ]"
        to="/find-a-doctor"
        >Find a Doctor</NuxtLink
      >
      <NuxtLink
        :class="['nav-link', { 'active-link': $route.path === '/services' }]"
        to="/services"
        >Services</NuxtLink
      >
      <NuxtLink
        :class="['nav-link', { 'active-link': $route.path === '/about' }]"
        to="/about"
        >About Us</NuxtLink
      >
    </div>
    <div class="flex-1 flex justify-end items-center m-auto">
      <p class="mr-4 text-gray-600">Login</p>
      <button @click="$router.push('/')" class="cta-button">
        Schedule An Appointment
      </button>
    </div>
  </nav>
</template>

<style scoped>
.nav-link {
  @apply mx-4 text-gray-600 hover:text-blue-700 transition-all font-medium;
}

.active-link {
  @apply text-blue-700;
}
</style>
