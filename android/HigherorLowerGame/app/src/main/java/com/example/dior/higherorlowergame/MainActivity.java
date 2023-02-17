package com.example.dior.higherorlowergame;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import java.util.Random;

public class MainActivity extends AppCompatActivity {
    int misteryInt;

    public int generateRandInt(){
        Random randInt = new Random();
        int x;
        x = randInt.nextInt(20) + 1;
        return x;
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        misteryInt = generateRandInt();
    }

    public void guessClicked(View view){
        EditText et = findViewById(R.id.editText);
        int guessInt = Integer.parseInt(et.getText().toString());

        String msg;

        if (guessInt > misteryInt) {
            msg = "Lower!";
        } else if (guessInt < misteryInt) {
            msg = "Higher!";
        } else {
            msg = "Bingo! You can play again!";
            misteryInt = generateRandInt();
        }

        Toast.makeText(this, msg, Toast.LENGTH_LONG).show();
    }
}
