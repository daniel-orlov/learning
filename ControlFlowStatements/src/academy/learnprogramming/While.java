package academy.learnprogramming;

public class While {

    public static boolean isEvenNumber(int number) {
        if (number <= 0) {
            return false;
        }
        return number % 2 == 0;
    }
}
