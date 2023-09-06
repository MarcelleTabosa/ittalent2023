#!/usr/bin/env bash

echo "==== For1"
for (( i = 0; i < 10; i++ )); do
    echo $i
done

echo "==== For2"
for i in $(seq 10)
do
    echo $i
done

echo "==== For3"
Fruits=('Laranja'
'Ameixa'
'Manga')
for i in ${Fruits[@]}
do
    echo $i
done

echo "==== While"
count=0
while [[ $count -lt ${#Fruits[@]} ]]; do
    echo $count
    count=$((count+1))
done