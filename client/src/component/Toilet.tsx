import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Toilet() {
	let y = 1;
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y)
  			getToilet();
		y = 0;
	}, []);

	function getToilet () {
		api.Post("/toilet", {
		}).then((resp: any) => {
		});		
	};

return (
	<div>
		<div></div>
		<div className="d-flex p-2">
			<button onClick={getToilet}>Get Fortune</button>;
		</div>
	</div>
	)
}

export default Toilet;
