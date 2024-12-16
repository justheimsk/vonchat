import { StrictMode, useEffect } from 'react';
import { createRoot } from 'react-dom/client';
import App from './pages/app/app';
import './globals.scss';
import { Modal } from '@/components/Modal/Modal';
import { vonchat } from '@/lib/Application';
import { useLibState } from '@/lib/state/Hook';
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
	const modal = useLibState(vonchat.state.reducers.ui).modal;

	useEffect(() => {
		function printError() {
			vonchat.logs.send('error', 'Unknown error.');
		}

		window.addEventListener('error', () => printError());
		window.addEventListener('unhandledrejection', () => printError());

		vonchat.logs.send('info', 'Listening to app-level unknown errors.');
	}, []);

	return (
		<StrictMode>
			<Modal
				onClose={() =>
					modal?.onClose?.() === false ? null : vonchat.ui.closeModal()
				}
				active={modal?.title !== undefined}
				title={modal?.title || ''}
				description={modal?.description || ''}
				buttons={modal?.buttons || []}
			/>
			<RouterProvider router={router} />
		</StrictMode>
	);
}

if (app) {
	createRoot(app).render(<Root />);
}
