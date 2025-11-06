Mini Kanban — Desafio Técnico Fullstack (React + Go)

Este projeto é um Mini Kanban desenvolvido como parte de um desafio técnico Fullstack, utilizando React no frontend e Go (Golang) no backend.
O objetivo é permitir que o usuário crie, edite, mova e exclua tarefas entre três colunas: A Fazer, Em Progresso e Concluídas — com persistência local em arquivo JSON.


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

🚀 Como Executar o Projeto
🧩 Pré-requisitos

Certifique-se de ter instalados:

Node.js
 (versão 16 ou superior)

Go (Golang)
 (versão 1.22 ou superior)

 🐹 1. Rodar o Backend (Go)

Abra o terminal na pasta backend:

cd backend


(Opcional) inicialize o módulo Go:

go mod init desafio-kanban-backend


Execute o servidor:

go run main.go handlers.go models.go


O backend iniciará em:

Servidor iniciado em :8080


➜ API disponível em http://localhost:8080

Endpoints disponíveis:

GET /tasks — lista todas as tarefas

POST /tasks — cria uma nova tarefa

PUT /tasks/{id} — atualiza uma tarefa

DELETE /tasks/{id} — remove uma tarefa

🗂️ As tarefas são salvas automaticamente no arquivo tasks.json para persistência local.

⚛️ 2. Rodar o Frontend (React)

Em outro terminal, vá até a pasta frontend:

cd frontend


Instale as dependências:

npm install


Inicie o servidor de desenvolvimento:

npm start


O app abrirá automaticamente em:
➜ http://localhost:3000

Certifique-se de que o backend está rodando em http://localhost:8080 — o frontend se comunica com essa API via HTTP.

💡 Rodar os dois simultaneamente (opcional)

Você pode usar o pacote concurrently no frontend para rodar os dois com um só comando:

npm install concurrently --save-dev


E adicionar no package.json do frontend:

"scripts": {
  "dev": "concurrently \"npm start\" \"cd ../backend && go run main.go handlers.go models.go\""
}


Depois é só rodar:

npm run dev

⚙️ Funcionalidades Principais

✅ Frontend (React)

Três colunas fixas: A Fazer, Em Progresso, Concluídas

Adicionar tarefas (título obrigatório, descrição opcional)

Editar, excluir e mover tarefas entre colunas

Interface intuitiva e responsiva

Drag and Drop funcional (com react-beautiful-dnd)

Feedbacks visuais de loading e erro

✅ Backend (Go)

Endpoints RESTful (GET, POST, PUT, DELETE)

Armazenamento em memória com persistência em arquivo JSON

Validação de título e status

Configuração de CORS para acesso pelo frontend

Código limpo e modular (main.go, handlers.go, models.go)

🧠 Decisões Técnicas

Frontend com React Hooks para simplicidade e clareza.

Armazenamento local das tarefas via API Go (arquivo tasks.json) — solução leve e funcional para o desafio.

react-beautiful-dnd para experiência fluida de drag-and-drop entre colunas.

CORS configurado no backend (Access-Control-Allow-Origin: *) para facilitar o desenvolvimento local.

Arquitetura limpa separando responsabilidades (UI, API, modelos, handlers).

🔄 Fluxo de Operação
1️⃣ User Flow (Diagrama Obrigatório)

O usuário:

Abre a aplicação (http://localhost:3000
)

Visualiza as 3 colunas do Kanban

Clica em + Nova Tarefa

Preenche título e descrição (opcional)

A tarefa aparece na coluna “A Fazer”

Pode arrastar para outras colunas, editar ou excluir

📄 Ver diagrama completo em /docs/user-flow.png

2️⃣ Data Flow (Opcional, mas incluso)

O Frontend (React) envia requisições REST (fetch) para o Backend (Go).

O Backend valida, atualiza em memória e salva no arquivo tasks.json.

O Frontend atualiza a UI com os dados retornados pela API.

Os dados persistem mesmo após reiniciar o servidor.

📄 Ver diagrama em /docs/data-flow.png

💾 Persistência dos Dados

As tarefas são armazenadas em memória e também salvas em backend/tasks.json para persistência simples.
Ao reiniciar o servidor, o arquivo é recarregado automaticamente.

🧰 Estrutura da API (exemplo de uso)
➤ Criar tarefa (POST)

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

🚧 Limitações Conhecidas

Sem autenticação (todas as ações são públicas).

Persistência local (arquivo JSON) não é adequada para produção.

Não há testes automatizados (ainda).

CORS liberado para todas as origens (ideal restringir em produção).

🔮 Melhorias Futuras

 Adicionar autenticação JWT

 Migrar persistência para banco de dados (ex: PostgreSQL ou SQLite)

 Implementar testes unitários e de integração

 Criar Docker Compose para subir frontend + backend com 1 comando

 Melhorar feedbacks visuais (toasts, loaders)

 Adicionar filtros e busca de tarefas

🧾 Documentação

📄 docs/user-flow.png — Diagrama obrigatório do fluxo do usuário
📄 docs/data-flow.png — Diagrama opcional mostrando a troca de dados entre frontend, backend e arquivo JSON

🧱 .gitignore

O projeto inclui um .gitignore configurado para:

Ignorar node_modules/, build/ e tasks.json

Ignorar arquivos de IDEs (.vscode/, .idea/)

Ignorar caches e logs temporários

🧩 Estrutura Recomendada no GitHub
desafio-fullstack-veritas/
├── backend/
├── frontend/
├── docs/
├── .gitignore
└── README.md

🤝 Autor

Arlen Filho
Desenvolvedor Frontend / Fullstack
💼 Projeto desenvolvido para avaliação técnica — Desafio Veritas
📧 Contato: [seu email aqui]

✅ Status do Projeto

✔️ Entregue — 100% dos requisitos atendidos (incluindo bônus)

CRUD completo

Persistência em arquivo JSON

Drag and Drop entre colunas

Documentação e diagramas