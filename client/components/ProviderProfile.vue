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
    // Join all elements with a comma except the last one
    return (
      services.slice(0, -1).join(', ') +
      ', and ' +
      services[services.length - 1]
    );
  } else {
    // Return the single element directly if there's only one
    return services[0];
  }
}

function formatPhoneNumber(phoneNumber: string): string {
  return phoneNumber.replace(/(\d{3})(\d{3})(\d{4})/, '$1-$2-$3');
}
</script>

<template>
  <a
    href="#"
    class="flex flex-col items-center rounded-lg border border-gray-200 bg-white shadow hover:bg-gray-100 md:max-w-xl md:flex-row"
  >
    <img
      class="h-96 w-full rounded-t-lg object-cover md:h-auto md:w-48 md:rounded-none md:rounded-s-lg"
      src="/img/fiend.png"
      alt=""
    />
    <div class="flex flex-col justify-between p-4 leading-normal">
      <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900">
        {{ provider.name }}
      </h5>
      <div class="flex space-x-2">
        <p>{{ formattedServices(services) }}</p>
      </div>
      <p>{{ formatPhoneNumber(provider.phone) }}</p>
    </div>
  </a>
</template>

<style scoped>
/* Custom CSS for truncating the overflow */
.truncate-last-item::after {
  content: '';
  display: inline-block;
  width: 100%;
}
</style>
