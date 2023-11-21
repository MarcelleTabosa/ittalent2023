FROM ubuntu:latest

WORKDIR /

COPY atividade.sh /

RUN apt-get update && apt-get install curl traceroute -y

RUN chmod +x atividade.sh

CMD ["./atividade.sh"]