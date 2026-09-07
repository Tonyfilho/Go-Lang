#### Aula18 BCrypt

https://www.youtube.com/watch?v=4vCb7jmwkzM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=120
https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/16_aplicacao/bcrypt/main.go

É uma maneira de encriptar senhas utilizando hashes.
x/crypto/bcrypt
GenerateFromPassword
CompareHashAndPassword
# Sem Go Playground! No terminal, dentro da pasta do seu projeto, execute:

# Tem que criar o Modulo
# 1. Inicialize o módulo na pasta atual
go mod init aula18

# 2. Agora instale o bcrypt
go get golang.org/x/crypto/bcrypt

# 3. Seu arquivo go.mod será criado/atualizado