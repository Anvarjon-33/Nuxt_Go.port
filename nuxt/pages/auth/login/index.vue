<template>
  <v-row>
    <v-col
        class="mx-auto pa-3 ma-3"
        cols="12"
        lg="4"
        md="6"
        sm="8"
        xl="2"
        xs="12"
    >
      <v-row class="border rounded-lg bg-brown-lighten-4">
        <v-col class="" cols="12">
          <v-row class="no-gutters">
            <v-col cols="6">
              <v-btn
                  block
                  color="primary"
                  prepend-icon="mdi-google"
                  @click="register_user('google')"
              >google
              </v-btn>
            </v-col>
            <v-col cols="6">
              <v-btn
                  block
                  color="secondary"
                  prepend-icon="mdi-github"
                  @click="register_user('github')"
              >github
              </v-btn>
            </v-col>
            <v-col cols="6">
              <v-btn
                  block
                  color="green-lighten-1"
                  prepend-icon="mdi-aws"
                  @click="register_user('yandex')"
              >yandex
              </v-btn>
            </v-col>
            <v-col cols="6">
              <v-btn
                  block
                  color="blue-lighten-2"
                  prepend-icon="mdi-amazon"
                  @click="register_user('amazon')"
              >amazon
              </v-btn>
            </v-col>
            <v-col cols="12">
              <v-btn
                  block
                  color="brown-darken-1"
                  prepend-icon="mdi-account"
                  @click="show.account = !show.account"
              >create own
              </v-btn>
            </v-col>
          </v-row>
        </v-col>
        <v-col v-if="show.account" cols="12">
          <v-sheet class="pb-2 bg-brown-lighten-4">
            <v-card class="bg-brown-lighten-5 pb-3">
              <v-card-title> Register via Email:</v-card-title>
              <v-card-subtitle>
                You can register via Social Network
              </v-card-subtitle>
              <v-card-text>
                <v-form ref="fields" submit.prevent="register_user">
                  <v-text-field
                      v-model="form.name"
                      :append-inner-icon="
                      fields?.items?.[0]?.isValid ? '' : 'mdi-alert-outline'
                    "
                      :rules="rule.text_required"
                      autofocus
                      placeholder="name"
                      type="text"
                  >
                  </v-text-field>
                  <v-text-field
                      v-model="form.last_name"
                      :append-inner-icon="
                      fields?.items?.[1]?.isValid ? '' : 'mdi-alert-outline'
                    "
                      :rules="rule.text_required"
                      placeholder="last_name"
                      type="text"
                  >
                  </v-text-field>
                  <v-text-field
                      v-model="form.email"
                      :append-inner-icon="
                      fields?.items?.[2]?.isValid ? '' : 'mdi-alert-outline'
                    "
                      :rules="rule.email"
                      placeholder="email"
                      type="text"
                  ></v-text-field>
                  <v-text-field
                      v-model="form.password"
                      :append-inner-icon="
                      show.password ? 'mdi-eye' : 'mdi-eye-off'
                    "
                      :type="show.password ? 'text' : 'password'"
                      placeholder="password"
                      @click:append-inner="show.password = !show.password"
                  >
                  </v-text-field>
                  <v-text-field
                      v-model="form.confirm_password"
                      :append-icon="
                      fields?.items?.[4]?.isValid ? '' : 'mdi-alert-outline'
                    "
                      :append-inner-icon="
                      show.confirm ? 'mdi-eye' : 'mdi-eye-off'
                    "
                      :rules="rule.password"
                      :type="show.confirm ? 'text' : 'password'"
                      placeholder="confirm_password"
                      @click:append-inner="show.confirm = !show.confirm"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-btn
                    color="green" variant="elevated" @click="register_user"
                >send
                </v-btn>
                <v-btn
                    class="ms-7"
                    color="warning"
                    variant="tonal"
                    @click="deny"
                >deny
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-sheet>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script setup>
const show = reactive({
  password: false,
  confirm: false,
  account: false,
});
const fields = ref();

const form = reactive({
  name: "Anvarjon",
  last_name: "Buranov",
  email: "Exmaple@gmail.com",
  password: "kadrkadr",
  confirm_password: "kadrkadr",
});

const deny = () => {
  let res = confirm("All data will be lost !");
  if (!res) return;
  for (let props in form) {
    form[props] = "";
  }
};

const register_user = async () => {
  try {
    let res = await $fetch("http://192.168.1.3:2222/csrf-header")
    for (let key in res) {
      useCookie(key, {
        sameSite: "lax",
        expires: new Date(Date.now() + 1000 * 30)
      }).value = res[key]
    }
  } catch (er) {
    console.log('something went wrong', er)
  }

  for (let props in form) {
    if (!form[props]) return;
    const {signUp} = useAuth();
    let res = await signUp({
      email: form.email,
      password: form.password,
      username: form.name + " " + form.last_name,
    });
    console.log(res.data);
  }
};

const rule = reactive({
  text_required: [
    (val) => !!val || "Required !",
    (val) => val.length > 2 || "Character length must be at last 2 !",
  ],
  email: [
    (val) => !!val || "Field Required !",
    (val) =>
        !!val.match(/[a-zA-Z0-9.*%±]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}/g) ||
        "Email no valid !",
  ],
  password: [
    (val) => !!val || "Enter Confirmation for password",
    (val) => val === form.password || "Password mismatch",
  ],
});
</script>

<style scoped></style>
