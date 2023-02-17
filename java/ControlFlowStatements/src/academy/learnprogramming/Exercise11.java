package academy.learnprogramming;

public class Exercise11 {
    public static boolean isPerfectNumber(int number){
        if (number < 1) return false;

        int sum = 0;
        int counter = 1;
        while (counter < number) {
            if (number % counter == 0) {
                sum += counter;
            }
            counter++;
        }
        return sum == number;
    }
}
