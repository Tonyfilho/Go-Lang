#### Aula19D  Concorrência – 4. Na prática: Condição de corrida

https://www.youtube.com/watch?v=XxG7qqJzDKk&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=128

Aqui vamos replicar a race condition mencionada no artigo anterior.
time.Sleep(time.Second) vs. runtime.Gosched()
go help → go help build → go run -race main.go
Como resolver? Mutex. 

# Resulmo vais goroutines leram uma variavel compartilhada e não conseguiram salvar na variavel