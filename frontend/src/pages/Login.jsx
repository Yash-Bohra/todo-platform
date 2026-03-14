import { useState, useContext } from "react"
import { loginUser } from "../api/api"
import { AuthContext } from "../context/AuthContext"
import { useNavigate, Link } from "react-router-dom"

export default function Login() {

  const { login } = useContext(AuthContext)
  const navigate = useNavigate()

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleSubmit = async (e) => {

    e.preventDefault()

    const res = await loginUser({ email, password })

    login(res.data.token)

    navigate("/dashboard")
  }

  return (
    <div className="auth-container">

      <h2>Login</h2>

      <form onSubmit={handleSubmit}>

        <input
          placeholder="Email"
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />

        <button>Login</button>

      </form>

      <Link to="/register">Register</Link>

    </div>
  )
}