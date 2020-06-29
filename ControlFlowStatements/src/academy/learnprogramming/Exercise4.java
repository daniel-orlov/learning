package academy.learnprogramming;

public class Exercise4 {

    public static boolean isPalindrome(int number){
        number = Math.abs(number);

        int tempNumber = number;

        int reverse = 0;
        int digit = 0;

        while (tempNumber > 0){
            digit = tempNumber % 10;
            reverse += digit;
            reverse *= 10;
            tempNumber /= 10;
            System.out.println("Current state of reverse: " + reverse);
            System.out.println("Current state of tempNumber: " + tempNumber);
            if (tempNumber < 10) {
                reverse += tempNumber;
                break;
            }
        }
        return reverse == number;
    }
}
