import { ErrorMagnifier } from '@/components/ErrorMagnifier/ErrorMagnifier';
import { Modal } from '@/components/Modal/Modal';
import { vonchat } from '@/lib/Application';
import { ModalBuilder } from '@/lib/core/ui/ModalBuilder';
import { useLibState } from '@/lib/state/Hook';
import { useEffect } from 'react';

export interface LayoutProps {
	children: React.ReactNode;
}

export function Layout(props: LayoutProps) {
	const modal = useLibState(vonchat.state.reducers.ui).modal;

	useEffect(() => {
		const read = localStorage.getItem('noticeRead');

		if (!read) {
			const modal = new ModalBuilder()
				.setTitle('Important Notice')
				.setDescription(`This project is still in its early stages of development, so you may encounter bugs, issues, or features that are incomplete or under construction. We are actively working to improve the experience and make the project more robust.
          \nIf you find any problems or have suggestions, we’d greatly appreciate it if you could report them on our GitHub repository. Your feedback and collaboration are essential for the growth and success of this open source project!
          \nAlso, if you enjoy coding and would like to contribute, feel free to explore the codebase and submit pull requests. Every contribution helps!
          \nLastly, if you like the project, consider leaving a star on our GitHub repository. It’s a small gesture, but it goes a long way in supporting the project and motivating us to keep improving.
          \nThank you for being part of this journey with us.`)
				.addButton('outline', 'Github', () => {
					window.open('https://github.com/justheimsk/vonchat', '_blank');
					return false;
				})
				.addButton('default', 'Ok', () => {})
				.setOnCloseCallback(() => {
					localStorage.setItem('noticeRead', 'true');
					return true;
				});

			vonchat.ui.createModal(modal);
		}
	}, []);

	return (
		<>
			<Modal
				onClose={() =>
					modal?.onClose?.() === false ? null : vonchat.ui.closeModal()
				}
				active={modal?.title !== undefined}
				title={modal?.title || ''}
				description={modal?.description || ''}
				buttons={modal?.buttons || []}
			/>
			<ErrorMagnifier />
			{props.children}
		</>
	);
}
