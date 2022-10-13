import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';

function Figlet() {
	let y = 1;
	const [figlet, setFiglet] = useState<string[]>([]);
	const [fileF, setFileF] = useState<string[]>([]);
	const [fileE, setFileE] = useState<string[]>([]);	
	
	const [checkedR, setCheckedR] = useState(false);
	const [checkedX, setCheckedX] = useState(false);
	const [checkedL, setCheckedL] = useState(false);
	const [checkedRR, setCheckedRR] = useState(false);
	const [checkedXX, setCheckedXX] = useState(false);
	const [checkedLL, setCheckedLL] = useState(false);
	const [checkedC, setCheckedC] = useState(false);
	const [checkedP, setCheckedP] = useState(false);
	const [checkedN, setCheckedN] = useState(false);
	const [checkedO, setCheckedO] = useState(false);
	const [checkedW, setCheckedW] = useState(false);
	const [checkedK, setCheckedK] = useState(false);
	const [checkedS, setCheckedS] = useState(false);
	const [checkedSS, setCheckedSS] = useState(false);
	const [checkedNN, setCheckedNN] = useState(false);
	const [checkedE, setCheckedE] = useState(false);
	const [checkedD, setCheckedD] = useState(false);
	const [checkedT, setCheckedT] = useState(false);
	const [checkedNNN, setCheckedNNN] = useState(false);
	const [message, setMessage] = useState("hello");
	const [checkedWW, setCheckedWW] = useState("");
	const [checkedM, setCheckedM] = useState("");
	const [checkedFFile, setCheckedFFile] = useState("");
	const [checkedEFile, setCheckedEFile] = useState("");
	const api = new Request("http://gopiko.fr/");	

	useEffect(() => {
		if (y) {
			getFileFiglet()
			y = 0;
		}
	}, []);

	function getFiglet () {
		api.Post("/figlet", {
			R: checkedR, X: checkedX, L: checkedL, RR: checkedRR, XX: checkedXX, LL: checkedLL, C: checkedC, P: checkedP, N: checkedN, O: checkedO,W: checkedW, K: checkedK, S: checkedS, SS: checkedSS, NN: checkedNN, E: checkedE, D: checkedD, T: checkedT, NNN: checkedNNN, Message: message, WW: checkedWW, M: checkedM, F: checkedFFile, CC: checkedEFile
		}).then((resp: any) => {
			if (resp.data && resp.data.StdoutResponse) setFiglet(resp.data.StdoutResponse)
		});
	};

	function getFileFiglet() {
		api.Post("/figletfile", {}).then((resp: any) => {
			setFileF(resp.data.F)
			setFileE(resp.data.E)
		})
	}

	function SelectFileFFiglet() {
		return (
			<Form.Select value={checkedFFile} onChange={(e) => {setCheckedFFile(e.target.value)}}>
				<option>ii</option>
				{fileF.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
			</Form.Select>
		);
	}

	function SelectFileEFiglet() {
		return (
			<Form.Select value={checkedEFile} onChange={(e) => {setCheckedEFile(e.target.value)}}>
				<option>ii</option>
				{fileE.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
			</Form.Select>
		);
}

	return (
		<>
			<pre>{figlet}</pre>
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedR(e.target.checked)}} checked={checkedR} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedX(e.target.checked)}} checked={checkedX} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedL(e.target.checked)}} checked={checkedL} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedRR(e.target.checked)}} checked={checkedRR} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedXX(e.target.checked)}} checked={checkedXX} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedLL(e.target.checked)}} checked={checkedLL} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedC(e.target.checked)}} checked={checkedC} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedP(e.target.checked)}} checked={checkedP} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedN(e.target.checked)}} checked={checkedN} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedO(e.target.checked)}} checked={checkedO} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedW(e.target.checked)}} checked={checkedW} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedK(e.target.checked)}} checked={checkedK} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedS(e.target.checked)}} checked={checkedS} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedSS(e.target.checked)}} checked={checkedSS} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedNN(e.target.checked)}} checked={checkedNN} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedE(e.target.checked)}} checked={checkedE} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedD(e.target.checked)}} checked={checkedD} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedT(e.target.checked)}} checked={checkedT} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedNNN(e.target.checked)}} checked={checkedNNN} />
			<NumericInput className="form-control"  value={checkedWW} onChange={(e)=>{if (e && e >= 1) setCheckedWW(e.toString());}} />
			<NumericInput className="form-control"  value={checkedM} onChange={(e)=>{if (e && e >= 1) setCheckedM(e.toString());}} />
			{SelectFileFFiglet()}
			{SelectFileEFiglet()}
			<Form.Control type="texte" onChange={(e)=>{setMessage(e.target.value)}}/>
			<button onClick={()=>{getFiglet()}}>ok</button>
		</>
	)
}

export default Figlet;
