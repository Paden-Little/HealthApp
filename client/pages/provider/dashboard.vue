<script setup lang="ts">
const provider = ref<Provider | null>();
const appointmentsArray = ref<Appointment[]>();

const logout = async () => {
  try {
    navigateTo('/provider/authentication/login');
    useAuth().logoutUser();
  } catch (error) {
    console.error('Failed to logout:', error);
  }
};

const loadProviderData = async () => {
  try {
    const user = useAuth().getProviderData();
    if (user) {
      user.then(data => {
        provider.value = data;
      });
    }
  } catch (error) {
    console.error('Failed to load provider data:', error);
    alert('Failed to load provider data');
    navigateTo('/provider/authentication/login');
  }
};

const getAppointments = async () => {
  try {
    const appointments = useAuth().getUserAppointment();
    if (appointments) {
      appointments.then(async data => {
        appointmentsArray.value = data;
        await getProviderNames();
      });
    }
  } catch (error) {
    console.error('Failed to load appointments:', error);
    alert('Failed to load appointments');
  }
};

async function getProviderNames() {
  if (appointmentsArray.value) {
    for (let index = 0; index < appointmentsArray.value.length; index++) {
      const element = appointmentsArray.value[index];
      let temp = useAuth().getProviderName(element.provider);
      if (temp) {
        temp.then(data => {
          element.provider = data;
        });
      }
    }
  }
}

onMounted(() => {
  loadProviderData();
  getAppointments();
});
</script>
<template>
  <div
    v-if="provider"
    class="m-auto mt-4 max-w-screen-lg rounded-lg border-[1px] border-gray-400 p-4 shadow"
  >
    <div class="flex">
      <div class="flex-1">
        <p class="text-xl font-bold text-gray-800">Your Information</p>
        <p class="mb-2 text-lg">
          {{ provider.firstname }} {{ provider.lastname }}
        </p>
        <div class="flex">
          <p class="font-semibold">Email:</p>
          <p class="ms-2">{{ provider.email }}</p>
        </div>
        <div class="flex">
          <p class="font-semibold">phone:</p>
          <p class="ms-2">{{ provider.phone }}</p>
        </div>
        <div class="flex">
          <p class="me-2 font-semibold">Languages:</p>
          <div v-for="language in provider.languages">
            {{ language }}
          </div>
        </div>
      </div>
      <div class="flex-1">
        <p class="text-lg font-bold text-gray-800">Appointments</p>
        <div
          v-for="appt in appointmentsArray"
          class="mt-2 rounded-lg border-[1px] border-gray-400 p-2"
        >
          <div class="flex">
            <p class="font-semibold">Date:</p>
            <p class="ms-2">
              {{ appt.date }} {{ appt.start_time }} - {{ appt.end_time }}
            </p>
          </div>
          <div>
            <p class="font-semibold">Reason:</p>
            <p class="ms-2">{{ appt.description }}</p>
          </div>
        </div>
      </div>
    </div>
    <button @click="logout()" class="cta-btn mt-4">logout</button>
  </div>
</template>
