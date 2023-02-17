package academy.learnprogramming;

public class Exercise13 {
    public static boolean canPack(int bigCount, int smallCount, int goal){
//        bigCount = 5 kg, smallCount = 1 kg
        if (bigCount < 0 || smallCount < 0 || goal < 0) {
            return false;
        } else if (goal > (bigCount*5 + smallCount)){
            return false;
        }
        for(int i = 0; i < bigCount; i++){
            if ((goal - 5) >= 0){
                goal -= 5;
            }
        }

        return (goal - smallCount <= 0);
    }
}
