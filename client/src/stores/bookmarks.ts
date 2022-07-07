import { defineStore } from "pinia"

import { NewBookmark, Bookmarks } from "@/entity/bookmarks"

import server from "@/services/server"

export const useBookmarksStore = defineStore("bookmarks", {
    state: () => ({
        loading: false,
        error: <any>null,
        bookmarks: <Bookmarks>[],
    }),

    actions: {
        async loadBookmarks() {
            this.loading = true
            server.getBookmarks()
                .then(bookmarks => {
                    if (bookmarks) {
                        this.bookmarks = bookmarks
                    } else {
                        this.bookmarks = []
                    }
                })
                .catch(error => this.error = error)
                .finally(() => this.loading = false)
        },

        async addBookmark(newBookmark: NewBookmark) {
            server.addBookmark(newBookmark)
                .then(bookmark => this.bookmarks.push(bookmark))
                .catch(error => this.error = error)
        },

        async deleteBookmark(bookmarkID: string) {
            server.deleteBookmark(bookmarkID)
                .then(() => this.bookmarks = this.bookmarks.filter(bookmark => bookmark.id != bookmarkID))
                .catch(error => this.error = error)
        },
    },
})