# Contributor Map GitHub Action

Automatically generate a world map showing where contributions to your repository are coming from.

![freeCodeCamp contributor map](docs/main.svg?1)

## Usage

Adding the following to your workflow will create a new file "contributor-map.svg" inside your repository

```yml
- name: Contributor Map
  uses: tunaitis/contributor-map@v1
```

The following example would generate a map showing contributions of the **facebook/react** repository and save it to the **data** folder.

```yml
- name: Contributor Map
  uses: tunaitis/contributor-map@v1
  with:
    repository: facebbook/react
    output: data/facebook-react.svg
```

## Inputs

### Parameters

|name|required|description|
|---|:-:|---|
|palette|no|A comma-separated list of HTML color codes to define a custom color scheme.|
|repository|no|Name of repository to use for generating the map (e.g, facebook/react, golang/go). Current repository name will be used if the parameter is not provided.|
|output|no|Name of the file where to save the map. |
