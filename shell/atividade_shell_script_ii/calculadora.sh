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
if  [ $OPCAO -eq 1 -o $OPCAO -eq 2 -o $OPCAO -eq 3  -o $OPCAO -eq 4 ]
then
    echo "Digite dois números"
    read NUM1
    read NUM2
fi

case $OPCAO in
1)
    echo "Opção escolhida: ADIÇÃO" 
    Adicao;;
2)
    echo "Opção escolhida: SUBTRAÇÃO"
    Subtracao;;
3)
    echo "Opção escolhida: MULTIPLICAÇÃO"
    Multiplicacao;;
4)
    echo "Opção escolhida: DIVISÃO"
    Divisao;;
*)
    echo "INVÁLIDO";;
esac


   
    