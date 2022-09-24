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
	const [beast, setBeast] = useState("");
	const [cow, setCow] = useState("");
	const [message, setMessage] = useState("hello");

	useEffect(() => {
		if (y) {
			getCow()
			getCowFile();
		}
		y = 0;
	}, []);

	function getCow () {
		api.Post("/cow", {
			Message: message,
		}).then((resp: any) => {
			setCow(resp.data.StdoutResponse)
			console.log(resp)
		});
	};
	
	function getCowFile() {
		api.Post("/cowfile", {}).then((resp: any) => {
			setFile(resp.data)
			console.log(resp.data)
		})
	}

	return (
		<>
			<Container fluid>
				<pre>{cow}</pre>
			</Container>
			<Container fluid>
				<Form.Select value={beast} onChange={(e) => {setBeast(e.target.value)}}>
					<option></option>
				 	{file.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
				</Form.Select>
			</Container>
		</>
	)
}

export default Cow;
