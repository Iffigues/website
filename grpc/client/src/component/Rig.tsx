import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Rig() {
	let y = 1;
	const [nbrRig, setNbrRig] = useState(1);
	const [checkedMan, setCheckedMan] = useState(false);
	const [checkedWoman, setCheckedWoman] = useState(false);
	const api = new Request("http://gopiko.fr/");
	
	useEffect(() => {
		if (y)
  			getRig();
		y = 0;
	}, []);

	function getRig () {
		api.Post("/rig", {
			man: checkedMan,
			woman: checkedWoman,
			nbr: nbrRig.toString()
		}).then((resp: any) => {
    			console.log(resp.data);
		});		
	};

return (
	<div>	
	<div className="d-flex p-2">
		<Form.Check 
			onChange={(e)=>{setCheckedMan(e.target.checked)}}
        		checked={checkedMan}
			type="switch"
        		id="man"
        		label="man"
     		 />
		<Form.Check 
        		onChange={(e)=>{setCheckedWoman(e.target.checked)}}
        		checked={checkedWoman}
			type="switch"
        		label="woman"
        		id="woman"
      		/>
		<NumericInput className="form-control" id="c-input" min={1} value={nbrRig} onChange={(e)=>{if (e && e >= 1) setNbrRig(e);}} />
		<button onClick={getRig}>Get Rig</button>;
		</div>
		<div></div>
	</div>
	)
}

export default Rig;
