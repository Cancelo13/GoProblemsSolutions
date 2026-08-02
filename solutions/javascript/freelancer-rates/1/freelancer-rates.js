// @ts-check


const HOURS_PER_DAY = 8;
const BILLED_DAYS_PER_MONTH = 22;

/**
 * The day rate, given a rate per hour
 *
 * @param {number} ratePerHour
 * @returns {number} the rate per day
 */

export function dayRate(ratePerHour){
    return HOURS_PER_DAY * ratePerHour
}

/**
 * Calculates the number of days in a budget, rounded down
 *
 * @param {number} budget: the total budget
 * @param {number} ratePerHour: the rate per hour
 * @returns {number} the number of days
 */
export function daysInBudget(budget, ratePerHour){
    return Math.floor(budget / dayRate(ratePerHour))
}

/**
 * Calculates the discounted rate for large projects, rounded up
 *
 * @param {number} ratePerHour
 * @param {number} numDays: number of days the project spans
 * @param {number} discount: for example 20% written as 0.2
 * @returns {number} the rounded up discounted rate
 */
export function priceWithMonthlyDiscount(ratePerHour, days, discountRate){
    return Math.ceil((Math.floor(days/BILLED_DAYS_PER_MONTH) * BILLED_DAYS_PER_MONTH * dayRate(ratePerHour) * (1 - discountRate)) + (days % BILLED_DAYS_PER_MONTH) * dayRate(ratePerHour))
}

