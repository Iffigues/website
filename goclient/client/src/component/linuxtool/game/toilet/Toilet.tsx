import React, {useRef, useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from '../../../Request'
import Save from '../../../Save'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Effect from '../../../effect/Effect';
import GetImg from '../../../Img'

const style = {
	lineHeight:"initial"
}

function Toilet() {
	let y = 1;
	const h2Ref = useRef<HTMLDivElement>(null);
	const h1Ref = useRef<HTMLPreElement>(null);
	const nohtml = useRef<HTMLDivElement>(null);
	const html = useRef<HTMLDivElement>(null);
	const api = new Request("http://gopiko.fr/");
	const [toilet, setToilet] = useState("");
	const [f,setF] = useState<string[]>([]);
	const [e, setE] = useState<string[]>([]);
	const [ff, setFf] = useState<string[]>([]);
	const [s, setS] = useState(false);
	const [ss, setSs] = useState(false);
	const [k, setK] = useState(false);
	const [w, setW] = useState(false);
	const [o, setO]	= useState(false);
	const [ffs, setFfs] = useState("")
  	const [fff, setFff] = useState("")
	const [ffff, setFfff] = useState("")
	const [ee, setEe] = useState("")
	const [message, Message] = useState("hello")
	const [l, setL] = useState("")
	const [ll, setLl] = useState("")
	const [checkedPercent, setCheckedPercent] = useState<any[]>([],);	
	const d = React.useRef<HTMLDivElement>(null);
	const dd = React.useRef<HTMLDivElement>(null);


	useEffect(() => {
		if (html.current)
			html.current.style.display = "none";
		if (nohtml.current)
			nohtml.current.style.display = "none";
	}, [])

	function SelectFileFortune(values :any, i :number) {
		return (
			<div key={i}>
			<Form.Select
				value={values.Fortune}
        		  	onChange={(e) => {
				let newArr = [...checkedPercent]; 
				newArr[i] = e.target.value
				setCheckedPercent(newArr)
			}}
			>
			<option>ii</option>
			{f.map((value :string, index :number) => {return <option value={value} key={index}>{value}</option>})}
		</Form.Select>
		<button onClick={() => {
			 let newArr = [...checkedPercent];
			 newArr.splice(i, 1);
			setCheckedPercent(newArr) 
		}}>X</button></div>
		);
	}

	function addPercent() {
		setCheckedPercent(oldArray => [...oldArray, ""]);	
	}

	
	function SelectPercent() {
		return (
			<>
				{
					checkedPercent.map((value :any, index :number) => {
						return (
							<div key={index}>
								{SelectFileFortune(value, index)}
							</div>
					)})
				}
			</>
 		);
	}


	useEffect(() => {
		if (y) {
  			getFileToilet();
			y = 0;
		}
	}, []);

	function getFileToilet() {
		api.Post("/toiletfile", {}).then((resp: any) => {
			setF(resp.data.F)
			setE(resp.data.E)
			setFf(resp.data.FF)
		})	
	}

	function save () {
		Save(l, "plain/text", "toilet.txt")
	}

	function getToilet () {
		api.Post("/toilet", {
			S: s, SS:ss, K: k, W: w, O: o, F: ffs, FF: fff, FFF: ffff, E: ee, F3:checkedPercent, Message: message
		}).then((resp: any) => {
			console.log(resp)
			if (resp.data && resp.data.StdoutResponse) {
				setL(resp.data.StdoutResponse)
				let s = document.getElementById("mes");
				let ss = document.getElementById('mos');
				if (ee == "html3" || ee == "html" || ee == "svg") {
					if (h2Ref.current)
						h2Ref.current.style.display = "block"
					if (h1Ref.current)
						h1Ref.current.style.display = "none"
					if (html.current)
						html.current.style.display = "block"
					if (nohtml.current)
						nohtml.current.style.display = "none"
				} else {
					if (h2Ref.current)
						h2Ref.current.style.display = "none"
					if (h1Ref.current)
						h1Ref.current.style.display = "block"
					if (html.current)
						html.current.style.display = "none"
					if (nohtml.current)
						nohtml.current.style.display = "block"
				}
			}
		});		
	};

	function getImg() {
		let y = "mos";
		if (ee == "html3" || ee == "html" || ee == "svg")
			y = "mes"	
	
		let e = document.getElementById(y);
		GetImg(e)
	}

	function Select1() {
		return (
			<Form.Select value={ffs} onChange={(e) => {setFfs(e.target.value)}}>
				<option value={""}>ii</option>
				{ff.map((value :string, index :number) => {return <option value={value} key={index}>{value}</option>})}
			</Form.Select>
		);
	}
	
	function Select2() {
		return (
			<Form.Select value={fff} onChange={(e) => {setFff(e.target.value)}}>
				<option value={""}>ii</option>
				{f.map((value :string, index :number) => {return <option value={value} key={index}>{value}</option>})}
			</Form.Select>
		);
	}

	function Select3() {
		return (
			<Form.Select value={ffff} onChange={(e) => {setFfff(e.target.value)}}>
				<option value={""}>ii</option>
				{f.map((value :string, index :number) => {return <option value={value} key={index}>{value}</option>})}
			</Form.Select>
		);
	}

	function Select4() {
		return (
			<Form.Select value={ee} onChange={(e) => {setEe(e.target.value)}}>
				<option value={""}>ii</option>
				{e.map((value :string, index :number) => {return <option value={value} key={index}>{value}</option>})}
			</Form.Select>
		);
	}	
		
	return (
		<div>
			<div>
				<div id="mes" ref={h2Ref}  style={
					{
  						 "display": "inline-flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
					}
				} dangerouslySetInnerHTML={{ __html: l}} />
				<div style={
					{
  						 "display": "flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
					}
				}><pre ref={h1Ref} id="mos">{l}</pre></div>
			</div>
			<Row>
				<Col><Form.Check type={"checkbox"} onChange={(e)=>{setS(e.target.checked)}} checked={s} /></Col>
				<Col><Form.Check type={"checkbox"} onChange={(e)=>{setSs(e.target.checked)}} checked={ss} /></Col>
				<Col><Form.Check type={"checkbox"} onChange={(e)=>{setK(e.target.checked)}} checked={k} /></Col>
				<Col><Form.Check type={"checkbox"} onChange={(e)=>{setW(e.target.checked)}} checked={w} /></Col>
				<Col><Form.Check type={"checkbox"} onChange={(e)=>{setO(e.target.checked)}} checked={o} /></Col>
			</Row>
			<Row>
			</Row>
			<Row>
				<Col>{Select1()}</Col>
				<Col>{Select4()}</Col>
			</Row>
			<Row>
				{SelectPercent()}
				<button onClick={addPercent}> add File</button>
	
			</Row>
			<Form.Control  as="textarea" rows={3} onChange={(e)=>{Message(e.target.value)}}/>	
			<div className="d-flex p-2">
				<button onClick={getToilet}>Get Fortune</button>;
				<button onClick={save}>Save</button>
				<button onClick={getImg}>get Image</button>
			</div>
			<div ref={nohtml}><Effect h1ref={h1Ref} /></div>
			<div ref={html}><Effect h1ref={h2Ref} /></div>
		</div>
	)
}

export default Toilet
