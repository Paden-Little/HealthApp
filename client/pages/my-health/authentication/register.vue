<script setup lang="ts">
definePageMeta({
  layout: false,
});
</script>
<template>
  <section class="bg-gray-50">
    <div
      class="mx-auto flex flex-col items-center justify-center px-6 py-8 md:h-screen lg:py-0"
    >
      <div class="w-full rounded-lg bg-white shadow sm:max-w-md md:mt-0 xl:p-0">
        <div class="space-y-4 p-6 sm:p-8 md:space-y-6">
          <h1
            class="text-xl font-bold leading-tight tracking-tight text-gray-900"
          >
            Create your Health Mark account
          </h1>
          <form class="space-y-4 md:space-y-6" action="#">
            <!-- Get Full Name -->
            <div>
              <div>
                <label for="first_name">First name</label>
                <input
                  type="text"
                  name="first_name"
                  id="first_name"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
              <div>
                <label for="middle_name">Middle name</label>
                <input
                  type="text"
                  name="middle_name"
                  id="middle_name"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                />
              </div>
              <div>
                <label for="last_name">Last name</label>
                <input
                  type="text"
                  name="last_name"
                  id="last_name"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
            </div>
            <div>
              <label
                for="gender"
                class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                >Biological gender</label
              >
              <select
                id="gender"
                class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500"
              >
                <option value="" disabled selected>Choose a gender</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
              </select>
            </div>
            <!-- Email -->
            <div>
              <label
                for="email"
                class="mb-2 block text-sm font-medium text-gray-900"
                >Email</label
              >
              <input
                type="email"
                name="email"
                id="email"
                class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                required="true"
              />
            </div>
            <!-- Password -->
            <div>
              <div>
                <label
                  for="password"
                  class="mb-2 block text-sm font-medium text-gray-900"
                  >Password</label
                >
                <input
                  type="password"
                  name="password"
                  id="password"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
              <div>
                <label
                  for="re-password"
                  class="mb-2 block text-sm font-medium text-gray-900"
                  >Re-enter password</label
                >
                <input
                  type="password"
                  name="re-password"
                  id="re-password"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
            </div>
            <!-- Address -->
            <div>
              <div>
                <label for="address">Address</label>
                <input
                  type="text"
                  name="address"
                  id="address"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
              <div>
                <label for="city">City</label>
                <input
                  type="text"
                  name="city"
                  id="city"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
              <label for="state">State</label>
              <select
                id="state"
                class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500"
                required="true"
              >
                <option value="" disabled selected>Choose a state</option>
                <option value="alabama">Alabama</option>
                <option value="alaska">Alaska</option>
                <option value="arizona">Arizona</option>
                <option value="arkansas">Arkansas</option>
                <option value="california">California</option>
                <option value="colorado">Colorado</option>
                <option value="connecticut">Connecticut</option>
                <option value="delaware">Delaware</option>
                <option value="florida">Florida</option>
                <option value="georgia">Georgia</option>
                <option value="hawaii">Hawaii</option>
                <option value="idaho">Idaho</option>
                <option value="illinois">Illinois</option>
                <option value="indiana">Indiana</option>
                <option value="iowa">Iowa</option>
                <option value="kansas">Kansas</option>
                <option value="kentucky">Kentucky</option>
                <option value="louisiana">Louisiana</option>
                <option value="maine">Maine</option>
                <option value="maryland">Maryland</option>
                <option value="massachusetts">Massachusetts</option>
                <option value="michigan">Michigan</option>
                <option value="minnesota">Minnesota</option>
                <option value="mississippi">Mississippi</option>
                <option value="missouri">Missouri</option>
                <option value="montana">Montana</option>
                <option value="nebraska">Nebraska</option>
                <option value="nevada">Nevada</option>
                <option value="new-hampshire">New Hampshire</option>
                <option value="new-jersey">New Jersey</option>
                <option value="new-mexico">New Mexico</option>
                <option value="new-york">New York</option>
                <option value="north-carolina">North Carolina</option>
                <option value="north-dakota">North Dakota</option>
                <option value="ohio">Ohio</option>
                <option value="oklahoma">Oklahoma</option>
                <option value="oregon">Oregon</option>
                <option value="pennsylvania">Pennsylvania</option>
                <option value="rhode-island">Rhode Island</option>
                <option value="south-carolina">South Carolina</option>
                <option value="south-dakota">South Dakota</option>
                <option value="tennessee">Tennessee</option>
                <option value="texas">Texas</option>
                <option value="utah">Utah</option>
                <option value="vermont">Vermont</option>
                <option value="virginia">Virginia</option>
                <option value="washington">Washington</option>
                <option value="west-virginia">West Virginia</option>
                <option value="wisconsin">Wisconsin</option>
                <option value="wyoming">Wyoming</option>
              </select>
              <div>
                <label for="zip">Zip code</label>
                <input
                  type="text"
                  name="zip"
                  id="zip"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 sm:text-sm"
                  required="true"
                />
              </div>
            </div>
            <button
              type="submit"
              class="w-full rounded-lg bg-primary-600 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-primary-700 focus:outline-none focus:ring-4 focus:ring-primary-300"
            >
              Create Account
            </button>
            <p class="text-sm font-light text-gray-500">
              Already have an account?
              <NuxtLink
                to="/my-health/authentication/login"
                class="ml-5 font-medium text-primary-600 hover:underline"
                >Sign In Here</NuxtLink
              >
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
