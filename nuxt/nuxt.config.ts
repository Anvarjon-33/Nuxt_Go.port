// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            link: [{rel: 'icon', type: 'image/png/ico', href: './favicon.ico'}]
        }
    },
    compatibilityDate: '2024-04-03',
    devtools: {enabled: false},
    imports: {autoImport: true},
    modules: [
        "vuetify-nuxt-module",
        "@pinia/nuxt",
        // @ts-ignore
        ["@sidebase/nuxt-auth"]
    ],
    devServer: {
        host: "http://" + process.env.NUXT_HOST
    },

    runtimeConfig: {
        google_secret: process.env.NUXT_GOOGLE_SECRET,
        git_secret: process.env.NUXT_GITHUB_SECRET,
        public: {
            git_id: process.env.NUXT_GITHUB_ID,
            google_id: process.env.NUXT_GOOGLE_ID,
        }
    },
    auth: {
        baseURL: 'http://192.168.1.3:2222/api/auth/',
        provider: {
            type: 'refresh',
            endpoints: {
                signIn: {path: 'login', method: 'post'},
                getSession: {path: 'session', method: 'post'},
                signOut: {path: 'logout', method: 'post'},
                signUp: {path: 'register', method: 'post'},
                // refresh: {path: 'login', method: 'post'},
            },
            pages: {
                login: "/auth/login"
            }
        },
    }
})
