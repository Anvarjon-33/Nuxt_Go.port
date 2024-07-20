<template>
  <h1>
    <fetch-from-url
        method="POST"
        url="http://192.168.1.3:2222"
    >
      <template #default="{data}">
        <v-row class="full-height overflow-y-auto">
          <v-col v-for="(v,k) in data" cols="12" lg="3" md="4" sm="6">
            <v-lazy :min-height="200" :options="{'threshold': 0}" transition="fade-transition">
              <v-card class="full-height" hover>
                <v-card-title>{{ v['last_name'] }}</v-card-title>
                <v-card-subtitle>{{ v['first_name'] }}</v-card-subtitle>
                <v-card-text>
                  Example Text -- {{ v['full_name'] }}
                </v-card-text>
                <v-card-actions>
                  <v-btn color="success">Example button-1</v-btn>
                  <v-btn color="warning">Example button-2</v-btn>
                </v-card-actions>
              </v-card>
            </v-lazy>
          </v-col>
        </v-row>
      </template>
    </fetch-from-url>
    The Index Page
  </h1>
</template>

<script lang="ts" setup>
</script>

<style scoped>
</style>
