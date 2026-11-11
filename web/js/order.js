/* Order builder and label designer. Edit prices in PRICE and the item lists in items. */
(function(){
  var WA="254712140006";
  var items={
    refill:[["5L refill","Ksh 50 · bring your container",34],["10L refill","Ksh 100 · bring your container",44],["20L refill","Ksh 200 · bring your container",56]],
    bottle:[["300ml bottle","Bottled drinking water",30],["500ml bottle","Bottled drinking water",38],["1L bottle","Bottled drinking water",48],["1.5L bottle","Bottled drinking water",56]]
  };
  items.container=[["20L hard bottle + water","Ksh 1,400 · bottle and water",56],["20L regular bottle","Ksh 450",56]];
  var PRICE={"5L refill":50,"10L refill":100,"20L refill":200,"20L hard bottle + water":1400,"20L regular bottle":450}, qty={}, tab="refill", mode="Delivery";
  var rows=document.getElementById("rows"), list=document.getElementById("sumList");
  function esc(s){return String(s).replace(/[&<>"]/g,function(c){return{"&":"&amp;","<":"&lt;",">":"&gt;",'"':"&quot;"}[c]})}
  function sw(){return window.LANG==="sw"}
  /* English text from the page dictionary, with a direct fallback for messages. */
  function tr(en){return (sw()&&window.__sw&&window.__sw[en])||en}
  function t(en,swText){return sw()?swText:en}
  function drawRows(){
    var source=items[tab]||[];
    rows.innerHTML=source.map(function(it){
      var k=it[0],h=it[2],w=Math.round(h*0.6);
      return '<div class="row"><div class="pic"><i style="--w:'+w+'px;--h:'+h+'px"></i></div><div class="nm">'+esc(k)+'<small>'+esc(it[1])+'</small></div><div class="step"><button data-k="'+esc(k)+'" data-d="-1" aria-label="'+(sw()?"Punguza ":"Less ")+esc(tr(k))+'">−</button><output>'+(qty[k]||0)+'</output><button data-k="'+esc(k)+'" data-d="1" aria-label="'+(sw()?"Ongeza ":"More ")+esc(tr(k))+'">+</button></div></div>';
    }).join("");
  }
  function drawSum(){
    var ks=Object.keys(qty).filter(function(k){return qty[k]>0});
    var tot=0,unpriced=false;
    ks.forEach(function(k){
      if(Object.prototype.hasOwnProperty.call(PRICE,k)){tot+=(PRICE[k]||0)*qty[k];}
      else{unpriced=true;}
    });
    if(!ks.length){list.innerHTML='<li class="empty">Nothing added yet.</li>';return}
    var html=ks.map(function(k){return '<li><span>'+esc(k)+'</span><b>× '+qty[k]+'</b></li>'}).join("");
    if(tot>0){html+='<li><span>Estimated total (listed items)</span><b>Ksh '+tot.toLocaleString("en-US")+'</b></li>'}
    if(unpriced){html+='<li class="empty">Bottled water prices are confirmed on WhatsApp.</li>'}
    list.innerHTML=html;
  }
  rows.addEventListener("click",function(e){
    var b=e.target.closest("button[data-k]"); if(!b)return;
    var k=b.dataset.k; qty[k]=Math.max(0,(qty[k]||0)+Number(b.dataset.d));
    b.parentNode.querySelector("output").textContent=qty[k]; drawSum();
  });
  /* Order tabs only. Gallery filters also use .tab and must not reset this panel. */
  document.querySelectorAll("#order .tabs .tab").forEach(function(tb){tb.addEventListener("click",function(){
    tab=tb.dataset.t; document.querySelectorAll("#order .tabs .tab").forEach(function(x){x.setAttribute("aria-pressed",x===tb)}); drawRows();
  })});
  document.querySelectorAll("#order .seg button").forEach(function(b){b.addEventListener("click",function(){
    mode=b.dataset.m; document.querySelectorAll("#order .seg button").forEach(function(x){x.setAttribute("aria-pressed",x===b)});
    document.getElementById("areaLbl").textContent=mode==="Delivery"?"Delivery area":"Name for pickup";
    var area=document.getElementById("area");
    area.dataset.ph0=mode==="Delivery"?"e.g. Ngara, Pangani, Eastleigh":"e.g. Your name";
    if(window.__setLang)window.__setLang(window.LANG||"en");
  })});
  function go(text){window.open("https://wa.me/"+WA+"?text="+encodeURIComponent(text),"_blank","noopener")}
  document.getElementById("sendOrder").addEventListener("click",function(e){
    e.preventDefault();
    var ks=Object.keys(qty).filter(function(k){return qty[k]>0});
    if(!ks.length){list.innerHTML='<li class="empty">Add at least one item first.</li>';return}
    var a=document.getElementById("area").value.trim();
    var tot=0,unpriced=false;
    ks.forEach(function(k){
      if(Object.prototype.hasOwnProperty.call(PRICE,k)){tot+=(PRICE[k]||0)*qty[k];}
      else{unpriced=true;}
    });
    var message=t("Hello Aquamor, I would like to order:\n","Habari Aquamor, ningependa kuagiza:\n");
    message+=ks.map(function(k){return "- "+tr(k)+" x "+qty[k]}).join("\n");
    if(tot>0){message+=(sw()?"\nJumla ya makadirio (bidhaa zilizo na bei): Ksh ":"\nEstimated total (listed items): Ksh ")+tot.toLocaleString("en-US")}
    if(unpriced){message+=(sw()?"\nBei za maji ya chupa zitathibitishwa kupitia WhatsApp.":"\nBottled water prices will be confirmed on WhatsApp.")}
    message+="\n"+tr(mode)+(a?": "+a:"");
    message+=(sw()?"\nTafadhali nithibitishie bei.":"\nPlease confirm the price.");
    go(message);
  });

  var ev=document.getElementById("ev"),msg=document.getElementById("msg"),t1=document.getElementById("t1"),t2=document.getElementById("t2"),lbl=document.getElementById("lbl"),col="#06476a";
  function fit(el,max,base,min){
    el.removeAttribute("textLength");
    var size=base;
    el.setAttribute("font-size",size);
    if(!el.getComputedTextLength)return;
    while(el.getComputedTextLength()>max&&size>min){size-=0.5;el.setAttribute("font-size",size)}
    // Extreme input still gets one final squeeze rather than being clipped.
    if(el.getComputedTextLength()>max){el.setAttribute("textLength",max);el.setAttribute("lengthAdjust","spacingAndGlyphs")}
  }
  function updateLabel(){
    t1.textContent=ev.value||t("Your event","Hafla yako");
    t2.textContent=msg.value||t("date or message","tarehe au ujumbe");
    fit(t1,54,12,4.5);
    fit(t2,54,8,4.5);
  }
  ev.addEventListener("input",updateLabel);
  msg.addEventListener("input",updateLabel);
  document.addEventListener("aquamor:lang",function(){drawRows();updateLabel()});
  document.querySelectorAll("#sw button").forEach(function(b){b.addEventListener("click",function(){
    col=b.dataset.c; lbl.setAttribute("fill",col); document.querySelectorAll("#sw button").forEach(function(x){x.setAttribute("aria-pressed",x===b)});
  })});
  document.getElementById("sendDesign").addEventListener("click",function(e){
    e.preventDefault();
    var names={"#06476a":"deep blue","#7a1f3d":"burgundy","#1d6b43":"green","#4b2c82":"purple","#1a1a1a":"black"};
    var namesSw={"#06476a":"bluu iliyokolea","#7a1f3d":"rangi ya divai","#1d6b43":"kijani","#4b2c82":"zambarau","#1a1a1a":"nyeusi"};
    var color=sw()?namesSw[col]:names[col];
    var notSet=t("(not set)","(hujaweka)");
    go(t("Hello Aquamor, I would like a branded bottle:\nEvent or name: ","Habari Aquamor, ningependa chupa yenye chapa yangu:\nHafla au jina: ")+(ev.value||notSet)+
      t("\nMessage: ","\nUjumbe: ")+(msg.value||notSet)+
      t("\nLabel color: ","\nRangi ya lebo: ")+color+
      t("\nBottle size: ","\nUkubwa wa chupa: ")+document.getElementById("sz").value+
      t("\nPlease tell me the quantity options and price.","\nTafadhali niambie idadi zinazopatikana na bei."));
  });
  drawRows();
  updateLabel();
  if(document.fonts&&document.fonts.ready){document.fonts.ready.then(updateLabel)}
})();
