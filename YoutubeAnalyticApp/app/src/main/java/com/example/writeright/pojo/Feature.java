package com.example.writeright.pojo;

public class Feature {
    private String name, description;
    private int drawableResource;

    public Feature(String name, String description, int drawableResource) {
        this.name = name;
        this.description = description;
        this.drawableResource = drawableResource;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public void setDrawableResource(int drawableResource) {
        this.drawableResource = drawableResource;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }

    public int getDrawableResource() {
        return drawableResource;
    }
}
