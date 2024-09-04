# Usando a imagem oficial do Golang
FROM golang:1.23

# Definindo o diretório de trabalho dentro do container
WORKDIR /app

# Copiando o código-fonte do projeto
COPY . .

# Compilando o projeto
RUN go build -o main .

# Comando para rodar a aplicação
CMD ["./main"]
