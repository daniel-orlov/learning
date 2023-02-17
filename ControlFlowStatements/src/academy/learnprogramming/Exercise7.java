package academy.learnprogramming;

public class Exercise7 {
    public static boolean hasSharedDigit(int x, int y){
        if (x < 10 || y < 10 || x > 99 || y > 99) return false;
        int firstDigitX, lastDigitX, firstDigitY, lastDigitY;

        firstDigitX = x % 10;
        lastDigitX = x / 10;
        firstDigitY = y % 10;
        lastDigitY = y / 10;

        return firstDigitX == firstDigitY || firstDigitX == lastDigitY || lastDigitX == firstDigitY || lastDigitX == lastDigitY;
    }
}
