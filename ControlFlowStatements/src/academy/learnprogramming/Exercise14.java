package academy.learnprogramming;

import static academy.learnprogramming.ForLoop.isPrime;

public class Exercise14 {

    public static int getLargestPrime(int number){
        int errorReturn = -1;

        if (number < 2) return errorReturn;

        boolean isPrime = true;

        for (int i = number; i > 2; i--){

            for (int j = 2; j <= i/2; j++){
                if (i % j == 0){
                    isPrime = false;
                    break;
                }
            }

            if (isPrime && number % i == 0){
                return i;
            }
        }
        return errorReturn;
    }
}
