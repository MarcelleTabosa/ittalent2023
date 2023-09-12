#!/usr/bin/env bash
#
# chaves_flags.sh - extrai usuarios do /etc/passwd *
#
# Autor:      Marcelle Tabosa
# Manutenção: Marcelle Tabosa
#
# ------------------------------------------------------------------------ #
#  Este programa irá extrair usuários do /etc/passwd, havendo a possibilidade
# de colocar em maiúsculo e em ordem alfabética
#
#   Exemplos:
#   $ ./chaves_flags.sh -s -m
#   Neste exemplo ficará em maiúsculo e em ordem alfabética
# ------------------------------------------------------------------------ #
# Histórico:
#
#   v1.0 09/12/2023 Marcelle:
#       -Adicionado -s, -h, -v
#   v1.1 09/12/2023 Marcelle:
#       
#       
# ------------------------------------------------------------------------ #
# Testado em:
#   bash version 5.0.17(1)-release (x86_64-pc-linux-gnu)
# ------------------------------------------------------------------------ #
# Agradecimentos:
#
# 	Mateus Müller - Disponibilizou o conteúdo em aula
# ------------------------------------------------------------------------ #

# ------------------------------- VARIÁVEIS ----------------------------------------- #
USUARIOS=$(cat /etc/passwd | cut -d : -f 1)
MENSAGEM_USO="
    ${basename $0} - [OPÇÕES]

    -h - Menu de ajuda
    -v - Versão
    -s - Ordernar saída
"
VERSAO="v1.0"
# ------------------------------------------------------------------------ #
# ------------------------------- EXECUÇÃO ----------------------------------------- #
# ./chaves_flags.sh
# ------------------------------------------------------------------------ #

#-----------------------------------------------------------------------------------------------------------------#
case "$1" in
    -h) echo "$MENSAGEM_USO" && exit 0 ;;
    -v) echo "$VERSAO" && exit 0 ;;
    -s) echo "$USUARIOS" | sort && exit 0;;
    *) echo "$USUARIOS" && exit 0
esac