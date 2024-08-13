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
            this.data = await $fetch(url, {..._params, ...this.params,})
        }
    },
    getters: {}
})
