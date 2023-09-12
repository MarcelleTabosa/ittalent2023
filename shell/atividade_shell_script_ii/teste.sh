#!/bin/bash

echo "Digite o nome da pasta"
read pasta

if [ -e $pasta ]
then 
    echo "Dir: $pasta, já existe, saindo."
    exit 0
else
    echo "Dir: $pasta, não existe, e será criado."
    mkdir $pasta
    [ -e $pasta ] && echo "Dir: $pasta criado com sucesso"
fi

for (( i = 1 ; i <= 10; i++ ))
do
    mkdir "$pasta/dir$i"
done