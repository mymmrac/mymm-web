<template>
    <div class="flex flex-col gap-2">
        <div class="flex gap-2">
            <back-home></back-home>
            <div class="m-box flex justify-between items-center gap-2 w-full">
                <div class="flex gap-2 items-center">
                    <i class="bi bi-journal-bookmark text-3xl"></i>
                    <p class="uppercase">Bookmarks</p>
                </div>
                <i class="bi bi-bookmark-plus text-3xl self-end m-hover-scale m-hover-highlight"></i>
            </div>
        </div>

        <div class="m-box flex items-center justify-between gap-3">
            <i class="bi bi-search text-3xl"></i>
            <input type="text" placeholder="Search..."
                   class="flex-1 border-0 border-b focus:ring-0 focus:border-green-400">
            <i class="bi bi-backspace text-3xl m-hover-scale m-hover-highlight"></i>
        </div>

        <div class="flex gap-2">
            <div class="m-box flex items-center gap-2 cursor-pointer m-hover-scale group">
                <i class="bi bi-grid text-2xl text-green-400 m-hover-highlight-group"></i>
                <p>All</p>
            </div>
            <div class="m-box flex items-center gap-2 cursor-pointer m-hover-scale group">
                <i class="bi bi-code-slash text-2xl m-hover-highlight-group"></i>
                <p>Dev</p>
            </div>

        </div>

        <div class="grid m-grid gap-2">
            <a v-for="bookmark in bookmarks" :href="bookmark.link" target="_blank" class="m-box m-item m-hover-scale">
                <img v-if="bookmark.icon" :src="bookmark.icon" alt="Icon"
                     class="border-0 rounded aspect-square w-1/2">
                <i v-else class="bi bi-question-square text-7xl"></i>
                <p class="mt-3">{{ bookmark.name }}</p>
            </a>
        </div>
    </div>
</template>

<script lang="ts" setup>
import BackHome from "@/components/BackHome.vue"
import { ref, Ref } from "vue"

type Bookmark = {
    name: string,
    link: string,
    icon?: string,
}

const bookmarks: Ref<Bookmark[]> = ref([
    {
        name: "GitHub",
        link: "https://github.com",
    },
    {
        name: "Google",
        link: "https://google.com",
    },
])

const api = "https://favicongrabber.com/api/grab/"
for (let i = 0; i < bookmarks.value.length; i++) {
    const url = new URL(bookmarks.value[i].link)

    fetch(api + url.hostname)
        .then(resp => resp.json())
        .then(data => {
            if (!data) {
                return
            }

            const icons: { src: string }[] = data.icons
            if (icons.length > 0) {
                bookmarks.value[i].icon = icons[0].src
            }
        })
}
</script>

<style lang="scss" scoped>
.m-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
}

.m-item {
    @apply aspect-square flex flex-col justify-center items-center;
}
</style>
