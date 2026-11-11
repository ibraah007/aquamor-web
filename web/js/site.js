/* Mobile navigation toggle. */
(function(){
  var btn=document.getElementById("navToggle");
  var menu=document.getElementById("navMenu");
  if(!btn||!menu)return;
  function close(){
    menu.classList.remove("open");
    btn.setAttribute("aria-expanded","false");
  }
  btn.addEventListener("click",function(){
    var open=menu.classList.toggle("open");
    btn.setAttribute("aria-expanded",open?"true":"false");
  });
  menu.addEventListener("click",function(e){
    if(e.target.closest("a"))close();
  });
  document.addEventListener("keydown",function(e){
    if(e.key==="Escape")close();
  });
  document.addEventListener("click",function(e){
    if(!btn.contains(e.target)&&!menu.contains(e.target))close();
  });
})();
