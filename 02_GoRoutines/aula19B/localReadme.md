#### Aula19B Concorrência – 2. Goroutines & WaitGroups 
# Continuação
https://www.youtube.com/watch?v=4jXSU2jw3Ag&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=126

Exemplo: código termina antes da go func executar.
Ou seja, precisamos de uma maneira pra "sincronizar" isso.
Ah, mas então... não.
Qualé então? sync.WaitGroup:
Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
func Add: "Quantas goroutines?"
func Done: "Deu!"
func Wait: "Espera todo mundo terminar."
Ah, mas então... sim!
Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()

Go Playground: https://play.golang.org/p/8iiqLX4sWc

# Como ja vimos o main() terminou antes da goroutine go Func01() e foi descartada e precisamos usar o WaitGroup  para controlar a execução



