import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Fortune() {
	let y = 1;
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
	const [checkedPercent, setCheckedPercent] = useState({} as Record<string, string>);
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y)
  			getFortune();
		y = 0;
	}, []);

	function getFortune () {
		api.Post("/fortune", {
		}).then((resp: any) => {
		});		
	};

return (
	<div>
		<div>
			<Form.Check 
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>
<Form.Check
				type={"checkbox"}
				onChange={(e)=>{setCheckedA(e.target.checked)}}
        			checked={checkedA}
          		/>

<Form.Control type="texte" placeholder="Enter email" />
<NumericInput
					className="form-control"
					id="c-input"
					min={1} value={checkedN}
					onChange={(e)=>{if (e && e >= 1) setCheckedN(e);}}
				/>
			<button onClick={getFortune}>Get Fortune</button>;
		</div>
	</div>
	)
}

export default Fortune;
