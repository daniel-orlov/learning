package com.example.dior.numbershapes;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {


    class Number {

        public boolean isSquare(int n){
            int squareNumber = 1;
            int x = 1;
            while (squareNumber <= n) {
                if (squareNumber == n) {
                    return true;
                }
//                squareNumber = (x*(x+1)*(2*x+1))/6;
                squareNumber = x*x;
                x++;
            }
            return false;
        }

        public boolean isTriangular(int n){
            int triangularNumber = 0;
            int x = 1;
            while (triangularNumber <= n) {
                if (triangularNumber == n) {
                    return true;
                }
                triangularNumber += x;
                x++;
            }
            return false;
        }
    }

    public void onClick(View view) {
        EditText et = findViewById(R.id.editText);
        String msg;

        if (et.getText().toString().isEmpty()) {
            msg = "Please, enter your number";
        } else {

            int inputValue = Integer.parseInt(et.getText().toString());
            if (inputValue == 0){
                msg = "Please, enter a positive number";
            } else {
                Number check = new Number();

                if (check.isSquare(inputValue) && check.isTriangular(inputValue)) {
                    msg = et.getText().toString() + " is both triangular and square";
                } else if (check.isSquare(inputValue)) {
                    msg = et.getText().toString() + " is square";
                } else if (check.isTriangular(inputValue)) {
                    msg = et.getText().toString() + " is triangular";
                } else {
                    msg = et.getText().toString() + " is neither triangular nor square";
                }
            }
        }
        Toast.makeText(this, msg, Toast.LENGTH_LONG).show();
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
}
