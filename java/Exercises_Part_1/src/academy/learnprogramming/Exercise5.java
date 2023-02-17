package academy.learnprogramming;

public class Exercise5 {
    public static boolean areEqualByThreeDecimalPlaces(double first, double second) {
        double result = first - second;
        if (Math.abs(result) <= 0.0009) {
            return true;
        }
        return false;
    }
}
