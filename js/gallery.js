/* Fading photo backgrounds (.bg), gallery filters and the enlarge (lightbox) view. */
(function(){
  try{
    var calm=window.matchMedia&&matchMedia("(prefers-reduced-motion: reduce)").matches;
    document.querySelectorAll(".bg").forEach(function(bg,n){
      var s=bg.children,i=0; if(!s.length)return; s[0].classList.add("on"); if(s.length<2||calm)return;
      setTimeout(function(){setInterval(function(){s[i].classList.remove("on");i=(i+1)%s.length;s[i].classList.add("on")},6000)},n*1500);
    });
    var tabs=document.querySelectorAll(".gallery .tab"),tiles=document.querySelectorAll(".g");
    tabs.forEach(function(t){t.addEventListener("click",function(){
      tabs.forEach(function(x){x.setAttribute("aria-pressed",x===t)});
      tiles.forEach(function(g){g.classList.toggle("hide",t.dataset.f!=="all"&&g.dataset.c.split(" ").indexOf(t.dataset.f)<0)});
    })});
    var lb=document.getElementById("lb");
    if(lb&&lb.showModal){
      tiles.forEach(function(g){g.addEventListener("click",function(){
        var img=lb.querySelector(".lbimg"),cap=g.querySelector("span"),label=cap?cap.textContent:"";
        img.className="lbimg p-"+g.dataset.k;
        // The visible caption is also the accessible name for the enlarged photo.
        img.dataset.al0=label;
        img.setAttribute("aria-label",label);
        document.getElementById("lbc").textContent=label;
        lb.showModal();
      })});
      lb.addEventListener("click",function(e){if(e.target===lb||e.target.classList.contains("x"))lb.close()});
    }
  }catch(e){console.error("Aquamor gallery:",e)}
})();
