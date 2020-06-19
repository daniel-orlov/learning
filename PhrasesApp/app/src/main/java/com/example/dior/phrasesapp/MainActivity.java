package com.example.dior.phrasesapp;

import android.media.MediaPlayer;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }



    public void onClick(View view) {
        MediaPlayer mP;
        Button phrase = (Button) view;

        int tappedButton = Integer.parseInt(phrase.getTag().toString());

        switch (tappedButton) {
            case 0: mP = MediaPlayer.create(this, R.raw.hello);
                mP.start();
                break;
            case 1: mP = MediaPlayer.create(this, R.raw.how_are_you);
                mP.start();
                break;
            case 2: mP = MediaPlayer.create(this, R.raw.i_speak_english);
                mP.start();
                break;
            case 3: mP = MediaPlayer.create(this, R.raw.i_am_from);
                mP.start();
                break;
            case 4: mP = MediaPlayer.create(this, R.raw.please);
                mP.start();
                break;
            case 5: mP = MediaPlayer.create(this, R.raw.my_name_is);
                mP.start();
                break;
            case 6: mP = MediaPlayer.create(this, R.raw.i_love_you);
                mP.start();
                break;
            case 7: mP = MediaPlayer.create(this, R.raw.thank_you_bye);
                mP.start();
                break;
        }
    }
}
