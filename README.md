# Mini Kanban — Desafio Técnico Fullstack (React + Go)

Este projeto é um **Mini Kanban** desenvolvido como parte de um **desafio técnico Fullstack**, utilizando **React** no frontend e **Go (Golang)** no backend.
O objetivo é permitir que o usuário crie, edite, mova e exclua tarefas entre três colunas: A Fazer, Em Progresso e Concluídas — com **persistência local em arquivo JSON**.

```
├── backend/              # API REST em Go
│   ├── main.go
│   ├── handlers.go
│   ├── models.go
│   └── tasks.json        # Gerado automaticamente
│
├── frontend/             # Aplicação React
│   ├── package.json
│   ├── public/
│   │   └── index.html
│   └── src/
│       ├── App.js
│       ├── api.js
│       ├── index.js/
│       ├── components/
│       │   ├── TaskCard.js
│       │   └── TaskForm.js
│       └── styles.css
│
├── docs/                 # Documentação
│   ├── user-flow.png     # Diagrama obrigatório
│   └── data-flow.png     # Diagrama opcional
│
└── README.md             # Este arquivo
```
---

## 🚀 Como Executar o Projeto
### 🧩 Pré-requisitos

Certifique-se de ter instalados:

```
- **Node.js** (versão 16 ou superior)

- **Go (Golang)** (versão 1.22 ou superior)
```
---

 ### 🐹 1. Rodar o Backend (Go)

Abra o terminal na pasta `backend` e inicialize o módulo (apenas na primeira vez):

```
go mod init desafio-kanban-backend
```

Isso cria o arquivo go.mod e prepara o ambiente Go.

### ⚛️ 2. Rodar o Frontend (React)

Agora, vá até a pasta frontend:
```
cd frontend
```

Instale as dependências:
```
npm install
```

Instale uma ferramenta para rodar o backend e o frontend ao mesmo tempo.
```
npm install concurrently --save-dev
```

Depois é só rodar:
```
npm run dev
```

 - O frontend abrirá automaticamente em http://localhost:3000
 - Certifique-se de que o backend está ativo em http://localhost:8080
---

## ⚙️ Funcionalidades Principais

#### ✅ Frontend (React)

 - Três colunas fixas: A Fazer, Em Progresso, Concluídas
 - Adicionar tarefas (título obrigatório, descrição opcional)
 - Editar, excluir e mover tarefas entre colunas
 - Interface intuitiva e responsiva
 - Drag and Drop funcional (`com react-beautiful-dnd`)

### ✅ Backend (Go)

 - Endpoints RESTful (`GET, POST, PUT, DELETE`)
 - Armazenamento em memória com persistência em arquivo JSON
 - Validação de título e status
 - Configuração de CORS para acesso pelo frontend
 - Código limpo e modular (`main.go, handlers.go, models.go`)

---
## 🧠 Decisões Técnicas

 - Frontend com React Hooks para simplicidade e clareza.
 - Armazenamento local das tarefas via API Go (arquivo `tasks.json`).
 - `react-beautiful-dnd` para experiência fluida de drag-and-drop entre colunas.
 - CORS configurado no backend (`Access-Control-Allow-Origin: *`) para facilitar o desenvolvimento local.
 - Arquitetura limpa separando responsabilidades (UI, API, modelos, handlers).

## 💾 Persistência dos Dados

As tarefas são armazenadas em memória e também salvas em `backend/tasks.json` para persistência simples.
 - Ao reiniciar o servidor, o arquivo é recarregado automaticamente.
 ---
 
## 🧰 Estrutura da API (exemplo de uso)
#### ➤ Criar tarefa (POST)
```
Request:

{
  "title": "Estudar React",
  "description": "Revisar hooks e componentes",
  "status": "todo"
}


Response:

{
  "id": "1730482312323123",
  "title": "Estudar React",
  "description": "Revisar hooks e componentes",
  "status": "todo",
  "createdAt": "...",
  "updatedAt": "..."
}
```
---

## 🔮 Melhorias Futuras

 - Adicionar autenticação JWT
  - Migrar persistência para banco de dados (ex: PostgreSQL ou SQLite)
  - Implementar testes unitários e de integração
  - Criar Docker Compose para subir frontend + backend com 1 comando
  - Melhorar feedbacks visuais (toasts, loaders)
  - Adicionar filtros e busca de tarefas
---

## 🧾 Documentação

 - 📄 `docs/user-flow.png` — Diagrama obrigatório do fluxo do usuário
 - 📄 `docs/data-flow.png` — Diagrama opcional mostrando a troca de dados entre frontend, backend e arquivo JSON
---

## 🤝 Autor

 - Arlen Filho
 - Desenvolvedor Frontend / Fullstack
 - 💼 Projeto desenvolvido para avaliação técnica — Desafio Veritas
 - 📧 Contato: arlen0filho04@gmail.com