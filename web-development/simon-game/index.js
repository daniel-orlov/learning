// Global variables
let gameIsOn = false;
let level = 0;
let gamePattern = [];
let userClickedPattern = [];

// ----------------- Main page triggers -----------------
// Game starts when any key is pressed
$(document).keypress(function () {
    checkGameState();
});

// Select the button
$(".btn").click(function () {
    handleUserClick($(this).attr("id"));
});


// --------------------- Game state ---------------------
// Restart the game
function startOver() {
    level = 0;
    gamePattern = [];
    gameIsOn = false;
}

// Check game state
function checkGameState() {
    if (!gameIsOn) {
        addNextColourToSequence();
        gameIsOn = true;
    }
}


// ----------------- Game logic -----------------
const buttonColours = ["red", "blue", "green", "yellow"];

// Pick a random colour.
function pickNextColour() {
    let randomNumber = Math.floor(Math.random() * 4);
    return buttonColours[randomNumber];
}

// Add the next colour to the game pattern.
function addNextColourToSequence() {

    level++;
    updateLevelTitle();

    let randomChosenColour = pickNextColour();

    gamePattern.push(randomChosenColour);

    flashButton(randomChosenColour);
    playSound(randomChosenColour);
}

// Check the answer
function checkAnswer(currentLevel) {
    console.log("checkAnswer: " + currentLevel);
    console.log("gamePattern: " + gamePattern[currentLevel]);
    console.log("userClickedPattern: " + userClickedPattern[currentLevel]);

    if (gamePattern[currentLevel] === userClickedPattern[currentLevel]) {
        console.log("success");

        if (userClickedPattern.length === gamePattern.length) {
            setTimeout(function () {
                wipeUserClickedPattern();
                addNextColourToSequence();
            }, 1000);
        }
    } else {
        console.log("failure");

        handleWrongAnswer();
    }
}


// ----------------- User input -----------------
// Handle the user clicking a button
function handleUserClick(userChosenColour) {
    console.log(userChosenColour);

    userClickedPattern.push(userChosenColour);

    playSound(userChosenColour);
    flashButton(userChosenColour);

    console.log(userClickedPattern);

    checkAnswer(userClickedPattern.length - 1);
}

// Handle wrong answer
function handleWrongAnswer() {
    // Play the wrong sound
    playSound("wrong");

    // Add the game-over class
    $("body").addClass("game-over");

    // Change the title
    $("#level-title").text("Game Over, Press Any Key to Restart");

    // Remove the game-over class after 200ms
    setTimeout(function () {
            $("body").removeClass("game-over");
        }, 200);

    // Restart the game
    startOver();
}

// Wipe the user clicked pattern
function wipeUserClickedPattern() {
    userClickedPattern = [];
}


// ----------------- UI -----------------
// Flash the button
function flashButton(currentColour) {
    // Add the pressed class
    $("#" + currentColour).addClass("pressed");

    // Remove the pressed class after 100ms
    setTimeout(function () {
        $("#" + currentColour).removeClass("pressed");
    }, 100);
}

// Play the sound
function playSound(name) {
    let audio = new Audio("sounds/" + name + ".mp3");
    audio.play();
}

// Update the level title
function updateLevelTitle() {
    $("#level-title").text("Level " + level);
}