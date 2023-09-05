#!/usr/bin/env bash

#[ "$1" -gt 10 ] && echo "$0 $$"


echo "Digite um número"
read NUMERO
[ $(($NUMERO%2)) == 0 ] && echo "PAR" || echo "IMPAR"
