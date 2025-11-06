import React, { useEffect, useState } from "react";
import { getTasks, createTask, updateTask, deleteTask } from "./api";
import TaskCard from "./components/taskcard";
import TaskForm from "./components/taskform";

const STATUSES = [
  { key: "todo", label: "A Fazer" },
  { key: "inprogress", label: "Em Progresso" },
  { key: "done", label: "Concluídas" },
];

function App() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [showForm, setShowForm] = useState(false);
  const [editing, setEditing] = useState(null);

  async function load() {
    setLoading(true);
    setError(null);
    try {
      const data = await getTasks();
      setTasks(data);
    } catch (err) {
      setError(err.message || "Erro ao carregar");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    load();
  }, []);

  const handleCreate = async (payload) => {
    setLoading(true);
    try {
      const created = await createTask(payload);
      setTasks((prev) => [...prev, created]);
      setShowForm(false);
    } catch (err) {
      setError(err.message || "Erro ao criar");
    } finally {
      setLoading(false);
    }
  };

  const handleUpdate = async (id, payload) => {
    setLoading(true);
    try {
      const updated = await updateTask(id, payload);
      setTasks((prev) => prev.map((t) => (t.id === id ? updated : t)));
      setEditing(null);
      setShowForm(false);
    } catch (err) {
      setError(err.message || "Erro ao atualizar");
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Confirmar exclusão?")) return;
    setLoading(true);
    try {
      await deleteTask(id);
      setTasks((prev) => prev.filter((t) => t.id !== id));
    } catch (err) {
      setError(err.message || "Erro ao excluir");
    } finally {
      setLoading(false);
    }
  };

  const moveTask = async (task, newStatus) => {
    setLoading(true);
    try {
      const updated = await updateTask(task.id, { status: newStatus });
      setTasks((prev) => prev.map((t) => (t.id === task.id ? updated : t)));
    } catch (err) {
      setError(err.message || "Erro ao mover tarefa");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app">
      <header>
        <h1>Mini Kanban</h1>
        <div>
          <button onClick={() => { setShowForm(true); setEditing(null); }}>+ Nova Tarefa</button>
          <button onClick={load} className="secondary">Atualizar</button>
        </div>
      </header>

      {loading && <div className="info">Carregando...</div>}
      {error && <div className="error">Erro: {error}</div>}

      <main>
        {STATUSES.map((col) => (
          <section key={col.key} className="column">
            <h2>{col.label}</h2>
            <div className="column-body">
              {tasks.filter(t => t.status === col.key).map(task => (
                <TaskCard
                  key={task.id}
                  task={task}
                  onEdit={() => { setEditing(task); setShowForm(true); }}
                  onDelete={() => handleDelete(task.id)}
                  onMove={(to) => moveTask(task, to)}
                />
              ))}
            </div>
          </section>
        ))}
      </main>

      {showForm && (
        <TaskForm
          task={editing}
          onClose={() => { setShowForm(false); setEditing(null); }}
          onSubmit={(payload) => {
            if (editing) {
              handleUpdate(editing.id, payload);
            } else {
              handleCreate(payload);
            }
          }}
        />
      )}

      <footer>
        <small>API: http://localhost:8080 — Dados persistidos em backend (tasks.json)</small>
      </footer>
    </div>
  );
}

export default App;
