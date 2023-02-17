package com.example.dior.loginscreen;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {

    public void clickFunction(View view){
        EditText usernameEditText = findViewById(R.id.usernameEditText);

        EditText passwordEditText = findViewById(R.id.passwordEditText);

        Log.i("Values", usernameEditText.getText().toString());

        Log.i("Values", passwordEditText.getText().toString());

        Toast.makeText(this, "Well met!", Toast.LENGTH_LONG).show();

    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
}
