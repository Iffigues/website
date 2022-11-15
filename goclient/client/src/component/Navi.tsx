import React, { useState, useRef, useLayoutEffect}  from 'react';
import { Navbar, Container, Nav, Form, FormControl,Button } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Image from 'react-bootstrap/Image'
import Dropdown from 'react-bootstrap/Dropdown';
import ReactLogo  from '../asset/logo.png';

const Lin = React.forwardRef((props, ref: React.Ref<HTMLDivElement>) => (
	<Dropdown>
		<Dropdown.Toggle variant="success" id="dropdown-basic">Play With Liunux</Dropdown.Toggle>		 
		<Dropdown.Menu>
        		<Dropdown.Item href="/fortune">Fortune</Dropdown.Item>
        		<Dropdown.Item href="/rig">Rig</Dropdown.Item>
        		<Dropdown.Item href="/toilet">Toilet</Dropdown.Item>
			<Dropdown.Item href="/figlet">Figlet</Dropdown.Item>
			<Dropdown.Item href="/cow">Cow</Dropdown.Item>
		</Dropdown.Menu>
	</Dropdown>
));


function Navi() {
return (
	
<Navbar bg="light" expand="lg">
  <Container fluid>
    <Navbar.Toggle aria-controls="navbarScroll" />
    <Navbar.Collapse id="navbarScroll">
        <Nav.Link href="/"><Image src={ReactLogo} alt="React Logo" fluid/></Nav.Link>
      <Nav>
	  <Lin />
      </Nav>
    </Navbar.Collapse>
  </Container>
</Navbar>
  );
}

export default Navi
