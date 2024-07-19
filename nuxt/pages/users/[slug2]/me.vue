<template>
  <h1>
    [wild] > me page : path -> {{ route.fullPath }}
  </h1>

</template>

<script lang="ts" setup>
const route = useRoute()
console.log(route.fullPath)
</script>

<style scoped>

</style>
