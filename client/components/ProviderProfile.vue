<script setup lang="ts">
const props = defineProps({
  provider: Object,
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
</script>

<template>
  <div class="flex border-b-[1px] border-gray-300 pb-6 pt-3">
    <div>
      <img
        class="w-28 rounded-full"
        src="/img/skillissue.jpg"
        alt="yeah it brokey"
      />
    </div>
    <div class="ms-4 w-80">
      <div class="flex justify-between">
        <p class="text-lg font-semibold">
          {{ provider!.name }}, {{ provider!.suffix }}
        </p>
        <div class="flex items-center justify-center text-sm">
          <Icon name="material-symbols:phone-enabled" />
          <p class="ms-2">{{ formattedPhoneNumber }}</p>
        </div>
      </div>
      <p class="text-md font-semibold text-gray-800">
        {{ formatServices(provider!.services) }}
      </p>
      <div class="mt-2">
        <p class="text-sm text-gray-600">{{ provider!.bio }}</p>
      </div>
    </div>
  </div>
</template>
