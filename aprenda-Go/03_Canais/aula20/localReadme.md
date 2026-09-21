##### Aula20 Package Organizations  8. Pacotes
## OBS tem q dar o "go mod init"

https://www.youtube.com/watch?v=SO-RFPSqD3c&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=138
https://rakyll.org/style-packages/ 

https://github.com/vkorbes/aprendago/tree/master/c%C3%B3digo/19_seu-ambiente-de-desenvolvimento/pacotes


Opção 1: uma pasta, vários arquivos.
package declaration em todos os arquivos
package scope: um elemento de um arquivo é acessível de todos os arquivos
imports tem file scope
Opção 2: separando por packages.
pastas diferentes
requer imports
para usar: package.Função()
Exportado vs. não-exportado, ou seja, visível vs. não-visível
Em Go não utilizamos os termos "público" e "privado" como em outras linguagens
É somente questão de capitalização
Com maiúscula: exportado, visível fora do package
Com minúscula: não exportado, não utilizável fora do package