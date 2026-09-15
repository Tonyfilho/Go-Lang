#### Aula21F Ex02

## Na documentação do Go é dito que quando tentamos ler um canal fechado ele responde com o valor zero do tipo do canal e um false no comma ok. 
Aqui  no exemplo, 	quando os canais 'par' e 'ímpar' são fechados, antes do true ser enviado para o 'quit', a goroutine da função receive pode acabar
realizando a leitura dos canais já fechados e recebendo um valor 0, daí como não há a verificação do comma ok neles, esse 0 é tratado como um valor válido.
Acredito que uma possível solução seria enviar o true para o canal 'quit' e depois disso, fechar os canais 'par' e 'ímpar', assim não precisaria verificar o comma ok dos canais 'par' e 'ímpar', pois a informação do canal canal 'quit' chegaria primeiro na função receive, e ela já pararia e não tentaria ler os demais canais.
