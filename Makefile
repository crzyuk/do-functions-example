deploy: 
	doctl serverless deploy .

invoke: 
	doctl serverless fn invoke crzy/${name}
