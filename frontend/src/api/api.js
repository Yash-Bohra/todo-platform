import axios from "axios"

const API = axios.create({
  baseURL: "http://localhost:8080"
})

export const registerUser = (data) => API.post("/auth/register", data)

export const loginUser = (data) => API.post("/auth/login", data)

export const getTodos = (token) =>
  API.get("/todos/", {
    headers: { Authorization: `Bearer ${token}` }
  })

export const createTodo = (todo, token) =>
  API.post("/todos/", todo, {
    headers: { Authorization: `Bearer ${token}` }
  })

export const deleteTodo = (id, token) =>
  API.delete(`/todos/${id}`, {
    headers: { Authorization: `Bearer ${token}` }
  })