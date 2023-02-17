package academy.learnprogramming;

public class Switch {

    public static void printSwitchValue(int switchValue) {
        switch(switchValue) {
            case 1:
                System.out.println("One");
                break;

            case 2:
                System.out.println("Two");
                break;

            case 3: case 4: case 5:
                System.out.println("Three, Four or Five");
                System.out.println("Actually it was " + switchValue);
                break;

            default:
                System.out.println("Forty two");
                break;
        }
    }

    public static void printLastPressedButton(char lastPressedButton) {

        switch (lastPressedButton) {
            case 'A':
                System.out.println("Aplha");
                break;

            case 'B':
                System.out.println("Bravo");
                break;

            case 'C':
                System.out.println("Charlie");
                break;

            case 'D':
                System.out.println("Delta");
                break;

            case 'E':
                System.out.println("Echo");
                break;

            default:
                System.out.println("Letter not found, over.");
                break;
        }
    }

    public static void printDayOfTheWeek(int dayNum){

        switch(dayNum){
            case 0:
                System.out.println("Monday");
                break;
            case 1:
                System.out.println("Tuesday");
                break;
            case 2:
                System.out.println("Wednesday");
                break;
            case 3:
                System.out.println("Thursday");
                break;
            case 4:
                System.out.println("Friday");
                break;
            case 5:
                System.out.println("Saturday");
                break;
            case 6:
                System.out.println("Sunday");
                break;
            default:
                System.out.println("Invalid day");
                break;
        }
    }
}
