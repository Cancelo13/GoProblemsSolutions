/// <reference path="./global.d.ts" />
// @ts-check

export function cookingStatus(timer){
    if (timer === undefined){
        return 'You forgot to set the timer.'
    }
    return timer ? 'Not done, please wait.' : 'Lasagna is done.';
}

export function preparationTime(layers, time = 2){
    return layers.length * time;
}

export function quantities(layers){
    let noodles = 0;
    let sauce = 0;
    for(let i = 0 ; i < layers.length ; i++){
        if(layers[i] == 'noodles'){
            noodles += 50;
        }
        else if(layers[i] == 'sauce'){
            sauce += 0.2;
        }
    }
    return {
        noodles: noodles,
        sauce: sauce,
    }
}

export function addSecretIngredient(friendList, list){
    list.push(friendList[friendList.length - 1])
}

export function scaleRecipe(receipe, number){
    let newReceipe = {}
    for(let i in receipe){
        newReceipe[i] = receipe[i] * number / 2;
    }
    return newReceipe;
}