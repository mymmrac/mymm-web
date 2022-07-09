import { defineStore } from "pinia"

import server from "@/services/server"

export const useAuthStore = defineStore("auth", {
    state: () => ({
        authorized: false,
        username: "",
        auth: "",
    }),

    actions: {
        async login(username: string, password: string): Promise<boolean> {
            this.username = username
            this.auth = window.btoa(username + ":" + password)

            return await server.login()
                .then(ok => {
                    if (ok) {
                        this.authorized = true
                    }
                    return ok
                })
        },

        logout() {
            this.authorized = false
            this.username = ""
            this.auth = ""
        },
    },

    getters: {
        authHeader(state) {
            return `Basic ${ state.auth }`
        },
    },
})
