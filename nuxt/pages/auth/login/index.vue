<template>
  <v-row>
    <v-col class="mx-auto pa-3 ma-3" cols="12" lg="4" md="6" sm="8" xl="2" xs="12">
      <v-row class="border rounded-lg">
        <v-col cols="12">
          <v-btn block color="primary" prepend-icon="mdi-google">google</v-btn>
        </v-col>
        <v-col cols="12">
          <v-btn block color="secondary" prepend-icon="mdi-github">github</v-btn>
        </v-col>
        <v-col cols="12">
          <v-sheet class="pt-3 px-10">
            <v-card>
              <v-card-title>
                Register via Email:
              </v-card-title>
              <v-card-subtitle>
                You can register via Social Network
              </v-card-subtitle>
              <v-card-text>
                <v-form submit.prevent="register">

                  <v-text-field
                      v-model="form.name"
                      :rules="rule.text_required"
                      autofocus
                      placeholder="name"
                      type="text"
                  >
                  </v-text-field>


                  <v-text-field
                      v-model="form.last_name" :rules="rule.text_required" placeholder="last_name" type="text"
                  ></v-text-field>

                  <v-text-field
                      v-model="form.email" :rules="rule.email" placeholder="email" type="text"
                  ></v-text-field>

                  <v-text-field
                      v-model="form.password" placeholder="password" type="password"
                  ></v-text-field>

                  <v-text-field
                      v-model="form.confirm_password"
                      :rules="rule.password"
                      placeholder="confirm_password" type="password"
                  ></v-text-field>

                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-btn color="green" @click="register_user">send</v-btn>
                <v-btn color="warning" @click="deny">deny</v-btn>
              </v-card-actions>
            </v-card>
          </v-sheet>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script lang="ts" setup>
const form = reactive({
  name: "Anvarjon",
  last_name: "Buranov",
  email: "Exmaple@gmail.com",
  password: "kadrkadr",
  confirm_password: "kadrkadr",
})

const valid = ref(false)

const deny = () => {
  let res = confirm("All data will be lost !")
  if (!res) return;
  for (let props in form) {
    // @ts-ignore
    form[props] = ""
  }
}

const register_user = async () => {
  for (let props in form) {
    // @ts-ignore
    if (!form[props]) return;
    const {signUp} = useAuth();
    let res = await signUp({
      email: form.email,
      password: form.password,
      username: form.name + " " + form.last_name
    })
    console.log(res.data)
  }
}

const rule = reactive({
  text_required: [
    (val: string) => !!val || 'Required !',
    (val: string) => val.length > 2 || 'Character length must be at last 2 !',
  ],
  email: [
    (val: string) => !!val || 'Field Required !',
    (val: string) => val.match(/[a-zA-Z0-9.*%±]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}/g) || 'Email no valid !'
  ],
  password: [
    (val: string) => !!val || 'Enter Confirmation for password',
    (val: string) => val === form.password || 'Password mismatch',
  ],
})
</script>

<style scoped>
</style>
