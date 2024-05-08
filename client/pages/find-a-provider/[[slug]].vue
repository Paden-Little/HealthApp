<template>
  <h1>{{ $route.path }}</h1>
  <p v-if="$route.params.slug != ''">
    this would display when searching for a doctor
  </p>
</template>
