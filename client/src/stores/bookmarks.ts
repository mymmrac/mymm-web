import { defineStore } from "pinia"

import { NewBookmark, Bookmarks, Bookmark } from "@/entity/bookmarks"

import server from "@/services/server"

export const useBookmarksStore = defineStore("bookmarks", {
    state: () => ({
        bookmarks: <Bookmarks>[],
    }),

    actions: {
        async loadBookmarks(): Promise<void> {
            return await server.getBookmarks()
                .then(bookmarks => {
                    if (bookmarks) {
                        this.bookmarks = bookmarks
                    } else {
                        this.bookmarks = []
                    }
                })
        },

        async addBookmark(newBookmark: NewBookmark): Promise<void> {
            return await server.addBookmark(newBookmark)
                .then(bookmark => {
                    this.bookmarks.push(bookmark)
                })
        },

        async updateBookmark(bookmark: Bookmark): Promise<void> {
            return await server.updateBookmark(bookmark)
                .then(() => {
                    this.bookmarks = this.bookmarks.map(b => b.id !== bookmark.id ? b : bookmark)
                })
        },

        async deleteBookmark(bookmarkID: string): Promise<void> {
            return await server.deleteBookmark(bookmarkID)
                .then(() => {
                    this.bookmarks = this.bookmarks.filter(bookmark => bookmark.id != bookmarkID)
                })
        },
    },
})