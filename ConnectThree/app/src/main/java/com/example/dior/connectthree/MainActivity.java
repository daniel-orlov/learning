package com.example.dior.connectthree;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.GridLayout;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import java.util.HashMap;
import java.util.Map;

public class MainActivity extends AppCompatActivity {

//    0 is Empty, 1 is Yellow, 2 is Red
    int activePlayer = 1;

    int[] gameState = {0, 0, 0, 0, 0, 0, 0, 0, 0};

    int[][] winStates = {{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}};

    boolean gameActive = true;

        public void dropIn(View view) {
        ImageView counter = (ImageView) view;

        int tappedCounter = Integer.parseInt(counter.getTag().toString());

        if (gameState[tappedCounter] == 0 && gameActive) {
            gameState[tappedCounter] = activePlayer;

            counter.setTranslationY(-1500);

            if (activePlayer == 1) {
                counter.setImageResource(R.drawable.yellow);
                activePlayer = 2;
            } else {
                counter.setImageResource(R.drawable.red);
                activePlayer = 1;
            }

            counter.animate().translationYBy(1500).rotation(1800).setDuration(600);




            for (int[] winState : winStates) {
                if (gameState[winState[0]] != 0 && gameState[winState[0]] == gameState[winState[1]] && gameState[winState[1]] == gameState[winState[2]]) {
                    TextView winnerTextView = (TextView) findViewById(R.id.winnerTextView);
                    Button replayButton = (Button) findViewById(R.id.replayButton);
                    if (gameState[winState[0]] == 1) {
                        winnerTextView.setText("Yellow wins!");
                    } else {
                        winnerTextView.setText("Red wins!");
                    }
                    winnerTextView.setVisibility(View.VISIBLE);
                    replayButton.setVisibility(View.VISIBLE);
                    gameActive = false;
                }
            }
        }
    }

    public void replay(View view) {
        TextView winnerTextView = (TextView) findViewById(R.id.winnerTextView);
        Button replayButton = (Button) findViewById(R.id.replayButton);
        winnerTextView.setVisibility(View.INVISIBLE);
        replayButton.setVisibility(View.INVISIBLE);

        GridLayout grid3x3 = (GridLayout) findViewById(R.id.grid3x3);

        for (int i = 0; i < grid3x3.getChildCount(); i++) {
            ImageView counter = (ImageView) grid3x3.getChildAt(i);
            counter.setImageDrawable(null);
        }

        for (int i = 0; i < gameState.length; i++) {
            gameState[i] = 0;
        }

        gameActive = true;
        activePlayer = 1;
    }


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }
}
