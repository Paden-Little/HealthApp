<script setup lang="ts">
const patient = ref<Patient>();

onBeforeMount(async () => {
  try {
    patient.value = await useAuth().getPatientData();
  } catch (error) {
    console.error(error);
  }
});
</script>
<template>
  <!-- TEMP TEST -->
  {{ patient }}
  <div v-if="patient">
    <h2 class="">Patient Information</h2>
    <p>Fullname: {{ patient.firstname + " " + patient.lastname }}</p>
    <!-- <p>Fullname: {{ patient.value.firstname + ' ' + patient.value.lastname}}</p>
    <p>Email: {{ patient.value.email }}</p>
    <p>Phone: {{ patient.value.phone }}</p>
    <p>Birthdate: {{ patient.value.birth }}</p>
    <p>Allergies</p>
    <ul>
      <li v-if="patient.value.allergies.count == 0">No allergies</li>
      <li v-else v-for="allergy in patient.value.allergies">
        <p>{{ allergy.name }}</p>
        <br>
        <p>{{ allergy.description }}</p>
      </li>
    </ul>
    <p>Current Perscriptions</p>
    <ul>
      <li v-if="patient.value.perscriptions.count == 0">No perscriptions</li>
      <li v-else v-for="persription in patient.value.perscriptions">
        <p>{{ persription.name }}</p>
        <p>Dosage: {{ persription.dosage }} {{ persription.freqency }}</p>
        <p>{{ persription.start }} - {{ persription.end }}</p>
      </li>
    </ul> -->
  </div>
  <div v-else>
    <p>
      something went wrong
    </p>
  </div>
</template>
