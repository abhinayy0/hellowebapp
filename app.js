const express = require('express');
const port = 5000;

const app = express();

app.get('/', function(req, res){
	res.send("Hello, World! from Javascript.")
})



app.listen(port, function(err){
	
	if(err){
	console.log("Error occured");
	}
});