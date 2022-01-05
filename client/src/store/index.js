import {createStore} from "vuex"
import systemService from "@/services/system";

export default createStore({
    state: {
        cpu: null,
        ram: null,
        load: null,
        disk: null,
    },
    mutations: {
        SET_CPU(state, cpu) {
            state.cpu = cpu
        },
        SET_RAM(state, ram) {
            state.ram = ram
        },
        SET_LOAD(state, load) {
            state.load = load
        },
        SET_DISK(state, disk) {
            state.disk = disk
        },
    },
    actions: {
        fetchCPU({commit}) {
            return systemService.getCPU()
                .then(response => {
                    commit("SET_CPU", response.data)
                })
                .catch(error => {
                    throw error
                })
        },
        fetchRAM({commit}) {
            return systemService.getRAM()
                .then(response => {
                    commit("SET_RAM", response.data)
                })
                .catch(error => {
                    throw error
                })
        },
        fetchLoad({commit}) {
            return systemService.getLoad()
                .then(response => {
                    commit("SET_LOAD", response.data)
                })
                .catch(error => {
                    throw error
                })
        },
        fetchDisk({commit}) {
            return systemService.getDisk()
                .then(response => {
                    commit("SET_DISK", response.data)
                })
                .catch(error => {
                    throw error
                })
        },
    },
    modules: {}
})
