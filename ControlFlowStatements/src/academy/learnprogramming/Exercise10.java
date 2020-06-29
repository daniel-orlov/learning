package academy.learnprogramming;

public class Exercise10 {
    public static void printFactors(int number) {
        if (number < 1) System.out.println("Invalid Value");

        int counter = number;
        while (counter > 0) {
            if (number % counter == 0) {
                System.out.println(counter);
            }
            counter--;
        }
    }
}
