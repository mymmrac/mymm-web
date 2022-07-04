<template>
    <header class="m-box container my-4 mx-auto flex justify-between">
        <router-link to="/" class="flex items-center">
            <i class="bi bi-x-diamond text-3xl mr-2 m-hover-highlight m-hover-scale"></i>
            <span class="hover:underline">mymm</span>
        </router-link>

        <div class="grid grid-flow-col gap-3">
            <i class="bi bi-person-circle text-3xl m-hover-scale m-hover-highlight"></i>
            <a href="https://github.com/mymmrac" target="_blank" class="m-hover-scale m-hover-highlight">
                <i class="bi bi-github text-3xl"></i>
            </a>
        </div>
    </header>
</template>

<script lang="ts" setup></script>

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
