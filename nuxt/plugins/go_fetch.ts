import type {FetchContext, FetchResponse} from "ofetch";

export default defineNuxtPlugin(app => {
    const csrfHeader = ref<string>("")
    const api = $fetch.create({
        baseURL: 'http://192.168.1.3:2222/',
        onRequest({request, options, error}) {
            const headers = options.headers ||= {}
            if (Array.isArray(headers)) {
                console.log("array")
                headers.push(['X-CSRF-Token', csrfHeader.value])
            } else if (headers instanceof Headers) {
                console.log("instance")
                headers.set('X-CSRF-Token', csrfHeader.value)
            } else {
                console.log("single")
                headers['X-CSRF-Token'] = csrfHeader.value
            }
        },
        onResponse(context: FetchContext & { response: FetchResponse<ResponseType> }): Promise<void> | void {
            const {response} = context
            let csrf = response.headers.get("X-CSRF-Token")
            if (!!csrf) {
                csrfHeader.value = csrf
            }
        },
    })
    return {
        provide: {
            api
        }
    }

})
