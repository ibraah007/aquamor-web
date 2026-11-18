/* Fading photo backgrounds (.bg), gallery filters and the enlarge (lightbox) view. */
(function(){
  try{
    var calm=window.matchMedia&&matchMedia("(prefers-reduced-motion: reduce)").matches;
    // Only the first photo of each banner is downloaded with the page. The
    // others keep their sources in data-src/data-srcset so a phone does not pay
    // for photos it never sees; each one loads just before it is shown and the
    // current photo stays on screen until the next one is ready.
    function load(pic){
      if(pic._loading)return pic._loading;
      var src=pic.querySelector("source[data-srcset]"),img=pic.querySelector("img");
      pic._loading=new Promise(function(resolve){
        if(!img){resolve();return}
        img.addEventListener("load",resolve,{once:true});
        img.addEventListener("error",resolve,{once:true});
        if(src&&src.hasAttribute("data-srcset")){src.srcset=src.getAttribute("data-srcset");src.removeAttribute("data-srcset")}
        if(img.hasAttribute("data-srcset")){img.srcset=img.getAttribute("data-srcset");img.removeAttribute("data-srcset")}
        if(img.hasAttribute("data-src")){img.src=img.getAttribute("data-src");img.removeAttribute("data-src")}
        if(img.complete)resolve();
      });
      return pic._loading;
    }
    document.querySelectorAll(".bg").forEach(function(bg,n){
      var s=[].slice.call(bg.children),i=0; if(!s.length)return;
      s[0].classList.add("on"); if(s.length<2||calm)return;
      var busy=false;
      function advance(){
        var j=(i+1)%s.length,pic=s[j];
        load(pic).then(function(){s[i].classList.remove("on");i=j;pic.classList.add("on");busy=false});
      }
      setTimeout(function(){setInterval(function(){if(!busy){busy=true;advance()}},5000)},n*400);
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
