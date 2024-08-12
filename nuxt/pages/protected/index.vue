<template>
  <v-row>
    <v-col class="mx-auto" cols="6">
      <v-card>
        <v-card-title>
          This is Page Must be Protected !
        </v-card-title>
        <v-card-subtitle>
          You are on the Protected Page
        </v-card-subtitle>
        <v-card-text>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Cumque dolor dolore earum ex laborum libero magnam
          nam, nesciunt nisi nobis officia quia quod ullam! Eum facere iusto porro tenetur velit.
        </v-card-text>
        <v-card-actions class="">
          <v-btn class="me-3" color="primary" variant="elevated">Send</v-btn>
          <v-btn color="secondary" variant="elevated">Deny</v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
definePageMeta({
  layout: "default",
  middleware: "auth",
});
</script>

<style scoped></style>
