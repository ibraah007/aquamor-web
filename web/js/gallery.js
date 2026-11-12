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
        var img=lb.querySelector(".lbimg");
        img.className="lbimg p-"+g.dataset.k;
        // Keep the English label so language switching can translate it back.
        img.dataset.al0=g.dataset.al0||g.getAttribute("aria-label");
        img.setAttribute("aria-label",g.getAttribute("aria-label"));
        document.getElementById("lbc").textContent=g.querySelector("span").textContent;
        lb.showModal();
      })});
      lb.addEventListener("click",function(e){if(e.target===lb||e.target.classList.contains("x"))lb.close()});
    }
  }catch(e){console.error("Aquamor gallery:",e)}
})();
