import './App.css'
import AuthenticationPage from './components/AuthenticationPage'
import ButtonPage from './components/ButtonPage'
import { useAuthenticationContext } from './hooks/useAuthenticationContext'

function App() {
  const authContext = useAuthenticationContext()

  if (authContext.authenticated) {
    return <ButtonPage />
  } else {
    return <AuthenticationPage />
  }
}

export default App
