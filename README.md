# serasa-challenge

Um projeto que tem como ideia criar um façade de uma aplicação legada que está sobrecarregada!

## Requisitos

É necessário somente o [docker-compose](https://docs.docker.com/compose/install/) para iniciar a aplicação

## Iniciando o projeto

Para iniciar o projeto basta executar o comando a seguir na raiz do projeto

    make init

Obs: as portas que serão ocupadas pelos serviços são as `5000` e `8529`. Certifique-se que essas
portas não estejam ocupadas na sua máquina!

## Endpoints

Como requisitado, há dois endpoints, um para ocorrer a sicronização dos dados com o mainframe e outro
para realizar a leitura de negativações a partir de um CPF

### Atualize informações do mainframe

Endpoint para realizar o processo de sicronismo do serviço com o mainframe.

Metodo `POST`

Rota `/v1/update`

**Exemplo**

    curl --location --request POST 'http://localhost:5000/v1/update'    

**Resposta**

Status `204`

### Ler negativações por documento

Endpoint para realizar a leitura de negativações a partir do CPF.

Metodo `GET`

Rota `/v1/negativations`

Parâmetro `cpf`

**Exemplo**

    curl --location --request GET 'http://localhost:5000/v1/negativation?cpf=51537476467'    

**Resposta**

Status `200`

Corpo 

```json
[
    {
        "companyDocument": "59291534000167",
        "companyName": "ABC S.A.",
        "customerDocument": "51537476467",
        "value": 1235.23,
        "contract": "bc063153-fb9e-4334-9a6c-0d069a42065b",
        "debtDate": "2015-11-13T20:32:51-03:00",
        "inclusionDate": "2020-11-13T20:32:51-03:00"
    },
    {
        "companyDocument": "77723018000146",
        "companyName": "123 S.A.",
        "customerDocument": "51537476467",
        "value": 400,
        "contract": "5f206825-3cfe-412f-8302-cc1b24a179b0",
        "debtDate": "2015-10-12T20:32:51-03:00",
        "inclusionDate": "2020-10-12T20:32:51-03:00"
    }
]
```

## Contato

[LinkedIn](https://www.linkedin.com/in/ednailsonvb/) | ednailsoncunha@gmail.com

