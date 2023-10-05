#imagem base
FROM golang:alpine
#local onde serão copiados os arquivos da aplicação
WORKDIR /api
#copia os arquivos de dependência
COPY go.mod go.sum ./
#baixa as dependências do arquivo de configuração
RUN go mod download && go mod verify
#copia os arquivos da aplicação para a raiz
COPY . ./
#builda a imagem
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/api
#expõe a porta
EXPOSE 9080
#define o comando para executar a aplicação
CMD [ "api" ]
