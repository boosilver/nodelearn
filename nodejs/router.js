const express = require('express') 
const date = express()
const controller = require('./controller')

date.post('/asumdate',  (req, res)=>{
    // var sum = controller.sum()
    // res.sendStatus(sum)
    var sumdate = controller.a(req.body.a)
    res.send(sumdate);
    console.log(`a : ${req.body.a}`);
});

date.get('/sumdate',(req,res) =>{
    var sum = controller.sum()    
    res.send(sum)
});


date.get('/newdate', (req, res) => {

    var newdate = controller.newdate()
    res.send(newdate)
});

module.exports =date