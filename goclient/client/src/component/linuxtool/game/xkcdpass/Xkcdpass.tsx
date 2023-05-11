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
 
function Xkcdpass() {
	const api = new Request("http://gopiko.fr/");	
	const [xkcdpass, setXkcdpass] = useState("");
	const [min, setMin] = useState("");
	const [max, setMax] = useState("");
	const [verbose, setVerbose] = useState(false);
	const [numwords, setNumwords] = useState("");
	const [validchars, setValidChars] = useState("");
	const [acrostic, setAcrostic] = useState("");
	const [count, setCount] = useState("");
	const [delimiter, setDelimiter] = useState("");
	const [separators, setSeparator] = useState("");
	const [cases, setCases] = useState("");	

	function save () {
		Save(xkcdpass, "plain/text", "cow.txt")
	}
	function getImg() {
		let e = document.getElementById("coco");
		GetImg(e)
	}

	function getXkcdpass () {
		api.Post("/xkcdpass", {
			Verbose: verbose,
			Max: max,
			Min: min,
			Numwords: numwords,
			ValidChars: validchars,
			Acrostic: acrostic ,
			Count: count,
			Delimiter: delimiter,
			Separator: separators,
			Case: cases,
		}).then((resp: any) => {
			if(resp.data && resp.data.StdoutResponse) setXkcdpass(resp.data.StdoutResponse)
		}).catch(function (error :any) {});
	};
	

	return (
		<>
			<div>		
				<div id="coco" style={{"textAlign": "initial",
  						 "display": "inline-flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
						}
						}>{xkcdpass}</div>
			</div>
			<Container fluid>
				<Row>
				<Col><Form.Check style={{width:"100%"}} onChange={(e)=>{setVerbose(e.target.checked)}} checked={verbose} type="switch"/></Col>
				</Row>
				<Row>
				<Col><Form.Control type="texte" onChange={(e)=>{setMin(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setMax(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setNumwords(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setValidChars(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setAcrostic(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setDelimiter(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setSeparator(e.target.value)}}/></Col>
				<Col><Form.Control type="texte" onChange={(e)=>{setCases(e.target.value)}}/></Col>
				</Row>
				<button onClick={getXkcdpass}>get Xkcdpass</button>
				<button onClick={save}>Save</button>
				<button onClick={getImg}>get Image</button>
			</Container>
		</>
	)
}

export default Xkcdpass;
