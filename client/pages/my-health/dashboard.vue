<script setup lang="ts">
const patient = ref<Patient | null>();
const appointmentsArray = ref<Appointment[]>();

const logout = async () => {
  try {
    navigateTo('/my-health/authentication/login');
    useAuth().logoutUser();
  } catch (error) {
    console.error('Failed to logout:', error);
  }
};

const loadPatientData = async () => {
  try {
    const user = useAuth().getPatientData();
    if (user) {
      user.then(data => {
        patient.value = data;
      });
    }
  } catch (error) {
    console.error('Failed to load patient data:', error);
    alert('Failed to load patient data');
    navigateTo('/my-health/authentication/login');
  }
};

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
  loadPatientData();
  getAppointments();
});
</script>
<template>
  <div
    v-if="patient"
    class="m-auto mt-4 max-w-screen-lg rounded-lg border-[1px] border-gray-400 p-4 shadow"
  >
    <div class="flex">
      <div class="flex-1">
        <p class="text-xl font-bold text-gray-800">Your Information</p>
        <p class="mb-2 text-lg">
          {{ patient.firstname }} {{ patient.lastname }}
        </p>
        <div class="flex">
          <p class="font-semibold">Sex:</p>
          <p class="ms-2">{{ patient.gender }}</p>
        </div>
        <div class="flex">
          <p class="font-semibold">Email:</p>
          <p class="ms-2">{{ patient.email }}</p>
        </div>
        <div class="flex">
          <p class="font-semibold">phone:</p>
          <p class="ms-2">{{ patient.phone }}</p>
        </div>
        <div class="flex">
          <p class="font-semibold">Allergies:</p>
          <div v-for="allergy in patient.allergies">
            {{ allergy }}
          </div>
        </div>
        <div class="flex">
          <p class="me-2 font-semibold">Languages:</p>
          <div v-for="language in patient.language">
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
