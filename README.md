<h1 align="center">
  <img alt="Logo" src="./docs/img/logo.png">
</h1>

<h1 align="center">Jornada Milhas</h1>
<p align = "center">Desafio 7 de back-end da Alura</p>

<p align="center">
  <a href="#-tecnologia">Tecnologia</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-projeto">Projeto</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-como-executar">Como Executar</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

## Challenge
Acompanhe os requistos do projeto e as tarefas planejadas através dos seguintes quadros do Trello:

- **[Semana 1](https://trello.com/b/OnuqDQ3A/alurachallengebackend7-semana-1)**
- **[Semana 2](https://trello.com/b/lxgEDut9/alurachallengebackend7-semana-2)**
- **[Semana 3](https://trello.com/b/Cuh1vI9X/alurachallengebackend7-semana-3)**
## ✨ Tecnologia

The Project was develop as using the following techs:
- [Go](https://go.dev/)
- [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
- [sqlc](https://sqlc.dev/)
- [Migrate](https://github.com/golang-migrate/migrate)
- [chi v5](https://go-chi.io/#/)
- [viper](https://github.com/spf13/viper)
- [Swagger](https://github.com/swaggo/swag)
- [OpenAI ChatGPT](https://openai.com/)


## 💻 Projeto
A Jornada Milha API é uma aplicação backend desenvolvida para gerenciar destinos e depoimentos relacionados a viagens. Ela permite que os usuários criem, atualizem, excluam e pesquisem destinos, além de gerenciar depoimentos associados a esses destinos. Uma característica única da API é a integração com a inteligência artificial ChatGPT da OpenAI, que é usada para gerar descrições de texto para destinos quando o campo de descrição é deixado vazio ou nulo.

### API Endpoints
O endpoint da API estão descritos abaixo

#### **Destinos (Destinations)**

1. **Listar Destinos**
    - **Endpoint**: **`GET /api/v1/destinations`**
    - **Descrição**: Retorna uma lista de todos os destinos disponíveis.
2. **Obter Destino**
    - **Endpoint**: **`GET /api/v1/destinations/{id}`**
    - **Descrição**: Retorna os detalhes de um destino específico pelo ID.
3. **Criar Destino**
    - **Endpoint**: **`POST /api/v1/destinations`**
    - **Descrição**: Cria um novo destino com as informações fornecidas. Se o campo de descrição estiver vazio, a IA ChatGPT será usada para gerar uma descrição.
4. **Atualizar Destino**
    - **Endpoint**: **`PUT /api/v1/destinations/{id}`**
    - **Descrição**: Atualiza as informações de um destino existente pelo ID.
5. **Excluir Destino**
    - **Endpoint**: **`DELETE /api/v1/destinations/{id}`**
    - **Descrição**: Exclui um destino específico pelo ID.

#### **Depoimentos (Testimonials)**

1. **Listar Depoimentos**
    - **Endpoint**: **`GET /api/v1/testimonials`**
    - **Descrição**: Retorna uma lista de todos os depoimentos disponíveis.
2. **Obter Depoimento**
    - **Endpoint**: **`GET /api/v1/testimonials/{id}`**
    - **Descrição**: Retorna os detalhes de um depoimento específico pelo ID.
3. **Criar Depoimento**
    - **Endpoint**: **`POST /api/v1/testimonials`**
    - **Descrição**: Cria um novo depoimento com as informações fornecidas.
4. **Atualizar Depoimento**
    - **Endpoint**: **`PUT /api/v1/testimonials/{id}`**
    - **Descrição**: Atualiza as informações de um depoimento existente pelo ID.
5. **Excluir Depoimento**
    - **Endpoint**: **`DELETE /api/v1/testimonials/{id}`**
    - **Descrição**: Exclui um depoimento específico pelo ID.

#### **Documentação (Swagger)**

1. **Documentação da API**
    - **Endpoint**: **`GET /api/v1/docs`**
    - **Descrição**: Acessa a documentação interativa da API através do Swagger.

### Funcionalidades
* Gerenciamento de Destinos: Crie, atualize, exclua e pesquise destinos, incluindo informações como nome, descrição e localização.
* Gerenciamento de Depoimentos: Gerencie depoimentos associados a destinos, incluindo a criação, atualização e exclusão de depoimentos.
* Geração Automática de Descrições: Utilize a inteligência artificial ChatGPT para gerar descrições de texto para destinos de forma automática.

## 🚀 Como Executar

### Opção 1
1. Clone o repositório
2. Mude para o diretório do projeto
3. Instale as dependências do Go, `go mod tidy `
4. Configure o banco de dados MySQL e atualize o arquivo de configuração
5. Compile a aplicação:
   * `go build cmd/api/main.go
6. Execute a aplicação:
    * `./main`
7. A aplicação agora deve estar rodando na porta especificada (a padrão é 8080):

### Opção 2
1. Clone o repositório
2. Mude para o diretório do projeto
3. Configure o banco de dados MySQL e atualize o arquivo de configuração
4. Compile e execute a aplicação usando o Docker Compose:
     * `docker-compose up --build -d`
     
A aplicação agora deve estar rodando na porta especificada (a padrão é 8080).

Escolha a opção de instalação que melhor atende às suas necessidades e comece a usar a Jornda Milhas API.

Também implementamos testes para garantir que todas as funcionalidades estejam funcionando bem. Realizamos os testes usando o seguinte comando:
```bash
go test -v ./...
```

## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---
## Author

Made with ♥ by Rafael 👋🏻


[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)
