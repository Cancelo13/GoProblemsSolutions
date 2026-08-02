export function frontDoorResponse(line){
    return line[0];
}

export function frontDoorPassword(word){
    word = word.toLowerCase()
    return word.charAt(0).toUpperCase() + word.slice(1);
}

export function backDoorResponse(line){
    line = line.trim()
    return line[line.length - 1]
}

export function backDoorPassword(word){
    return frontDoorPassword(word) + ', please'
}