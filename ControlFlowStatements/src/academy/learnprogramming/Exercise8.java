package academy.learnprogramming;

public class Exercise8 {
    public static boolean hasSameLastDigit(int a, int b, int c){
        if (isValid(a) && isValid(b) && isValid(c)){
            int rmA, rmB, rmC;
    //        rm = rightmost
            rmA = a % 10;
            rmB = b % 10;
            rmC = c % 10;
            return rmA == rmB || rmA == rmC || rmB == rmC;
        }
        return false;
    }

    public static boolean isValid(int a){
        return a >= 10 && a <= 1000;
    }

}
