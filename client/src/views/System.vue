<template>
    <div class="flex flex-col gap-2">
        <div class="flex gap-2">
            <back-home></back-home>
            <div class="m-box flex items-center gap-2 w-full">
                <i class="bi bi-cpu text-3xl"></i>
                <p class="uppercase">System</p>
            </div>
        </div>

        <div class="m-box">
            System
        </div>
    </div>
</template>

<script lang="ts" setup>
import BackHome from "@/components/BackHome.vue"
</script>

<style lang="scss" scoped></style>
