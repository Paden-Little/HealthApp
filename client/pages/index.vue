<script setup lang="ts">
import { initFlowbite } from 'flowbite';

definePageMeta({
  layout: false,
});

onMounted(() => {
  initFlowbite();
});
</script>

<template>
  <section class="bg-slate-600 bg-cover">
    <Navbar bgColor="" textColor="text-white" />
    <div
      class="mx-auto grid max-w-screen-xl px-4 pb-8 pt-8 lg:grid-cols-12 lg:gap-8 lg:pb-0 lg:pt-16 xl:gap-0"
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
        <form class="max-w-lg">
          <div class="flex">
            <label
              for="search-dropdown"
              class="sr-only mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Your Email</label
            >
            <button
              id="dropdown-button"
              data-dropdown-toggle="dropdown"
              class="z-10 inline-flex flex-shrink-0 items-center rounded-s-lg border border-gray-300 bg-gray-100 px-4 py-2.5 text-center text-sm font-medium text-gray-900 hover:bg-gray-200 focus:outline-none focus:ring-4 focus:ring-gray-100 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-700"
              type="button"
            >
              All Specialties
              <svg
                class="ms-2.5 h-2.5 w-2.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 10 6"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="m1 1 4 4 4-4"
                />
              </svg>
            </button>
            <div
              id="dropdown"
              class="z-10 hidden w-44 divide-y divide-gray-100 rounded-lg bg-white shadow dark:bg-gray-700"
            >
              <ul
                class="py-2 text-sm text-gray-700 dark:text-gray-200"
                aria-labelledby="dropdown-button"
              >
                <li>
                  <button
                    type="button"
                    class="inline-flex w-full px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  ></button>
                </li>
                <li>
                  <button
                    type="button"
                    class="inline-flex w-full px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    Templates
                  </button>
                </li>
                <li>
                  <button
                    type="button"
                    class="inline-flex w-full px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    Design
                  </button>
                </li>
                <li>
                  <button
                    type="button"
                    class="inline-flex w-full px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    Logos
                  </button>
                </li>
              </ul>
            </div>
            <div class="relative w-full">
              <input
                type="search"
                id="search-dropdown"
                class="z-20 block w-full rounded-e-lg border border-s-2 border-gray-300 border-s-gray-50 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:border-s-gray-700 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500"
                placeholder="Search Mockups, Logos, Design Templates..."
                required
              />
              <button
                type="submit"
                class="absolute end-0 top-0 h-full rounded-e-lg border border-blue-700 bg-blue-700 p-2.5 text-sm font-medium text-white hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
              >
                <svg
                  class="h-4 w-4"
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
                <span class="sr-only">Search</span>
              </button>
            </div>
          </div>
        </form>
      </div>
      <div class="hidden lg:col-span-5 lg:mt-0 lg:flex">
        <img src="/img/doctor.png" alt="mockup" />
      </div>
    </div>
  </section>
  <section class="bg-white">
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
          <h3 class="mb-2 text-xl font-bold">Primary Care</h3>
          <img src="/img/primarycare.jpg" class="h-38 bg-cover" />
          <p class="p-4 text-gray-500">
            Staying well means caring for your health every day. Our primary
            care providers help you and your family set healthy habits, stay on
            top of routine screenings, and prevent and manage disease. They also
            treat most diseases except for very rare or unusual ones.
          </p>
        </div>
        <div class="rounded-md bg-white p-4 drop-shadow-2xl">
          <h3 class="mb-2 text-xl font-bold">Pediatrics</h3>
          <img src="/img/pediatric.jpg" class="h-38 bg-contain" />
          <p class="p-4 text-gray-500">
            We provide a long list of pediatric services, in collaboration with
            Primary Children's Hospital, to care for your child from infancy to
            adulthood.
          </p>
        </div>
        <div class="rounded-md bg-white p-4 drop-shadow-2xl">
          <h3 class="mb-2 text-xl font-bold">Heart Health</h3>
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
  <footer class="bg-white p-4 md:p-8 lg:p-10">
    <div class="mx-auto max-w-screen-xl text-center">
      <a
        href="#"
        class="flex items-center justify-center text-2xl font-semibold text-gray-900"
      >
        HealthApp
      </a>
      <p class="my-6 text-gray-500">
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
