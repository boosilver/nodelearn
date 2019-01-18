const express = require('express')
const app = express()
const controller = require('./controller')

app.get('/', function (req, res){
    res.send('Hello You Are In Router API Part (router.js)')
})

app.get('/gettime', (req, res) => {
    var servertime = controller.serverdate()
    res.send(servertime)

})

app.get('/sumab/:a',(req,res) =>{
    console.log(`a : ${req.params.a}`);
    var sumab = controller.isa(req.params.a)
    res.send(sumab.toString())
    
}
)

app.post('/sumsum/', (req, res) => {
    var sumplus = controller.datesum(req.body.a)
    res.send(sumplus)
    res.json(sumplus)
})

module.exports =app