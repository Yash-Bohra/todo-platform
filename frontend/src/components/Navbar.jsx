import { useContext } from "react"
import { AuthContext } from "../context/AuthContext"
import { useNavigate } from "react-router-dom"

export default function Navbar() {

  const { logout } = useContext(AuthContext)
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate("/")
  }

  return (
    <div className="navbar">
      <h2>Todo Platform</h2>
      <button onClick={handleLogout}>Logout</button>
    </div>
  )
}