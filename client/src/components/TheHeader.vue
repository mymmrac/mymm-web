<template>
    <header
        class="bg-white container my-4 mx-auto py-2 px-3 rounded shadow grid grid-cols-3 items-center"
    >
        <router-link to="/" class="flex items-center">
            <i class="bi bi-x-diamond text-3xl mr-2 hover:text-green-400"></i>
            <span>mymm</span>
        </router-link>
        <div class="justify-self-center grid grid-flow-col gap-2" id="header-nav">
            <router-link to="/">Home</router-link>
            <router-link to="/utils">Utils</router-link>
        </div>
        <div class="justify-self-end">
            <a href="https://github.com/mymmrac" target="_blank">
                <i class="bi bi-github text-3xl hover:text-green-400"></i>
            </a>
        </div>
    </header>
</template>

<script>
export default {
    name: "TheHeader",
};
</script>

<style scoped lang="scss">
#header-nav {
    .router-link-exact-active {
        @apply text-green-400 hover:no-underline;
    }

    a {
        @apply hover:underline;
    }
}
</style>
