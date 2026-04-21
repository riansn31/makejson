# Go JSON File Manager

Este projeto é uma aplicação simples em Go (Golang) que demonstra como interagir com o terminal para receber dados do utilizador, processá-los em formato JSON e gerir ficheiros no sistema operativo.

## 🚀 Funcionalidades

- **Interação via Terminal:** Solicita o nome e o endereço do utilizador.
- **Processamento JSON:** Converte os dados recolhidos num objeto JSON serializado.
- **Persistência Local:** Guarda as informações num ficheiro chamado `data.json`.
- **Leitura e Output:** Lê o conteúdo do ficheiro guardado e exibe-o no ecrã para verificação.

## 🛠️ Tecnologias Utilizadas

- **Go (Golang):** Linguagem de programação.
- **Pacote `encoding/json`:** Para a codificação dos dados.
- **Pacote `os`:** Para operações de escrita e leitura de ficheiros.
- **Pacote `fmt`:** Para entrada e saída de dados formatada.

## 📋 Pré-requisitos

Para correr este projeto, necessita de ter o Go instalado. Pode verificar a instalação correndo:

```bash
go version
```

## 🔧 Como utilizar

1. Crie um ficheiro chamado `main.go` e cole o código fonte.
2. Abra o terminal na pasta onde guardou o ficheiro.
3. Execute o comando:

```bash
go run main.go
```

Insira o seu nome e endereço quando solicitado.

## 📂 Exemplo de Saída (`data.json`)

O ficheiro gerado terá um formato semelhante a este:

```json
{
  "address": "Rua Principal 123",
  "name": "Utilizador"
}
```

## 📝 Notas Importantes

- **Captura de Dados:** O uso de `fmt.Scanln` tem uma limitação: ele termina a leitura ao encontrar o primeiro espaço. Para capturar frases completas (ex: nomes, apelidos ou moradas longas), recomenda-se o uso do `bufio.NewScanner`.
- **Permissões:** O ficheiro é criado com a permissão `0644`, que é o padrão seguro para ficheiros de texto legíveis no sistema.

---

_Desenvolvido para fins de aprendizagem sobre manipulação de I/O em Go._
