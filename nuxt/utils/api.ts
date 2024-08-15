export function goApi() {
    if (tryUseNuxtApp()) {
        const {$api} = useNuxtApp()
        return $api
    }
}
