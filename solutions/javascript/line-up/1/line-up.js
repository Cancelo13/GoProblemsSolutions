export function format(name, number){
    console.log(number % 10)
    let rank = "th"
    if (number % 10 == 1 && number % 100 != 11){
        rank = "st"
    }
    else if(number % 10 == 2 && number % 100 != 12){
        rank = "nd"
    }
    else if(number % 10 == 3 && number % 100 != 13){
        rank = "rd"
    }
    return name + ", you are the " + number + rank + " customer we serve today. Thank you!"
}