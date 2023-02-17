package academy.learnprogramming;

public class StringToNums {
    public static void main(String[] args) {
        String numberAsString = "2020";
        System.out.println("Number As String = " + numberAsString);

        int number = Integer.parseInt(numberAsString);
        System.out.println("Number = " + number);

        numberAsString += 1;
        number += 1;

        System.out.println("Number As String + 1 = " + numberAsString);
        System.out.println("Number + 1 = " + number);
    }

}
