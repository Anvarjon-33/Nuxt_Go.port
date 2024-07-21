<template>
  <slot :data name="data"></slot>
</template>

<script lang="ts" setup>
const props = defineProps({
  url: {
    type: String,
    default: "/"
  },
  method: {
    type: String,
    default: "GET"
  },
  live: {
    type: Boolean,
    default: false
  }
})

const data: any = await $fetch(props.url, {
  method: props.method as any
})

console.log(data.message)

</script>

<style scoped>
</style>
