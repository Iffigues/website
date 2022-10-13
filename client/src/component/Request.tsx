const axios = require('axios').default;
axios.defaults.headers.post['Content-Type'] ='application/x-www-form-urlencoded';
class Request {
	private Api: any;

	constructor(baseUrl: string) {
		this.Api = axios.create({baseURL: `http://gopiko.fr:9090/`});	
  	}

	Post(endpoint: string, data: any) {
		const headers = {
       			 'Content-Type': 'application/json'
   		 };
		return this.Api.post(endpoint, JSON.stringify(data), headers);
	}
}

export default Request
