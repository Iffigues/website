import React, {useRef, useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from '../../../Request'
import Save from '../../../Save'
import GetImg from '../../../Img'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import { Markup } from 'interweave';
import Effect from '../../../effect/Effect';
 
function Cow() {
	const h1Ref = useRef<HTMLPreElement>(null);
	const api = new Request("http://gopiko.fr/");	
	const [xkcdpass, setXkcdpass] = useState("");
	const [min, setMin] = useState("");
	const [max, setMax] = useState("");
	const [Verbose, setVerbose] = useState(false);
	const [numwords, setNumwords] = useState("");
	const [validchars, setValidChars] = useState("");
	const [acrostic, setAcrostic] = useState("");
	const [count, setCount] = useState("");
	const [delimiter, setDelimiter] = useState("");
	const [Separator, setSeparator] = useState("");
	const [case, Case] = useState("");	

	function save () {
		Save(xkcdpass, "plain/text", "cow.txt")
	}
	function getImg() {
		let e = document.getElementById("coco");
		GetImg(e)
	}

	function getCow () {
		api.Post("/xkcdpass", {
			Separator: separator,
			Verbose: verbose,
			Max: max,
			Min: min,
			numwords: numwords,
			ValidChars: validchars,
			Acrostic: acrostic ,
			Count: count,
			Delimiter: delimiter,
			Separator: separator,
			Case: case,
		}).then((resp: any) => {
			if(resp.data && resp.data.StdoutResponse) setXkcdpass(resp.data.StdoutResponse)
		}).catch(function (error :any) {});
	};
	

	return (
		<>
			<div>		
				<div id="coco" ref={h1Ref}  style={{"textAlign": "initial",
  						 "display": "inline-flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
						}
						}>{xkcdpass}</div>
			</div>
			<Container fluid>
				<Row>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedThink(e.target.checked)}} checked={checkedThink} type="switch"/></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedL(e.target.checked)}} checked={checkedL} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedB(e.target.checked)}} checked={checkedB} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedD(e.target.checked)}} checked={checkedD} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedG(e.target.checked)}} checked={checkedG} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedP(e.target.checked)}} checked={checkedP} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedS(e.target.checked)}} checked={checkedS} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedT(e.target.checked)}} checked={checkedT} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedW(e.target.checked)}} checked={checkedW} type="switch" /></Col>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setCheckedY(e.target.checked)}} checked={checkedY} type="switch" /></Col>
				</Row>
				<Row>
				<Col><Form.Select value={beast} onChange={(e) => {setBeast(e.target.value)}}> <option></option> {file.map((value :string, index :number) => {return <option value={value} key={index}>{value}</option>})}</Form.Select></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setCheckedE(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setCheckedTT(e.target.value)}}/></Col>
				</Row>
				<Form.Control type="texte" onChange={(e)=>{setMessage(e.target.value)}}/>
				<button onClick={getCow}>get Cow</button>
				<button onClick={save}>Save</button>
				<button onClick={getImg}>get Image</button>
			</Container>
			<Effect h1ref={h1Ref} />
		</>
	)
}

export default Cow;
