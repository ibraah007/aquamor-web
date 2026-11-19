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
    /* Each .bg is a fading slideshow. The controls are added to the section
       itself because .bg is hidden from screen readers. Auto-advance pauses
       while someone hovers, focuses or touches the section, and starts again
       when they leave; the arrows and a sideways swipe change the photo by
       hand. Only visible sections keep cycling, so off-screen photos are not
       downloaded for nothing. */
    document.querySelectorAll(".bg").forEach(function(bg,n){
      var s=[].slice.call(bg.children); if(!s.length)return;
      var host=bg.parentNode,total=s.length,i=0,busy=false,timer=null,paused=false,visible=true,touchTimer=null;
      s[0].classList.add("on");
      var vh=window.innerHeight||document.documentElement.clientHeight,rect=host.getBoundingClientRect();
      visible=rect.bottom>-200&&rect.top<vh+300;
      var counter=null;
      function schedule(delay){
        if(timer){clearTimeout(timer);timer=null}
        if(calm||total<2||paused||!visible||document.hidden)return;
        timer=setTimeout(function(){timer=null;step(1)},delay==null?5000:delay);
      }
      function show(j){
        if(j===i||busy)return;
        busy=true; var pic=s[j];
        load(pic).then(function(){
          s[i].classList.remove("on"); i=j; pic.classList.add("on");
          if(counter)counter.textContent=(i+1)+" / "+total;
          busy=false; schedule();
        });
      }
      function step(dir){
        if(total<2)return;
        show((i+dir+total)%total);
      }
      if(total>1){
        host.classList.add("has-bgslides");
        var ctl=document.createElement("div"); ctl.className="bgctl";
        ctl.setAttribute("role","group"); ctl.setAttribute("aria-label","Background photos");
        var prev=document.createElement("button"),next=document.createElement("button");
        prev.type="button"; next.type="button";
        prev.className="bgbtn bgprev"; next.className="bgbtn bgnext";
        prev.setAttribute("aria-label","Previous background photo");
        next.setAttribute("aria-label","Next background photo");
        prev.innerHTML='<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M15 5l-7 7 7 7"/></svg>';
        next.innerHTML='<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M9 5l7 7-7 7"/></svg>';
        counter=document.createElement("span"); counter.className="bgcount";
        counter.setAttribute("aria-hidden","true"); counter.textContent="1 / "+total;
        ctl.appendChild(prev); ctl.appendChild(counter); ctl.appendChild(next);
        host.appendChild(ctl);
        prev.addEventListener("click",function(){step(-1)});
        next.addEventListener("click",function(){step(1)});
        ctl.addEventListener("keydown",function(e){
          if(e.key==="ArrowLeft"){e.preventDefault();step(-1)}
          else if(e.key==="ArrowRight"){e.preventDefault();step(1)}
        });
        function lit(){
          host.classList.add("is-lit"); paused=true;
          if(timer){clearTimeout(timer);timer=null}
        }
        function dim(){
          if(host.contains(document.activeElement)&&document.activeElement!==document.body)return;
          host.classList.remove("is-lit"); paused=false; schedule();
        }
        host.addEventListener("pointerenter",function(e){if(e.pointerType==="mouse")lit()});
        host.addEventListener("pointerleave",function(e){if(e.pointerType==="mouse")dim()});
        host.addEventListener("focusin",lit);
        host.addEventListener("focusout",function(e){if(!host.contains(e.relatedTarget))dim()});
        host.addEventListener("touchstart",function(){
          lit();
          if(touchTimer)clearTimeout(touchTimer);
          touchTimer=setTimeout(dim,4000);
        },{passive:true});
        var sx=0,sy=0,tracking=false;
        host.addEventListener("pointerdown",function(e){
          if(e.pointerType==="mouse")return;
          if(!e.target.closest||e.target.closest("a,button,input,textarea,select,label,summary,video,iframe,.stage,.bottle"))return;
          tracking=true; sx=e.clientX; sy=e.clientY;
        },{passive:true});
        host.addEventListener("pointerup",function(e){
          if(!tracking)return; tracking=false;
          var dx=e.clientX-sx,dy=e.clientY-sy;
          if(Math.abs(dx)>45&&Math.abs(dx)>Math.abs(dy)*1.4)step(dx<0?1:-1);
        },{passive:true});
        host.addEventListener("pointercancel",function(){tracking=false},{passive:true});
        if("IntersectionObserver" in window){
          new IntersectionObserver(function(entries){
            entries.forEach(function(en){
              visible=en.isIntersecting;
              if(visible)schedule();
              else if(timer){clearTimeout(timer);timer=null}
            });
          },{rootMargin:"200px 0px"}).observe(host);
        }else{
          visible=true;
        }
        document.addEventListener("visibilitychange",function(){
          if(document.hidden){if(timer){clearTimeout(timer);timer=null}}
          else schedule();
        });
      }
      setTimeout(schedule,n*400);
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
