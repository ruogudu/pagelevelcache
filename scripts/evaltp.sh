#!/bin/bash

samples="4 8 16 32 64"
sizes="36078002 360780020 3607800200"
threads="1 2 4 8"

for size in $sizes
do
for snum in $samples
do
for th in $threads
do

echo Samples: $snum
echo Cache size: $size

go run main.go -e -t /ssd/msr/trace_2018_03_06_1m.json -m $snum -s $size -n $th -a h2 -ad 100000

done
done
done
