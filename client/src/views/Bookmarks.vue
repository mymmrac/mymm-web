<template>
    <div class="flex flex-col gap-2">
        <div class="flex gap-2">
            <back-home></back-home>
            <div class="m-box flex justify-between items-center gap-2 w-full">
                <div class="flex gap-2 items-center">
                    <i class="bi bi-journal-bookmark text-3xl"></i>
                    <p class="uppercase">Bookmarks</p>
                </div>
                <i class="bi bi-bookmark-plus text-3xl self-end cursor-pointer m-hover-scale m-hover-highlight"></i>
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
