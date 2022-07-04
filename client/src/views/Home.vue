<template>
    <div class="flex flex-wrap gap-4">
        <router-link to="/bookmarks" tag="div" class="m-box m-item m-hover-scale">
            <i class="bi bi-journal-bookmark"></i>
            <p>Bookmarks</p>
        </router-link>
        <router-link to="/system" tag="div" class="m-box m-item m-hover-scale">
            <i class="bi bi-cpu"></i>
            <p>System</p>
        </router-link>
    </div>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped>
.m-item {
    @apply w-[192px] h-[192px] flex flex-col justify-center items-center uppercase;

    i {
        @apply text-8xl mb-4;
    }
}
</style>
