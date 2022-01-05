import apiClient from "@/services/api";

const systemPrefix = "/system"

export default {
    getCPU() {
        return apiClient.get(`${systemPrefix}/cpu`)
    },
    getRAM() {
        return apiClient.get(`${systemPrefix}/ram`)
    },
    getLoad() {
        return apiClient.get(`${systemPrefix}/load`)
    },
    getDisk() {
        return apiClient.get(`${systemPrefix}/disk`)
    }
}