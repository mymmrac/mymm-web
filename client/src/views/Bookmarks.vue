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

        <div class="m-box">
            Bookmarks
        </div>
    </div>
</template>

<script lang="ts" setup>
import BackHome from "@/components/BackHome.vue"
</script>

<style lang="scss" scoped></style>
