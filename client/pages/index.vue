<script setup lang="ts">
definePageMeta({
  layout: false,
});
</script>

<template>
  <section class="bg-[url(/img/hero.png)] bg-cover">
    <HomeNavbar />
    <div
      class="mx-auto grid max-w-screen-xl px-4 pt-8 lg:grid-cols-12 lg:gap-8 lg:pt-16 xl:gap-0"
    >
      <div class="mr-auto place-self-center lg:col-span-7">
        <h1
          class="mb-4 max-w-2xl text-4xl font-extrabold leading-none tracking-tight text-white md:text-5xl xl:text-6xl"
        >
          Your health,
        </h1>
        <h1
          class="mb-4 max-w-2xl text-4xl font-extrabold leading-none tracking-tight text-white md:text-5xl xl:text-6xl"
        >
          is our priority.
        </h1>
        <p
          class="mb-6 max-w-2xl font-light text-gray-50 md:text-lg lg:mb-8 lg:text-xl"
        >
          Book Your Appointment Today for Expert Medical Care. Discover
          Healthcare Solutions Tailored to you, and get Personalized Care from
          Our Experienced Medical Team
        </p>
        <a
          href="#"
          class="mr-3 inline-flex items-center justify-center rounded-lg bg-primary-700 px-5 py-3 text-center text-base font-medium text-white hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:focus:ring-primary-900"
        >
          Get started
        </a>
      </div>
      <div class="hidden lg:col-span-5 lg:mt-0 lg:flex">
        <img src="/img/tricyclen.png" alt="mockup" />
      </div>
    </div>
  </section>
  <section class="bg-white dark:bg-gray-900">
    <div class="mx-auto max-w-screen-xl px-4 py-8 sm:py-16 lg:px-6">
      <div class="mb-8 max-w-screen-md lg:mb-16">
        <h2 class="mb-4 text-4xl font-extrabold tracking-tight text-gray-900">
          Explore Our Services
        </h2>
        <p>We offer a variety of services covering all aspects of life</p>
      </div>
      <div
        class="space-y-8 md:grid md:grid-cols-2 md:gap-12 md:space-y-0 lg:grid-cols-3"
      >
        <div class="rounded-md bg-white p-4 drop-shadow-2xl">
          <h3 class="mb-2 text-xl font-bold dark:text-white">Primary Care</h3>
          <img src="/img/primarycare.jpg" class="h-38 bg-cover" />
          <p class="p-4 text-gray-500">
            Staying well means caring for your health every day. Our primary
            care providers help you and your family set healthy habits, stay on
            top of routine screenings, and prevent and manage disease. They also
            treat most diseases except for very rare or unusual ones.
          </p>
        </div>
        <div class="rounded-md bg-white p-4 drop-shadow-2xl">
          <h3 class="mb-2 text-xl font-bold dark:text-white">Pediatrics</h3>
          <img src="/img/pediatric.jpg" class="h-38 bg-contain" />
          <p class="p-4 text-gray-500">
            We provide a long list of pediatric services, in collaboration with
            Primary Children's Hospital, to care for your child from infancy to
            adulthood.
          </p>
        </div>
        <div class="rounded-md bg-white p-4 drop-shadow-2xl">
          <h3 class="mb-2 text-xl font-bold dark:text-white">Heart Health</h3>
          <img src="/img/hearthealth.jpg" />
          <p class="p-4 text-gray-500">
            Centered at the Cardiovascular Center at University of Utah
            Hospital, our heart and vascular services include heart care, heart
            surgery, and treatment for all types of heart disease.
          </p>
        </div>
      </div>
    </div>
  </section>
  <footer class="bg-white p-4 dark:bg-gray-800 md:p-8 lg:p-10">
    <div class="mx-auto max-w-screen-xl text-center">
      <a
        href="#"
        class="flex items-center justify-center text-2xl font-semibold text-gray-900"
      >
        HealthApp
      </a>
      <p class="my-6 text-gray-500 dark:text-gray-400">
        Built by a group of nerds for a projects class over the span of 5 weeks.
      </p>
      <ul class="mb-6 flex flex-wrap items-center justify-center text-gray-900">
        <li>
          <a href="#" class="mr-4 hover:underline md:mr-6">About</a>
        </li>
        <li>
          <a href="#" class="mr-4 hover:underline md:mr-6">Contact</a>
        </li>
        <li>
          <a
            href="https://robertshirts.com"
            class="mr-4 hover:underline md:mr-6"
            >Robbie Shirts</a
          >
        </li>
        <li>
          <a
            href="http://www.wildlifeforyou.com/"
            class="mr-4 hover:underline md:mr-6"
            >WFY</a
          >
        </li>
        <li>
          <a
            href="https://media.discordapp.net/attachments/1097926007035535381/1239095446379696138/image.png?ex=6641ad0f&is=66405b8f&hm=bbed2a63b69df7c5dc43083b1d7b7f1b10cf6188bd23f24018d2b129a127b9a3&=&format=webp&quality=lossless&width=785&height=982"
            class="mr-4 hover:underline md:mr-6"
            >Residents</a
          >
        </li>
      </ul>
      <span class="text-sm text-gray-500 sm:text-center">© 2024-2024 </span>
      <a
        href="https://media.discordapp.net/attachments/1177754818576658572/1239035466708422718/20240511_1826440.jpg?ex=66417533&is=664023b3&hm=56efe8e148e23c2faa6e1c62d67a098fa3ee0229bdd154a116790b521001ba7d&=&format=webp&width=1309&height=982"
        class="text-sm text-gray-500 hover:underline"
        >Image of Kellen with Dog™</a
      >
    </div>
  </footer>
</template>
