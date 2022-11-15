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

function Color<T>(props: ClickableListProps<T>) {
const h2Ref = useRef<HTMLDivElement>(null);
const h3Ref = useRef<HTMLDivElement>(null);
const h4Ref = useRef<HTMLDivElement>(null);
const [backgroundColor, setBackgroundColor] = useColor("hex", "#121212");
const [textColor, setTextColor] = useColor("hex", "#121212");
const [t, setT] = useState(false);
const [b, setB] = useState(false);
useEffect(() => {
	if (h2Ref.current)
		h2Ref.current.style.display = "none";
	if (h3Ref.current)
		h3Ref.current.style.display = "none";
	if (h4Ref.current)
		h4Ref.current.style.display = "none";
}, [])
return (
	<Container>
		<Row>
			<button  className="FancyButton" onClick={
				()=>{
					if (h2Ref.current)
						if (h2Ref.current.style.display == "none")
							h2Ref.current.style.display = "inline-flex";
						else
							h2Ref.current.style.display = "none";
				}
			}>Background Color</button>
		</Row>
		<Container ref={h2Ref}>
				<Container><Row>
					<button   onClick={
							()=>{
								if (h3Ref.current)
									if (h3Ref.current.style.display == "none")
										h3Ref.current.style.display = "inline-flex";
									else
										h3Ref.current.style.display = "none";
						 }
						}>Background Color</button>
				</Row></Container>
				<Container>
				<Row ref={h3Ref} >
					<Form.Check type="switch" label="active color" id="disabled-custom-switch"  onChange={(e) => {
						setB(!b)
						props.h1ref.current.style.backgroundColor = null
						if (!b)
							changeBackgroundColor(props.h1ref.current, backgroundColor.hex)
					}} />
					<ColorPicker width={456} height={228} color={backgroundColor} onChange={(e) => {
						if (b)
							if (props && props.h1ref) changeBackgroundColor(props.h1ref.current, e.hex)
						setBackgroundColor(e)
					}} hideHSV dark />;
				</Row></Container>
			<Container><Row>
					<button   onClick={
							()=>{
								if (h4Ref.current)
									if (h4Ref.current.style.display == "none")
										h4Ref.current.style.display = "inline-flex";
									else
										h4Ref.current.style.display = "none";
							}
						}>Text Color</button>
				</Row></Container>
				<Container>
				<Row  ref={h4Ref}>
					<Form.Check type="switch" checked={t} label="active color" id="disabled-custom-switch" onChange={(e) => {
						setT(!t)
						props.h1ref.current.style.color = null
						if (!t)
							changeTextColor(props.h1ref.current, textColor.hex)
					}}/>
					<ColorPicker width={456} height={228} color={textColor} onChange={(e) => {
						if (t)
							if (props && props.h1ref) changeTextColor(props.h1ref.current, e.hex)
						setTextColor(e)
					}} hideHSV dark />;
				</Row></Container>

		</Container>
	</Container>
)
}

export default Color;
