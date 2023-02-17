package com.example.dior.listview2;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.Toast;

import java.util.ArrayList;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        ListView myListView = findViewById(R.id.myListView);

        final ArrayList<String> countries = new ArrayList<String>();
        countries.add("Russia");
        countries.add("Turkey");
        countries.add("Tunisia");
        countries.add("Germany");
        countries.add("France");
        countries.add("Denmark");
        countries.add("UAE");
        countries.add("India");
        countries.add("Finland");
        countries.add("Sweden");
        countries.add("UK");
        countries.add("Argentine");
        countries.add("Chile");
        countries.add("Spain");
        countries.add("Estonia");
        countries.add("Italy");
        countries.add("the Netherlands");

        ArrayAdapter<String> arrayAdapter = new ArrayAdapter<String>(this, android.R.layout.simple_list_item_1, countries);

        myListView.setAdapter(arrayAdapter);

        myListView.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
                Toast.makeText(getApplicationContext(), countries.get(i), Toast.LENGTH_SHORT).show();
            }
        });
    }
}
