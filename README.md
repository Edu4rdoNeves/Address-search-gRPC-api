# Projeto de Busca de Endereço por CEP

Este é um projeto desenvolvido em Go que fornece um serviço para buscar endereços com base em CEPs. Utiliza a concorrência do Go e o gRPC para fornecer respostas rápidas e eficientes.

## Funcionalidades

- Busca de endereço por CEP em duas APIs diferentes de forma concorrente.
- Resposta rápida, utilizando a API que retorna mais rapidamente.
- Mecanismo de timeout para evitar esperas prolongadas.

## Estrutura do Projeto

- **bussiness**: Contém a lógica de negócios relacionada à busca de endereços.
- **dto**: Define as estruturas de transferência de dados utilizadas para comunicação.
- **model**: Define os modelos de dados utilizados internamente no projeto.
- **pb**: Contém os arquivos gerados pelo protocol buffer.
- **proto**: Contém o arquivo .proto que define o serviço gRPC.
- **server**: Implementação do servidor gRPC.
- **service**: Implementação do serviço de busca de endereço.
- **utils**: Utilitários auxiliares, como funções para chamadas HTTP.
- **vendor**: Dependências do projeto.
- **main.go**: Ponto de entrada da aplicação.

## Como Executar

1. Certifique-se de ter o Go instalado em seu sistema.
2. Clone este repositório.
3. Instale as dependências do projeto:
  1. Digite o comando `go mod tidy` para garantir que o seu `go.mod` esteja atualizado com todas as dependências necessárias.
  2. Em seguida, se desejar, você pode executar `go mod vendor` para atualizar o diretório vendor com as dependências necessárias.
  3. Execute o comando `go run main.go` para iniciar o servidor.
  4. Compile o projeto `go build`

## Utilizando o Evans

- Para interagir com o serviço gRPC, recomenda-se o uso do Evans. Siga as instruções abaixo para baixar e executar o Evans:
1. Baixe o Evans do repositório oficial: https://github.com/ktr0731/evans.
2. Extraia o arquivo baixado e navegue até o diretório onde o Evans foi baixado.
3. Execute o Evans no modo REPL, especificando o arquivo proto e o host e porta do servidor gRPC:
   1- `./evans -r -p 50051 -t [caminho para o arquivo .proto]`
   2- Certifique-se de substituir `[caminho para o arquivo .proto]` pelo caminho correto para o seu arquivo .proto.

- Se, ao iniciar o Evans, os métodos não forem encontrados, será criado um arquivo denominado `.evans.toml.` É necessário preencher as informações deste arquivo com base na estrutura do seu projeto.

## Agradeço a todos pela visita e pelo interesse neste projeto. Qualquer contribuição, feedback ou sugestão é sempre bem-vindo!
