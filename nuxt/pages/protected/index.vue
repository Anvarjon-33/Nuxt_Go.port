import defineNuxtConfig from '../../nuxt.config';
<template></template>

<script setup>
definePageMeta({
  layout: "default",
  middleware: "auth",
});
</script>

<style scoped></style>
