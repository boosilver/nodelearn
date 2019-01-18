//Part สำหรับคำนวณโดยเฉพาะจ้าไว้ให้ router เรียก

function timeserver() {
    var d = new Date().toISOString().slice(0,10);
    return 'server time is :' + d;
}

function sumdate(a) {
    var a,sum;
    var y = new Date()
    A = y.getFullYear(),
    B = (y.getMonth() + 1),
    C = y.getDate(),
    sum = parseFloat((A+B+C)/callA(a)).toFixed(4)
    return 'sum is :' + (sum);
}

function callA(a){
    var a
    console.log(`a : ${a}`);
      a = (a);
      sumA = a
      return sumA ;
}

module.exports = { //ส่วนของการนำค่าออกไปให้กับ router
    serverdate: timeserver,
    datesum: sumdate,
    isa : callA
}