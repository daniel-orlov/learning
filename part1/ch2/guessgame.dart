import 'dart:io';
import 'dart:math';

void main() {
  print('Let\'s play!');
  playGuessGame();

  print('Thanks for playing!');
  /// Ask the user if they want to play again
  print('Do you want to play again? (y/n)');
  String? answer = stdin.readLineSync();
  if (answer == 'y') {
    main();
  }
}

void playGuessGame() {
  /// Ask a user how many tries they want
  print('How many tries do you want?');
  String? input = stdin.readLineSync();
  int? tries = int.tryParse(input!);


  print('Okay, let\'s guess the number from 1 to 100 in $tries tries');
  print('Enter your guess.');

  /// Pick a random number between 1 and 100
  var number = new Random().nextInt(100) + 1;

  /// Read the user's input from the command line four times
  for (var i = 0; i < tries!; i++) {
    var guessString = stdin.readLineSync();
    var guess = int.parse(guessString!);

    if (guess == number) {
      print('You win!');
      return;
    }

    print('Try another time...');
    /// Tell the user if the guess was too high or too low
    if (guess > number) {
      print('Your guess was too high');
    } else {
      print('Your guess was too low');
    }
  }

  /// The user has run out of tries, tell them the number
  print('You almost got it! The number was $number');
}