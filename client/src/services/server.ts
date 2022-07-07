import axios from "axios"

import { Bookmark, NewBookmark, Bookmarks } from "@/entity/bookmarks"

const serverAPI = axios.create({
    baseURL: "http://localhost:8080",
    withCredentials: false,
    headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
    },
})

let authStore: any = null

export default {
    init(auth: any) {
        authStore = auth
    },

    async getBookmarks() {
        return serverAPI.get<Bookmarks>("/bookmarks")
            .then(response => response.data)
    },

    async addBookmark(newBookmark: NewBookmark) {
        return serverAPI.post<Bookmark>("/bookmarks", newBookmark, {
            headers: { "Authorization": authStore.authHeader },
        })
            .then(response => response.data)
    },

    async deleteBookmark(bookmarkID: string) {
        return serverAPI.delete("/bookmarks", {
            headers: { "Authorization": authStore.authHeader },
            data: bookmarkID,
        })
    },
}