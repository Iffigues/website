import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Figlet() {
	let y = 1;
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y)
  			getFiglet();
		y = 0;
	}, []);

	function getFiglet () {
		api.Post("/figlet", {
		}).then((resp: any) => {
		});		
	};

return (
	<div>
		<div></div>
		<div className="d-flex p-2">
			<button onClick={getFiglet}>Get Fortune</button>;
		</div>
	</div>
	)
}

export default Figlet;
