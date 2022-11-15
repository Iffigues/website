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
import Fortune from './component/linuxtool/game/fortune/Fortune'
import Cow from './component/linuxtool/game/cow/Cow'
import Rig from './component/linuxtool/game/rig/Rig'
import MyComponent from './component/Test'
import Home from './component/Home'
import Toilet from './component/linuxtool/game/toilet/Toilet'
import Figlet from './component/linuxtool/game/figlet/Figlet'
import Navi from './component/Navi'
import Request from './component/Request'

function App() {
	return (
	<div className="App">
		<Navi />
			<div style={{minHeight: "100"}}>
	    			<BrowserRouter>
					<Routes>
						<Route path="/" 	element={<Home />}	/>
						<Route path="/fortune" 	element={<Fortune />}	/>
						<Route path="/rig" 	element={<Rig />} 	/>
						<Route path="/cow" 	element={<Cow />} 	/>
						<Route path="/figlet" 	element={<Figlet />} 	/>
						<Route path="/toilet" 	element={<Toilet />} 	/>
					</Routes>
				</BrowserRouter>
			</div>
		<footer></footer>
	</div>
  );
}

export default App;
