package academy.learnprogramming;

public class Main {

    public static void main(String[] args) {

        boolean gameOver = true;
        int currentScore = 500;
        int currentLevel = 14;
        int finalBonus = 100;


        if (currentScore == 500) {
            System.out.println("Your score is " + currentScore);
        } else {
            System.out.println("Your code is below 500");
        }

//        CHALLENGE
        currentScore = 10000;
        currentLevel = 8;
        finalBonus = 200;

        if (gameOver) {
            int finalScore = currentScore + (currentLevel * finalBonus);
            System.out.println("Final score: " + finalScore);
        }

//        REFACTOR

//        printFinalScore(gameOver, currentScore, currentLevel, finalBonus);

        printFinalScore(true, 14568, 76, 150);

        int finalScore2;
        finalScore2 = getFinalScore(true, 654, 5, 15);
        System.out.println("Final score: " + finalScore2);


//        CHALLENGE 2

        int player1Position = calculateHighScorePosition(1500);
        int player2Position = calculateHighScorePosition(900);
        int player3Position = calculateHighScorePosition(400);
        int player4Position = calculateHighScorePosition(50);

        displayHighScorePosition("Player1", player1Position);
        displayHighScorePosition("Player2", player2Position);
        displayHighScorePosition("Player3", player3Position);
        displayHighScorePosition("Player4", player4Position);






    }

//        METHODS
    public static void printFinalScore(boolean gameOver, int currentScore, int currentLevel, int finalBonus) {
        if (gameOver) {
            int finalScore = currentScore + (currentLevel * finalBonus);
            System.out.println("Final score: " + finalScore);
        }
    }

    public static int getFinalScore(boolean gameOver, int currentScore, int currentLevel, int finalBonus) {
        int finalScore = currentScore;
        if (gameOver) {
            finalScore = currentScore + (currentLevel * finalBonus);
        }
        return finalScore;
    }

    public static void displayHighScorePosition(String playerName, int position) {
        System.out.println(playerName + " managed to get into position " +
                position + " on the high-score table.");
    }

    public static int calculateHighScorePosition(int playerScore) {
        if (playerScore >= 1000) {
            return 1;
        } else if (playerScore >= 500) {
            return 2;
        } else if (playerScore >= 100) {
            return 3;
        }
        return 4;
    }

}

