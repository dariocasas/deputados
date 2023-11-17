# Deputados

Sistema constituído de um servidor local golang, conectado a mongoDB Atlas e, via grpc, com um applicativo Flutter desktop.

O que faz?

- Consulta API da Câmara dos Deputados colhendo os ids de cada deputado e cria um índice.
- Cria o DB, iterando cada id, agrupando segundo parâmetro quantidade de rotinas concorrentes e, concorrentemente, por cada id realiza duas tarefas:
  - consulta API recolhendo as informações de cada deputado.
  - descarrega pagina web do deputado do site da Câmara dos Deputados, parsea o HTML e recolhe mais dados.
- Mostra a lista de deputados, com foto, link para a página e redes sociais.

## Vídeo

- [Vídeo YouTube](https://www.youtube.com/watch?v=F4aH6iiBiTY)

## Referências

- [Pipelines by Sameer Ajmani](https://go.dev/blog/pipelines)

## Requisitos

  Golang, Flutter, criar um projeto no [MongoDB Atlas](https://www.mongodb.com/atlas/database)

## Estrutura

### Server (golang)

  *Caminho*: /src

  *Configuração*: criar /src/cmd/server/.env

~~~env
MONGODB_HOST = "mongodb+srv://databaseUser:password@cluster0.w0ga1rt.mongodb.net/"
DEPS_GRPC_PORT=50051
DEPS_GRPC_ADDR="localhost"   
~~~

### Frontend (Flutter desktop windows)

  *Caminho*: /frontend/deputados

  *Configuração*: criar /frontend/deputados/.env

~~~env
DEPS_GRPC_PORT=50051
DEPS_GRPC_ADDR="localhost"   
~~~

### Ferramentas

- build.bat
- run.bat

## Todo

- GRPC Health Checking
- Web Client. Proxy? gRPC-Web?
