import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';

function Cow() {
	let y = 1;
	const api = new Request("http://gopiko.fr/");	
	const [file, setFile] = useState<string[]>([]);

	useEffect(() => {
		if (y)
		getCowFile();
		y = 0;
	}, []);
	
	function getCowFile() {
		api.Post("/cowfile", {}).then((resp: any) => {
			setFile(resp.data)
		})
	}

	return (
		<Container fluid>
		</Container>
	)
}

export default Cow;
