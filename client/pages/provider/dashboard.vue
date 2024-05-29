<script setup lang="ts">
const provider = ref<Provider | null>();

const logout = async () => {
  try {
    useAuth().logoutUser();
    navigateTo('/provider/authentication/login');
  } catch (error) {
    console.error('Failed to logout:', error);
  }
};

const loadProviderData = async () => {
  try {
    const user = useAuth().getProviderData();
    if (user) {
      user.then((data) => {
        provider.value = data;
      })
    }
  } catch (error) {
    console.error('Failed to load provider data:', error);
    alert('Failed to load provider data');
    navigateTo('/provider/authentication/login');
  }
};

onMounted(() => {
  let pid = useCookie("pid");
  console.log(pid.value)
  loadProviderData();
});
</script>
<template>

  <button @click="logout()" class="cta-btn">logout</button>
</template>
