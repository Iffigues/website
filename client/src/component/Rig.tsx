import React, { useState, useEffect}  from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import NumericInput from 'react-numeric-input';
import Container from 'react-bootstrap/Container';
import Request from './Request'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';

function Rig() {
	let y = 1;
	const [nbrRig, setNbrRig] = useState(1);
	const [checkedMan, setCheckedMan] = useState(false);
	const [checkedWoman, setCheckedWoman] = useState(false);
	const [data, setData] =  useState<any[]>([]);
	
	const listData =  data.map((n) =>
			<Col style={{paddingBottom:"2px"}}>
				<Card style={{ width: '18rem' }}>
					<Card.Body>
						<Card.Header>{n.Name}</Card.Header>
						<Card.Text>{n.Addr}</Card.Text>
						<Card.Text>{n.Postal}</Card.Text>
						<Card.Footer>{n.Tel}</Card.Footer>
					</Card.Body>
				</Card>
			</Col>
  	);	

	const api = new Request("http://gopiko.fr/");	

	useEffect(() => {
		if (y)
  			getRig();
		y = 0;
	}, []);

	function save () {
		let json = JSON.stringify(data);
		const blob = new Blob([json], { type: "application/json" });
  		const href = URL.createObjectURL(blob);
		const link = document.createElement("a");
  		link.href = href;
  		link.download =	"save.json";
  		document.body.appendChild(link);
  		link.click();
		document.body.removeChild(link);
		URL.revokeObjectURL(href);
	}

	function getRig () {
		api.Post("/rig", {
			man: checkedMan,
			woman: checkedWoman,
			nbr: nbrRig.toString()
		}).then((resp: any) => {
			console.log(resp.data)
    			setData(resp.data)
		});		
	};

return (
	<Container fluid>
		<Row>
			<Col>
				<p style={{width:"100%"}}>Man</p>
				<Form.Check  
					onChange={(e)=>{setCheckedMan(e.target.checked)}}
        				checked={checkedMan}
					type="switch"
        				id="man"
					style={{width:"100%"}}
     				 />
				<p style={{width:"100%"}}>Woman</p>
				<Form.Check 
					style={{width:"100%"}}
        				onChange={(e)=>{setCheckedWoman(e.target.checked)}}
        				checked={checkedWoman}
					type="switch"
      				/>
				<p style={{width:"100%"}}>Number</p>
				<NumericInput 
					className="form-control" 
					id="c-input" 
					min={1} value={nbrRig} 
					onChange={(e)=>{if (e && e >= 1) setNbrRig(e);}} 
				/>
				<button onClick={getRig} style={{width:"100%"}}>Get Rig</button>
				<button onClick={save} style={{width:"100%"}}>Save</button>
			</Col>
			<Col xs={10}>
				<Row xs={"auto"}>
				{listData}
				</Row>
			</Col>
		</Row>
	</Container>
	)
}

export default Rig;
