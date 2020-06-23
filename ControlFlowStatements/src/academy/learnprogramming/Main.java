package academy.learnprogramming;

import static academy.learnprogramming.DigitSumChallenge.sumDigits;
import static academy.learnprogramming.ForLoop.*;
import static academy.learnprogramming.Switch.*;
import static academy.learnprogramming.While.isEvenNumber;

public class Main {



    public static void main(String[] args) {
	// SWITCH

//        int switchValue = 4;
//        printSwitchValue(switchValue);
//
//        char lastPressedButton = 'D';
//        printLastPressedButton(lastPressedButton);
//
//        int dayNum = 5;
//        printDayOfTheWeek(dayNum);

    // FOR LOOP for(int; termination; increment){}

//        double amount = 12000d;
//
//        for (int i=2; i<9; i++){
//            System.out.println(amount + " at " + i + "% interest = " + calculateInterest(12000, i));
//        }
//
//        System.out.println("__________________________________________");
//
//        for (int i=8; i>1; i--){
//            System.out.println(amount + " at " + i + "% interest = " + calculateInterest(12000, i));
//        }
//
//        CHALLENGE

//        int primeNumbersFound = 0;
//
//        for(int i=1; i < 1000; i++){
//            if (isPrime(i)) {
//                System.out.println("Prime found: " + i);
//                primeNumbersFound++;
//            } else if (primeNumbersFound == 12) {
//                break;
//            }
//        }
//
//        int sum = 0;
//        int count = 0;
//
//        for(int i=1; i <= 1000; i++){
//            if (isDivisibleByThreeAndFive(i)){
//                System.out.println("This number is divisible by both 3 and 5: " + i);
//                sum += i;
//                count++;
//            } else if (count == 5) {
//                System.out.println("The sum of  the first " + count + " numbers, divisible by 3 & 5 = " + sum);
//                break;
//            }
//        }

        // WHILE STATEMENT
//        int count = 0;
//        while (count <= 7){
//            count++;
//            System.out.println("Count is " + count);
//
//        }
//
//        System.out.println("_________________________________");
//
//        do {
//            System.out.println("Count is " + count);
//            count--;
//        } while (count > 0);

        int startNumber = 7;
        int endNumber = 23;

//        while (startNumber <= endNumber){
//            startNumber++;
//            if(!isEvenNumber(startNumber)){
//                continue;
//            }
//            System.out.println("This is even number: " + startNumber);
//        }

//        int evenNumbersFound = 0;
//
//        while (startNumber <= endNumber && evenNumbersFound < 5){
//            startNumber++;
//            if(!isEvenNumber(startNumber)){
//                continue;
//            }
//            System.out.println("This is even number: " + startNumber);
//            evenNumbersFound++;
//        }
//        System.out.println("Total even numbers found: " + evenNumbersFound);


//        DIGIT SUM CHALLENGE
//        System.out.println(sumDigits(11));
//        System.out.println("___________________");
        System.out.println(sumDigits(123));
        System.out.println("___________________");
//        System.out.println(sumDigits(1234));
//        System.out.println("___________________");
//        System.out.println(sumDigits(12345));
    }
}
