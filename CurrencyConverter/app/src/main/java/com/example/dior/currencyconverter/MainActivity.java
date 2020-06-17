package com.example.dior.currencyconverter;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {

    public void onClickUSDuoBINConversion(View view){
        EditText editText = findViewById(R.id.editText);

        String amountInUsd = editText.getText().toString();

        double amountInUsdDouble = Double.parseDouble(amountInUsd);

        double amountInBynDouble = amountInUsdDouble * 2.39;

        String amountInByn = String.format("%.2f", amountInBynDouble);

        Toast.makeText(this, (amountInUsd + " USD is " + amountInByn + " BYN"), Toast.LENGTH_LONG).show();

    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
}
