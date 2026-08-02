
export function canExecuteFastAttack(knightIsAwake){
    return !knightIsAwake
}

export function canSpy(knightIsAwake, archerIsAwake, prisonerIsAwake){
    return knightIsAwake || archerIsAwake || prisonerIsAwake
}

export function canSignalPrisoner(archerIsAwake, prisonerIsAwake){
    return prisonerIsAwake && !archerIsAwake
}

export function canFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, dog){
    if(dog && !archerIsAwake){
        return true
    }
    return prisonerIsAwake && !knightIsAwake && !archerIsAwake
}
