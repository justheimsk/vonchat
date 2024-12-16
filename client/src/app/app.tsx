import { useEffect } from 'react';
import { createRoot } from 'react-dom/client';
import App from './pages/app/app';
import './globals.scss';
import { vonchat } from '@/lib/Application';
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

function Root() {
	useEffect(() => {
		function printError() {
			vonchat.logs.send('error', 'Unknown error.');
		}

		window.addEventListener('error', () => printError());
		window.addEventListener('unhandledrejection', () => printError());

		vonchat.logs.send('info', 'Listening to app-level unknown errors.');
	}, []);

	return (
		<>
			<RouterProvider router={router} />
		</>
	);
}

if (app) {
	createRoot(app).render(<Root />);
}
