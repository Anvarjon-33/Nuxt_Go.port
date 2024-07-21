<template>
  <slot :data="g.data" name="data"></slot>
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

const g = gl()
g._fetch('/data')

</script>

<style scoped>
</style>
