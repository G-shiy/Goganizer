# 🐹 Goganizer

Um utilitário de linha de comando escrito em **Go**, projetado para **organizar arquivos automaticamente** com base em regras definidas pelo usuário.  
O Goganizer lê um conjunto de extensões e seus diretórios de destino e move os arquivos correspondentes para a pasta correta.

Ideal para organizar diretórios bagunçados, downloads, documentos, imagens e qualquer ambiente com alto volume de arquivos.


## 🚀 Funcionalidades

- Organização automática por extensão.
- Leitura de regras via arquivo JSON externo ou embutido no binário.
- Criação automática de diretórios quando necessário.
- Suporte a arquivos sem extensão.
- Rápido, leve e multiplataforma (Linux, macOS, Windows).
- Sem dependências externas.

## 📁 Estrutura do Projeto

```bash
goganizer/
│
├── main.go
├── handlers/
│    └── rules.go
├── rules/
│    └── rules.json
└── README.md
```

## ⚙️ Como Funciona

1. O programa lê o arquivo `rules/rules.json` (ou usa regras embutidas, se compilado assim).  
2. Varre o diretório atual (ou outro indicado como argumento).  
3. Identifica a extensão de cada arquivo.  
4. Move o arquivo para o diretório correspondente.  
5. Cria automaticamente diretórios que não existam.  
## 📜 Exemplo de `rules.json`

```json
{
  ".txt": "text",
  ".md": "markdown",
  ".pdf": "pdf",
  ".docx": "docs",
  ".csv": "spreadsheets",
  ".xlsx": "spreadsheets",
  ".png": "images",
  ".jpg": "images",
  ".jpeg": "images",
  ".mp4": "videos",
  ".mp3": "audio",
  ".zip": "archives",
  ".rar": "archives",
  "": "no_extension"
}
```

## 📥 Instalação

### 1. Via `go install` (recomendado)

Se você já tem o Go instalado:

```bash
go install github.com/G-shiy/goganizer/cmd/goganizer@latest
```

> O binário será instalado automaticamente no diretório `$GOBIN` (ou `$GOPATH/bin`) e poderá ser executado de qualquer lugar.

Para rodar na pasta atual:

```bash
goganizer
```

Você também pode organizar outro diretório passando o caminho como argumento:

```bash
goganizer /caminho/para/pasta
```

---

### 2. Compilando manualmente

Clone o repositório:

```bash
git clone https://github.com/G-shiy/goganizer.git cd goganizer
```

Compile o programa:

```bash
go build -o goganizer main.go
```

E execute:

```bash
./goganizer
```

> No Windows, o executável será `goganizer.exe`.

---

## 💡 Dicas

- Mantenha seu `rules.json` atualizado com as extensões que você mais utiliza.
    
- O Goganizer cria as pastas automaticamente apenas se elas ainda não existirem.
    
- Arquivos sem extensão podem ser organizados em uma pasta específica (`no_extension` por padrão).****
