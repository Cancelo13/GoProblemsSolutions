export function getItem(cards, pos){
    return cards[pos];
};

export function setItem(cards, pos, newCard){
    cards[pos] = newCard;
    return cards;
};

export function insertItemAtTop(cards, newCard){
    cards.push(newCard)
    return cards;
};

export function removeItem(cards, pos){
    cards.splice(pos, 1);
    return cards;
};

export function removeItemFromTop(cards){
    cards.pop();
    return cards
};

export function insertItemAtBottom(cards, newCard){
    cards.unshift(newCard)
    return cards;
};

export function removeItemAtBottom(cards){
    cards.shift()
    return cards;
};

export function checkSizeOfStack(cards, stackSize){
    return cards.length == stackSize
};