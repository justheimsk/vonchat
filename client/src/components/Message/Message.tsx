import './Message.scss';

export default function Message() {
	return (
		<>
			<div className="message">
				<div className="message__avatar" />
				<div className="message__data">
					<div className="message__data__infos">
						<span>vonderheimsk</span>
						<small>Today at 2:51 pm</small>
					</div>
					<span>
						Lorem ipsum dolor sit amet consectetur adipisicing elit.
						Exercitationem nobis voluptatum repellendus similique quidem
						blanditiis deleniti illum ratione voluptas. Vero in assumenda
						eveniet laborum ea quis fuga corrupti expedita atque!
					</span>
				</div>
			</div>
		</>
	);
}
