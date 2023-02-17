package academy.learnprogramming;

public class Exercise7 {
    public static boolean hasTeen(int first, int second, int third) {
        return isTeen(first) || isTeen(second) || isTeen(third);
    }
    public static boolean isTeen(int num) {
        return num >= 13 && num <= 19;
    }
}
