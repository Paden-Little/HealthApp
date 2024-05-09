<template>
  <nav class="flex px-20 py-3">
    <div class="flex-1 flex items-center m-auto">
      <p>HealthApp</p>
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
      <p class="mr-8 text-gray-600">Login</p>
      <button @click="$router.push('/')" class="cta-button">
        Book an Appointment
      </button>
    </div>
  </nav>
</template>

<style scoped>
.nav-link {
  @apply mx-4 text-gray-600 hover:text-[#0352CD] transition-all font-medium;
}

.active-link {
  @apply text-[#0352CD];
}
</style>
