<script setup lang="ts">
  const patientId = '00664891-25b4-4e4b-80c2-4cdfa48eabe8'
  const patient = await useFetch('http://localhost/patient/' + patientId);
</script>
<template>
  <h1>Patient Dashboard</h1>
  <p>{{ patient.firstname }} {{ patient.lastname }}</p>
  <p>{{ patient.age }}</p>
</template>
z