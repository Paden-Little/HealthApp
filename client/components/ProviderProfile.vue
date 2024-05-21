<script setup lang="ts">
type Provider = {
  id: string;
  name: string;
  suffix: string;
  email: string;
  phone: string;
  bio: string;
  languages: Array<string>;
  services: Array<string>;
};

const services = ['dermatology', 'general care', 'your mother'];

defineProps({
  provider: {
    type: Object as () => Provider,
    required: true,
  },
});

function formattedServices(services: Array<string>): string {
  if (services.length > 1) {
    return (
      services.slice(0, -1).join(', ') +
      ', and ' +
      services[services.length - 1]
    );
  } else {
    return services[0];
  }
}

function formatPhoneNumber(phoneNumber: string): string {
  return phoneNumber.replace(/(\d{3})(\d{3})(\d{4})/, '$1-$2-$3');
}
</script>

<template>
  <div
    class="items-center rounded-lg bg-gray-50 shadow dark:border-gray-700 dark:bg-gray-800 sm:flex"
  >
    <a href="#">
      <img
        class="w-full rounded-lg sm:rounded-none sm:rounded-l-lg"
        src="/img/skillissue.jpg"
        :alt="provider.name"
      />
    </a>
    <div class="p-5">
      <h3
        class="text-xl font-bold tracking-tight text-gray-900 dark:text-white"
      >
        <a href="#">{{ provider.name }}</a>
      </h3>
      <span class="text-gray-500"
        >{{ formattedServices(provider.services) }}
      </span>
      <p class="mb-4 mt-3 font-light text-gray-500">
        Am I just unwell these days? I'm not myself And it feels like I'll never
        change Too late for me Nothing's healing the hurt, only worsens the pain
      </p>
    </div>
  </div>
</template>

<style scoped>
/* Custom CSS for truncating the overflow */
.truncate-last-item::after {
  content: '';
  display: inline-block;
  width: 100%;
}
</style>
