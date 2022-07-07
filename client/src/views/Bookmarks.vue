<template>
    <div class="flex flex-col gap-2">
        <div class="flex gap-2">
            <back-home></back-home>
            <div class="m-box flex justify-between items-center gap-2 w-full">
                <div class="flex gap-2 items-center">
                    <i class="bi bi-journal-bookmark text-3xl"></i>
                    <p class="uppercase">Bookmarks</p>
                </div>
                <i class="bi bi-bookmark-plus text-3xl self-end m-hover-scale m-hover-highlight"
                   @click="showAddModal = true"></i>
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

        <div class="m-box" v-if="loading">
            Loading...
        </div>
        <div class="m-box" v-if="error">
            Error: {{ error }}
        </div>

        <div class="grid m-grid gap-2">
            <a v-for="bookmark in bookmarks" :key="bookmark.id" :href="bookmark.link" target="_blank"
               class="m-box m-item m-hover-scale relative">
                <img v-if="bookmark.iconLink" :src="bookmark.iconLink" alt="Icon"
                     class="border-0 rounded aspect-square w-1/2">
                <i v-else class="bi bi-question-square text-7xl"></i>
                <p class="mt-3">{{ bookmark.name }}</p>
                <i class="bi bi-trash absolute top-2 right-2 m-hover-highlight m-hover-scale"
                   @click="bookmarksStore.deleteBookmark(bookmark.id)"></i>
            </a>
        </div>

        <modal-box :shown="showAddModal" @closed="showAddModal = false" title="New Bookmark" close-button>
            <form class="grid grid-cols-1">
                <label>
                    Name
                    <input type="text" v-model="newBookmark.name">
                </label>
                <label>
                    Link
                    <input type="text" v-model="newBookmark.link">
                </label>
                <label>
                    Category
                    <select v-model="newBookmark.category">
                        <option disabled selected>None</option>
                        <option value="dev">Dev</option>
                    </select>
                </label>

                <button @click.prevent="bookmarksStore.addBookmark(newBookmark)">Add</button>
            </form>
        </modal-box>

        <modal-box :shown="showDeleteModal" title="Do you want to delete?" close-button>
            <div class="flex justify-center gap-4">
                <button class="rounded border bg-red-400 px-3 py-2">Delete</button>
                <button class="rounded border bg-gray-200 px-3 py-2" @click="showDeleteModal = false">Close</button>
            </div>
        </modal-box>
    </div>
</template>

<script lang="ts" setup>
import BackHome from "@/components/BackHome.vue"
import ModalBox from "@/components/ModalBox.vue"

import { ref, Ref } from "vue"
import { storeToRefs } from "pinia"

import { NewBookmark } from "@/entity/bookmarks"
import { useBookmarksStore } from "@/stores/bookmarks"
import { useAuthStore } from "@/stores/auth"

const authStore = useAuthStore()
authStore.login("mymmrac", "pass")

const bookmarksStore = useBookmarksStore()
bookmarksStore.loadBookmarks()

const { bookmarks, loading, error } = storeToRefs(bookmarksStore)

let showAddModal = ref(false)
let newBookmark: Ref<NewBookmark> = ref({
    name: "",
    link: "",
    category: "",
})

let showDeleteModal = ref(false)
</script>

<style lang="scss" scoped>
.m-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
}

.m-item {
    @apply aspect-square flex flex-col justify-center items-center;
}
</style>
