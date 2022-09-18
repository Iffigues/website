import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

interface Percent {
   [key: string]: string;
}

function Fortune() {
	let y = 1;
	const [fortune, setFortune] = useState("");
	const [checkedA, setCheckedA] = useState(false);
	const [checkedC, setCheckedC] = useState(false);
	const [checkedF, setCheckedF] = useState(false);
	const [checkedE, setCheckedE] = useState(false);
	const [checkedL, setCheckedL] = useState(false);
	const [checkedO, setCheckedO] = useState(false);
	const [checkedS, setCheckedS] = useState(false);
	const [checkedI, setCheckedI] = useState(false);
	const [checkedU, setCheckedU] = useState(false);
	const [checkedM, setCheckedM] = useState("");
	const [checkedN, setCheckedN] = useState(1);
	const [checkedPercent, setCheckedPercent] = useState<Percent>({});
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y)
  			getFortune();
			getFileFortune();
		y = 0;
	}, []);

	function getFileFortune() {
		api.Post("/filefortune", {}).then((resp: any) => {
			console.log(resp.data)
		})
	}

	function getFortune () {
		api.Post("/fortune", {
			A: checkedA,
			C: checkedC,
			E: checkedE,
			F: checkedF,
			L: checkedL,
			O: checkedO,
			S: checkedS,
			I: checkedI,
			U: checkedU,
			M: checkedM,
			N: checkedN,
			Percent: getPercent(),
			
		}).then((resp: any) => {
			setFortune(resp.data.StdoutResponse)
		});		
	};

	function getPercent() {
		Object.keys(checkedPercent).map((item,index) => {
			
		})
	}

return (
	<div>
		<Container fluid><pre>{fortune}</pre></Container>
		<Container fluid>
<Row>
<Col>
			<p>A</p>
			<Form.Check 
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
</Col>
<Col>
<p>C</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedC(e.target.checked)}}
        			checked={checkedC}
          		/></Col>
<Col>
<p>E</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedE(e.target.checked)}}
        			checked={checkedE}
          		/></Col>
<Col>
<p>F</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedF(e.target.checked)}}
        			checked={checkedF}
          		/></Col>
<Col>
<p>L</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedL(e.target.checked)}}
        			checked={checkedL}
          		/></Col>
<Col>
<p>O</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedO(e.target.checked)}}
        			checked={checkedO}
          		/></Col>
<Col>
<p>S</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedS(e.target.checked)}}
        			checked={checkedS}
          		/></Col>
<Col>
<p>I</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedI(e.target.checked)}}
        			checked={checkedI}
          		/></Col>
<Col>
<p>U</p>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedU(e.target.checked)}}
        			checked={checkedU}
          		/></Col>
</Row>
<Form.Control 
	type="texte" 
	onChange={(e)=>{setCheckedM(e.target.value)}
}
 />
<NumericInput
					className="form-control"
					id="c-input"
					min={1} value={checkedN}
					onChange={(e)=>{if (e && e >= 1) setCheckedN(e);}}
				/>
			<button>add Percent</button>
			<button onClick={getFortune}>Get Fortune</button>
		</Container>
	</div>
	)
}

export default Fortune;
