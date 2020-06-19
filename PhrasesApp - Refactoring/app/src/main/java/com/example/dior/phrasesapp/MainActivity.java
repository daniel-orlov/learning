package com.example.dior.phrasesapp;

import android.media.MediaPlayer;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;

import static android.support.v4.os.LocaleListCompat.create;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }



    public void onClick(View view) {
        Button phrase = (Button) view;

        String tappedButton = phrase.getTag().toString();

        MediaPlayer mediaPlayer = MediaPlayer.create(this, getResources().getIdentifier(tappedButton, "raw", getPackageName()));

        mediaPlayer.start();
    }
}
