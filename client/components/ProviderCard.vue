<script setup lang="ts">
const props = defineProps({
  provider: {
    type: Object as () => Provider,
  },
});

function formatServices(services: string[]): string {
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

const formattedPhoneNumber = props.provider!.phone.replace(
  /(\d{3})(\d{3})(\d{4})/,
  '$1-$2-$3',
);

function getNextTwoWeeks(): Date[] {
  const startDate = new Date();
  startDate.setDate(startDate.getDate() - 13);

  const dates = [];
  while (startDate <= new Date()) {
    dates.push(new Date(startDate));
    startDate.setDate(startDate.getDate() + 1);
  }

  return dates;
}

function formatDate(date: Date): string {
  const options: Intl.DateTimeFormatOptions = {
    month: 'short',
    day: 'numeric',
  };
  return new Date(date).toLocaleDateString(undefined, options);
}
</script>

<template>
  <div
    class="flex w-full justify-between border-b-[1px] border-gray-300 pb-6 pt-3"
  >
    <div class="flex">
      <div>
        <img
          class="w-20 rounded-full"
          src="/img/skillissue.jpg"
          alt="yeah it brokey"
        />
      </div>
      <div class="ms-4">
        <div class="flex justify-between">
          <p class="text-lg font-semibold">
            {{ provider!.firstname + ' ' + provider!.lastname }},
            {{ provider!.suffix }}
          </p>
        </div>
        <p class="text-md text-sm font-semibold text-gray-800">
          {{ formatServices(provider!.services) }}
        </p>
        <div class="mt-1 flex items-center align-middle tracking-tight">
          <div class="flex items-center text-sm text-gray-500">
            <Icon name="ic:baseline-email" />
            <p class="ms-2">{{ provider!.email }}</p>
          </div>
          <div
            class="ms-4 flex items-center justify-start text-sm text-gray-500"
          >
            <Icon name="material-symbols:phone-enabled" />
            <p class="ms-1">{{ formattedPhoneNumber }}</p>
          </div>
        </div>
        <div class="mt-4 w-72">
          <p class="text-sm text-gray-500">{{ provider!.bio }}</p>
        </div>
      </div>
    </div>
    <div class="ms-2 flex flex-col">
      <p class="mb-2 font-semibold tracking-tight">Available Dates:</p>
      <div class="grid grid-cols-7">
        <div
          v-for="(date, index) in getNextTwoWeeks()"
          :key="index"
          class="mb-1 me-1 rounded border-[1px] border-gray-300 px-2 py-1 text-sm shadow-lg"
        >
          <p>{{ formatDate(date) }}</p>
          <p class="mt-2">no</p>
          <p>appts</p>
        </div>
      </div>
    </div>
    <div class="ms-8 flex flex-col items-end justify-center">
      <NuxtLink
        to="'/find-a-provider/' + provider.firstname + '-' provider.lastname"
        class="secondary-btn w-44"
        >View Profile</NuxtLink
      >
      <NuxtLink
        to="'/book-appointment/' + provider.id"
        class="cta-btn mt-2 w-44"
        >Book Appointment</NuxtLink
      >
    </div>
  </div>
</template>
