package com.example.writeright;

import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import android.os.Bundle;

import com.example.writeright.adapters.FeatureViewAdapter;
import com.example.writeright.pojo.Feature;

import java.util.ArrayList;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        RecyclerView featureRecyclerView = findViewById(R.id.feature_recyclerview);
        featureRecyclerView.setLayoutManager(new LinearLayoutManager(this));
        FeatureViewAdapter featureViewAdapter = new FeatureViewAdapter(getFeatureArrayList());
        featureRecyclerView.setAdapter(featureViewAdapter);
    }

    private ArrayList<Feature> getFeatureArrayList() {
        ArrayList<Feature> features = new ArrayList<>(3);

        features.add(new Feature(getString(R.string.summarize_text), getString(R.string.summarize_description), R.drawable.summarise));
        features.add(new Feature(getString(R.string.sentiment_text), getString(R.string.sentiment_description), R.drawable.sentiment));
        features.add(new Feature(getString(R.string.correctness_text), getString(R.string.correctness_decription), R.drawable.correctness));

        return features;
    }
}