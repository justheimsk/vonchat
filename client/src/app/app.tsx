import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import App from './pages/app/app'
import './globals.scss'
import {createBrowserRouter, createRoutesFromElements, Route, RouterProvider} from 'react-router-dom'
import Auth from './pages/auth/auth'
import {Provider} from 'react-redux'
import store from '@/store/store'

const router = createBrowserRouter(createRoutesFromElements(
  <>
    <Route path="/" element={<App />} />
    <Route path="/auth" element={<Auth />} />
  </>
))

const app = document.getElementById('root');
if(app) {
  createRoot(app).render(
    <StrictMode>
      <Provider store={store}>
        <RouterProvider router={router} />
      </Provider>
    </StrictMode>,
  )
}
