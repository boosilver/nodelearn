const express = require('express')
const app = express()
var bodyParser = require('body-parser');

var port = process.env.PORT || 7777;
// parse application/json
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true
}));

app.get('/', function (req, res) {
    res.send('Hello This is HomePage Enjoy in HEll');
    
});

app.get('/index', function (req, res) {
    res.send('<h1>This is index page</h1>');
});

app.use('/myapi', require('./router'));



app.listen(port, function() {
    console.log('Starting node.js on port ' + port);
});
