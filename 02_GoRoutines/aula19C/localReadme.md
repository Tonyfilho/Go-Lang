#### Aula19 C Concorrência – 3. Race Condition (Discussão: Condição de corrida)

https://www.youtube.com/watch?v=0qGILXmLfMM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=127

Agora vamos dar um mergulho na documentação:
https://go.dev/doc/effective_go
https://go.dev/doc/effective_go#concurrency


# https://pt.wikipedia.org/wiki/Condi%C3%A7%C3%A3o_de_corrida
Uma condição de corrida é uma falha num sistema ou processo em que o resultado do processo é inesperadamente dependente da sequência ou sincronia doutros eventos. Apesar de ser conhecido em português por 'condição de corrida' uma tradução melhor seria 'condição de concorrência' pois o problema está relacionado justamente ao gerenciamento da concorrência entre processos teoricamente simultâneos.O fenômeno pode ocorrer em sistemas eletrônicos, especialmente em circuitos lógicos, e em programas de computador, especialmente no uso de multitarefa ou computação distribuída.

# tradução Yield = Redimento, produzir, render, fornecer, dar se por vencido, submeter-se, troca de processo, etc ...
O que é yield? runtime.Gosched() é o tempo que o processador de 1 core, usa para entercalar entre as execuções.
# Exemplo de Race condition: Observe o o incremento da variavel não acompanham o momento real da execução veja VAR que é compratilhada entre 2 goroutines
        Função 1       var     Função 2
         Lendo: 0   →   0
         Yield          0   →   Lendo: 0
         var++: 1               Yield
         Grava: 1   →   1       var++: 1
                        1   ←   Grava: 1
         Lendo: 1   ←   1
         Yield          1   →   Lendo: 1
         var++: 2               Yield
         Grava: 2   →   2       var++: 2
                        2   ←   Grava: 2
# E é por isso que vamos ver mutex, atomic e, por fim, channels. Com isto a execução fica bloqueada ate que seja salvo o valor real
## MUTEX Faz Com isto a execução fica bloqueada ate que seja salvo o valor real muito parecido com threads
## ATOMIC São blocos de Mutex, 
## CHANNELS(Canais) é o mais usando em GO
# Programação concorrente é um tópico extenso e aqui há espaço apenas para alguns destaques específicos do Go.

A programação concorrente em muitos ambientes é dificultada pelas sutilezas necessárias para implementar o acesso correto a variáveis ​​compartilhadas. O Go incentiva uma abordagem diferente, na qual os valores compartilhados são passados ​​por meio de canais e, na verdade, nunca são compartilhados ativamente por threads de execução separadas. Apenas uma goroutine tem acesso ao valor em um dado momento. Condições de corrida de dados não podem ocorrer, por design. Para incentivar essa forma de pensar, resumimos tudo em um slogan:

# Obs: Não se comunique compartilhando memória; em vez disso, compartilhe memória comunicando-se.

Essa abordagem pode ser levada ao extremo. Contagens de referência podem ser melhor implementadas colocando um mutex em torno de uma variável inteira, por exemplo. Mas, como uma abordagem de alto nível, usar canais para controlar o acesso facilita a escrita de programas claros e corretos.