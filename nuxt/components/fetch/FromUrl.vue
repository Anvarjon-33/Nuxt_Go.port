<template>
  <slot :data="data['message']"></slot>
</template>

<script lang="ts" setup>
const props = defineProps({
  url: {
    type: String,
    default: "/"
  },
  method: {
    type: String,
    default: "POST"
  },
  live: {
    type: Boolean,
    default: false
  }
})

const data = await $fetch(props.url, {
  method: props.method as any
})
console.log(data)
</script>

<style scoped>
</style>
