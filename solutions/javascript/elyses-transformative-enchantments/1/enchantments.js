// @ts-check

/**
 * Double every card in the deck.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} deck with every card doubled
 */
export function seeingDouble(cards){
    return cards.map((card) =>{
        return card * 2;
    });
}

/**
 *  Creates triplicates of every 3 found in the deck.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} deck with triplicate 3s
 */
export function threeOfEachThree(cards){
    for(let i = 0; i < cards.length ; i++){
        if(cards[i] === 3){
            cards.splice(i, 0, 3, 3);
            i += 3;
        }
    }
    return cards;
}

/**
 * Extracts the middle two cards from a deck.
 * Assumes a deck is always 10 cards.
 *
 * @param {number[]} deck of 10 cards
 *
 * @returns {number[]} deck with only two middle cards
 */
export function middleTwo(cards){
    return cards.slice(4, 6);
}

/**
 * Moves the outside two cards to the middle.
 *
 * @param {number[]} deck with even number of cards
 *
 * @returns {number[]} transformed deck
 */

export function sandwichTrick(cards){
    let el = Number(cards.splice(0,1));
    let el2 = Number(cards.splice(cards.length - 1,1));
    cards.splice(cards.length / 2, 0, el2, el)
    return cards;
}
/**
 * Removes every card from the deck except 2s.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} deck with only 2s
 */
export function twoIsSpecial(cards){
    return cards.filter((card) =>{
        return card === 2;
    });
}

/**
 * Returns a perfectly order deck from lowest to highest.
 *
 * @param {number[]} deck shuffled deck
 *
 * @returns {number[]} ordered deck
 */
export function perfectlyOrdered(cards){
    return cards.sort((cardA, cardB) =>{
        if(cardA < cardB){
            return -1;
        }
        else if(cardA > cardB){
            return 1;
        }
        return 0;
    });
}
/**
 * Reorders the deck so that the top card ends up at the bottom.
 *
 * @param {number[]} deck
 *
 * @returns {number[]} reordered deck
 */
export function reorder(cards){
    return cards.reverse();
}
