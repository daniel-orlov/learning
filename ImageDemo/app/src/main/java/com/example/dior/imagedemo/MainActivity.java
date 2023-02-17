package com.example.dior.imagedemo;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;

public class MainActivity extends AppCompatActivity {

    ImageView image;
    boolean flag = true;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        image = (ImageView) findViewById(R.id.thumbUp);
    }


    public void onClick(View view) {
        if(flag)
        {
            image.setImageResource(R.drawable.thumbsup);
            flag = false;
        }
        else
        {
            image.setImageResource(R.drawable.thumbsdown);
            flag = true;
        }
    }
}
