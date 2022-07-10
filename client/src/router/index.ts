import { storeToRefs } from "pinia"
import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

import { useAuthStore } from "@/stores/auth"

import Home from "@/views/Home.vue"
import Bookmarks from "@/views/Bookmarks.vue"
import System from "@/views/System.vue"
import NotFound from "@/views/NotFound.vue"

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "home",
        component: Home,
    },
    {
        path: "/bookmarks",
        name: "bookmarks",
        component: Bookmarks,
    },
    {
        path: "/system",
        name: "system",
        component: System,
    },

    {
        path: "/:pathMatch(.*)*",
        alias: "/404",
        name: "not-found",
        component: NotFound,
    },
]

const publicPages = ["home", "bookmarks", "not-found"]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { top: 0 }
        }
    },
})

router.beforeEach(async (to) => {
    if (!to.name) {
        return { name: "home" }
    }

    const authRequired = !publicPages.includes(to.name.toString())
    const { authorized } = storeToRefs(useAuthStore())

    if (authRequired && !authorized.value) {
        return { name: "home" }
    }
})

export default router