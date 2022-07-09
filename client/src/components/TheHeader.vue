<template>
    <header class="m-box container my-4 mx-auto flex justify-between">
        <router-link to="/" class="flex items-center">
            <i class="bi bi-x-diamond text-3xl mr-2 m-hover-highlight m-hover-scale"></i>
            <span class="hover:underline">mymm</span>
        </router-link>

        <div class="grid grid-flow-col gap-3">
            <i v-if="!authorized" @click="showLoginModal = true"
               class="bi bi-person-circle text-3xl m-hover-scale m-hover-highlight"></i>
            <i v-else @click="authStore.logout()"
               class="bi bi-box-arrow-left text-3xl m-hover-scale m-hover-highlight"></i>

            <a href="https://github.com/mymmrac" target="_blank" class="m-hover-scale m-hover-highlight">
                <i class="bi bi-github text-3xl"></i>
            </a>
        </div>

        <modal-box :shown="showLoginModal" @closed="closeLoginModal" title="Login" close-button>
            <form class="grid items-center grid-cols-[0.5fr_1fr] p-4 space-y-2">
                <label class="block">Username</label>
                <input type="text" class="m-input" v-model="username" placeholder="...">

                <label class="block">Password</label>
                <input type="password" class="m-input" v-model="password" placeholder="...">

                <div class="col-span-2 h-4"></div>
                <div v-show="loginError" class="col-span-2 py-2 text-center text-red-500">
                    {{ loginError }}
                </div>
                <button class="col-span-2 m-hover-highlight m-hover-scale rounded py-2 m-shadow" @click.prevent="login">
                    Sign In
                </button>
            </form>
        </modal-box>
    </header>
</template>

<script lang="ts" setup>
import ModalBox from "@/components/ModalBox.vue"
import { useAuthStore } from "@/stores/auth"
import axios, { AxiosError } from "axios"
import { storeToRefs } from "pinia"
import { Ref, ref } from "vue"

const authStore = useAuthStore()
let { authorized } = storeToRefs(authStore)

let showLoginModal = ref(false)
let username = ref("")
let password = ref("")
let loginError: Ref<string> = ref("")

function login() {
    if (username.value.length < 1) {
        loginError.value = "Username can't be empty"
        return
    }

    if (password.value.length < 1) {
        loginError.value = "Password can't be empty"
        return
    }

    authStore.login(username.value, password.value)
        .then(ok => {
            if (ok) {
                closeLoginModal()
            } else {
                loginError.value = "Username or password is not correct"
            }
        })
        .catch((error: Error | AxiosError) => {
            if (axios.isAxiosError(error)) {
                if (error.request.status === 401) {
                    loginError.value = "Username or password is not correct"
                } else {
                    loginError.value = error.message
                }
            } else {
                loginError.value = error.message
            }
        })
}

function closeLoginModal() {
    showLoginModal.value = false
    username.value = ""
    password.value = ""
    loginError.value = ""
}
</script>

<style lang="scss" scoped>
#header-nav {
    .router-link-exact-active {
        @apply text-green-400 hover:no-underline;
    }

    a {
        @apply hover:underline;
    }
}
</style>
