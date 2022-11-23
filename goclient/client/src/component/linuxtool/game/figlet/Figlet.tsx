import React, { useState, useEffect, useRef}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Save from '../../../Save'
import GetImg from '../../../Img'
import Request from '../../../Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import Effect from '../../../effect/Effect';

function Figlet() {
	const [y, setY] = useState(1);
	const [figlet, setFiglet] = useState<string>("");
	const [fileF, setFileF] = useState<string[]>([]);
	const [fileE, setFileE] = useState<string[]>([]);	
	const h1Ref = useRef<HTMLPreElement>(null);	
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
		if (y) getFileFiglet()
		setY(0);
	}, []);
	
	function getFiglet () {
		api.Post("/figlet", {
			R: checkedR, X: checkedX, L: checkedL, RR: checkedRR, XX: checkedXX, LL: checkedLL, C: checkedC, P: checkedP, N: checkedN, O: checkedO,W: checkedW, K: checkedK, S: checkedS, SS: checkedSS, NN: checkedNN, E: checkedE, D: checkedD, T: checkedT, NNN: checkedNNN, Message: message, WW: checkedWW, M: checkedM, F: checkedFFile, CC: checkedEFile
		}).then((resp: any) => {
			if (resp.data && resp.data.StdoutResponse) setFiglet(resp.data.StdoutResponse)
		}).catch(function (error :any) {});
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
				<option value={""}>ii</option>
				{fileF.map((value :string, index :number) => {return <option value={value} key={value}>{value}</option>})}
			</Form.Select>
		);
	}

	function SelectFileEFiglet() {
		return (
			<Form.Select value={checkedEFile} onChange={(e) => {setCheckedEFile(e.target.value)}}>
				<option value={""} >ii</option>
				{fileE.map((value :string, index :number) => {return <option value={value} key={value}>{value}</option>})}
			</Form.Select>
		);
	}
	function save () {
		Save(figlet, "plain/text", "figlet.txt")
	}

	function getImg() {
                let e = document.getElementById("fig");
                GetImg(e)
        }

	return (
		<>
			<pre ref={h1Ref} id="fig" style={{
  						 "display": "inline-flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
					}}>{figlet}</pre>
			<Row>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedR(e.target.checked)}} checked={checkedR} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedX(e.target.checked)}} checked={checkedX} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedL(e.target.checked)}} checked={checkedL} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedRR(e.target.checked)}} checked={checkedRR} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedXX(e.target.checked)}} checked={checkedXX} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedLL(e.target.checked)}} checked={checkedLL} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedC(e.target.checked)}} checked={checkedC} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedP(e.target.checked)}} checked={checkedP} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedN(e.target.checked)}} checked={checkedN} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedO(e.target.checked)}} checked={checkedO} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedW(e.target.checked)}} checked={checkedW} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedK(e.target.checked)}} checked={checkedK} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedS(e.target.checked)}} checked={checkedS} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedSS(e.target.checked)}} checked={checkedSS} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedNN(e.target.checked)}} checked={checkedNN} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedE(e.target.checked)}} checked={checkedE} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedD(e.target.checked)}} checked={checkedD} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedT(e.target.checked)}} checked={checkedT} /></Col>
			<Col><Form.Check type={"checkbox"} onChange={(e)=>{setCheckedNNN(e.target.checked)}} checked={checkedNNN} /></Col>
			</Row>
			<Row>
			<Col><NumericInput className="form-control"  value={checkedWW} onChange={(e)=>{if (e && e >= 1) setCheckedWW(e.toString());}} /></Col>
			<Col><NumericInput className="form-control"  value={checkedM} onChange={(e)=>{if (e && e >= 1) setCheckedM(e.toString());}} /></Col>
			<Col>{SelectFileFFiglet()}</Col>
			<Col>{SelectFileEFiglet()}</Col>
			</Row>
			<Form.Control type="texte" onChange={(e)=>{setMessage(e.target.value)}}/>
			<button onClick={()=>{getFiglet()}}>ok</button>
			<button onClick={save} >Save</button>
			<button onClick={getImg}>get Image</button>
			<Effect h1ref={h1Ref} />
		</>
	)
}

export default Figlet;
