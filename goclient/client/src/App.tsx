import React  from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';
import './asset/App.css';
import Fortune from './component/linuxtool/game/fortune/Fortune'
import Cow from './component/linuxtool/game/cow/Cow'
import Rig from './component/linuxtool/game/rig/Rig'
import Home from './component/Home'
import Toilet from './component/linuxtool/game/toilet/Toilet'
import Figlet from './component/linuxtool/game/figlet/Figlet'
import Xkcdpass from './component/linuxtool/game/xkcdpass/Xkcdpass'
import Banner from './component/linuxtool/game/banner/Banner'
import Navi from './component/Navi'

function App() {
	return (
	<div className="App">
		<Navi />
			<div>
	    			<BrowserRouter>
					<Routes>
						<Route path="/" 	element={<Home />}	/>
						<Route path="/fortune" 	element={<Fortune />}	/>
						<Route path="/rig" 	element={<Rig />} 	/>
						<Route path="/cow" 	element={<Cow />} 	/>
						<Route path="/figlet" 	element={<Figlet />} 	/>
						<Route path="/toilet" 	element={<Toilet />} 	/>
						<Route path="/xkcdpass" element={<Xkcdpass />}	/>
						<Route path="/banner" element={<Banner />}	/>
					</Routes>
				</BrowserRouter>
			</div>
		<footer></footer>
	</div>
  );
}

export default App;
