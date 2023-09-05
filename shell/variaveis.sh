#!/usr/bin/env bash

NOME="Marcelle 
Tabosa"
echo "$NOME"

NUM_1=26
NUM_2=27
TOTAL=$((NUM_1+NUM_2))

echo "$TOTAL"

SAIDA_CAT="$(cat /etc/passwd | grep marcelle)"
echo "$SAIDA_CAT"

echo "Parametro 1: $1"
echo "Parametro 2: $2"

echo "Todos os parâmetros: $*"

echo "Quantidade de parâmetros: $#"

echo "Último parâmetros: $?"

echo "PID: $$"

echo "Nome do script: $0"
