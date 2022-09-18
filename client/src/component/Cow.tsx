import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Cow() {
	let y = 1;
	const [isThink, setIsThink] = useState(false);
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y)
  			getCow();
		y = 0;
	}, []);

	function getCow () {
		let endpoint = "cow"
		if (isThink) {
			endpoint = "cowthink"
		}
	
		api.Post("/cow", {
		}).then((resp: any) => {
		});		
	};

return (
	<div>
		<div></div>
		<div className="d-flex p-2">
			<button onClick={getCow}>Get Fortune</button>;
		</div>
	</div>
	)
}

export default Cow;
