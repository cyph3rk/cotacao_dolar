# cotacao_dolar




* Primeira correção: 23/06/2025 14:46

Olá, Diego!

Esperamos que esteja tudo certo por aí!

O timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
>> Inlcui a correção, eu tinha feito mas durante uma busca incansável por um segmentFault, acabei perdendo.

E utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado.
>> Esse eu não entendi muito bem, o cliente já está com esse timeout. Está no arquivo cliente/cliente.go linha 20

Reveja esse ponto e quando terminar, envie o projeto novamente para correção.

Bom trabalho!

TODO:
- [x] Arquivo de saida do cliente não sobrescrever a saida
- [ ] Criar um arquivo separado com logs inclusive erros
- [x] Criar um arquivo de configuração os valores constantes como portas e timeouts.







***************


* Primeiro subir a aplicação servidor.

```sh
go run main.go
```

* Depois testar rodando a aplicação cliente.

```sh
go run cliente/main.go
```
