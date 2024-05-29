<script setup lang="ts">
const patient = ref<Patient | null>();

const logout = async () => {
  try {
    useAuth().logoutUser();
    navigateTo('/my-health/authentication/login');
  } catch (error) {
    console.error('Failed to logout:', error);
  }
};

const loadPatientData = async () => {
  try {
    const user = useAuth().getPatientData();
    if (user) {
      user.then((data) => {
        patient.value = data;
      })
    }
  } catch (error) {
    console.error('Failed to load patient data:', error);
    alert('Failed to load patient data');
    navigateTo('/my-health/authentication/login');
  }
};

onMounted(() => {
  let pid = useCookie("pid");
  console.log(pid.value)
  loadPatientData();
});
</script>
<template>
  <div v-if="patient">
    <h2 class="text-lg font-bold">Patient Information</h2>
    <p>Fullname: {{ patient.firstname + ' ' + patient.lastname }}</p>
    <p>Email: {{ patient.email }}</p>
    <p>Birthdate: {{ patient.birth }}</p>
    <p>Allergies</p>
    <ul>
      <!-- <li v-if="patient.allergies.length = 0">No allergies</li>
      <li v-else v-for="allergy in patient.allergies">
        <p>{{ allergy.name }}</p>
        <br />
        <p>{{ allergy.description }}</p>
      </li> -->
    </ul>
    <p>Perscriptions</p>
    <ul>
      <!-- <li v-if="patient.prescriptions.length = 0">No perscriptions</li>
      <li v-else v-for="perscription in patient.prescriptions">
        <p>{{ perscription.name }}</p>
        <p>Dosage: {{ perscription.dosage }} {{ perscription.frequency }}</p>
        <p>{{ perscription.start }} - {{ perscription.end }}</p>
      </li> -->
    </ul>
  </div>
  <div v-else>
    <p>something went wrong - No user...</p>
  </div>
  <button @click="logout()" class="cta-btn">logout</button>
</template>
