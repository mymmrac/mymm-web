import server from "@/services/server"
import { useAuthStore } from "@/stores/auth"
import { createApp } from "vue"
import { createPinia } from "pinia"

import App from "./App.vue"
import router from "@/router"

import "bootstrap-icons/font/bootstrap-icons.css"
import "./assets/scss/index.scss"

const pinia = createPinia()

createApp(App)
    .use(pinia)
    .use(router)
    .mount("#app")

const authStore = useAuthStore()
server.init(authStore)

