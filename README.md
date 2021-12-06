![freeCodeCamp contributor map](docs/main.svg?1)

# Contributor Map GitHub Action

Automatically generate a world map showing where contributions to your repository are coming from.

## Usage

Adding the following to an existing workflow would create a new "contributor-map.svg" file inside your repository.

```yml
- name: Contributor Map
  uses: tunaitis/contributor-map@v1
```

The following example would generate a map showing contributions of the **facebook/react** repository and save it to your repository's **data** folder.

```yml
- name: Contributor Map
  uses: tunaitis/contributor-map@v1
  with:
    repository: facebbook/react
    output: data/facebook-react.svg
```

Example of the action inside a workflow that wouild run the action every Monday at 8AM UTC to create a map.

```yml
name: Create Contributor Map

on:
  schedule:
    - cron: "0 8 * * 1"
    
jobs:
  build-and-push-image:
    runs-on: ubuntu-latest
    
    steps:
      - name: Contributor Map
        uses: tunaitis/contributor-map@v1 
```

## Inputs

### Parameters

|name|required|description|
|---|:-:|---|
|palette|no|A comma-separated list of HTML color codes to define a custom color scheme.|
|repository|no|Name of repository to use for generating the map (e.g, facebook/react, golang/go). Current repository name will be used if the parameter is not provided.|
|output|no|Name of the file where to save the map. |
