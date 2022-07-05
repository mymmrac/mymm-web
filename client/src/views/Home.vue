<template>
    <div>
        <div class="grid m-grid gap-2">
            <router-link v-for="item in menu" :to="item.link" tag="div" class="m-box m-item m-hover-scale">
                <i class="bi" :class="`bi-${item.icon}`"></i>
                <p>{{ item.name }}</p>
            </router-link>
        </div>
    </div>
</template>

<script lang="ts" setup>
type MenuButton = {
    name: string,
    link: string,
    icon: string,
}

const menu: MenuButton[] = [
    {
        name: "Bookmarks",
        link: "/bookmarks",
        icon: "journal-bookmark",
    },
    {
        name: "System",
        link: "/system",
        icon: "cpu",
    },
]
</script>

<style lang="scss" scoped>
.m-grid {
    grid-template-columns: repeat(auto-fill, minmax(192px, 1fr));
}

.m-item {
    @apply aspect-square flex flex-col justify-center items-center uppercase;

    i {
        @apply text-8xl mb-4;
    }
}
</style>
