<template>
  <nav class="flex px-20 py-3">
    <div class="m-auto flex flex-1 items-center">
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
    <div class="flex flex-1 items-center justify-end">
      <NuxtLink
        class="mr-8 font-semibold text-black"
        to="/my-health/authentication/login"
        >Login</NuxtLink
      >
      <button @click="$router.push('/')" class="cta-button">
        Book an Appointment
      </button>
    </div>
  </nav>
</template>

<style scoped>
.nav-link {
  @apply mx-4 font-semibold text-black transition-all hover:text-[#0352CD];
}

.active-link {
  @apply text-[#344983];
}
</style>
