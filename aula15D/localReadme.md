#### Aula15D UnMarshal E Marshal com as Funções NewDecoder NewEncoder# Continuação do video depois do minuto 10
https://www.youtube.com/watch?v=mcbj-wy8Ro8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116
https://cs.opensource.google/go/go/+/refs/tags/go1.27.0:src/encoding/json/v2_stream.go;l=38


Marshal/unmarshal vs. encoder/decoder
Marshal vai pra uma variável
Encoder "vai direto"

## Resumo, todo Encoder que eu fizer é para IR direto para Interface os.Stdout, sem variaveis
	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(toniFilho) //{"Nome":"Tony","SobreNome":"Filho","Idade":40,"Profissao":"Developer","ContaBancaria":1000}

Com Encoder: https://play.golang.org/p/Pgwr0O07aL

## Aqui podemos fazer Encoder de forma mais dinamica

func NewDecoder(r io.Reader) *Decoder

func NewDecoder(r io.Reader) *Decoder
NewDecoder retorna um novo decodificador que lê de r.

O decodificador introduz seu próprio buffer e pode ler dados de r além dos valores JSON solicitados.

func (*Decoder) Buffered ¶
adicionado em go1.1
func (dec *Decoder) Buffered() io.Reader
Buffered retorna um leitor dos dados restantes no buffer não lido, que pode conter zero ou mais bytes. Esses são os dados já consumidos da entrada io.Reader, mas ainda não lidos por uma chamada Decoder.Decode ou Decoder.Token. Podem conter bytes que não formam um JSON válido, pois ainda não foram validados de acordo com a gramática JSON. A quantidade exata de dados em buffer é um detalhe de implementação do Decoder e pode mudar ao longo do tempo.

É responsabilidade do chamador concatenar este buffer com o restante do leitor de entrada para obter a sequência completa de bytes após o último valor JSON decodificado.

O leitor é válido até a próxima chamada a `Decoder.Decode` ou `Decoder.Token`.

func (*Decoder) Decode ¶
func (dec *Decoder) Decode(v any) error
Decode lê o próximo valor codificado em JSON de sua entrada e o armazena no valor apontado por `v`.

Consulte a documentação de `Unmarshal` para obter detalhes sobre a conversão de JSON em um valor Go.


