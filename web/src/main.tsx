import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import axios from 'axios'
import { AtuhenticationProvider } from './hooks/useAuthenticationContext.tsx'

const queryClient = new QueryClient()

axios.defaults.baseURL = import.meta.env.PROD ? 'https://api.drknap.org/api' : 'http://localhost:30420/api'
axios.defaults.withCredentials = true

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <AtuhenticationProvider>
        <App />
      </AtuhenticationProvider>
      <ReactQueryDevtools client={queryClient} initialIsOpen={true} />
    </QueryClientProvider>
  </StrictMode>
)
