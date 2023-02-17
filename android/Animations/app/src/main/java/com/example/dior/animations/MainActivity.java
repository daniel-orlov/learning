package com.example.dior.animations;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.ImageView;

public class MainActivity extends AppCompatActivity {

    boolean flagSunOut = true;

//    public void fade(View view) {
//        Log.i("Info", "Image tapped");
//
//        ImageView sunImageView = findViewById(R.id.sunImageView);
//
//        ImageView moonImageView = findViewById(R.id.moonImageView);
//
//        if (flagSunOut) {
//            sunImageView.animate().alpha(0).setDuration(2250);
//
//            moonImageView.animate().alpha(1).setDuration(2250);
//
//            flagSunOut = false;
//        } else {
//            sunImageView.animate().alpha(1).setDuration(2250);
//
//            moonImageView.animate().alpha(0).setDuration(2250);
//
//            flagSunOut = true;
//        }
//    }

    public void translate(View view) {
        Log.i("Info", "Image tapped");

        ImageView sunImageView = findViewById(R.id.sunImageView);

        ImageView moonImageView = findViewById(R.id.moonImageView);


        if (flagSunOut) {
            sunImageView.animate().translationYBy(2000).alpha(0).setDuration(3000);

            moonImageView.animate().alpha(1).setDuration(3000);

            flagSunOut = false;
        } else {
            moonImageView.animate().alpha(0).setDuration(3000);

            sunImageView.animate().translationYBy(-2000).alpha(1).rotation(270).setDuration(3000);

            flagSunOut = true;
        }
    }

//    public void translate(View view) {
//        Log.i("Info", "Image tapped");
//
//        ImageView sunImageView = findViewById(R.id.sunImageView);
//
//        ImageView moonImageView = findViewById(R.id.moonImageView);
//
//
//        if (flagSunOut) {
//            sunImageView.animate().translationYBy(2000).translationXBy(-2000).rotation(360).scaleX(0.2f).scaleY(0.2f).setDuration(2500);
//
//            moonImageView.animate().translationYBy(-2000).translationXBy(-2000).scaleX(1).scaleY(1).alpha(1).setDuration(2500);
//
//            flagSunOut = false;
//        } else {
//            moonImageView.animate().translationYBy(2000).translationXBy(-2000).alpha(0).scaleX(0.2f).scaleY(0.2f).setDuration(2500);
//
//            sunImageView.animate().translationYBy(-2000).translationXBy(2000).rotation(360).scaleY(1).alpha(1).setDuration(2500);
//
//            flagSunOut = true;
//        }
//    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
}
