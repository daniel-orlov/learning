package academy.learnprogramming;

public class Exercise4 {
    public static boolean isLeapYear(int year) {
        if ((year >= 1 && year <= 9999) && ((year % 400 == 0 ) || (year % 4 == 0 && year % 100 != 0))) {
            return true;
        }
        return false;
    }
}
