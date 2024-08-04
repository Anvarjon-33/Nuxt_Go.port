<template>
  <form action="http://192.168.1.3:2222/test/" method="post">
    <v-checkbox-btn name="colors[]" value="red">red</v-checkbox-btn>
    <v-checkbox-btn name="colors[]" value="yellow">yellow</v-checkbox-btn>
    <v-checkbox-btn name="colors[]" value="black">black</v-checkbox-btn>
    <v-checkbox-btn name="colors[]" value="warning">warning</v-checkbox-btn>
    <v-checkbox-btn name="colors[]" value="success">success</v-checkbox-btn>
    <v-btn type="submit">SUBMIT</v-btn>
  </form>
</template>

<script setup>
</script>

<style scoped>
</style>
