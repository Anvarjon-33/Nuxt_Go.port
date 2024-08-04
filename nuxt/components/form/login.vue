<template>
  <v-card class="mx-auto mt-10 pa-2" width="500">
    <v-card-title>
      <v-row>
        <v-col v-for="(provider, name) in providers" :key="name">
          <v-btn
            class="my-5 pa-2 w-100"
            color="secondary"
            @click="signIn(name)"
          >
            <v-icon start>mdi-{{ name }}</v-icon>
            <v-spacer class="mx-3"></v-spacer>
            {{ name }}
          </v-btn>
        </v-col>
      </v-row>
    </v-card-title>

    <v-card-text>
      <v-sheet class="mx-auto">
        <v-form fast-fail @submit.prevent>
          <v-text-field v-model="firstName" label="First name"></v-text-field>

          <v-text-field v-model="lastName" label="Last name"></v-text-field>

          <v-btn block class="mt-2" type="submit">Submit</v-btn>
        </v-form>
      </v-sheet>
    </v-card-text>
  </v-card>
</template>

<script setup>
const firstName = ref();
const lastName = ref();
const { signIn, getProviders } = useAuth();
const providers = await getProviders();
</script>
