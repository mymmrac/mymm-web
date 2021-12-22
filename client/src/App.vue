<template>
    <the-header></the-header>

    <main class="container mx-auto flex-grow">
        <router-view v-slot="{ Component }">
            <transition name="scale" mode="out-in">
                <component :is="Component"/>
            </transition>
        </router-view>
    </main>

    <the-footer></the-footer>
</template>

<script>
import TheHeader from "@/components/TheHeader.vue";
import TheFooter from "@/components/TheFooter.vue";

export default {
    name: "App",
    components: {TheFooter, TheHeader},
};
</script>

<style lang="scss">
#app {
    @apply flex flex-col min-h-screen px-2 sm:px-0 bg-gray-100;
}

.scale-enter-active,
.scale-leave-active {
    transition: all 0.3s ease;
}

.scale-enter-from,
.scale-leave-to {
    opacity: 0;
    transform: scale(0.98);
}
</style>
