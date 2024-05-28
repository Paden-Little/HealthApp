<script setup lang="ts">

definePageMeta({
  layout: false,
});

const email = ref('');
const password = ref('');
// set provider ref
const provider = ref();

function login() {
  if (email.value === '' || password.value === '') {
    alert('Please fill in all fields');
    return;
  }

  provider.value = useAuth().getProviderData()
}
</script>
<template>
  <section class="bg-gray-50">
    <div
      class="mx-auto flex flex-col items-center justify-center px-6 py-8 md:h-screen lg:py-0"
    >
      <div
        class="w-full rounded-lg bg-white shadow sm:max-w-md md:mt-0 xl:p-0"
      >
        <div class="space-y-4 p-6 sm:p-8 md:space-y-6">
          <h1
            class="text-xl font-bold leading-tight tracking-tight text-gray-900"
          >
            Sign in to your Health Mark provider account
          </h1>
          <div class="space-y-4 md:space-y-6">
            <div>
              <label class="mb-2 block text-sm font-medium text-gray-900"
                >Email</label
              >
              <input
                v-model="email"
                type="email"
                name="email"
                class="focus:border-primary-600 focus:ring-primary-600 block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 sm:text-sm"
                required="true"
              />
            </div>
            <div>
              <label
                for="password"
                class="mb-2 block text-sm font-medium text-gray-900"
                >Password</label
              >
              <input
                v-model="password"
                type="password"
                name="password"
                class="focus:border-primary-600 focus:ring-primary-600 block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 sm:text-sm"
                required="true"
              />
            </div>
            <!-- Align center -->
            <button
              type="button"
              class="mb-2 me-2 rounded-lg bg-blue-700 px-5 py-2.5 text-sm font-medium text-white hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300"
              @click="login()"
            >
              Sign in
            </button>
            <p
              class="flex justify-around text-sm font-light text-gray-500"
            >
              Don’t have an account yet?
              <NuxtLink
                to="/provider/authentication/register"
                class="font-medium text-primary-600 hover:underline"
                >Create Provider Account</NuxtLink
              >
            </p>
            <p
              class="flex justify-around text-sm font-light text-gray-500"
            >
              Not a provider?
              <NuxtLink
                to="/my-health/authentication/login"
                class="font-medium text-primary-600 hover:underline"
                >Sign In Here</NuxtLink
              >
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>