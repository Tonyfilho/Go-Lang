#### aula16 A Interface WRITER
https://www.youtube.com/watch?v=S4hEdA0RPVI&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116

https://pkg.go.dev/io#Writer

## A interface writer do pacote io.
# OBS O Tipo File Implementa a Interface Writer, por isto que Escrevemos
## Recebe um Slice de [] Bytes e retorna um Numero e um Erro quando houver, senão um nil
# Para implementar precisamos de uma Função que recebe o Metodo Write(p []byte) (n int, err error)
type Writer interface {
	Write(p []byte) (n int, err error)
}

# type Writer interface { Write(p []byte) (n int, err error) } 
Recebemos aqui um WRITE 
pkg os:   func (f *File) Write(b []byte) (n int, err error)

# Ex: pkg json: func NewEncoder(w io.Writer) *Encoder Esta função recebe um io.WRITER e retorna um Ponteiro de *Encoder

"Println [...] writes to standard output."
func Println [...] return Fprintln(os.Stdout, a...)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
Stdout: NewFile(uintptr(syscall.Stdout), "/dev/stdout") (Google: Standard streams)
func NewFile(fd uintptr, name string) *File
func (f *File) Write(b []byte) (n int, err error)
Exemplo:
Println
Fprintln os.Stdout
io.WriteString os.Stdout
Ou:
func Dial(network, address string) (Conn, error)
type Conn interface { [...] Write(b []byte) (n int, err error) [...] }

Resumo
1º Temos a Interface Write que tem o Metodo Write type Writer interface { Write(p []byte) (n int, err error) }
2º temos uma função: func (f *File) Write(b []byte) (n int, err error) que implementa este metodo       da Interface Write
3º Por fim temos uma Função que recebe dados pela Interface io.Write e retorna um Pondeiro      : func NewEncoder(w io.Writer) *Encoder
A Grande questão é que esta Interface não é so usada Imprimir, é Implementada em Muita coisa em GO!
Ex: É a Função de Conexão func DIAL que implementa a Interface Conn , mas a  interface Conn implementa o metodo WRITE:
Função func Dial(network, address string) (Conn, error)
O Seja TUDO que tem o METODO WRITE pode ser recebido como ARGUMENTO
Ate uma conexão de REDE é passado como WRITE


