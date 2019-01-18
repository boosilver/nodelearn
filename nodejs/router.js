const express = require('express') 
const app = express()
const controller = require('./controller')

app.post('/asumdate',  (req, res)=>{
    // var sum = controller.sum()
    // res.sendStatus(sum)
    var sumdate = controller.a(req.body.a)
    res.send(sumdate);
    console.log(`a : ${req.body.a}`);
});

app.get('/sumdate',(req,res) =>{
    var sum = controller.sum()    
    res.send(sum)
});


app.get('/newdate', (req, res) => {

    var newdate = controller.newdate()
    res.send(newdate)
});

module.exports =app