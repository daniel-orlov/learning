package academy.learnprogramming;

public class Exercise12 {
    public static void numberToWords(int number){
        if (number < 0) {
            System.out.println("Invalid Value");
            return;
        }

        int reversedNum = reverse(number);
        int originalLength = getDigitCount(number);
        int reversedLength = getDigitCount(reversedNum);

        String finalString = "";
        int digit = 0;

        for (int i=0; i < originalLength; i++) {
            digit = reversedNum % 10;
            reversedNum /= 10;
            switch (digit) {
                case 0:
                    finalString += "Zero ";
                    break;
                case 1:
                    finalString += "One ";
                    break;
                case 2:
                    finalString += "Two ";
                    break;
                case 3:
                    finalString += "Three ";
                    break;
                case 4:
                    finalString += "Four ";
                    break;
                case 5:
                    finalString += "Five ";
                    break;
                case 6:
                    finalString += "Six ";
                    break;
                case 7:
                    finalString += "Seven ";
                    break;
                case 8:
                    finalString += "Eight ";
                    break;
                case 9:
                    finalString += "Nine ";
                    break;
            }
        }
        System.out.println(finalString);
    }

    public static int getDigitCount(int number){
        if (number < 0) return -1;
        String numString = Integer.toString(number);
        return numString.length();
    }

    public static int reverse(int number){
        int absNumber = Math.abs(number);
        if (absNumber < 10) return number;

        int tempNumber = absNumber;

        int reverse = 0;
        int digit;

        while (true){
            digit = tempNumber % 10;
            reverse += digit;
            reverse *= 10;
            tempNumber /= 10;
//            System.out.println("Current state of reverse: " + reverse);
//            System.out.println("Current state of tempNumber: " + tempNumber);
            if (tempNumber < 10) {
                reverse += tempNumber;
                break;
            }
        }

//        System.out.println(reverse);

        if (number < 0) {
            System.out.println(reverse);
            return reverse * (-1);
        }
        return reverse;
    }
}
