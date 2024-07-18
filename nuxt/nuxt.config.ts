// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2024-04-03',
    devtools: {enabled: false},
    imports: {autoImport: true},
    modules: ["vuetify-nuxt-module", "@sidebase/nuxt-auth", ["@pinia/nuxt", {}]],
    auth: {
        globalAppMiddleware: true,
    },
    devServer: {
        url: "http://localhost:3000"
    },
    runtimeConfig: {
        google_secret: process.env.NUXT_GOOGLE_SECRET,
        git_secret: process.env.NUXT_GITHUB_SECRET,
        public: {
            git_id: process.env.NUXT_GITHUB_ID,
            google_id: process.env.NUXT_GOOGLE_ID
        }
    }
})
