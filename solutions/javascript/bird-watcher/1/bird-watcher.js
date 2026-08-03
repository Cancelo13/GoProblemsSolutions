// @ts-check

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birds){
    let sum = 0;
    for(let i = 0 ; i < birds.length ; i++){
        sum += birds[i];
    }
    return sum;
}


/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birds, week){
    let sum = 0;
    for(let i = (week - 1) * 7 ; i < (week * 7)  ; i++){
        sum += birds[i];
    }
    return sum;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {void} should not return anything
 */
export function fixBirdCountLog(birds){
    for(let i = 0 ; i < birds.length ; i += 2){
        birds[i]++;
    }
    return birds;
}
