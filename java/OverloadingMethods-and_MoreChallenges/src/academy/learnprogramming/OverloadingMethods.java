package academy.learnprogramming;

public class OverloadingMethods {
    public static double calcFeetAndInchesToCentimeters(double feet, double inches) {
        if (feet < 0 || inches < 0 || inches > 12) {
            return -1d;
        }
        double allInInches = inches + 12d * feet;
        return allInInches * 2.54;
    }

    public static double calcFeetAndInchesToCentimeters(double inches) {
        if (inches < 0) {
            System.out.println("Invalid parameters");
            return -1d;
        }
        return inches * 2.54;
    }
}
