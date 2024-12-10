import './Modal.scss';

export interface ModalProps {
	title: string;
	active: boolean;
	onClose: () => void;
}

export function Modal(props: ModalProps) {
	return (
		<>
			<div
				onClick={() => props.onClose?.()}
				className={`modal__overlay ${props.active ? 'modal__overlay--active' : ''}`}
			/>
			<div className={`modal ${props.active ? 'modal--active' : ''}`}>
				<h3>{props.title}</h3>
			</div>
		</>
	);
}
