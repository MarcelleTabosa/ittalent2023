#!/usr/bin/env bash
#
# organizacao-com-cabecalho.sh - printa *
#
# Autor:      Marcelle Tabosa
# Manutenção: Marcelle Tabosa
#
# ------------------------------------------------------------------------ #
#  Este programa irá imprimir * no console
# ------------------------------------------------------------------------ #
# Histórico:
#
#   v1.0 08/09/2023 Marcelle:
#       - Copia do programa
# ------------------------------------------------------------------------ #
# Testado em:
#   bash version 5.0.17(1)-release (x86_64-pc-linux-gnu)
# ------------------------------------------------------------------------ #
# Agradecimentos:
#
# 	Mateus Müller - Disponibilizou o arquivo com o código
# ------------------------------------------------------------------------ #

# ------------------------------- VARIÁVEIS ----------------------------------------- #
# COMECA, ATE
# ------------------------------------------------------------------------ #

# ------------------------------- TESTES ----------------------------------------- #
# N/A
# ------------------------------------------------------------------------ #

# ------------------------------- FUNÇÕES ----------------------------------------- #
# N/A
# ------------------------------------------------------------------------ #

# ------------------------------- EXECUÇÃO ----------------------------------------- #
# ./organizacao-com-cabecalho.sh
# ------------------------------------------------------------------------ #

#-----------------------------------------------------------------------------------------------------------------#
comeca=0 ate=100
[ $comeca -ge $ate ] && exit 1

for i in $(seq $comeca $ate)
do 
    for j in $(seq $i $ate)
    do 
        printf "*"
    done
    printf "\n"
done



#-----------------------------------------------------------------------------------------------------------------#

