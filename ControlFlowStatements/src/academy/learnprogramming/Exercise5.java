package academy.learnprogramming;

public class Exercise5 {
    public static int sumFirstAndLastDigit(int number){
        if (number < 0){
            return -1;
        } else if (number < 10){
            return number * 2;
        }
        int lastDigit = number % 10;
        int properDivisor = 1;
        int numLength = (Integer.toString(number)).length();
        int properDivisorLength = 0;
        while (properDivisorLength < numLength){
            properDivisor *= 10;
            properDivisorLength = (Integer.toString(properDivisor)).length();
        }
        int firstDigit = number / properDivisor;
        return firstDigit + lastDigit;
    }
}
