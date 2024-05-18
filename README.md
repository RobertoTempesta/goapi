# API de Estudos

Este é um projeto de uma API de estudos, desenvolvida utilizando as seguintes tecnologias:

- **GoLanguage**: Linguagem de programação utilizada para o desenvolvimento da API.
- **Swagger**: Ferramenta para documentação e teste da API.
- **SQLite**: Banco de dados utilizado para armazenamento de dados.
- **Gin**: Framework web utilizado para criar a API.

## Autor

- **Roberto Junior Tempesta**

## Instalação

1. Clone o repositório:
   ```sh
   git clone https://github.com/RobertoTempesta/seurepositorio.git
2. Navegue até o diretório do projeto:
    ```sh
    cd seurepositorio
3. Instale as dependências:
    ```sh
    go mod tidy

## Configuração

1. Crie o banco de dados SQLite:
    ```sh
    sqlite3 openings.db < schema.sql
2. Execute a aplicação:
    ```sh
    go run main.go

## Uso
A documentação da API está disponível via Swagger. Para acessá-la, inicie a aplicação e navegue até:

http://localhost:8080/swagger/index.html

## Endpoints

- *GET* /v1/opening;
- *POST* /v1/opening;
- *DELETE* /v1/opening;
- *PUT* /v1/opening;
- *GET* /v1/openings;

## Vídeo de Referência
Video utilizado para estudos: https://www.youtube.com/live/L6gk7FHBNkM?si=f_EcK5deInTp0sl7