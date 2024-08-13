export const gl = defineStore("global_POST_data", {
    state: () => ({
        data: {} as any,
        params: {
            method: 'get',
            url: "/",
            headers: {
                "X-Requested-With": "XMLHttpRequest"
            }
        }
    }),
    actions: {
        async _fetch(url: string, _params = false as {}) {
            if (url.slice(0, 4) != 'http') {
                url = "http://192.168.1.3:2222" + url;
            }
            const {data, error} = await useAsyncData('modules', () => goApi()(url, {..._params, ...this.params}))
            this.data = data
        }
    },
    getters: {}
})

/*
export const interceptor = defineStore("fetch with state", {
    state: () => ({}),
    actions: {},
    getters: {}
})

let res = useFetch("/some_data", {
    onRequest(context: FetchContext): Promise<void> | void {
    },
    onResponse(context: FetchContext & { response: FetchResponse<R> }): Promise<void> | void {
    },
    onRequestError(context: FetchContext & { error: Error }): Promise<void> | void {
    },
    onResponseError(context: FetchContext & { response: FetchResponse<R> }): Promise<void> | void {
    }
})
*/
