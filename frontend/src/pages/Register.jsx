import { useState } from "react"
import { registerUser } from "../api/api"
import { useNavigate } from "react-router-dom"

export default function Register() {

  const navigate = useNavigate()

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleSubmit = async (e) => {

    e.preventDefault()

    await registerUser({ email, password })

    navigate("/")
  }

  return (
    <div className="auth-container">

      <h2>Register</h2>

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

        <button>Register</button>

      </form>

    </div>
  )
}