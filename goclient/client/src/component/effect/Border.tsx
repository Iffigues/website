import React, {useRef, useState, useEffect}  from 'react';
import { ColorPicker, useColor } from "react-color-palette";
import "react-color-palette/lib/css/styles.css";
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Container from 'react-bootstrap/Container';

type ClickableListProps<T> = {
  h1ref: any;
};

function changeBackgroundColor(el: HTMLElement, value: string) {
  el.style.backgroundColor = value;
}

function changeTextColor(el: HTMLElement, value: string) {
  el.style.color = value;
}

function Border<T>(props: ClickableListProps<T>) {
const [backgroundColor, setBackgroundColor] = useColor("hex", "#121212");
const [textColor, setTextColor] = useColor("hex", "#121212");
return (
	<Container>
		<Row>
			<button  className="FancyButton" onClick={
				()=>{
					let s = document.getElementById("bob");
					if (s) {
						if (s.style.display =='none')
							s.style.display='inline-flex';
						else
							s.style.display='none';
					}
				}
			}>Background Color</button>
		</Row>
		<Container id="bob" style={{display:"none"}}>
				<Container><Row>
					<button   onClick={
						()=>{
							let s = document.getElementById("bol");
								if (s) { 
									if (s.style.display =='none')
										s.style.display='inline-flex';
									else
										s.style.display='none';
								}
							}
						}>Background Color</button>
				</Row></Container>
				<Container>
				<Row id="bol" style={{display:"none"}}>
				</Row></Container>
			<Container><Row>
					<button   onClick={
						()=>{
							let s = document.getElementById("bil");
								if (s) {
									if (s.style.display =='none')
										s.style.display='inline-flex';
									else
										s.style.display='none';
								}
							}
						}>Text Color</button>
				</Row></Container>
				<Container>
				<Row id="bil" style={{display:"none"}}>
				</Row></Container>

		</Container>
	</Container>
)
}

export default Border;
