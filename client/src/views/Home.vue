<template>
    <div class="bg-white p-2 shadow rounded">
        <h1>CPU</h1>
        <ol v-if="cpu">
            <li v-for="(core, index) in cpu.cores" :key="index">{{ core }}</li>
        </ol>
        <p v-else>Loading...</p>
    </div>
</template>

<script>
export default {
    name: "Home",
    data() {
        return {
            polling: null
        }
    },
    methods: {
        startPolling() {
            this.polling = setInterval(() => {
                this.$store.dispatch("fetchCPU")
                    .catch(error => {
                        console.log(error)
                    })
            }, 3000)
        }
    },
    mounted() {
        this.startPolling()
    },
    beforeUnmount() {
        clearInterval(this.polling)
        this.$store.commit("SET_CPU", null)
    },
    computed: {
        cpu() {
            return this.$store.state.cpu
        }
    }
}
</script>

<style scoped lang="scss"></style>
