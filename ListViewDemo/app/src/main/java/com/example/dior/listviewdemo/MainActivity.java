package com.example.dior.listviewdemo;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.ListView;

import java.util.ArrayList;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        ListView myListView = findViewById(R.id.myListView);

        final ArrayList<String> athletes = new ArrayList<String>();
        athletes.add("Brent Fikowski");
        athletes.add("Mat Fraser");
        athletes.add("Annie Thorisdottir");
        athletes.add("Kristen Holte");
        athletes.add("Lukas Hogberg");
        athletes.add("Patrick Vellner");
        athletes.add("Laura Horvath");
        athletes.add("James Newberry");
        athletes.add("Scott Panchik");
        athletes.add("Saxon Panchik");
        athletes.add("Bjorgvin Gudmundsson");
        athletes.add("Rich Froning");
        athletes.add("Jacob Heppner");
        athletes.add("Sara Sigmundsdottir");

        final ArrayAdapter<String> arrayAdapter = new ArrayAdapter<String>(this, android.R.layout.simple_list_item_1, athletes);

        myListView.setAdapter(arrayAdapter);

        myListView.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
                Log.i("ItemPressed", athletes.get(i));
            }
        });

    }
}
