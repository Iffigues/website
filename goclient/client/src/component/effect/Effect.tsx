import React, {useRef, useState, useEffect}  from 'react';
import { ColorPicker, useColor } from "react-color-palette";
import "react-color-palette/lib/css/styles.css";
import Color from "./Color"
import Border from "./Border"
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';

type ClickableListProps<T> = {
  h1ref: any;
};

function changeInnerText(el: HTMLElement, value: string) {
  el.style.backgroundColor = value;
}

function Effect<T>(props: ClickableListProps<T>) {
	const h2Ref = useRef<HTMLDivElement>(null);
	useEffect(() => {
		if (h2Ref.current)
			h2Ref.current.style.display = "none";
	}, [])
	return (
		<div>
			<Row>
				<button  className="FancyButton" onClick={
					()=>{
						if (h2Ref.current)
							if (h2Ref.current.style.display == "none")
								h2Ref.current.style.display = "block";
							else
								h2Ref.current.style.display = "none";
					}
				}> Effect </button>
			</Row>
			<div>
				<div ref={h2Ref}>
					<Row ><Color h1ref={props.h1ref} /></Row>
				</div>
			</div>
		</div>
	)
}

export default Effect;
