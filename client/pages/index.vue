<script setup lang="ts">
definePageMeta({
  layout: false,
});
</script>

<template>
  <div class="bg-[url(/img/hero.png)]">
    <Navbar />
    <section>
      <div class="mx-auto grid max-w-screen-xl grid-cols-12 gap-0 pt-16">
        <div class="col-span-7 mr-auto place-self-center">
          <h1 class="mb-4 max-w-2xl text-6xl font-extrabold text-black">
            Your Health,
          </h1>
          <h1
            class="mb-4 max-w-2xl text-6xl font-extrabold leading-none tracking-tight text-black"
          >
            Our Priority
          </h1>
          <p
            class="mb-6 max-w-2xl font-normal text-black md:text-lg lg:mb-8 lg:text-xl"
          >
            Someone put something else here
          </p>
          <button class="cta-button">Book an Appointment</button>
          <button class="glass-button">Find a Doctor</button>
        </div>
        <div class="lg:col-span-5 lg:mt-0 lg:flex">
          <img src="/img/doctor.png" alt="image broke" />
        </div>
      </div>
    </section>
    <section class="bg-white"></section>
  </div>
</template>
