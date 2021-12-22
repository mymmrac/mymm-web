<template>
    <footer
        class="bg-white container my-4 mx-auto py-2 px-3 rounded shadow grid grid-cols-3 items-center"
    >
        <div>
            &copy; 2021
            <a
                href="https://github.com/mymmrac"
                target="_blank"
                class="hover:underline"
            >@mymmrac</a
            >
        </div>
        <div class="justify-self-center">
            <router-link to="/">
                <i class="bi bi-x-diamond text-3xl hover:text-green-400"></i>
            </router-link>
        </div>
        <div class="justify-self-end grid grid-flow-col gap-2">
            <router-link class="hover:underline" to="/">Home</router-link>
            <router-link class="hover:underline" to="/utils">Utils</router-link>
        </div>
    </footer>
</template>

<script>
export default {
    name: "TheFooter",
};
</script>

<style scoped lang="scss"></style>
