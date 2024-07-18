<template>
  <h1>
    I Am Protected
  </h1>
</template>

<script lang="ts" setup>
definePageMeta({
  middleware: ["auth"],
})
</script>

<style scoped>
</style>
