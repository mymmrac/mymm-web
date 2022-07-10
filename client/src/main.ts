import server from "@/services/server"
import { useAuthStore } from "@/stores/auth"
import { createApp, watch } from "vue"
import { createPinia } from "pinia"

import App from "./App.vue"
import router from "@/router"

import "bootstrap-icons/font/bootstrap-icons.css"
import "./assets/scss/index.scss"

const pinia = createPinia()

const savedAuth = localStorage.getItem("auth")
if (savedAuth) {
    pinia.state.value.auth = JSON.parse(savedAuth)
}

watch(() => pinia.state.value.auth, (state) => {
    localStorage.setItem("auth", JSON.stringify(state))
}, { deep: true })

createApp(App)
    .use(pinia)
    .use(router)
    .mount("#app")

const authStore = useAuthStore()
server.init(authStore)
