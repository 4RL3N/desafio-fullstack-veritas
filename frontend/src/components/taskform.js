import React, { useState, useEffect } from "react";

export default function TaskForm({ task, onClose, onSubmit }) {
  const [title, setTitle] = useState(task ? task.title : "");
  const [description, setDescription] = useState(task ? task.description : "");
  const [status, setStatus] = useState(task ? task.status : "todo");
  const [error, setError] = useState(null);

  useEffect(() => { setError(null); }, [title, description]);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!title.trim()) {
      setError("Título é obrigatório.");
      return;
    }
    onSubmit({ title: title.trim(), description: description.trim(), status });
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h3>{task ? "Editar Tarefa" : "Nova Tarefa"}</h3>
        {error && <div className="error">{error}</div>}
        <form onSubmit={handleSubmit}>
          <label>
            Título *
            <input value={title} onChange={(e) => setTitle(e.target.value)} />
          </label>
          <label>
            Descrição
            <textarea value={description} onChange={(e) => setDescription(e.target.value)} />
          </label>
          <label>
            Status
            <select value={status} onChange={(e) => setStatus(e.target.value)}>
              <option value="todo">A Fazer</option>
              <option value="inprogress">Em Progresso</option>
              <option value="done">Concluídas</option>
            </select>
          </label>
          <div className="form-actions">
            <button type="submit">{task ? "Salvar" : "Criar"}</button>
            <button type="button" onClick={onClose} className="secondary">Fechar</button>
          </div>
        </form>
      </div>
    </div>
  );
}
