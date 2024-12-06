import { FaSearch } from 'react-icons/fa';
import './Input.scss';

export interface InputProps extends React.HTMLProps<HTMLInputElement> {}

export default function Input(props: InputProps) {
	return (
		<>
			<div className="input-container">
				<input {...props} />
				<i>
					<FaSearch />
				</i>
			</div>
		</>
	);
}
