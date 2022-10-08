import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import { Markup } from 'interweave'; 
function Cow() {
	let y = 1;
	const api = new Request("http://gopiko.fr/");	
	const [file, setFile] = useState<string[]>([]);
	const [beast, setBeast] = useState("");
	const [cow, setCow] = useState("");
	const [message, setMessage] = useState("hello");
	const [checkedThink, setCheckedThink] = useState(false);
	const [checkedL, setCheckedL] = useState(false);
	const [checkedB, setCheckedB] = useState(false);
	const [checkedD, setCheckedD] = useState(false);	
	const [checkedG, setCheckedG] = useState(false);
	const [checkedP, setCheckedP] = useState(false);
	const [checkedS, setCheckedS] = useState(false);
	const [checkedT, setCheckedT] = useState(false);
	const [checkedW, setCheckedW] = useState(false);
	const [checkedE, setCheckedE] = useState("");
	const [checkedTT, setCheckedTT] = useState("");
	const [checkedY, setCheckedY] = useState(false);

	useEffect(() => {
		if (y) {
			getCow()
			getCowFile();
		}
		y = 0;
	}, []);

	function getCow () {
		api.Post("/cow", {
			Think: checkedThink,
			L: checkedL,
			B: checkedB,
			D: checkedD,
			G: checkedG,
			P: checkedP,
			S: checkedS,
			W: checkedW,
			E: checkedE,
			TT: checkedTT,
			F: beast,
			Y: checkedY,
			Message: message,
		}).then((resp: any) => {
			if(resp.data && resp.data.StdoutResponse) setCow(resp.data.StdoutResponse)
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
			<div>
				<pre style={{"textAlign": "initial"}}>{cow}</pre>
			</div>
			<Container fluid>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedThink(e.target.checked)}}
        				checked={checkedThink}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedL(e.target.checked)}}
        				checked={checkedL}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedB(e.target.checked)}}
        				checked={checkedB}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedD(e.target.checked)}}
        				checked={checkedD}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedG(e.target.checked)}}
        				checked={checkedG}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedP(e.target.checked)}}
        				checked={checkedP}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedS(e.target.checked)}}
        				checked={checkedS}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedT(e.target.checked)}}
        				checked={checkedT}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedW(e.target.checked)}}
        				checked={checkedW}
					type="switch"
      				/>
				<Form.Check
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedY(e.target.checked)}}
        				checked={checkedY}
					type="switch"
      				/>
				<Form.Select value={beast} onChange={(e) => {setBeast(e.target.value)}}>
					<option></option>
				 	{file.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
				</Form.Select>
				<Form.Control type="texte" onChange={(e)=>{setCheckedE(e.target.value)}}/>
				<Form.Control type="texte" onChange={(e)=>{setCheckedTT(e.target.value)}}/>
				<Form.Control type="texte" onChange={(e)=>{setMessage(e.target.value)}}/>
				<button onClick={getCow}>get Cow</button>
			</Container>
		</>
	)
}

export default Cow;
