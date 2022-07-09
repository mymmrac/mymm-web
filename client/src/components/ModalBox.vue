<template>
    <div v-if="show" @click="close" ref="modal" @keydown.esc="close" tabindex="0"
         class="fixed bg-gray-700/25 top-0 left-0 bottom-0 right-0 grid place-content-center p-2 z-50">
        <div class="m-box relative" @click.stop>
            <div v-if="title" class="flex justify-between items-center gap-4">
                <p class="text-lg">{{ title }}</p>
                <i v-if="closeButton" @click="close" class="bi bi-x-square text-xl m-hover-scale m-hover-highlight"></i>
            </div>
            <slot></slot>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Ref, ref, watch } from "vue"

const emit = defineEmits(["closed"])

const props = defineProps<{
    shown?: boolean,
    title?: string,
    closeButton?: boolean,
}>()

const show = ref(props.shown)
watch(() => props.shown, (isShown) => {
    show.value = isShown
})

const modal: Ref<HTMLInputElement | null> = ref(null)
watch(modal, (m) => {
    m?.focus()
})

function close() {
    emit("closed")
}
</script>

<style lang="scss" scoped></style>
