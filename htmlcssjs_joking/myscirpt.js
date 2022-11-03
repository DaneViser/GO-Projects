
window.addEventListener("keydown", function(event) {
    if (event.key == "p"){
        changeCount()
    }
});
muda = 0
function changeCount(){
    muda++;
    if(muda > 10){muda = 0};
    var counter = document.getElementById("count")
    counter.innerHTML = muda + "/10"
    anime({
        targets: counter,
        color: ["#000","#FFA500"],
        borderColor:["#000","#FFA500","#000"],   
        scale:[1,1.5,0.9,1],
       
        
        easing: 'linear'
      });
      
     
    
}