#!/bin/bash
time parallel -a ./topics.txt -j 240 bisquitt-pub -h 127.0.0.1 -p 1883 -t {} -m "{}" 