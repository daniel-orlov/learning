function checkGuestIsAllowed(guests, guest) {
    // Log guests array
    console.log(guests);

    // Log guest
    console.log(guest);

    let result = guests.includes(guest);
    console.log(result);

    return result;
}

/*
checkGuestIsAllowed(['John', 'Marta', 'Mike'], 'Marta'); // true
checkGuestIsAllowed(['John', 'Marta', 'Mike'], 'Bob'); // false
 */

function interactiveGuestCheck(guests) {
    let guest = prompt('Enter your name');
    let result = checkGuestIsAllowed(guests, guest);
    alert(result);
}

interactiveGuestCheck(['John', 'Marta', 'Mike']);