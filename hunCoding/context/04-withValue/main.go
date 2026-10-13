package main

import (
	"context"
	"fmt"
)

func main() {
	/**como não temos contexto real, vamos usar o contexto background */
	/**Podemo usar o background() e todo() pois são conext vazios */
	/**A Grande diferença entre WithDeadline e WithTimeout é que o WithTimeout recebe um tempo de duração */

    /***Aqui seria como criamos uma variavel onde fosse algo parecido com Map */
	/**Ou seja  o valor é o any data typescript*/
 var test map[ interface{}]interface{} = make(map[interface{}]interface{})
	test["name"] = "Tony"
	test["age"] = 30	
/*	é isto que temos no WithValue() ele cria um contexto com valores que podemos passar para outras funções */
ctx := context.WithValue(context.Background(), "minhaChave", 20)

	/**Chamamos a função printUntilCanceled passando o contexto*/
	printUntilCanceled(ctx)

/**Lembrando que esta valor é consistente e não muda, pode ser usado em várias partes 
do código e pode somente ser lido ou clonado e adicionado*/
/**Podemos pegar o valor do contexto usando o método Value() pela CHAVE*/
fmt.Println("Valor do contexto", ctx.Value("minhaChave"))

/**Clone o contexto e adicione um novo valor*/
/**Nesta caso observe que não estou usando Background(), pois ja temos um contexto */
ctx2 := context.WithValue(ctx, "outraChave", 	40)
fmt.Println("Valor do contexto clonado", ctx2.Value("outraChave"))







}

/*Criamores um função q será um go routine no main**/

func printUntilCanceled(ctx context.Context) {
	
		fmt.Println("Printing withtimeout..", ctx.Value("minhaChave"))
	

}
