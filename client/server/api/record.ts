export default defineEventHandler(async (event) => {
  return await fetch('http://localhost:3000');
});
