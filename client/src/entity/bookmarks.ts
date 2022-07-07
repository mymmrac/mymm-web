export type Bookmark = {
    id: string,
    name: string,
    link: string,
    category: string,
    iconLink?: string,
}

export type Bookmarks = Bookmark[]

export type NewBookmark = {
    name: string,
    link: string,
    category: string,
}
