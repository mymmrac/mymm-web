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

    async login(): Promise<boolean> {
        return await serverAPI.get("/login", { headers: { "Authorization": authStore.authHeader } })
            .then(response => response.status == 202)
    },

    async getBookmarks(): Promise<Bookmarks> {
        return await serverAPI.get<Bookmarks>("/bookmarks")
            .then(response => response.data)
    },

    async addBookmark(newBookmark: NewBookmark): Promise<Bookmark> {
        return await serverAPI.post<Bookmark>("/bookmarks", newBookmark, {
            headers: { "Authorization": authStore.authHeader },
        })
            .then(response => response.data)
    },

    async updateBookmark(bookmark: Bookmark): Promise<void> {
        return await serverAPI.put("/bookmarks", bookmark, {
            headers: { "Authorization": authStore.authHeader },
        })
    },

    async deleteBookmark(bookmarkID: string): Promise<void> {
        return await serverAPI.delete("/bookmarks", {
            headers: { "Authorization": authStore.authHeader },
            data: bookmarkID,
        })
    },
}