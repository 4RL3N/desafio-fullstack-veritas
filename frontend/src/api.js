const API_BASE = process.env.REACT_APP_API_BASE || "http://localhost:8080";

async function request(path, options = {}) {
  const res = await fetch(`${API_BASE}${path}`, {
    headers: { "Content-Type": "application/json" },
    ...options,
  });
  if (!res.ok && res.status !== 204) {
    const text = await res.text();
    throw new Error(text || res.statusText || "Erro na requisição");
  }
  if (res.status === 204) return null;
  return res.json();
}

export const getTasks = () => request("/tasks");
export const createTask = (payload) =>
  request("/tasks", { method: "POST", body: JSON.stringify(payload) });
export const updateTask = (id, payload) =>
  request(`/tasks/${id}`, { method: "PUT", body: JSON.stringify(payload) });
export const deleteTask = (id) =>
  request(`/tasks/${id}`, { method: "DELETE" });
