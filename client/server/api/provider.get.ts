export default defineEventHandler(async event => {
  const data = await $fetch('http://localhost/provider');
  return data;
});
