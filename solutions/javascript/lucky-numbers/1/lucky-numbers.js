// @ts-check

/**
 * Calculates the sum of the two input arrays.
 *
 * @param {number[]} array1
 * @param {number[]} array2
 * @returns {number} sum of the two arrays
 */
export function twoSum(partA, partB){
    let sA = ""
    let sB = ""
    for(let i = 0 ; i < partA.length ; i++){
        sA += partA[i];
    }
    for(let i = 0 ; i < partB.length ; i++){
        sB += partB[i];
    }
    return Number(sA) + Number(sB);
}

/**
 * Checks whether a number is a palindrome.
 *
 * @param {number} value
 * @returns {boolean} whether the number is a palindrome or not
 */
export function luckyNumber(number){
    let s = String(number);
    let n = s.length;
    for(let i = 0 ; i < n / 2 ; i++){
        if(s[i] != s[n - i - 1]){
            return false;
        }
    }
    return true;
}
/**
 * Determines the error message that should be shown to the user
 * for the given input value.
 *
 * @param {string|null|undefined} input
 * @returns {string} error message
 */
export function errorMessage(userInput){
    let s = Number(userInput)
    if(userInput === null || userInput === undefined || userInput === ''){
        return 'Required field'
    }
    if(!s){
       return 'Must be a number besides 0'
    }
    return ''
}
