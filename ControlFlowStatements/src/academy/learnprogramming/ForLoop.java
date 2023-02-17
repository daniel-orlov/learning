package academy.learnprogramming;

public class ForLoop {

    public static double calculateInterest(double amount, double interestRate) {
        return (amount * (interestRate/100));
    }

    public static boolean isPrime(int n){

        if(n == 1) {
            return false;
        }

        for (int i=2; i <= n/2; i++) {
            if (n % i == 0) {
                return false;
            }
        }

        return true;
    }

    public static boolean isDivisibleByThreeAndFive(int num){
        if (num % (3*5) == 0){
            return true;
        }
        return false;
    }


}
