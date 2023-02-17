// Get random number between 1 and 6
function rollDice() {
    return Math.floor(Math.random() * 6) + 1;
}

// Pick random image for the left dice and display it
let leftDice = rollDice();
document.querySelector(".img1").setAttribute("src", "images/dice" + leftDice + ".png");

// Pick random image for the right dice and display it
let rightDice = rollDice();
document.querySelector(".img2").setAttribute("src", "images/dice" + rightDice + ".png");


// Display the winner
if (leftDice > rightDice) {
    document.querySelector("h1").textContent = "🚩 Player 1 Wins!";
} else if (leftDice < rightDice) {
    document.querySelector("h1").textContent = "Player 2 Wins! 🚩";
} else {
    document.querySelector("h1").textContent = "Draw!";
}