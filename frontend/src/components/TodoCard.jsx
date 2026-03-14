export default function TodoCard({ todo, onDelete }) {

  return (
    <div className="todo-card">

      <h4>{todo.title}</h4>

      <button
        onClick={() => onDelete(todo.id)}
      >
        Delete
      </button>

    </div>
  )
}