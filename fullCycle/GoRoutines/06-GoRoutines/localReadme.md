#### FullCicles GoRoutines Ex 06 - Entendendo Go Routines 

https://www.youtube.com/watch?v=B4NL0rMvXMg

## Como um Webservice em Go é mais rápido do que outras liguagens.

Entenderemos como um webservice consegue rodar com tantas requisiões e ser mais rapido do que outros , é por causa das GoRoutines!

Resumo, este For de entrada, pode fazer o Look quando o canal estiver Vazio, ou seja o for fica´ra esperando
O Worker pegará a a mensagem e esperará 1 segundo, somente depois dará permissão for de abastecimente 
Mas a medida q formos adicionando Worker teremos mais Threads e com mais Threads mais Processamento
Um Thread em outras liguam custa 1 Mega, aqui no Go custa 2 kbits
