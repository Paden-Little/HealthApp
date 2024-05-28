export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', 'nuxt-icon'],
  build: {
    transpile: ['@vuepic/vue-datepicker'],
  },
  tailwindcss: {
    cssPath: '~/assets/css/tailwind.css',
  },
  devServer: {
    // STOP FUCKING PORT FORWARDING TO 3000 PLEASE I NEED IT FOR FRONTEND
    port: 1337,
  },
  // Hey Ethan jsyk removing this will break the entire app
  routeRules: {
    '/api/**': {
      proxy: 'http://localhost:80/**',
    },
  },
});
