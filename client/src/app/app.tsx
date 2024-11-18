import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import App from './pages/home.tsx'
import './globals.scss'

const app = document.getElementById('root');
if(app) {
  createRoot(app).render(
    <StrictMode>
      <App />
    </StrictMode>,
  )
}
