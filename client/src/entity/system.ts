export type CPU = {
    cores: number[],
}

export type Load = {
    load1: number,
    load5: number,
    load15: number,
}

export type RAM = {
    total: number,
    free: number,
    used: number,
    usedPercent: number,
}

export type Swap = {
    total: number,
    free: number,
    used: number,
    usedPercent: number,
}

export type Uptime = {
    bootTime: number,
    uptime: number,
}

export type Disk = {
    path: string,
    fsType: string,
    total: number,
    free: number,
    used: number,
    usedPercent: number,
}

export type System = {
    cpu: CPU,
    load: Load,
    ram: RAM,
    swap: Swap,
    uptime: Uptime,
    disk: Disk,
}
