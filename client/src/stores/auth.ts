import { defineStore } from "pinia"

export const useAuthStore = defineStore("auth", {
    state: () => ({
        username: "",
        auth: "",
    }),

    actions: {
        login(username: string, password: string) {
            this.username = username
            this.auth = window.btoa(username + ":" + password)
        },

        logout() {
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
