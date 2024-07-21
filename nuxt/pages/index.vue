<template>
  <!--  <fetch-from-url>-->
  <!--  </fetch-from-url>-->
  <form>
    <v-text-field v-model="name" name="first_name" placeholder="name"></v-text-field>
    <v-text-field v-model="last_name" name="last_name" placeholder="last_name"></v-text-field>
    <v-btn>Send</v-btn>
    <fetch-from-url method="POST" url="http://192.168.1.3:2222/data">
      <template #data="{data}">

        <v-row>
          <v-col
              v-for="(val,key) in data['message']" class="height-full" cols="12" lg="3" md="4" sm="6" xl="2"
              xs="12"
          >
            <v-sheet class="pa-1 ma-1" min-height="300">
              <v-lazy :options="{'threshold': .1}" min-height="200" transition="transition-fade">
                <v-card hover>
                  <v-card-title>
                    {{ val['full_name'] }}
                  </v-card-title>
                  <v-card-subtitle>
                    {{ val['first_name'] }}
                  </v-card-subtitle>
                  <v-card-text>
                    {{ val['last_name'] }}
                  </v-card-text>
                  <v-card-actions>
                    <v-btn color="warning" density="compact">send</v-btn>
                    <v-btn color="success" density="compact">deny</v-btn>
                  </v-card-actions>
                </v-card>
              </v-lazy>
            </v-sheet>
          </v-col>
        </v-row>

      </template>
    </fetch-from-url>
  </form>
</template>

<script lang="ts" setup>
const name = ref()
const last_name = ref()

watch([name, last_name], () => {
  g._fetch('/data', {method: 'get'})
})
</script>

<style scoped>
</style>
