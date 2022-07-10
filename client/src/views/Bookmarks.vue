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
            <input type="text" placeholder="Search..." class="flex-1 m-input" v-model="searchInput">
            <i class="bi bi-backspace text-3xl m-hover-scale m-hover-highlight" @click="searchInput = ''"></i>
        </div>

        <div class="flex flex-wrap gap-2">
            <div v-for="category in categories" :key="category.value" @click="selectedCategory = category.value"
                 class="m-box flex items-center gap-2 cursor-pointer m-hover-scale group">
                <i class="bi text-2xl m-hover-highlight-group transition duration-400"
                   :class="`bi-${category.icon} ${selectedCategory === category.value ? 'text-green-400' : ''}`"></i>
                <p>{{ category.name }}</p>
            </div>
        </div>

        <div class="m-box" v-if="bookmarksError">
            Error: {{ bookmarksError }}
        </div>

        <transition-group tag="div" class="grid m-grid gap-2 relative">
            <div v-for="bookmark in displayedBookmarks" :key="bookmark.id"
                 class="m-box m-item m-hover-scale relative group">
                <a :href="bookmark.link" target="_blank" class="flex flex-col justify-center items-center">
                    <img v-if="bookmark.iconLink" :src="bookmark.iconLink" alt="Icon"
                         class="border-0 rounded aspect-square w-1/2">
                    <i v-else class="bi bi-question-square text-7xl"></i>
                    <p class="mt-3">{{ bookmark.name }}</p>
                </a>

                <div class="absolute top-2 right-2 flex flex-col">
                    <i class="bi bi-trash hidden group-hover:block m-hover-highlight m-hover-scale"
                       v-if="authorized" @click.stop="askToDeleteBookmark(bookmark)"></i>

                    <i class="bi bi-pencil-square hidden group-hover:block m-hover-highlight m-hover-scale"
                       v-if="authorized" @click.stop="openEditModal(JSON.parse(JSON.stringify(bookmark)))"></i>
                </div>

                <i class="bi absolute bottom-1 left-2" :class="`bi-${getCategory(bookmark.category).icon}`"></i>
            </div>
        </transition-group>

        <modal-box :shown="showAddModal" @closed="closeNewBookmark" title="New Bookmark" close-button>
            <form class="grid items-center grid-cols-[0.5fr_1fr] p-4 space-y-2">
                <label class="block">Name</label>
                <input type="text" class="m-input" v-model="newBookmark.name" placeholder="...">

                <label class="block">Link</label>
                <input type="text" class="m-input" v-model="newBookmark.link" placeholder="...">

                <label class="block">Category</label>
                <select v-model="newBookmark.category" class="m-input">
                    <option v-for="category in categories.slice(1)" :key="category.value" :value="category.value">
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

        <modal-box :shown="showEditModal" @closed="closeEditBookmark" title="Edit Bookmark" close-button>
            <form class="grid items-center grid-cols-[0.5fr_1fr] p-4 space-y-2">
                <label class="block">Name</label>
                <input type="text" class="m-input" v-model="editBookmark.name" placeholder="...">

                <label class="block">Link</label>
                <input type="text" class="m-input" v-model="editBookmark.link" placeholder="...">

                <label class="block">Category</label>
                <select v-model="editBookmark.category" class="m-input">
                    <option v-for="category in categories.slice(1)" :key="category.value" :value="category.value">
                        {{ category.name }}
                    </option>
                </select>

                <label class="block">Icon Link</label>
                <input type="text" class="m-input" v-model="editBookmark.iconLink" placeholder="...">

                <div class="col-span-2 h-4"></div>
                <div v-show="editBookmarkError" class="col-span-2 py-2 text-center text-red-500">
                    {{ editBookmarkError }}
                </div>
                <button class="col-span-2 m-hover-highlight m-hover-scale rounded py-2 m-shadow disabled:text-gray-400"
                        @click.prevent="updateBookmark" :disabled="editLoading">
                    {{ editLoading ? "Updating..." : "Update" }}
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

import { computed, ComputedRef, ref, Ref } from "vue"
import { storeToRefs } from "pinia"

import { useAuthStore } from "@/stores/auth"
import { useBookmarksStore } from "@/stores/bookmarks"
import { Bookmark, NewBookmark, Categories, Category, Bookmarks } from "@/entity/bookmarks"

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
        name: "All",
        value: "all",
        icon: "grid",
    },
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
        name: "Assets",
        value: "assets",
        icon: "image",
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

let selectedCategory: Ref<string> = ref("all")
let searchInput: Ref<string> = ref("")

const displayedBookmarks: ComputedRef<Bookmarks> = computed(() => {
    let resultingBookmarks = bookmarks.value.slice()

    if (selectedCategory.value !== "all") {
        resultingBookmarks = resultingBookmarks.filter(bookmark => bookmark.category === selectedCategory.value)
    }

    if (searchInput.value !== "") {
        resultingBookmarks = resultingBookmarks.filter(bookmark => {
            const bookmarkInfo = bookmark.name + " " + bookmark.link

            let ok = false

            const keywords = searchInput.value.split(" ")
            for (let i = 0; i < keywords.length; i++) {
                if (!bookmarkInfo.includes(keywords[i])) {
                    ok = true
                    break
                }
            }

            return !ok
        })
    }

    return resultingBookmarks
})

let showAddModal = ref(false)
let newBookmark: Ref<NewBookmark> = ref({
    name: "",
    link: "",
    category: "",
})
let newBookmarkError: Ref<string> = ref("")
let addLoading = ref(false)

function validateBookmark(name: string, link: string, category: string): string | null {
    if (name.length < 1) {
        return "Name can't be empty"
    }

    if (link.length < 1) {
        return "Link can't be empty"
    }

    try {
        const url = new URL(link)
        if (!url.hostname) {
            return "Link is not a valid URL"
        }
    } catch (error) {
        return "Link can't be parsed to URL"
    }

    if (category.length < 1) {
        return "Category should be selected"
    }

    if (!categories.value.find(c => c.value == category)) {
        return "Category should in list of defined categories"
    }

    return null
}

function addBookmark() {
    newBookmarkError.value = ""
    const nb = newBookmark.value

    const err = validateBookmark(nb.name, nb.link, nb.category)
    if (err) {
        newBookmarkError.value = err
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

let showEditModal = ref(false)
let editBookmark: Ref<Bookmark> = ref({
    id: "",
    name: "",
    link: "",
    category: "",
})
let editBookmarkError: Ref<string> = ref("")
let editLoading = ref(false)

function openEditModal(bookmark: Bookmark) {
    showEditModal.value = true
    editBookmark.value = bookmark
}

function updateBookmark() {
    const bookmark = editBookmark.value
    const err = validateBookmark(bookmark.name, bookmark.link, bookmark.category)
    if (err) {
        editBookmarkError.value = err
        return
    }

    editLoading.value = true
    bookmarksStore.updateBookmark(bookmark)
        .then(closeEditBookmark)
        .catch((error) => {
            editBookmarkError.value = "Error: " + error
            return
        })
        .finally(() => {
            editLoading.value = false
        })
}

function closeEditBookmark() {
    showEditModal.value = false
    editBookmarkError.value = ""
    editBookmark.value = {
        id: "",
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
