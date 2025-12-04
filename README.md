# goganizer

Um utilitário de linha de comando escrito em **Go**, projetado para **organizar arquivos automaticamente** com base em regras definidas pelo usuário.
O goganizer lê um conjunto de extensões e seus diretórios de destino e move os arquivos correspondentes para a pasta correta.

Ideal para organizar diretórios bagunçados, pastas de downloads, documentos, imagens e qualquer ambiente com alto volume de arquivos.

---

## 🚀 Funcionalidades

- Organização automática por extensão.
- Leitura de regras via arquivo JSON externo.
- Criação automática de diretórios.
- Suporte a arquivos sem extensão.
- Rápido, leve e multiplataforma (Linux, macOS, Windows).
- Não usa dependências externas.

---

## 📁 Estrutura do Projeto
```bash
goganizer/
│
├── main.go
├── rules/
│    └── rules.json
└── README.md
```

---

## ⚙️ Como Funciona

1. O programa lê o arquivo `rules/rules.json`.
2. Varre o diretório atual (ou outro indicado).
3. Identifica a extensão de cada arquivo.
4. Move o arquivo para o diretório correspondente.
5. Cria automaticamente diretórios que não existam.

---

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

## 📥 Instalação e Uso

### 1. Clone o repositório

```bash
git clone https://github.com/G-shiy/goganizer.git
cd goganizer
```
### 2. Compile o programa ou use o binário pré-compilado

```bash
go build -o goganizer main.go
```
### 3. Execute o programa

```bash
./goganizer
```
