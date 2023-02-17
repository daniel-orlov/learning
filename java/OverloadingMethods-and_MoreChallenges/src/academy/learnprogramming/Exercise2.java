package academy.learnprogramming;

public class Exercise2 {
    public static void printYearsAndDays(long minutes) {
        if (minutes < 0) {
            System.out.println("Invalid Value");
        } else {
            long rawDays = minutes / (60L * 24L);
            long years = rawDays / 365L;
            long days = rawDays % 365L;
            System.out.println(minutes + " min = " + years + " y and " + days + " d");
        }
    }
}
