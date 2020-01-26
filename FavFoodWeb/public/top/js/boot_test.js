window.onload = function(){
    var donut = this.document.getElementById("donut")
    var deg = 0;
    this.setInterval(()=>{
        deg++;
        if(deg > 360){ deg = 0; }
        donut.style.transform = "rotate(" + deg + "deg)";
    }, 50);
}