#### Aula21B 2. Canais direcionais & Utilizando canais
https://www.youtube.com/watch?v=vYYHoKLb_8I&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=148

Canais podem ser direcionais.
E isso serve pra...?
Um send channel e um receive channel são tipos diferentes. Isso permite que os type-checking mechanisms do compilador façam com que não seja possível, por exemplo, escrever num canal de leitura.
https://stackoverflow.com/questions/13596186/whats-the-point-of-one-way-channels-in-goCanais 
## bidirecionals (send & receive)
# send chan←
error: "invalid operation: ←cs (receive from send-only type chan← int)"
# receive ←chan
error: "invalid operation: cr ← 42 (send to receive-only type ←chan int)"
Exemplo: https://play.golang.org/p/TlcSm8bHkW
A seta sempre aponta para a esquerda.
