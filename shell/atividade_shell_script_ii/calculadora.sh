#!/bin/bash

Adicao(){
    echo "$NUM1 + $NUM2 = $[$NUM1 + $NUM2]"
}
Subtracao(){
    echo "$NUM1 - $NUM2 = $[$NUM1 - $NUM2]"
}
Multiplicacao(){
    echo "$NUM1 * $NUM2 = $[$NUM1 * $NUM2]" 
}
Divisao(){
    echo "$NUM1 / $NUM2 = $[$NUM1 / $NUM2]" 
}

echo "Escolha uma opção"
echo "1 - ADIÇÃO"
echo "2 - SUBTRAÇÃO"
echo "3 - MULTIPLICAÇÃO"
echo "4 - DIVISÃO"

read OPCAO
if  [ $OPCAO -ge 1 -o $OPCAO -le 4 ]
then
    echo "Digite dois números"
    read NUM1
    read NUM2
fi

case $OPCAO in
1)
    Adicao;;
2)
    Subtracao;;
3)
    Multiplicacao;;
4)
    Divisao;;
*)
    echo "INVÁLIDO";;
esac


   
    