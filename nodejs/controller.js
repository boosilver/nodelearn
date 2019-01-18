function date() {
    var day =new Date();
    return 'date : '+day.getFullYear()+"-"+(day.getMonth()+1) +"-"+ day.getDate();
}

function sumdate() {
    var day = new Date();
    var sumdate = day.getFullYear()+(day.getMonth()+1) + day.getDate();
    sumdate = parseFloat(sumdate)
    return 'sumdate : '+sumdate;
}

function suma(a){
    var a,sum;
    a = parseFloat(a);
    var day = new Date();
    var sum = day.getFullYear()+(day.getMonth()+1) + day.getDate();
    sum = parseFloat(sum)
    var total = parseFloat(sum/a).toFixed(4);
    return 'sumdate : '+sum+"/"+a+"="+total;
}




module.exports = {
    newdate: date,
    sum : sumdate,
    a : suma,
}