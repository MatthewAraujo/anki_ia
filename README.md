# Projeto Anki PDF - Gerador de Perguntas e Alternativas

Este é um projeto que visa processar arquivos PDF, extrair seu conteúdo e gerar perguntas com alternativas baseadas no conteúdo extraído. Utiliza o modelo de linguagem ChatGPT para gerar as perguntas e alternativas, as quais são armazenadas em um banco de dados para posterior uso.

## Funcionalidades

- **Upload de Arquivos PDF**: O sistema permite que os usuários façam upload de arquivos PDF.
- **Extração de Texto**: O texto é extraído do PDF utilizando bibliotecas específicas.
- **Geração de Perguntas e Alternativas**: O ChatGPT é utilizado para gerar perguntas com alternativas baseadas no conteúdo extraído.
- **Armazenamento em Banco de Dados**: As perguntas e alternativas geradas são armazenadas no banco de dados.
- **Suporte a Vários Tipos de Pergunta**: O sistema suporta perguntas de múltipla escolha, com opções de resposta.
  
## Tecnologias Utilizadas

- **Go**: Linguagem principal utilizada para o backend.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar PDFs, perguntas e alternativas.
- **ChatGPT/OpenAI**: Utilizado para gerar perguntas e alternativas a partir do texto extraído dos PDFs.
- **GitHub Actions**: CI/CD para automação do deploy e testes.
  
## Arquitetura

- **PDFs**: São carregados no sistema e processados para extração de texto.
- **Perguntas**: São geradas dinamicamente a partir do conteúdo extraído.
- **Alternativas**: São associadas às perguntas e armazenadas com a informação de qual alternativa é a correta.
- **Banco de Dados**: Armazena os PDFs, perguntas e alternativas associadas.

## Estrutura de Banco de Dados

### Tabelas

- **users**: Tabela de usuários que possuem permissão para subir e acessar os PDFs e suas respectivas perguntas.
- **pdfs**: Tabela que armazena as informações sobre os PDFs, incluindo status de processamento e conteúdo extraído.
- **questions**: Tabela que armazena as perguntas geradas a partir do conteúdo dos PDFs.
- **options**: Tabela que armazena as alternativas associadas às perguntas.

```sql
-- Tabela de usuários
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de PDFs
CREATE TABLE pdfs (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    filename VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status ENUM('pending', 'processed', 'failed') DEFAULT 'pending',
    text_content TEXT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Tabela de perguntas
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    pdf_id INT NOT NULL,
    question_text TEXT NOT NULL,
    FOREIGN KEY (pdf_id) REFERENCES pdfs(id) ON DELETE CASCADE
);

-- Tabela de opções de respostas
CREATE TABLE options (
    id SERIAL PRIMARY KEY,
    question_id INT NOT NULL,
    option_key CHAR(1) NOT NULL, -- A, B, C, D
    option_text TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);
```
## Fluxo de Trabalho

1. O usuário faz upload de um arquivo PDF.
2. O sistema extrai o texto do PDF.
3. O texto extraído é enviado ao ChatGPT, que gera as perguntas e alternativas.
4. As perguntas e alternativas são inseridas no banco de dados.
5. O usuário pode consultar as perguntas e respostas geradas a partir do PDF.
