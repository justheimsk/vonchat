import type { ModalButton } from '@/lib/types/Modal';
import './Modal.scss';

export interface ModalProps {
	title: string;
	description: string;
	buttons: ModalButton[];
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
				<div className="modal--padding">
					<h3>{props.title}</h3>
				</div>
				<div className="modal__content modal--padding">{props.description}</div>
				<div className="modal__footer modal--padding">
					{props.buttons.map((bn) => (
						<button
							onClick={() =>
								bn.callback?.() === false ? null : props.onClose?.()
							}
							type="button"
							key={bn.label}
							className={`modal__footer__button modal__footer__button--${bn.type}`}
						>
							{bn.label}
						</button>
					))}
				</div>
			</div>
		</>
	);
}
