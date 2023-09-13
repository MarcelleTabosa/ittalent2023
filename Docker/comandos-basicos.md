# Docker

##### Lista as imagens 
- docker images

##### Lista os containers ativos 
- docker ps

##### Lista os containers ativos e inativos
- docker ps -a

##### Baixa a imagem, caso não tenha local, e executa o container
- docker container run hello-world

## Comando run faz pull da imagem, cria container, inicializa o container e executa em modo iterativo

#### Comando run faz pull da imagem, cria container, inicializa o container, executa em modo iterativo e executa o bash --version

- docker container run debian bash --version

##### Remove um container pelo nome ou ID
- docker rm debian

##### Cria e executa o cantainer no modo iterativo e o terminal
- docker run -it debian bash

##### Cria e executa e nomeia o container no modo iterativo e o terminal
- docker run --name mydeb -it debian bash

##### Inicializa o container no modo iterativo e o terminal
- docker start -ai mydeb

#### Baixa, executa e mapeia a porta
- docker run -p 8080:80 nginx

#### Exibe os logs 
- docker logs vibrant_goldwasser

#### Mapea volumes
- docker run -p 8080:80 -v $(pwd)/html:/usr/share/nginx/html nginx
