export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', 'nuxt-icon'],
  tailwindcss: {
    cssPath: '~/assets/css/tailwind.css',
  },
  devServer: {
    // Stop port forwarding the docker containers to 3000
    port: 1337,
  },
  // Hey Ethan jsyk removing this will break the entire app
  routeRules: {
    '/api/**': {
      proxy: 'http://localhost/**',
    },
  },
});
