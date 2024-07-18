<template>
  <v-card>
    <v-layout>
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
          <v-btn v-if="noLogged" href="/login" icon="mdi-login"></v-btn>
          <v-btn v-else icon="mdi-logout" @click="signOut()"></v-btn>
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

      <v-main style="height: 500px;">
        <!--    Content Here    -->
        <slot></slot>
      </v-main>
    </v-layout>
  </v-card>
</template>

<script lang="ts" setup>
const {status, signOut} = useAuth()
const noLogged = ref(false)
noLogged.value = status.value == "unauthenticated"
let drawer = false;
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
