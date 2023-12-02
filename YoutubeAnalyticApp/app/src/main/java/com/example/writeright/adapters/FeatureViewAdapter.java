package com.example.writeright.adapters;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.writeright.R;
import com.example.writeright.pojo.Feature;

import java.util.ArrayList;

public class FeatureViewAdapter extends RecyclerView.Adapter<FeatureViewHolder> {

    private ArrayList<Feature> features;

    public FeatureViewAdapter(ArrayList<Feature> features) {
        this.features = features;
    }

    @NonNull
    @Override
    public FeatureViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.feature_card, parent, false);
        return new FeatureViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull FeatureViewHolder holder, int position) {
        Feature feature = features.get(position);

        holder.featureHeader.setText(feature.getName());
        holder.featureDescription.setText(feature.getDescription());
        holder.featureImageView.setImageResource(feature.getDrawableResource());
    }

    @Override
    public int getItemCount() {
        return features.size();
    }
}

class FeatureViewHolder extends RecyclerView.ViewHolder {
    TextView featureHeader, featureDescription;
    ImageView featureImageView;

    public FeatureViewHolder(@NonNull View itemView) {
        super(itemView);

        featureHeader = itemView.findViewById(R.id.feature_header);
        featureDescription = itemView.findViewById(R.id.feature_description);
        featureImageView = itemView.findViewById(R.id.feature_img);
    }
}
