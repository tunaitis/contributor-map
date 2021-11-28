#!/usr/bin/env bash

rm -rf .out
mkdir .out

declare -a reps=("golang/go" "rails/rails" "ruby/ruby" "facebook/react" "Homebrew/brew" "dotnet/core" "microsoft/vscode" "scala/scala" "nodejs/node" "mongodb/mongo")

for rep in "${reps[@]}"
do
  INPUT_CACHE=1 INPUT_REPOSITORY=$"$rep" INPUT_OUTPUT=.out/${rep//\//-}.svg go run cmd/contributor-map/main.go
  echo "<h4>${rep}</h4><img src=\"${rep//\//-}.svg\" />" >> .out/preview.html
done
