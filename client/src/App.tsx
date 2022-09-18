import React, { useState, useRef, useLayoutEffect, useEffect}  from 'react';
import logo from './logo.svg';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import { Navbar, Container, Nav, Form, FormControl,Button } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Image from 'react-bootstrap/Image'
import Dropdown from 'react-bootstrap/Dropdown';
import './asset/App.css';
import Rig from './component/Rig'
import Navi from './component/Navi'
import Fortune from './component/Fortune'
import Request from './component/Request'

function App() {
	
	let y = 1;	
	const [fortune, setFortune] = useState("");
	const api = new Request("http://gopiko.fr/");
	useEffect(() => {
		if (y)
  			getFortune();
		y = 0;
	}, []);

	function getFortune () {
		api.Post("/fortune", {
		}).then((resp: any) => {
			setFortune(resp.data.StdoutResponse)
		});
	};
	return (
	<div className="App">
		<Navi />	
	    		<BrowserRouter>
				<Routes>
					<Route path="/fortune" element={<Fortune />} />
					<Route path="/rig" element={<Rig />} />
				</Routes>
			</BrowserRouter>
		<footer></footer>
	</div>
  );
}

export default App;
