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
    if(e.key!=="Escape"||!menu.classList.contains("open"))return;
    var hadFocus=menu.contains(document.activeElement);
    close();
    if(hadFocus)btn.focus();
  });
  document.addEventListener("click",function(e){
    if(!btn.contains(e.target)&&!menu.contains(e.target))close();
  });
})();

/* Scroll reveal. Elements get .reveal only after this script confirms they are
   off-screen, so a failed script still leaves the page readable and nothing
   flashes during load. */
(function(){
  var motion=window.matchMedia?window.matchMedia("(prefers-reduced-motion: reduce)"):null;
  if(!("IntersectionObserver" in window)||(motion&&motion.matches))return;
  var selectors=[
    "section > .wrap > .eyebrow",
    "section > .wrap > h2",
    "section > .wrap > .lead",
    "section .band > .wrap > .eyebrow",
    "section .band > .wrap > h2",
    "section .band > .wrap > .lead",
    "#order .build",
    "#planner .plan",
    "#how .steps > li",
    "#how .who",
    "#design .dgrid > *",
    "#cleaning .cgrid > *",
    "#gallery .grid > .g",
    "#gallery .clips > .clip",
    "#process .steps > li",
    "#about .aboutgrid > *",
    "#faq details",
    "#contact .box > div",
    "#contact .map"
  ].join(",");
  var nodes=[].slice.call(document.querySelectorAll(selectors));
  var reveal=[];
  var viewport=window.innerHeight||document.documentElement.clientHeight;
  nodes.forEach(function(el){
    if(el.closest("[hidden]"))return;
    var rect=el.getBoundingClientRect();
    if(rect.top<viewport*0.92&&rect.bottom>0)return;
    var delay=0;
    var prev=el.previousElementSibling;
    while(prev){
      if(prev.classList.contains("reveal"))delay++;
      prev=prev.previousElementSibling;
    }
    if(delay>6)delay=6;
    el.style.setProperty("--reveal-delay",(delay*60)+"ms");
    el.classList.add("reveal");
    if(el.tagName==="H2")el.classList.add("reveal-title");
    else if(el.classList.contains("g")||el.classList.contains("stage"))el.classList.add("reveal-zoom");
    reveal.push(el);
  });
  if(!reveal.length)return;
  var observer=new IntersectionObserver(function(entries){
    entries.forEach(function(entry){
      if(!entry.isIntersecting)return;
      entry.target.classList.add("is-in");
      observer.unobserve(entry.target);
    });
  },{rootMargin:"0px 0px -6% 0px",threshold:0.08});
  reveal.forEach(function(el){observer.observe(el)});
  function revealInView(){
    var vh=window.innerHeight||document.documentElement.clientHeight;
    reveal.forEach(function(el){
      var rect=el.getBoundingClientRect();
      if(rect.top<vh*0.92&&rect.bottom>0){
        el.classList.add("is-in");
        observer.unobserve(el);
      }
    });
  }
  function showAll(){
    reveal.forEach(function(el){
      el.classList.add("is-in");
      observer.unobserve(el);
    });
  }
  function onMotionChange(e){if(e.matches)showAll()}
  if(motion){
    if(motion.addEventListener)motion.addEventListener("change",onMotionChange);
    else if(motion.addListener)motion.addListener(onMotionChange);
  }
  /* A hash link or a restored scroll position can land on hidden elements
     before the observer has reported them. */
  window.addEventListener("load",revealInView);
  window.addEventListener("hashchange",revealInView);
  window.addEventListener("pageshow",revealInView);
})();

/* Scroll progress bar. */
(function(){
  var fill=document.getElementById("scrollFill");
  if(!fill)return;
  var ticking=false;
  function update(){
    var d=document.documentElement;
    var max=d.scrollHeight-d.clientHeight;
    fill.style.width=(max>0?(d.scrollTop/max)*100:0)+"%";
    ticking=false;
  }
  function queue(){
    if(ticking)return;
    ticking=true;
    requestAnimationFrame(update);
  }
  window.addEventListener("scroll",queue,{passive:true});
  window.addEventListener("resize",queue);
  window.addEventListener("load",update);
  update();
})();
