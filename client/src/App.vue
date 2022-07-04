<template>
    <the-header></the-header>

    <main class="container mx-auto flex-grow relative grid grid-flow-row gap-2">
        <router-view v-slot="{ Component }">
            <transition name="scale" mode="out-in">
                <component :is="Component"/>
            </transition>
        </router-view>
    </main>

    <the-footer></the-footer>
</template>

<script lang="ts" setup>
import TheHeader from "@/components/TheHeader.vue"
import TheFooter from "@/components/TheFooter.vue"
</script>

<style lang="scss">
#app {
    @apply flex flex-col min-h-screen px-2 sm:px-0 bg-gray-100;
}
</style>
