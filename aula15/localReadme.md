

#### Aula15 Documentação JSON
https://www.youtube.com/watch?v=jnnIgvV0_yA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=113

https://pkg.go.dev/encoding/json
https://go.dev/blog/json
https://pkg.go.dev/encoding/json#example-Marshal
# temos o  que converte Json para Go 
https://mholt.github.io/json-to-go/ 


Já entendemos ponteiros, já entendemos métodos. Já temos o conhecimento necessário para começar a utilizar a standard library.
Nesse vídeo faremos uma orientação sobre como abordar a documentação.
Essa aula não foi preparada. Vai ser tudo ao vivo no improviso pra vocês verem como funciona o processo.
golang.org → Documents → Package Documentation 
godoc.org → encoding/json
files
examples
funcs
types
methods

type Message struct {
    Name string
    Body string
    Time int64
}

m := Message{"Alice", "Hello", 1294706395881547000}

we can marshal a JSON-encoded version of m using json.Marshal:

# Nesta instancia recebemos o Bite e o Error
b, err := json.Marshal(m)

# Decoding A Func retorna mutiliplos valores, Data e Error
To decode JSON data we use the Unmarshal function.
func Unmarshal(data []byte, v interface{}) error

# We must first create a place where the decoded data will be stored
var m Message

# and call json.Unmarshal, passing it a []byte of JSON data and a pointer to m
err := json.Unmarshal(b, &m)

# If b contains valid JSON that fits in m, after the call err will be nil and the data from b will have been stored in the struct m, as if by an assignment like:
m = Message{
    Name: "Alice",
    Body: "Hello",
    Time: 1294706395881547000,
}

### Relembrando o operador NIL
## Operador NIL em GO
# Em resumo: nil em Go é um valor válido que indica "não aponta para nada", enquanto undefined em JS é a ausência de valor propriamente dita. São conceitos diferentes com propósitos diferentes!

Go (nil apenas para):
Ponteiros (*T)
Slices ([]T)
Maps (map[T]T)
Channels (chan T)
Functions
Interfaces