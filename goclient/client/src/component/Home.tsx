import React, { useState, useEffect}  from 'react';
import Container from 'react-bootstrap/Container';
import Request from './Request'

function Home() {
	const [y, setY] = useState(1)
	const [fortune, setFortune] = useState("");
	const api = new Request("http://gopiko.fr/");

	useEffect(() => {
		if (y) 
  			getFortune();
		setY(0)
	}, []);

	function getFortune () {
		api.Post("/fortune", {
		}).then((resp: any) => {
			if (resp.data && resp.data.StdoutResponse) setFortune(resp.data.StdoutResponse)
		});
	};

return (
	<Container fluid>
				<pre style={{"textAlign": "initial"}}>{fortune}</pre>
	</Container>
	)
}

export default Home;
