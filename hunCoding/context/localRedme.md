#### Aula sobre Contexts

https://www.youtube.com/watch?v=SujUMmy9BtQ

https://github.com/HunCoding/golang-basics/tree/main/context


# É um tipo de variavel q começa vazia e depois de setada pode ser configurada p não mudar , somente pode ser clonada e editada o clone
# Com isto poderá carregar valores 

# Por baixo o Contexts usa Chanels 

###  Temos 3 tipos de contexts em Go
Deadline onde setamos um empo maximo de funcionamento.
Cancellation Signals onde mandamos um sinal de execução onde o processo seja sincrono e assíncrono termine.
Request-scoped values onde enviamos um valor como se fosse um Objeto ou Map e que depois de setado Não muda mais o valor, so pode ser clonado e adicionado ao clone
é muito bom para dados assícronos, tipo o observable

### Deadlines com o metodo withdeadline() 
è um dos metodo de Contexts terminar  o codigo e saber se terminou por pelo fim da execução ou por timeout do contexto 
Quando usar: Quando queremos determinar ate que data ou hora o codigo pode ser executado
Ex: time.now mais 10 minutos
# Obs:  Apesar de ter o cancelamento automatico o WithdeadLine retorna p context e um cancel(), o cancel() é uma opção extra.

### Deadline com o metodo  WithTimeOut()
Podemos ter um tempo maximo onde passamos a duração final, não passamos horas mais sim quanto tempo Ex: daqui a 5 minutos
# Obs:  WithTimeOut() retorna p context e um cancel(), o cancel() é uma opção extra caso queira cancelar dentro de um função.

### Deadline com  o metodo WithCancel() 
Podemos mandar um sinal para todos os processos os  quais estão utilizando  o contexto para que cancelem ou Ignorem o sinal
# Obs: Continua a executar o contexto ate a chegada do sinal de cancelamento. Ou seja executo ate receber o cancel()

#### Deadline com metodo WithValue()
Este é contexto onde temos um valor no formato de objeto ou Map q depois de setado não muda, Ideal para síncronos e assíncronos
# Obs: Neste caso o retorno é context em forma de value






