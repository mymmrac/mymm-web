<template>
    <div class="grid sm:place-content-center">
        <div class="space-y-2">
            <div class="m-box flex flex-col items-center justify-center sm:w-[512px] h-[512px]">
                <i class="bi bi-question-square text-8xl"></i>
                <p class="text-gray-700 my-6">Looks like something is missing?</p>
                <p>404 - Page Not Found</p>
            </div>
            <div class="m-box sm:w-[512px]">
                <i class="bi bi-chevron-double-left"></i>
                <router-link to="/" class="hover:underline">Go back to Home</router-link>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped></style>
