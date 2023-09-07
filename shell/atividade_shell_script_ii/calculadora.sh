#!/bin/bash

echo "Escolha uma opção"
echo "1 - ADIÇÃO"
echo "2 - SUBTRAÇÃO"
echo "3 - MULTIPLICAÇÃO"
echo "4 - DIVISÃO"

read OPCAO
if  [ $OPCAO -eq 1 -o $OPCAO -eq 2 -o $OPCAO -eq 3  -o $OPCAO -eq 4 ]
then
    echo "Digite dois números"
    read NUM1
    read NUM2
fi

case $OPCAO in
1)
    echo "Opção escolhida: ADIÇÃO"
    echo "$NUM1 + $NUM2 = $[$NUM1 + $NUM2]" ;;
2)
    echo "Opção escolhida: SUBTRAÇÃO"
    echo "$NUM1 - $NUM2 = $[$NUM1 - $NUM2]" ;;
3)
    echo "Opção escolhida: MULTIPLICAÇÃO"
    echo "$NUM1 * $NUM2 = $[$NUM1 * $NUM2]" ;;
4)
    echo "Opção escolhida: DIVISÃO"
    echo "$NUM1 / $NUM2 = $[$NUM1 / $NUM2]" ;;
*)
    echo "INVÁLIDO";;
esac