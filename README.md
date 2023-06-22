# Teste de Stress com T50 em Go

Este projeto em Go foi desenvolvido para realizar testes de stress em aplicações web utilizando a ferramenta T50. O T50 é uma ferramenta de teste de carga e estresse que permite medir o desempenho e a capacidade de uma aplicação em lidar com um alto volume de tráfego.

## Pré-requisitos

Antes de executar o projeto, você precisará ter instalado o Go em sua máquina. Para obter o Go, visite o site oficial do Go (https://golang.org) e siga as instruções de instalação para o seu sistema operacional.

## Como usar

1. Faça o clone deste repositório para a sua máquina local:

   ```shell
   git clone https://github.com/osamucaaaa/stress.git
   ```

2. Navegue até o diretório do projeto:

   ```shell
   cd stress
   ```

3. Ajuste as configurações do teste de stress no arquivo `main.go`:

   - `numRequests`: Número total de requisições a serem feitas durante o teste.
   - `concurrency`: Número de requisições concorrentes a serem executadas ao mesmo tempo.
   - `targetURL`: URL do alvo a ser testado.

4. Execute o projeto:

   ```shell
   go run main.go
   ```

   O projeto iniciará o teste de stress, realizando as requisições para o URL especificado. Após a conclusão do teste, será exibido o total de requisições feitas.

## Personalização

Este projeto é um esboço básico de um teste de stress em Go usando o T50. Você pode personalizá-lo de acordo com suas necessidades, adicionando lógica adicional para processar as respostas, tratamento de erros, configurações avançadas de requisições, monitoramento de métricas, entre outros.

Sinta-se à vontade para explorar o código-fonte e adaptá-lo de acordo com os requisitos específicos do seu teste de stress.

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema, tiver sugestões ou melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Agradecimentos

Agradecemos ao projeto T50 (https://github.com/secretsquirrel/the-backdoor-factory) por fornecer a ferramenta de teste de stress utilizada neste projeto.
