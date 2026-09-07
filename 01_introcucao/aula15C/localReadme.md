#### Aula15 C UnMarshal (Desordenando) Json
# De Json para GO
https://www.youtube.com/watch?v=mcbj-wy8Ro8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116
https://pkg.go.dev/encoding/json#Unmarshal
https://mholt.github.io/json-to-go/

E agora o contrário.
JSON-to-Go
Marshal/unmarshal vs. encoder/decoder
Marshal vai pra uma variável
Tags
Encoder "vai direto"
# `json:"Nome"`  Isto são as Tags Encoder
Ex de uso das Tags Encoder no Campo Profissão que recebe Trabalho, poderia ser o contrario tb
Go Playground: https://play.golang.org/p/l6wbuLu1NS
Com Encoder: https://play.golang.org/p/Pgwr0O07aL

## Precisamos Criar uma função de UnMarshal
# O retorno vem Nulo ou em um Ponteiro `v` , tem que haver os mesmo campos, caso contrario teremos error `errors.ErrUnsupported`

func Unmarshal(data []byte, v any) error

A função `Unmarshal` analisa os dados codificados em JSON e armazena o resultado no valor apontado por `v`. Se `v` for nulo ou não for um ponteiro, `Unmarshal` retorna um erro `InvalidUnmarshalError`.

A função `Unmarshal` utiliza o inverso das codificações usadas por `Marshal`, alocando mapas, fatias e ponteiros conforme necessário, com as seguintes regras adicionais:

Para desserializar JSON em um ponteiro, `Unmarshal` primeiro trata o caso em que o JSON é um literal JSON nulo. Nesse caso, `Unmarshal` define o ponteiro como nulo. Caso contrário, `Unmarshal` desserializa o JSON para o valor apontado pelo ponteiro. Se o ponteiro for nulo, `Unmarshal` aloca um novo valor para ele apontar.

A entrada JSON é decodificada de acordo com as seguintes regras:

Se o tipo de valor implementar `jsonv2.UnmarshalerFrom`, o método `UnmarshalJSONFrom` será chamado para decodificar o valor JSON. Se o método retornar `errors.ErrUnsupported`, a entrada será decodificada de acordo com as regras subsequentes.

Se o tipo de valor implementar `Unmarshaler`, o método `UnmarshalJSON` será chamado para decodificar o valor JSON, inclusive quando a entrada for um JSON nulo.

Se o valor implementar `encoding.TextUnmarshaler` e a entrada for uma string JSON, o método `UnmarshalText` será chamado com a string sem aspas.