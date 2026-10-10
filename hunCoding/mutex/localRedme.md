#### Aula de Mutex e explicação de uso entre Mutex e Chanel

# Observei que na versão atual ja corrigiu o problema de falta de mutex

https://github.com/HunCoding

https://www.youtube.com/watch?v=JUlz804xgtM


# Um mutex é como uma chave de banheiro em um posto de gasolina, garantindo que apenas uma pessoa possa usar o banheiro de cada vez E que ninguém mais possa usar o banheiro até que o ocupante atual termine e a chave seja devolvida. 

No ex: sem Mutex o valor do contador nunca será o mesmo, pois cada Go Routines terá um valor, é justamente o que não queremos. Isto gera um problema. Veja o exemplo

Quando usamos o Mutex, somente 1 Go Routine por vez irá poder ler e gravar na Go Routine.
Temos a opçáo do podemos deixar mais de um Go Routine LER ao mesmo tempo.

# O Processo é Padrão da Seção ou Bloco Critica:
por Lock() ANTES
A variavel ou Metodo ou Função que vai ler ou gravar na variavel 
por Unlock() DEPOIS.


# O Processo usando Defer é
por Lock() ANTES
defer Unlock() que vai ser rodados no final ou quando terminar 
A variavel ou Metodo ou Função que vai ler ou gravar na variavel 

# Sync RMutex RLock()

RMutex permite que mais de uma Go Routine passa LER ao mesmo tempo a Variavel que esteja em Lock()

Obs: Somente 1 Go Routine consegue escrever por vez, Ler caso seja usado RMutex será permitido, mas Gravar ou escrever não.

O problema que não saberá se as go routine terá o valor correto da variavel ou não.
