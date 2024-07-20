<template id="app">
  <v-app>
    <v-container class="py-3">
      <div class="display-2">Endless scrolling with v-lazy</div>
      <h5><span v-text="visiblePosts"></span> of <span>{{ posts.length }}</span> posts shown</h5>
      <v-row v-if="posts.length" class="fill-height overflow-y-auto">
        <v-col v-for="(post, index) in posts" cols="12" lg="3" md="4" sm="6">
          <v-sheet class="fill-height" color="transparent" min-height="250">
            <v-lazy
                v-model="post.isActive" :options="{
                                threshold: .5
                            }" class="fill-height"
            >
              <v-card class="fill-height" hover>
                <v-card-text>
                  <v-row :key="index" @click="">
                    <v-col class="text-sm-left text-center" cols="12" sm="10"> #{{ (index + 1) }}
                      <h2 v-html="post.title"></h2>
                      <div v-html="post.body"></div>
                    </v-col>
                  </v-row>
                </v-card-text>
              </v-card>
            </v-lazy>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data() {
    return {
      posts: []
    }
  },
  computed: {
    visiblePosts() {
      return this.posts.filter(p => p.isActive).length
    }
  },
  methods: {
    addPosts() {
      axios.get('https://jsonplaceholder.typicode.com/posts')
          .then(response => {
            this.posts = response.data
          })
    }
  },
  created() {
    this.addPosts()
  }
})


</script>
