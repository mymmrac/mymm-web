import {createRouter, createWebHistory} from "vue-router"
import Home from "@/views/Home.vue"
import Utils from "@/views/Utils.vue"

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home
    },
    {
        path: "/utils",
        name: "Utils",
        component: Utils
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return {top: 0}
        }
    },
})

export default router
