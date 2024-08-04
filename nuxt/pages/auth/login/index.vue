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
      <v-row class="border rounded-lg">
        <v-col cols="12">
          <v-btn block color="primary" prepend-icon="mdi-google">google</v-btn>
        </v-col>
        <v-col cols="12">
          <v-btn
              block color="secondary" prepend-icon="mdi-github"
          >github
          </v-btn>
        </v-col>
        <v-col cols="12">
          <v-sheet class="pt-3 px-10">
            <v-card>
              <v-card-title> Register via Email:</v-card-title>
              <v-card-subtitle>
                You can register via Social Network
              </v-card-subtitle>
              <v-card-text>
                <v-form ref="fields" submit.prevent="register_user">
                  <v-text-field
                      v-model="form.name"
                      :append-inner-icon="fields?.items?.[0]?.isValid ? '':'mdi-alert-outline'"
                      :rules="rule.text_required"
                      autofocus
                      placeholder="name"
                      type="text"
                  >
                  </v-text-field>
                  <v-text-field
                      v-model="form.last_name"
                      :append-inner-icon="fields?.items?.[1]?.isValid ? '':'mdi-alert-outline'"
                      :rules="rule.text_required"
                      placeholder="last_name"
                      type="text"
                  >
                  </v-text-field>
                  <v-text-field
                      v-model="form.email"
                      :append-inner-icon="fields?.items?.[2]?.isValid ? '':'mdi-alert-outline'"
                      :rules="rule.email"
                      placeholder="email"
                      type="text"
                  ></v-text-field>
                  <v-text-field
                      v-model="form.password"
                      placeholder="password"
                      type="password"
                  ></v-text-field>
                  <v-text-field
                      v-model="form.confirm_password"
                      :append-inner-icon="fields?.items?.[4]?.isValid ? '':'mdi-alert-outline'"
                      :rules="rule.password"
                      placeholder="confirm_password"
                      type="password"
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

<script setup>
const fields = ref()

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
