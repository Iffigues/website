import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Toilet() {
	let y = 1;
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

	const d = React.useRef<HTMLDivElement>(null);
	const dd = React.useRef<HTMLDivElement>(null);

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

	function getToilet () {
		console.log(f, e, ff)
		api.Post("/toilet", {
			S: s,
			SS:ss,
			K: k,
			W: w,
			O: o,
			F: ffs,
			FF: fff,
			FFF: ffff,
			E: ee,
			Message: message
		}).then((resp: any) => {
			setL(resp.data.StdoutResponse)
			console.log("zz", resp.data)
			let s = document.getElementById("mes");
			let ss = document.getElementById('mos');
			if (ee == "html3" || ee == "html") {
				if (s) {
					s.style.display='block';
				}
				if (ss) {
					ss.style.display='none';
				}
			} else {
				if (s) {
					s.style.display='none';
				}
				if (ss) {
					ss.style.display='block';
				}
			}
		});		
	};

	function Select1() {
		return (
			<Form.Select
				value={ffs}
	        	  	onChange={(e) => {setFfs(e.target.value)}}
			>
				<option>ii</option>
				{ff.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
			</Form.Select>
		);
	}
	
	function Select2() {
		return (
			<Form.Select
				value={fff}
	        	  	onChange={(e) => {setFff(e.target.value)}}
			>
				<option>ii</option>
				{f.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
			</Form.Select>
		);
	}

	function Select3() {
		return (
			<Form.Select
				value={ffff}
	        	  	onChange={(e) => {setFfff(e.target.value)}}
			>
				<option>ii</option>
				{f.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
			</Form.Select>
		);
	}

	function Select4() {
		return (
			<Form.Select
				value={ee}
	        	  	onChange={(e) => {setEe(e.target.value)}}
			>
				<option>ii</option>
				{e.map((value :string, index :number) => {return <option value={value}>{value}</option>})}
			</Form.Select>
		);
	}	
	

	return (
		<div>
			<div>
					<div id="mes" dangerouslySetInnerHTML={{ __html: l}} />
					<pre id="mos">{l}</pre>
			</div>
			<Form.Check type={"checkbox"} onChange={(e)=>{setS(e.target.checked)}} checked={s} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setSs(e.target.checked)}} checked={ss} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setK(e.target.checked)}} checked={k} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setW(e.target.checked)}} checked={w} />
			<Form.Check type={"checkbox"} onChange={(e)=>{setO(e.target.checked)}} checked={o} />
			{Select1()}
			{Select2()}
			{Select3()}
			{Select4()}
			<Form.Control type="texte" onChange={(e)=>{Message(e.target.value)}}/>	
			<div className="d-flex p-2">
				<button onClick={getToilet}>Get Fortune</button>;
			</div>
		</div>
	)
}

export default Toilet;
