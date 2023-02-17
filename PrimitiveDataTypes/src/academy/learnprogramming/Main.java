package academy.learnprogramming;

public class Main {

    public static void main(String[] args) {

//        INTEGER

        int myValue = 9001;

        int myMinIntValue = Integer.MIN_VALUE;
        int myMaxIntValue = Integer.MAX_VALUE;

        System.out.println("Integer minimum value is: " + myMinIntValue);
        System.out.println("Integer maximum value is: " + myMaxIntValue);

//        Causing Overflow:
        System.out.println("Busting maximum value: " + (myMaxIntValue + 1));

//        Causing Underflow:
        System.out.println("Busting minimum value: " + (myMinIntValue - 1));

//        BYTE

        byte myMinByteValue = Byte.MIN_VALUE;
        byte myMaxByteValue = Byte.MAX_VALUE;

        System.out.println("Byte minimum value is: " + myMinByteValue);
        System.out.println("Byte maximum value is: " + myMaxByteValue);

//        SHORT

        short myMinShortValue = Short.MIN_VALUE;
        short myMaxShortValue = Short.MAX_VALUE;

        System.out.println("Short minimum value is: " + myMinShortValue);
        System.out.println("Short maximum value is: " + myMaxShortValue);

//        LONG

        long myLongValue = 100500L;

        long myMinLongValue = Long.MIN_VALUE;
        long myMaxLongValue = Long.MAX_VALUE;

        System.out.println("Long minimum value is: " + myMinLongValue);
        System.out.println("Long maximum value is: " + myMaxLongValue);

//        PRACTICE: CASTING

        int myIntValue = (myMaxIntValue / 47);

        byte myByteValue = (byte) (myMinByteValue / 47);

//        CHALLENGE

        byte challengeByteVar = 9;
        short challengeShortVar = 12;
        int challengeIntVar = 314;

        long challengeLongResult = 50000L + 10L *
                (challengeByteVar + challengeIntVar + challengeShortVar);

        System.out.println(challengeLongResult);

//        FLOAT AND DOUBLE

        float myMinFloatValue = Float.MIN_VALUE;
        float myMaxFloatValue = Float.MAX_VALUE;

        System.out.println("Float minimum value is: " + myMinFloatValue);
        System.out.println("Float maximum value is: " + myMaxFloatValue);

        double myMinDoubleValue = Double.MIN_VALUE;
        double myMaxDoubleValue = Double.MAX_VALUE;

        System.out.println("Double minimum value is: " + myMinDoubleValue);
        System.out.println("Double maximum value is: " + myMaxDoubleValue);

        int myIntVar = 13 / 5;
        float myFloatVar = 13.85f / 5f;
        double myDoubleVar = -13.456d / -5d; //double is better

        System.out.println(myIntVar);
        System.out.println(myFloatVar);
        System.out.println(myDoubleVar);

//        CHALLENGE: POUND TO KILOS

        double onePound = 0.4359237;
        double weightInKg = 70.3;

        double weightInLbs = weightInKg / onePound;

        System.out.println(weightInKg + " in kgs is " + weightInLbs + " in lbs");


//        CHAR

        char myChar = 'D';
        char myUnicodeChar = '\u0147';

        System.out.println(myChar);
        System.out.println(myUnicodeChar);



//        BOOLEAN

        boolean myTruly = true;
        boolean myFalsy = false;

        boolean isBooleanTrue = true;


//        STRING (NOT A PRIMITIVE TYPE, BUT ANYWAY)

        String myString = "Dan";

        String myNumericString = "123.456";
        String myOtherNumericString = "456.789";

        String myStupidAttemptToSumTwoString = myNumericString + myOtherNumericString;

        System.out.println("And this is how you concatenate: " + myStupidAttemptToSumTwoString);

    }
}
