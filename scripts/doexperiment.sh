#!/bin/bash

samples="4 8 16 32 64"
types="p o"
sizes="464680120 4646801200 46468012000"

run=$1
batch=0

echo Run: $run

sleep 5

for snum in $samples
do
for type in $types
do
for size in $sizes
do

echo Samples: $snum
echo Type: $type
echo Cache size: $size

go run main.go -t /ssd/msr/trace_2018_03_06_30m.json -$type -s $size -a LFU -g 100000 -r 0 -m $snum > "run_${run}_batch_${batch}.txt"
go run main.go -t /ssd/msr/trace_2018_03_06_30m.json -$type -s $size -a LRU -g 100000 -r 0 -m $snum > "run_${run}_batch_$(($batch+1)).txt"
go run main.go -t /ssd/msr/trace_2018_03_06_30m.json -$type -s $size -a hyperbolic -g 100000 -r 0 -m $snum > "run_${run}_batch_$(($batch+2)).txt"
go run main.go -t /ssd/msr/trace_2018_03_06_30m.json -$type -s $size -a h1 -g 100000 -r 0 -m $snum > "run_${run}_batch_$(($batch+3)).txt"
go run main.go -t /ssd/msr/trace_2018_03_06_30m.json -$type -s $size -a h2 -g 100000 -r 0 -m $snum > "run_${run}_batch_$(($batch+4)).txt"
go run main.go -t /ssd/msr/trace_2018_03_06_30m.json -$type -s $size -a h2 -ad 100000 -g 100000 -r 0 -m $snum > "run_${run}_batch_$(($batch+5)).txt"

let batch=batch+10

done
done
done


