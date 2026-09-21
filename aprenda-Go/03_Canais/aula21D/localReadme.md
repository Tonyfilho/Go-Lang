#### Aula21D  Canais – 3. Range e close
https://www.youtube.com/watch?v=B1UArMoYDJ0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=152
Range:
gofunc com for loop com send e close(chan)
recebe com range chan

Código: https://play.golang.org/p/_g5IEjSkh1

# O problema está na execução das goroutines
Seu código não imprime os valores porque as goroutines não têm tempo para executar antes do programa terminar.

# O que são goroutines?
São "threads leves" que executam em paralelo

Quando main termina, todas as goroutines são mortas

## Linha do Tempo → Sem WaitGroup 
## Tempo →
# main: inicia → cria canal → inicia goroutine1 → inicia goroutine2 → FIM
goroutine1: inicia → envia 0 → envia 1 → ... (não termina)
goroutine2: inicia → espera dados do canal → ... (não termina)

## Linha do Tempo → Com WaitGroup 
## Tempo →
# main: inicia → cria canal → inicia goroutine1 → inicia goroutine2 → espera (WAIT) 
goroutine1: inicia → envia 0 → envia 1 → ... → close → done!
goroutine2: inicia → recebe 0 → recebe 1 → ... → range termina → done!
                                                                      
main: wait desbloqueia → FIM (só depois que ambas terminaram)