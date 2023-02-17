package academy.learnprogramming;

public class Exercise3 {
    public static void printEqual(int first, int second, int third) {
        if (first < 0 || second < 0 || third < 0) {
            System.out.println("Invalid Value");
        } else if (first == second && second == third){
            System.out.println("All numbers are equal");
        } else if (first != second && second != third && third != first) {
            System.out.println("All numbers are different");
        } else {
            System.out.println("Neither all are equal or different");
        }
    }
}
