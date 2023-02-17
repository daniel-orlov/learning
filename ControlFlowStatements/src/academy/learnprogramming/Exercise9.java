package academy.learnprogramming;

public class Exercise9 {
    public static int getGreatestCommonDivisor(int first, int second){
        if (first < 10 || second < 10) return -1;

        int minNum = Math.min(first, second);
        int maxNum = Math.max(first, second);

        int counter = minNum;

        while (counter > 0){
            if (maxNum % counter == 0 && minNum % counter == 0){
                return counter;
            }
            counter--;
        }
        return 1;
    }
}
