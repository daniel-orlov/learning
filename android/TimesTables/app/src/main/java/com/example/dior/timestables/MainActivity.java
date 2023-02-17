package com.example.dior.timestables;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.widget.ArrayAdapter;
import android.widget.ListAdapter;
import android.widget.ListView;
import android.widget.SeekBar;
import java.util.ArrayList;
import java.util.Arrays;

import static java.util.Arrays.asList;

public class MainActivity extends AppCompatActivity {

    ListView timesTable;

    public void generateTimesTable(int seekBarCurrentValue) {
        ArrayList<String> times = new ArrayList <String>();
        for (int j = 1; j < 10; j++) {;
            times.add(Integer.toString(j * seekBarCurrentValue));
        }
        ArrayAdapter<String> arrayAdapter = new ArrayAdapter<>(this, android.R.layout.simple_list_item_1, times);
        timesTable.setAdapter(arrayAdapter);
    }


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        timesTable = findViewById(R.id.timesTable);
        final SeekBar selectMultiplier = findViewById(R.id.seekBar);

        int min = 1;
        int max = 26;

        selectMultiplier.setMax(max);
        selectMultiplier.setMin(min);
        selectMultiplier.setProgress(min);

        generateTimesTable(min);

        selectMultiplier.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int i, boolean b) {

                generateTimesTable(i);

            }

            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {

            }

            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {

            }
        });

    }
}
