# Hackaton

Equipe

- João Anderson
- Marcelle Tabosa

### Link das imagens

- https://hub.docker.com/repository/docker/marcelletabosa/hackaton/general

### Comandos utilizados

- vagrant init
- vagrant up
- vagrant ssh ubuntu
- vagrant ssh centos
- git clone https://github.com/ittalent2023-2/go-mongo-crud-rest-api.git
- touch Dockerfile
- nano Dockerfile
- docker build -t marcelletabosa/hackaton:1.2 . 
- docker push marcelletabosa/hackaton:1.2
- touch docker-compose.yml
- nano docker-compose.yml
- docker compose up -d
- curl -X POST 'http://192.168.56.10:9080/users' -H "Content-Type: application/json" -d '{"name": "Bob", "email": "bob@gmail.com", "password": "ilovealice"}'
- curl -X GET 'http://192.168.56.10:9080/users/bob@gmail.com'
- curl -X PUT 'http://192.168.56.10:9080/users/bob@gmail.com' -H "Content-Type: application/json" -d '{"password": "loveyoualice"}'
- curl -X DELETE 'http://192.168.56.10:9080/users/bob@gmail.com'