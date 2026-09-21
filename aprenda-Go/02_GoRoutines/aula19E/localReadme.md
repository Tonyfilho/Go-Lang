#### Aula19E  Concorrência –  5. Mutex
https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=129
https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/18_concorrencia/06_mutex/main.go
# https://pkg.go.dev/sync#Mutex


Agora vamos resolver a race condition do programa anterior utilizando mutex.
# Mutex é mutual exclusion, exclusão mútua.
# Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
Na prática:
type Mutex
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
RWMutex

# Um Mutex é um bloqueio de exclusão mútua. O valor zero para um Mutex representa um mutex desbloqueado.

Um Mutex não deve ser copiado após o primeiro uso.

Na terminologia do modelo de memória do Go, a n-ésima chamada a `Mutex.Unlock` "sincroniza antes" da m-ésima chamada a `Mutex.Lock` para qualquer n < m. Uma chamada bem-sucedida a `Mutex.TryLock` é equivalente a uma chamada a `Lock`. Uma chamada malsucedida a `TryLock` não estabelece nenhuma relação de "sincroniza antes".


