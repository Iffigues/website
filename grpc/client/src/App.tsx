import React, { useState, useRef, useLayoutEffect}  from 'react';
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

const Hello = React.forwardRef((props, ref: React.Ref<HTMLDivElement>) => (
	<>
	</>
));

function App() {
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
