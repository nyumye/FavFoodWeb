window.onload = function(){
    var pekapeka = this.document.getElementById("pekapeka-img")
    var deg = 0;
    this.setInterval(()=>{
        deg++;
        if(deg > 360){ deg = 0; }
        pekapeka.style.transform = "rotate(" + deg + "deg)";
    }, 50);
}