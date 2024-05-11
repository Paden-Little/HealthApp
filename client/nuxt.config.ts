export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss"],
  app: {
    pageTransition: { name: "page", mode: "out-in" },
  },
  devServer: {
    port: 3001,
  },
});
