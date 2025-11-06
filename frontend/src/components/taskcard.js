import React from "react";

const STATUS_LABEL = {
  todo: "A Fazer",
  inprogress: "Em Progresso",
  done: "Concluídas",
};

export default function TaskCard({ task, onEdit, onDelete, onMove }) {
  const canMoveLeft = task.status !== "todo";
  const canMoveRight = task.status !== "done";

  return (
    <div className="task-card">
      <div className="task-header">
        <strong>{task.title}</strong>
        <div className="task-actions">
          <button onClick={onEdit}>Editar</button>
          <button onClick={onDelete} className="danger">Excluir</button>
        </div>
      </div>
      {task.description && <p className="task-desc">{task.description}</p>}
      <div className="task-footer">
        <small>{STATUS_LABEL[task.status]}</small>
        <div>
          <button disabled={!canMoveLeft} onClick={() => onMove(task.status === "inprogress" ? "todo" : "inprogress")}>⮜</button>
          <button disabled={!canMoveRight} onClick={() => onMove(task.status === "todo" ? "inprogress" : "done")}>⮞</button>
        </div>
      </div>
    </div>
  );
}
