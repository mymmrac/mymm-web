<template>
    <div class="flex flex-col gap-2">
        <div class="flex gap-2">
            <back-home></back-home>
            <div class="m-box flex items-center gap-2 w-full">
                <i class="bi bi-cpu text-3xl"></i>
                <p class="uppercase">System</p>
            </div>
        </div>

        <div class="m-box" v-if="systemErr">
            Error: {{ systemErr }}
        </div>

        <div class="m-box" v-if="systemInfo">
            <pre>{{ JSON.stringify(systemInfo, null, 4) }}</pre>
        </div>
    </div>
</template>

<script lang="ts" setup>
import BackHome from "@/components/BackHome.vue"
import { System } from "@/entity/system"
import server from "@/services/server"
import { Ref, ref } from "vue"

let systemInfo: Ref<System> = ref(<System>{})
let systemErr: Ref<string> = ref("")
server.getSystemInfo()
    .then(system => {
        systemInfo.value = system
    })
    .catch(error => {
        systemErr = error
    })
</script>

<style lang="scss" scoped></style>
