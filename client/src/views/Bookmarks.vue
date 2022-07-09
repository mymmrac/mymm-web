<template>
    <div class="flex flex-col gap-2">
        <div class="flex gap-2">
            <back-home></back-home>
            <div class="m-box flex justify-between items-center gap-2 w-full">
                <div class="flex gap-2 items-center">
                    <i class="bi bi-journal-bookmark text-3xl"></i>
                    <p class="uppercase">Bookmarks</p>
                </div>
                <i v-if="authorized" class="bi bi-bookmark-plus text-3xl self-end m-hover-scale m-hover-highlight"
                   @click="showAddModal = true"></i>
            </div>
        </div>

        <div class="m-box flex items-center justify-between gap-3">
            <i class="bi bi-search text-3xl"></i>
            <input type="text" placeholder="Search..."
                   class="flex-1 m-input">
            <i class="bi bi-backspace text-3xl m-hover-scale m-hover-highlight"></i>
        </div>

        <div class="flex gap-2">
            <div class="m-box flex items-center gap-2 cursor-pointer m-hover-scale group">
                <i class="bi bi-grid text-2xl text-green-400 m-hover-highlight-group"></i>
                <p>All</p>
            </div>
            <div v-for="category in categories" :key="category.value"
                 class="m-box flex items-center gap-2 cursor-pointer m-hover-scale group">
                <i class="bi text-2xl m-hover-highlight-group" :class="`bi-${category.icon}`"></i>
                <p>{{ category.name }}</p>
            </div>

        </div>

        <div class="m-box" v-if="bookmarksError">
            Error: {{ bookmarksError }}
        </div>

        <div class="grid m-grid gap-2">
            <div v-for="bookmark in bookmarks" :key="bookmark.id" class="m-box m-item m-hover-scale relative group">
                <a :href="bookmark.link" target="_blank" class="flex flex-col justify-center items-center">
                    <img v-if="bookmark.iconLink" :src="bookmark.iconLink" alt="Icon"
                         class="border-0 rounded aspect-square w-1/2">
                    <i v-else class="bi bi-question-square text-7xl"></i>
                    <p class="mt-3">{{ bookmark.name }}</p>
                </a>

                <i class="bi bi-trash absolute top-2 right-2 hidden group-hover:block m-hover-highlight m-hover-scale"
                   v-if="authorized" @click.stop="askToDeleteBookmark(bookmark)"></i>

                <i class="bi absolute bottom-1 left-2" :class="`bi-${getCategory(bookmark.category).icon}`"></i>
            </div>
        </div>

        <modal-box :shown="showAddModal" @closed="closeNewBookmark" title="New Bookmark" close-button>
            <form class="grid items-center grid-cols-[0.5fr_1fr] p-4 space-y-2">
                <label class="block">Name</label>
                <input type="text" class="m-input" v-model="newBookmark.name" placeholder="...">

                <label class="block">Link</label>
                <input type="text" class="m-input" v-model="newBookmark.link" placeholder="...">

                <label class="block">Category</label>
                <select v-model="newBookmark.category" class="m-input">
                    <option v-for="category in categories" :key="category.value" :value="category.value">
                        {{ category.name }}
                    </option>
                </select>

                <div class="col-span-2 h-4"></div>
                <div v-show="newBookmarkError" class="col-span-2 py-2 text-center text-red-500">
                    {{ newBookmarkError }}
                </div>
                <button class="col-span-2 m-hover-highlight m-hover-scale rounded py-2 m-shadow disabled:text-gray-400"
                        @click.prevent="addBookmark" :disabled="addLoading">
                    {{ addLoading ? "Adding..." : "Add" }}
                </button>
            </form>
        </modal-box>

        <modal-box :shown="showDeleteModal" @closed="showDeleteModal = false" title="Do you want to delete?"
                   close-button>
            <div class="p-4">
                <div v-show="deleteBookmarkError" class="col-span-2 py-2 text-center text-red-500">
                    {{ deleteBookmarkError }}
                </div>
                <button class="m-hover-scale rounded py-2 px-3 w-full m-shadow text-red-500" @click="deleteBookmark">
                    Delete
                </button>
            </div>

        </modal-box>
    </div>
</template>

<script lang="ts" setup>
import BackHome from "@/components/BackHome.vue"
import ModalBox from "@/components/ModalBox.vue"

import { ref, Ref } from "vue"
import { storeToRefs } from "pinia"

import { useAuthStore } from "@/stores/auth"
import { useBookmarksStore } from "@/stores/bookmarks"
import { Bookmark, NewBookmark, Categories, Category } from "@/entity/bookmarks"

const authStore = useAuthStore()
let { authorized } = storeToRefs(authStore)

const bookmarksStore = useBookmarksStore()
const { bookmarks } = storeToRefs(bookmarksStore)

let bookmarksError: Ref<string> = ref("")
bookmarksStore.loadBookmarks()
    .catch(error => {
        bookmarksError.value = error
    })

const categories: Ref<Categories> = ref([
    {
        name: "Dev",
        value: "dev",
        icon: "code-slash",
    },
    {
        name: "Utils",
        value: "utils",
        icon: "paperclip",
    },
    {
        name: "Converters",
        value: "converters",
        icon: "recycle",
    },
])

function getCategory(category: string): Category {
    return categories.value.find(c => c.value == category)!
}

let showAddModal = ref(false)
let newBookmark: Ref<NewBookmark> = ref({
    name: "",
    link: "",
    category: "",
})
let newBookmarkError: Ref<string> = ref("")
let addLoading = ref(false)

function addBookmark() {
    newBookmarkError.value = ""
    const nb = newBookmark.value

    if (nb.name.length < 1) {
        newBookmarkError.value = "Name can't be empty"
        return
    }

    if (nb.link.length < 1) {
        newBookmarkError.value = "Link can't be empty"
        return
    }

    try {
        const url = new URL(nb.link)
        if (!url.hostname) {
            newBookmarkError.value = "Link is not a valid URL"
            return
        }
    } catch (error) {
        newBookmarkError.value = "Link can't be parsed to URL"
        return
    }

    if (nb.category.length < 1) {
        newBookmarkError.value = "Category should be selected"
        return
    }

    addLoading.value = true
    bookmarksStore.addBookmark(nb)
        .then(closeNewBookmark)
        .catch((error) => {
            newBookmarkError.value = "Error: " + error
            return
        })
        .finally(() => {
            addLoading.value = false
        })
}

function closeNewBookmark() {
    showAddModal.value = false
    newBookmarkError.value = ""
    newBookmark.value = {
        name: "",
        link: "",
        category: "",
    }
}

let showDeleteModal = ref(false)
let deleteBookmarkData: Ref<Bookmark> = ref(<Bookmark>{})
let deleteBookmarkError: Ref<string> = ref("")

function askToDeleteBookmark(bookmark: Bookmark) {
    deleteBookmarkError.value = ""
    showDeleteModal.value = true
    deleteBookmarkData.value = bookmark
}

function deleteBookmark() {
    bookmarksStore.deleteBookmark(deleteBookmarkData.value.id)
        .then(() => {
            showDeleteModal.value = false
        })
        .catch((error) => {
            deleteBookmarkError.value = error
        })
}
</script>

<style lang="scss" scoped>
.m-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
}

.m-item {
    @apply aspect-square grid;
}
</style>
