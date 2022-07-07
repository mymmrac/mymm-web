export type Bookmark = {
    id: string,
    name: string,
    link: string,
    category: string,
    iconLink?: string,
}

export type Bookmarks = Bookmark[]

export type Category = {
    name: string,
    value: string,
    icon: string,
}

export type Categories = Category[]

export type NewBookmark = {
    name: string,
    link: string,
    category: string,
}
