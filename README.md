# Contributor Map Action

![freeCodeCamp contributor map](docs/main.svg?1)

<p align="center">
  <sup>The image above generated using contributor data from freeCodeCamp/freeCodeCamp repository.</sup>
</p>
  
## Introduction

Contributor Map is a GitHub action that automatically generates an SVG world map with countries colored according to the number of received code contributions. 

It uses GitHub API to get a list of repository contributors and their profile information. The location field from the public profile is used to determine the country from which the contribution came. 

The action can generate a world map for any public GitHub repository.

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

## Credits

* SVG World Map by Cherkash and others, [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0), via Wikimedia Commons
* City database by [GeoNames.org](https://geonames.org), [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)
* Country and Region database from [dr5hn/countries-states-cities-database](https://github.com/dr5hn/countries-states-cities-database), [ODbL-1.0](https://opendatacommons.org/licenses/odbl/1-0/)
