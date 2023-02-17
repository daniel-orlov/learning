package academy.learnprogramming;

public class DigitSumChallenge {

    public static int sumDigits(int number){
        if (number < 10) {
            return -1;
        }
        String numString = Integer.toString(number);
        int sum = 0;
        int digit = 0;

        for (int i = 0; i < numString.length(); i++) {
            digit = number % 10;
            sum += digit;
            number /= 10;
        }

        return sum;
    }
}
