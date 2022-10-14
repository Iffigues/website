import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

function Home() {
	let y = 1;
	const [fortune, setFortune] = useState("");
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y) 
  			getFortune();
		y = 0;
	}, []);

	function getFortune () {
		api.Post("/fortune", {
		}).then((resp: any) => {
			if (resp.data && resp.data.StdoutResponse) setFortune(resp.data.StdoutResponse)
		});
	};

return (
	<Container fluid>
				<pre style={{"textAlign": "initial"}}>{fortune}</pre>
	</Container>
	)
}

export default Home;
