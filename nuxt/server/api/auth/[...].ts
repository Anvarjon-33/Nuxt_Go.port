import {NuxtAuthHandler} from "#auth";
import GithubProvider from "next-auth/providers/github";
import GoogleProvider from "next-auth/providers/google";

export default NuxtAuthHandler({
    pages: {
        signIn: "/login"
    },
    providers: [
        // @ts-ignore
        GithubProvider.default({
            clientId: process.env.NUXT_GITHUB_ID,
            clientSecret: process.env.NUXT_GITHUB_SECRET,
        }),
        // @ts-ignore
        GoogleProvider.default({
            clientId: process.env.NUXT_GOOGLE_ID,
            clientSecret: process.env.NUXT_GOOGLE_SECRET,
            // authorization: {
            //     params: {
            //         prompt: "consent",
            //         access_type: "offline",
            //         response_type: "code"
            //     }
            // }
        })
    ],
})
