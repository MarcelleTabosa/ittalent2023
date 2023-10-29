# Nginx

### Intale o nginx em um container no docker

docker run -itd --name nginx -p 80:80 nginx

O comando baixa a imagem, caso ainda não tenha baixado, cria o container com a imagem do nginx, executa em modo iterativo e cria em modo detached, nomeia o container e mapea a porta padrão do nginx.

### Verifique se o container está sendo executado com o comando

docker ps

o comando lista os containers que estão sendo executados.

### Acesse a página padrão do nginx no navegador

http://localhost:80

### Verifique o caminho do arquivo de configuração do nginx

nginx -h

o comando acima lista alguns comandos e paths do nginx
