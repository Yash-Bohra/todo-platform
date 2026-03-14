import { useContext, useEffect, useState } from "react"
import { AuthContext } from "../context/AuthContext"
import { getTodos, createTodo, deleteTodo } from "../api/api"
import Navbar from "../components/Navbar"
import TodoCard from "../components/TodoCard"

export default function Dashboard() {

  const { token } = useContext(AuthContext)

  const [todos, setTodos] = useState([])
  const [title, setTitle] = useState("")

  const fetchTodos = async () => {

    const res = await getTodos(token)

    setTodos(res.data)
  }

  useEffect(() => {
    fetchTodos()
  }, [])

  const addTodo = async () => {

    await createTodo(
      { user_id: 1, title, completed: false },
      token
    )

    setTitle("")
    fetchTodos()
  }

  const removeTodo = async (id) => {

    await deleteTodo(id, token)

    fetchTodos()
  }

  return (
    <div>

      <Navbar />

      <div className="dashboard">

        <h2>Your Todos</h2>

        <div className="create-todo">

          <input
            placeholder="New Todo"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />

          <button onClick={addTodo}>Add</button>

        </div>

        {todos.map(todo => (
          <TodoCard
            key={todo.id}
            todo={todo}
            onDelete={removeTodo}
          />
        ))}

      </div>

    </div>
  )
}