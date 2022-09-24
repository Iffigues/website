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
import Fortune from './component/Fortune'
import Cow from './component/Cow'
import Rig from './component/Rig'
import Navi from './component/Navi'
import Request from './component/Request'

function App() {
	return (
	<div className="App">
		<Navi />	
	    		<BrowserRouter>
				<Routes>
					<Route path="/fortune" element={<Fortune />} />
					<Route path="/rig" element={<Rig />} />
					<Route path="/cow" element={<Cow />}/>
				</Routes>
			</BrowserRouter>
		<footer></footer>
	</div>
  );
}

export default App;
