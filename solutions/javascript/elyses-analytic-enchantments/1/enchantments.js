/**
 * Get the position (index) of the card in the given stack
 *
 * @param {number[]} stack
 * @param {number} card
 *
 * @returns {number} position of the card in the stack
 */
export function getCardPosition(cards, card){
    return cards.indexOf(card);
}

/**
 * Determine if the stack contains the card
 *
 * @param {number[]} stack
 * @param {number} card
 *
 * @returns {boolean} true if card is in the stack, false otherwise
 */
export function doesStackIncludeCard(cards, card){
    return cards.includes(card);
}

/**
 * Determine if each card is even
 *
 * @param {number[]} stack
 *
 * @returns {boolean} true if all cards are even, false otherwise
 */
export function isEachCardEven(cards){
    return cards.every((card) =>{
        return card % 2 === 0;
    });
}

/**
 * Check if stack contains odd-value card
 *
 * @param {number[]} stack
 *
 * @returns {boolean} true if the array contains odd card, false otherwise
 */
export function doesStackIncludeOddCard(cards){
    return cards.some((card) =>{
        return card % 2 !== 0;
    })
};

/**
 * Get the first odd card from the stack
 *
 * @param {number[]} stack
 *
 * @returns {number | undefined} the first odd value
 */
export function getFirstOddCard(cards){
    return cards.find((card) =>{
        return card % 2 !== 0;
    });
}

/**
 * Determine the position of the first card that is even
 *
 * @param {number[]} stack
 *
 * @returns {number} position of the first card that is even
 */
export function getFirstEvenCardPosition(cards){
    return cards.findIndex((card) => {
        return card % 2 == 0;
    });
}
