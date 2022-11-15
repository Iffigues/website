import React, { useState, useEffect, useRef}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from '../../../Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import GetImg from '../../../Img'
import Save from '../../../Save'
import Effect from '../../../effect/Effect';

interface Percent {
   [key: string]: string;
}

function Fortune() {
	let y = 1;
	const h1Ref = useRef<HTMLPreElement>(null);
	const [f, setF] = useState(false)
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
	const [checkedN, setCheckedN] = useState("1");
	const [checkedPercent, setCheckedPercent] = useState<any[]>([],);
	const [file, setFile] = useState<string[]>([]);
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y) {
  			getFortune();
			getFileFortune();
		}
		y = 0;
	}, []);

	function addPercent() {
		setCheckedPercent(oldArray => [...oldArray, {Fortune:"", Percent: "", percents: 0}]);	
	}

	function getFileFortune() {
		api.Post("/fortunefile", {}).then((resp: any) => {
			setFile(resp.data)
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
			N: checkedN.toString(),
			Percent: checkedPercent,
			
		}).then((resp: any) => {
			if (resp.data && resp.data.StdoutResponse) setFortune(resp.data.StdoutResponse)
		}).catch(function (error :any) {});
	};

	function SelectFileFortune(values :any, i :number) {
		return (
			<Form.Select
				value={values.Fortune}
        		  	onChange={(e) => {
				let newArr = [...checkedPercent]; 
				newArr[i].Fortune = e.target.value
				setCheckedPercent(newArr)
			}}
			>
			<option>ii</option>
			{file.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
		</Form.Select>
		);
	}


	function SelectPercent() {
		return (
			<>
				{
					checkedPercent.map((value :any, index :number) => {
						return (
							<div>
								{SelectFileFortune(value, index)}
								<NumericInput
									value={value.percents}
									onChange={(e)=>{
										let newArr = [...checkedPercent];
										newArr[index].percents = e;
										if (e) newArr[index].Percent = '' + e;
											setCheckedPercent(newArr)
											value.Percents = e
									}}	
								/>
							</div>
					)})
				}
			</>
 		);
	}
	function save () {
		Save(fortune, "plain/text", "fortune.txt")
	}

	function getImg() {
		let e = document.getElementById("dd");
		GetImg(e)
	}

	return (
		<>
		<Container fluid>
			<Row>
				<Col>
					<Row>
						<Col>
							<p>A</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedA(e.target.checked)}} checked={checkedA} />
						</Col>
						<Col>
							<p>C</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedC(e.target.checked)}} checked={checkedC} />
						</Col>
						<Col>
							<p>E</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedE(e.target.checked)}} checked={checkedE} />
						</Col>
						<Col>
							<p>F</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedF(e.target.checked)}} checked={checkedF} />
						</Col>
						<Col>
							<p>L</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedL(e.target.checked)}} checked={checkedL} />
						</Col>
						<Col>
							<p>O</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedO(e.target.checked)}} checked={checkedO} />
						</Col>
						<Col>
							<p>S</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedS(e.target.checked)}} checked={checkedS} />
						</Col>
						<Col>
							<p>I</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedI(e.target.checked)}} checked={checkedI} />
						</Col>
						<Col>
							<p>U</p>
							<Form.Check type={"checkbox"} onChange={(e)=>{setCheckedU(e.target.checked)}} checked={checkedU} />
						</Col>
					</Row>
					<Form.Control type="texte" onChange={(e)=>{setCheckedM(e.target.value)}}/>
					<NumericInput className="form-control"  value={checkedN} onChange={(e)=>{if (e && e >= 1) setCheckedN(e.toString());}} />
					{SelectPercent()}
					<button onClick={addPercent}> add File</button>
					<button onClick={getFortune}>Get Fortune</button>
					<button onClick={save}>Save</button>
					<button onClick={getImg}>get Image</button>
				</Col>
				<Col xs={10}>
					<pre ref={h1Ref} id="dd" style={{"textAlign": "initial",
  						 "display": "inline-flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
					}}>{fortune}</pre>
				</Col>
			</Row>
		</Container>
		<Effect h1ref={h1Ref} />
		</>
	)
}

export default Fortune;
