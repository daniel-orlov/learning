package com.example.dior.eggtimer;

import android.media.MediaPlayer;
import android.os.CountDownTimer;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.SeekBar;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    TextView timerTextView;
    SeekBar timerSeekBar;
    CountDownTimer countDownTimer;
    Boolean counterIsActive = false;
    Button startButton;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        startButton = findViewById(R.id.startButton);

        timerSeekBar = findViewById(R.id.timerSeekBar);
        timerTextView = findViewById( R.id.timerTextView);

        int max = 601;
        int start = 30;
        int min = 1;

        timerSeekBar.setMax(max);
        timerSeekBar.setProgress(start);
        timerSeekBar.setMin(min);

        timerSeekBar.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int i, boolean b) {
                displayTime(i);
                Log.i("Info", Integer.toString(i));
            }

            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {

            }

            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {

            }
        });
    }


    public void displayTime(int timeInSeconds) {
        int minuteInt = timeInSeconds / 60;
        int secondInt = timeInSeconds % 60;

        String minuteString = Integer.toString(minuteInt);
        String secondString = Integer.toString(secondInt);

        if (secondInt <= 9){
            secondString = "0" + secondString;
        }

        timerTextView.setText(minuteString + ":" + secondString);
    }

    public void resetTimer() {
        timerTextView.setText("1:30");
        timerSeekBar.setProgress(90);
        timerSeekBar.setEnabled(true);
        countDownTimer.cancel();
        startButton.setText("Start");
        counterIsActive = false;
    }

    public void startButtonClicked(View view) {

        if (counterIsActive) {

            resetTimer();

        } else {

            counterIsActive = true;
            timerSeekBar.setEnabled(false);

            startButton.setText("Stop");

            countDownTimer = new CountDownTimer(timerSeekBar.getProgress() * 1000 + 100, 1000) {
                @Override
                public void onTick(long milliSecondsLeftLong) {
                    displayTime((int) milliSecondsLeftLong / 1000);
                }

                @Override
                public void onFinish() {
                    MediaPlayer mediaPlayer = MediaPlayer.create(getApplicationContext(), R.raw.foghorn_daniel_simon);
                    mediaPlayer.start();
                    resetTimer();
                }
            }.start();

            Log.i("Info", "Button Clicked");
        }
    }
}
