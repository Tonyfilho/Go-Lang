#### Aula02 Sobre Packages e importações como funciona

# 📁 Estrutura que vamos criar
pacotes/
├── go.mod                    ← módulo: pacotes
├── main.go                   ← package main
└── math/                     ← subpacote
    └── math.go               ← package math


# 2. Inicialize o módulo (cria o go.mod)
Dentro da pasta pacotes, execute:

bash
go mod init pacotes
Isso vai criar o arquivo go.mod com este conteúdo:

# 💡 O nome pacotes é o nome do módulo. Você pode usar qualquer nome, mas evite espaços e caracteres especiais. O ideal seria algo como github.com/seu-usuario/pacotes, mas para estudo local, pacotes já serve.

Dentro de pacotes/math/, crie o arquivo math.go:

# ⚠️ Atenção: O nome do pacote (package math) é o que você vai usar no import do main.go. Como o pacote se chama math, o caminho do import será pacotes/math


Podemos ter mais de uma arquivo no mesmo Packge, o Go vai acha-lo

## Obs: Podemos ter acesso, a Variaveis, Funções dentro dos packages

# Lembrando que inicio de nome de variaveis , funções e metodos  Upcase = public e DownCase = privite


# Documentação e uso do GoDoc
 Documentação em Go Ex:  da função criada, tem que começar com nome da Função iqualmente a declaração.
# Isto a biblioteca GoDoc reconhece este padrão e faz a documentação para nos
Ex:
// Soma faz a soma de 2 numeros e retorna um int
func Soma(a, b int) int {

	return a + b
}
