#!/bin/sh

for f in $(ls *.go); do
  go install $f
done
