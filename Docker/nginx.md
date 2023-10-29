# Nginx

### Intale o nginx em um container no docker

docker run -itd --name nginx -p 80:80 nginx

O comando baixa a imagem, caso ainda não tenha baixado, cria o container com a imagem do nginx, executa em modo iterativo e cria em modo detached, nomeia o container e mapea a porta padrão do nginx.

