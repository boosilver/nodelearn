const express = require('express') 
const app = express()
const controller = require('./controller')


app.get('/gettime', (req, res) => {

    var datetime = controller.ttt()
    // console.log(`datetime: ${datetime}`)
    // controller.getServerTime
    res.send(datetime)
})

app.get('/sumab/:a/:b',(req,res) =>{
    console.log(`a : ${req.params.a}`);
    console.log(`b : ${req.params.b}`);
    var sumab = controller.ab(req.params.a,req.params.b)
    console.log(`sum : ${sumab}`);
    res.send(sumab.toString())
    
}
)

app.post('/postsumab',(req,res)=>{

    console.log(`a : ${req.body.a}`);
    console.log(`b : ${req.body.b}`);
    var resual= controller.evenodd(req.body.a,req.body.b)
    console.log(`resual is : ${JSON.stringify(resual)}`);
    
    res.json(resual.sumc) //jsonรูปแบบของข้อมูลที่เป็น tree เก็บแบบ key value
     



}

)



module.exports =app