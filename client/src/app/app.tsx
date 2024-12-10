import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './pages/app/app';
import './globals.scss';
import {
	Route,
	RouterProvider,
	createBrowserRouter,
	createRoutesFromElements,
} from 'react-router-dom';
import Auth from './pages/auth/auth';

const router = createBrowserRouter(
	createRoutesFromElements(
		<>
			<Route path="/" element={<App />} />
			<Route path="/auth" element={<Auth />} />
		</>,
	),
);

const app = document.getElementById('root');
if (app) {
	createRoot(app).render(
		<StrictMode>
			<RouterProvider router={router} />
		</StrictMode>,
	);
}
