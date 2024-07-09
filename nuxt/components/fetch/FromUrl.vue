<template>
  <slot :data="g.data" name="data"></slot>
</template>

<script lang="ts" setup>
const props = defineProps({
  url: {
    type: String,
    default: "/data"
  },
  GET: {
    type: Boolean,
    default: true
  },
  POST: {
    type: Boolean,
    default: false
  },
  live: {
    type: Boolean,
    default: false
  }
})

const g = gl()
g.params.method = "get"
g.params.url = props.url
g._fetch('/data')
</script>

<style scoped>
</style>
