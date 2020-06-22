package academy.learnprogramming;

public class Challenge {

    public static String getDurationString(long rawMinutes, long seconds) {
        if (rawMinutes < 0 || seconds < 0 || seconds > 59) {
            return "Invalid value";
        }
        long hours = rawMinutes / 60L;
        long minutes = rawMinutes % 60L;

        String finalResult = "";

        if (hours < 10) {
            finalResult += "0" + hours + "h ";
        } else {
            finalResult += hours + "h ";
        }

        if (minutes < 10) {
            finalResult += "0" + minutes + "m ";
        } else {
            finalResult += minutes + "m ";
        }

        if (seconds < 10) {
            finalResult += "0" + seconds + "s";
        } else {
            finalResult += seconds + "s";
        }

        return finalResult;
    }

    public static String getDurationString(long rawSeconds) {
        if (rawSeconds < 0) {
            return "Invalid value";
        }
        long rawMinutes = rawSeconds / 60L;
        long hours = rawMinutes / 60L;
        long minutes = rawMinutes % 60L;
        long seconds = rawSeconds % 60L;

        String finalResult = "";

        if (hours < 10) {
            finalResult += "0" + hours + "h ";
        } else {
            finalResult += hours + "h ";
        }

        if (minutes < 10) {
            finalResult += "0" + minutes + "m ";
        } else {
            finalResult += minutes + "m ";
        }

        if (seconds < 10) {
            finalResult += "0" + seconds + "s";
        } else {
            finalResult += seconds + "s";
        }

        return finalResult;
    }
}
