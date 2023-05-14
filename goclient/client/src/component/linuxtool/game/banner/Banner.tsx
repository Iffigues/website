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
	const [Message, setMessage] =  useState<string[]>([""]);

	function save () {
		Save(xkcdpass, "plain/text", "cow.txt")
	}
	function getImg() {
		let e = document.getElementById("coco");
		GetImg(e)
	}

	function getBanner () {
		api.Post("/banner", {
			Message: Message,
		}).then((resp: any) => {
			if(resp.data && resp.data.StdoutResponse) setXkcdpass(resp.data.StdoutResponse)
		}).catch(function (error :any) {});
	};

	function addButton() {
		setMessage([...Message, ""]);
	}

	function handleUpdate(index: any, newValue: any) {
    		const newArray = [...Message];
    		newArray[index] = newValue;
    		setMessage(newArray);
 	 }

	function handleDelete(index: number) {
    		const newArray = Message.filter((_, i) => i !== index);
    		setMessage(newArray);
  	}
	
	function Mes() {
		return (
			<>
				{
					Message.map((value :any, index :number) => {
						return (
							<div>
								<Form.Control type="texte" onChange={(e)=>{
									handleUpdate(index, e.target.value)
								}}/>
								<button onClick={() => handleDelete(index)}>Delete</button>
							</div>
					)})
				}
			</>
 		);
	}

	return (
		<>
			<div>

				<pre  id="coco" style={{"textAlign": "initial",
  						 "display": "inline-flex",
 						 "alignItems": "center",
  						 "justifyContent": "center",
					}}>{xkcdpass}</pre>	
			</div>
			{Mes()}
			<button onClick={addButton}>Add</button>
			<Container fluid>
				<button onClick={getBanner}>get Banner</button>
				<button onClick={save}>Save</button>
				<button onClick={getImg}>get Image</button>
			</Container>
		</>
	)
}

export default Xkcdpass;
