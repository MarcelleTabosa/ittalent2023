#!/bin/bash

echo "Digite o tamanho que deseja"
read TAMANHO
SENHA=`</dev/random tr -dc A-Z a-z 0-9 | head -c${TAMANHO}`
echo "Sua senha $SENHA"