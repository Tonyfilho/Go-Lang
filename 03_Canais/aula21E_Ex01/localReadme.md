#### Aula21E   Canais – 4. Select 
https://www.youtube.com/watch?v=dp8s5jAc7h0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=149

# Select é como switch, só que pra canais, e não é sequencial.

"A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
Na prática:
# Exemplo 1:
Duas go funcs enviando X/2 numeros cada uma pra um canal
For loop X valores, select case ←x
Go Playground:
1. https://play.golang.org/p/xC3e1wBxgv






