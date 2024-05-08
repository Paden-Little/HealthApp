export default defineNuxtRouteMiddleware((to, from) => {
  if (to.path == '/my-health') {
    return navigateTo('/my-health/authentication/login');
  }
});
