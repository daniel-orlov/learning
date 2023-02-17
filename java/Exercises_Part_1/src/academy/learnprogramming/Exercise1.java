package academy.learnprogramming;

public class Exercise1 {
    public static long toMilesPerHour(double kilometersPerHour) {
        long result = -1L;
        double kilometersInMile = 1.609;

        if (kilometersPerHour >= 0) {
            double preResult = kilometersPerHour / kilometersInMile;
            result = Math.round(preResult);
            return result;
        }
        return result;
    }
    public static void printConversion(double kilometersPerHour) {
        long milesInKilometers = toMilesPerHour(kilometersPerHour);
        if (milesInKilometers != -1) {
            System.out.println(kilometersPerHour + " km/h = " +
                    milesInKilometers + " mi/h");
        } else {
            System.out.println("Invalid Value");
        }
    }
}
