<script setup lang="ts">
import axios from 'axios';
const pid = useCookie('patId');
const jwt = useCookie('jwt');
const urlRef = ref(`http://localhost/patient/${pid.value}`);
const patient = ref();

const getHeaders = () => ({
  Authorization: `Bearer ${jwt.value}`,
});

const getPatient = async () => {
  try {
    const response = await axios.get(urlRef.value, {
      headers: getHeaders(),
    });
    console.log(response.data);
  } catch (error) {
    console.error('Failed to get patient:', error);
  }
};

onMounted(() => {
  patient.value = getPatient();
  patient.value = {
    firstname: 'John',
    lastname: 'Doe',
    email: 'testing email',
    phone: '123-456-7890',
    birth: '01/01/2000',
  }
});
</script>
<template>
  <!-- TEMP TEST -->
  {{ patient }}

  <h2 class="">Patient Information</h2>
  <p>Fullname: {{ patient.value.firstname + ' ' + patient.value.lastname}}</p>
  <p>Email: {{ patient.value.email }}</p>
  <p>Phone: {{ patient.value.phone }}</p>
  <p>Birthdate: {{ patient.value.birth }}</p>
  <!-- <p>Allergies</p>
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
</template>
