import './App.css'
import AuthenticationPage from './components/AuthenticationPage'
import ButtonPage from './components/ButtonPage'
import StatsPage from './components/StatsPage'
import { useAuthenticationContext } from './hooks/useAuthenticationContext'
import { MenuProvider, useMenuContext } from './hooks/useMenuContext'

function App() {
  const authContext = useAuthenticationContext()

  if (authContext.isLoading) {
    return <></>
  }

  if (authContext.authenticated) {
    return <MenuProvider><Router /></MenuProvider>
  } else {
    return <AuthenticationPage />
  }
}

export default App

function Router() {
  const ctx = useMenuContext()
  switch (ctx.page) {
    case "button":
      return <ButtonPage />
    case "stats":
      return <StatsPage />
    default:
      return <h1>404</h1>
  }
}