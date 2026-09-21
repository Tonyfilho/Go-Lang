# 01 Strings são sequencias de bytes.

String são Imutáveis.
Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
Na prática:
%v %T
Raw string literals
Conversão para slice of bytes: []byte(x)
%#U, %#x
Go Playground: https://play.golang.org/p/dt2x1ies5b & https://play.golang.org/p/PpDnspiyA_7
https://blog.golang.org/strings

OBS: cada item é um byte


# 02 Constantes

São valores imutáveis.
Podem ser tipadas ou não:
const oi = "Bom dia"
const oi string = "Bom dia"
As não tipadas só terão um tipo atribuido a elas quando forem usadas.
Ex. qual o tipo de 42? int? uint? float64?
Ou seja, é uma flexibilidade conveniente.
Na prática: int, float, string.
const x = y
const ( x = y )

# 03 IOTA

golang.org/ref/spec
Numa declaração de constantes, o identificador iota representa números sequenciais.
Na prática.
iota, iota + 1, a = iota b c, reinicia em cada const, _
Go Playground: https://play.golang.org/p/eSrwoQjuYR

# 04 Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
https://play.golang.org/p/7MOnbhx4R4
  / bit-hacking-with-go  
Fim da sessão. Massa!