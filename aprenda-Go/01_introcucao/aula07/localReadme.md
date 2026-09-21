#### Maps

## compações

## 3. Comparação com outras linguagens
# Linguagem	Estrutura Chave-Valor	Set (Valores Únicos)
# Go	map[K]V	Não tem nativo (simula com map[K]struct{})
# Python	dict	set
# JavaScript	Object / Map	Set
# Java	HashMap<K,V>	HashSet<T>
# C#	Dictionary<TKey,TValue>	HashSet<T>
# Rust	HashMap<K,V>	HashSet<T>

# 01 Maps Introdução

Utiliza o formato key:value.
E.g. nome e telefone
Performance excelente para lookups.
map[key]value{ key: value }
Acesso: m[key]
Key sem value retorna zero. Isso pode trazer problemas.
Para verificar: comma ok idiom.
v, ok := m[key]
ok é um boolean, true/false
Na prática: if v, ok := m[key]; ok { }
Para adicionar um item: m[v] = value
Maps não tem ordem.
Go Playground: https://play.golang.org/p/JXDdJan8Ev


# 02 Maps  range & deletando

Range: for k, v := range map { }
Reiterando: maps não tem ordem e um range usará uma ordem aleatória.
Go Playground: https://play.golang.org/p/6zEMfIP-AE
delete(map, key)
Deletar uma key não-existente não retorna erros!
Go Playground: https://play.golang.org/p/0uuIicU3Zz