#!/bin/bash

for i in {1..16}; do
    go run main.go  /ssd/msr/trace_2018_03_06_1m.json 50000000000 $i > res_50G_$i.txt
done
