<template>
  The Login Page:
  <form-login></form-login>
</template>

<script lang="ts" setup>
definePageMeta({
  auth: {
    unauthenticatedOnly: true,
    navigateAuthenticatedTo: "/"
  }
})
</script>

<style scoped>
</style>
