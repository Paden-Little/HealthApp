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

// insert get appointment function here for provider
const getAppointments = async () => {
  try {
    const appointments = useAuth().getUserAppointment();
    if (appointments) {
      appointments.then(data => {
        appointmentsArray.value = data;
      });
    }
  } catch (error) {
    console.error('Failed to load appointments:', error);
    alert('Failed to load appointments');
  }
};

onMounted(() => {
  loadProviderData();
});
</script>
<template>
  <div v-if="provider">
    <div>
      <img :src="provider.image" />
    </div>
    <aside>
      <h2 class="text-lg font-bold">Provider Information</h2>
      <p>
        Full Name:
        {{
          provider.firstname + ' ' + provider.lastname + ' ' + provider.suffix
        }}
      </p>
      <p>Email: {{ provider.email }}</p>
      <p>Phone: {{ provider.phone }}</p>
      <p>Bio: {{ provider.bio }}</p>
      <p>Services:</p>
      <ul class="">
        <li v-for="service in provider.services">
          {{ service }}
        </li>
      </ul>
      <p>Languages:</p>
      <ul>
        <li v-for="language in provider.languages">
          {{ language }}
        </li>
      </ul>
    </aside>
  </div>
  <div v-else>
    <p class="">Something went wrong - No provider found.</p>
  </div>
  <div>
    <h2>Appointments</h2>
    <ul>
      <li v-for="appointment in appointmentsArray">
        <p>
          {{ appointment.date }} {{ appointment.startTime }} -
          {{ appointment.endTime }}
        </p>
        <p>{{ appointment.patient }}</p>
        <p>{{ appointment.service }}</p>
        <p>{{ appointment.description }}</p>
      </li>
    </ul>
  </div>
  <button @click="logout()" class="cta-btn">logout</button>
</template>
