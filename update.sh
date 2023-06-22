#!/bin/bash

git fetch
git branch -v
git merge origin/main

go run .


# sudo docker build -t darwin --build-arg USER_AGENT="TOTO" .
# sudo docker run -d -p 8080:8080 --name darwin darwin
