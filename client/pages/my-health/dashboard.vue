<script setup lang="ts">
const patient = ref<Patient | null>();

const loadPatientData = async () => {
  try {
    const data = await useAuth().getPatientData();
    patient.value = data;
  } catch (error) {
    console.error('Failed to load patient data:', error);
    navigateTo('/my-health/authentication/login');
  }
};

onMounted(loadPatientData);
</script>
<template>
  <div v-if="patient">
    <h2 class="">Patient Information</h2>
    <p>Fullname: {{ patient.firstname + ' ' + patient.lastname }}</p>
    <p>Email: {{ patient.email }}</p>
    <p>Birthdate: {{ patient.birth }}</p>
    <p>Allergies</p>
    <ul>
      <li v-if="patient.allergies.length === 0">No allergies</li>
      <li v-else v-for="allergy in patient.allergies">
        <p>{{ allergy.name }}</p>
        <br />
        <p>{{ allergy.description }}</p>
      </li>
    </ul>
    <p>Perscriptions</p>
    <ul>
      <li v-if="patient.prescriptions.length === 0">No perscriptions</li>
      <li v-else v-for="perscription in patient.prescriptions">
        <p>{{ perscription.name }}</p>
        <p>Dosage: {{ perscription.dosage }} {{ perscription.frequency }}</p>
        <p>{{ perscription.start }} - {{ perscription.end }}</p>
      </li>
    </ul>
  </div>
  <div v-else>
    <p>something went wrong - No user...</p>
  </div>
</template>
