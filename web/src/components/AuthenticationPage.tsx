import { useMutation } from "@tanstack/react-query"
import axios from "axios"
import { useState } from "react"
import { useAuthenticationContext } from "../hooks/useAuthenticationContext"

function AuthenticationPage() {
    const [code, setCode] = useState("")
    const [newUserName, setNewUserName] = useState("")
    const [newUserCode, setNewUserCode] = useState("")
    const authContext = useAuthenticationContext()

    const loginMutation = useMutation({
        mutationFn: () => axios.post("/user/auth",
            { "code": code },
            {
                headers: { "Content-Type": "multipart/form-data" },
                withCredentials: true
            }),
        onSuccess: () => { authContext.setAuthenticated(true) },
    })

    const registerMutation = useMutation({
        mutationFn: () => axios.post("/user/new",
            { "name": newUserName, "code": newUserCode },
            {
                headers: { "Content-Type": "multipart/form-data" },
                withCredentials: true
            }),
        onSuccess: () => { authContext.setAuthenticated(true) },
    })

    const handleLoginSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        loginMutation.mutate()
    }

    const handleRegisterSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        registerMutation.mutate()
    }

    return (
        <div className="authentication-page">
            <form onSubmit={handleLoginSubmit}>

                <input
                    type="text"
                    name="code"
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => setCode(e.target.value)}
                    value={code} />

                <button>login</button>
            </form>
            <form onSubmit={handleRegisterSubmit}>

                <input
                    type="text"
                    name="name"
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => setNewUserName(e.target.value)}
                    value={newUserName} />

                <input
                    type="text"
                    name="code"
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => setNewUserCode(e.target.value)}
                    value={newUserCode} />

                <button>register</button>
            </form>
        </div>
    )
}

export default AuthenticationPage