#!/bin/bash

git fetch
git branch -v
git merge origin/main

export PATH=$PATH:/usr/local/go/bin
go run .