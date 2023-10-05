# imagem base com golang
FROM golang:latest

# pasta de trabalho onde será copiado os arquivos da aplicação
WORKDIR /app

# copia or arquivos que contém as dependências
COPY app-search-ip/go.mod app-search-ip/go.sum ./

# baixa as dependências
RUN go mod download && go mod verify

# copia os arquivos da aplicação para a pasta de trabalho
COPY app-search-ip .

# cria o arquivo binário dentro do diretório /usr/local/bin/app
RUN go build -o /usr/local/bin/app