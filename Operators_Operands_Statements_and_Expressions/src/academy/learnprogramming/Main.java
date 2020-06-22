package academy.learnprogramming;

public class Main {

    public static void main(String[] args) {

        int result = 1 + 2; //1 + 2 = 3

        result = result - 1; //3 - 1 = 2

        result = result * 10; //2 * 10 = 20

        result = result / 5; //20 / 5 = 4

//        SAME BUT SHORTER
        int total = 5;

        total++; //5 + 1 = 6
        total += 4; //6 + 4 = 10
        total *= 5; //10 * 5 = 50
        total /= 10; //50 / 10 = 5
        total--; //5 - 1 = 4

//        IF STATEMENT

        boolean isNuts = false;

        if (isNuts) {
            System.out.println("DIS NUTS!");
        }

        if (!isNuts)
            System.out.println("DIS NUTS!");
            System.out.println("Only this line gets printed because we omitted a code block.");

//        LOGICAL "AND" OPERATOR
        int topScore = 120;

        if (topScore < 100) {
            System.out.println("This is below a previous High Score");
        }

        int newTopScore = 146;

        if (newTopScore > topScore && topScore > 100) {
            System.out.println("A New High Score has been established");
        }
//        LOGICAL "OR" OPERATOR
        int someNumber = 90;

        if (someNumber < 100) {
            System.out.println("This is below a previous High Score (duh!)");
        }

        int someNewNumber = 77;

        if (someNewNumber > topScore || someNumber < 100) {
            System.out.println("One or both of the conditions validated to true");
        }

//        TERNARY OPERATOR

        boolean isSeenBefore = false;

        boolean wasSeenBefore = isSeenBefore ? true : false;
        if (!wasSeenBefore) {
            System.out.println("Wow! This is actually something new!");
        }

//        CHALLENGE

        double challengeDoubleOne = 20.00d;
        double challengeDoubleTwo = 80.00d;
        double challengeResult = (challengeDoubleOne + challengeDoubleTwo) * 100.00d;
        double challengeRemainder = challengeResult % 40.00d;
        boolean remainderExists = (challengeRemainder == 0) ? true : false;
        System.out.println(remainderExists);
        if (!remainderExists) {
            System.out.println("Got some remainder");
        }
//        } else {
//            System.out.println("No remainder");
//        }
    }
}
