export default defineNuxtRouteMiddleware((to, from) => {
  if (to.path == '/my-health') {
    // If unauthenticated send to login, else dashboard
    return navigateTo('/my-health/authentication/login');
  }
});
