<template>
  <v-app>
    <v-app-bar
        color="primary"
        density="compact"
        prominent
    >
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>My files</v-toolbar-title>

      <v-spacer></v-spacer>

      <template v-if="$vuetify.display.mdAndUp">
        <v-btn icon="mdi-magnify" variant="text"></v-btn>
        <v-btn icon="mdi-filter" variant="text"></v-btn>
      </template>

      <div>
        <v-btn icon="mdi-logout"></v-btn>
      </div>

      <v-btn icon="mdi-dots-vertical" variant="text"></v-btn>
    </v-app-bar>

    <v-navigation-drawer
        v-model="drawer"
        :location="$vuetify.display.mobile ? 'bottom' : undefined"
        temporary
    >
      <v-list
          :items="items"
      ></v-list>
    </v-navigation-drawer>

    <v-main style="min-height: 500px;">
      <!--    Content Here    -->
      <v-container>
        <slot></slot>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
let drawer = ref(false);
let group = null
let items = [
  {
    title: 'Foo',
    value: 'foo',
  },
  {
    title: 'Bar',
    value: 'bar',
  },
  {
    title: 'Fizz',
    value: 'fizz',
  },
  {
    title: 'Buzz',
    value: 'buzz',
  },
]

</script>
