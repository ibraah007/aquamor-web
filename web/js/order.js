/* Order builder and label designer. Edit prices in PRICE and the item lists in items. */
(function(){
  var WA="254712140006";
  var items={
    refill:[["5L refill","Ksh 50 · bring your container",34],["10L refill","Ksh 100 · bring your container",44],["20L refill","Ksh 200 · includes water",56]],
    bottle:[
      ["200ml bottle","Price on quotation",24],
      ["300ml bottle","Price on quotation",30],
      ["500ml bottle","Price on quotation",38],
      ["1L bottle","Price on quotation",48],
      ["1.5L bottle","Price on quotation",56],
      ["2L bottle","Price on quotation",64],
      ["3L bottle","Price on quotation",72],
      ["5L bottle","Price on quotation",80],
      ["10L bottle","Price on quotation",96],
      ["20L bottle","Price on quotation",112]
    ]
  };
  items.container=[["20L regular bottle","Ksh 450–550 · bottle + water",56],["20L hard bottle","Ksh 1,000–1,700 · bottle + water",56],["10L bottle","Ksh 200–250 · bottle + water",44],["5L bottle","Ksh 100 · bottle + water",34]];
  var PRICE={"5L refill":50,"10L refill":100,"20L refill":200,"20L regular bottle":450,"20L hard bottle":1400,"10L bottle":200,"5L bottle":100}, qty={}, tab="refill", mode="Delivery";
  /* Catalog IDs for online payment. Only items with confirmed prices are here. */
  var PRODUCTS={"5L refill":{id:3,option:"swap"},"10L refill":{id:2,option:"swap"},"20L refill":{id:1,option:"swap"},"20L hard bottle":{id:4,option:"new_purchase"},"20L regular bottle":{id:5,option:"new_purchase"}};
  var rows=document.getElementById("rows"), list=document.getElementById("sumList");
  var nameInput=document.getElementById("customerName"), phoneInput=document.getElementById("customerPhone"), orderError=document.getElementById("orderError");
  var LAST_KEY="aquamor-last-order", reorderBtn=document.getElementById("reorderBtn"), reorderNote=document.getElementById("reorderNote"), note="";
  function esc(s){return String(s).replace(/[&<>"]/g,function(c){return{"&":"&amp;","<":"&lt;",">":"&gt;",'"':"&quot;"}[c]})}
  function sw(){return window.LANG==="sw"}
  /* English text from the page dictionary, with a direct fallback for messages. */
  function tr(en){return (sw()&&window.__sw&&window.__sw[en])||en}
  function t(en,swText){return sw()?swText:en}
  /* Acknowledgment shown after any WhatsApp or email hand-off. */
  var toast=document.getElementById("toast"), toastTimer=null;
  function showThanks(){
    if(!toast)return;
    toast.textContent=t(
      "Thanks for reaching out. Your message is ready in WhatsApp or email; send it and Aquamor will get back to you.",
      "Asante kwa kuwasiliana. Ujumbe wako uko tayari kwenye WhatsApp au barua pepe; tuma na Aquamor atakujibu."
    );
    toast.hidden=false;
    requestAnimationFrame(function(){toast.classList.add("show")});
    if(toastTimer)clearTimeout(toastTimer);
    toastTimer=setTimeout(function(){
      toast.classList.remove("show");
      toastTimer=setTimeout(function(){toast.hidden=true},300);
    },7000);
  }
  document.addEventListener("click",function(e){
    var link=e.target.closest&&e.target.closest('a[href*="wa.me/"],a[href^="mailto:"]');
    if(link)showThanks();
  });
  /* Last order: saved on send, restored by the Reorder button. */
  function knownKeys(){var all={};Object.keys(items).forEach(function(g){items[g].forEach(function(it){all[it[0]]=1})});return all}
  function readLast(){
    try{
      var raw=localStorage.getItem(LAST_KEY); if(!raw)return null;
      var o=JSON.parse(raw); if(!o||typeof o!=="object"||!o.qty)return null;
      return o;
    }catch(e){return null}
  }
  function showNote(kind){
    note=kind||"";
    if(!reorderNote)return;
    reorderNote.hidden=!note;
    if(!note)return;
    reorderNote.textContent=note==="gone"?t("Your last order items are no longer available.","Bidhaa za oda yako ya mwisho hazipatikani tena."):
      note==="dropped"?t("Some items from your last order are no longer available.","Baadhi ya bidhaa za oda yako ya mwisho hazipatikani tena."):
      t("Last order restored. Check it, then send on WhatsApp.","Oda yako ya mwisho imerejeshwa. Iangalie, kisha itume kupitia WhatsApp.");
  }
  function saveLast(keys,area,savedQty,savedMode){
    try{
      var q={};
      keys.forEach(function(k){q[k]=(savedQty&&savedQty[k]!=null)?savedQty[k]:qty[k]});
      var saved={qty:q,mode:savedMode||mode,area:area,ts:Date.now()};
      localStorage.setItem(LAST_KEY,JSON.stringify(saved));
    }catch(e){}
  }
  function refreshReorder(){if(reorderBtn)reorderBtn.hidden=!readLast()}
  function applyLast(){
    var last=readLast(); if(!last)return;
    var known=knownKeys(),next={},dropped=false;
    Object.keys(last.qty||{}).forEach(function(k){
      var n=Math.floor(Number(last.qty[k]));
      if(n>0){ if(known[k])next[k]=n; else dropped=true; }
    });
    if(!Object.keys(next).length){showNote("gone");return}
    qty=next;
    var first=Object.keys(next)[0],target="refill";
    if(items.bottle.some(function(it){return it[0]===first}))target="bottle";
    else if(items.container.some(function(it){return it[0]===first}))target="container";
    tab=target;
    document.querySelectorAll("#order .tabs .tab").forEach(function(x){x.setAttribute("aria-pressed",x.dataset.t===target)});
    var m=last.mode==="Pickup"?"Pickup":"Delivery";
    var mb=document.querySelector('#order .seg [data-m="'+m+'"]');
    if(mb&&mb.getAttribute("aria-pressed")!=="true")mb.click();
    var area=document.getElementById("area");
    if(area&&typeof last.area==="string")area.value=last.area;
    drawRows(); drawSum();
    showNote(dropped?"dropped":"ok");
  }
  function drawRows(){
    var source=items[tab]||[];
    rows.innerHTML=source.map(function(it){
      var k=it[0],h=it[2],w=Math.round(h*0.6);
      return '<div class="row"><div class="pic"><i style="--w:'+w+'px;--h:'+h+'px"></i></div><div class="nm">'+esc(k)+'<small>'+esc(it[1])+'</small></div><div class="step"><button data-k="'+esc(k)+'" data-d="-1" aria-label="'+(sw()?"Punguza ":"Less ")+esc(tr(k))+'">−</button><output>'+(qty[k]||0)+'</output><button data-k="'+esc(k)+'" data-d="1" aria-label="'+(sw()?"Ongeza ":"More ")+esc(tr(k))+'">+</button></div></div>';
    }).join("");
  }
  function setOrderError(text){if(!orderError)return;orderError.hidden=!text;orderError.textContent=text||""}
  function orderDetails(){
    var name=(nameInput?nameInput.value:"").trim();
    var phone=(phoneInput?phoneInput.value:"").trim();
    var area=document.getElementById("area").value.trim();
    if(name.length<2){setOrderError(t("Please enter your name.","Tafadhali weka jina lako."));if(nameInput)nameInput.focus();return null}
    if(phone.replace(/\D/g,"").length<7){setOrderError(t("Please enter a valid phone number.","Tafadhali weka namba sahihi ya simu."));if(phoneInput)phoneInput.focus();return null}
    if(area.length<2){setOrderError(mode==="Delivery"?t("Please enter your delivery area.","Tafadhali weka eneo la kuletewa."):t("Please enter a pickup name or note.","Tafadhali weka jina la kuchukua au maelezo."));var areaEl=document.getElementById("area");if(areaEl)areaEl.focus();return null}
    return {name:name,phone:phone,area:area};
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
    if(tot>0){html+='<li><span>Estimated total (from listed prices)</span><b>Ksh '+tot.toLocaleString("en-US")+'</b></li>'}
    if(unpriced){html+='<li class="empty">Some items are quoted on request. We will confirm the price on WhatsApp.</li>'}
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
    setOrderError("");
    document.getElementById("areaLbl").textContent=mode==="Delivery"?"Delivery area":"Pickup name or note";
    var area=document.getElementById("area");
    area.dataset.ph0=mode==="Delivery"?"e.g. Ngara, Pangani, Eastleigh":"e.g. Pickup at 5pm";
    if(window.__setLang)window.__setLang(window.LANG||"en");
  })});
  [nameInput,phoneInput,document.getElementById("area")].forEach(function(el){if(el)el.addEventListener("input",function(){setOrderError("")})});
  if(reorderBtn)reorderBtn.addEventListener("click",applyLast);
  refreshReorder();
  function go(text){window.open("https://wa.me/"+WA+"?text="+encodeURIComponent(text),"_blank","noopener")}
  var lastSaved={sig:"",at:0}, saving={};
  function saveOrder(keys,details){
    if(!window.fetch)return;
    var priced=true;
    keys.forEach(function(k){if(!PRODUCTS[k])priced=false});
    if(!priced)return;
    var payload={
      customer_name:details.name,
      customer_phone:details.phone,
      delivery_estate:mode==="Delivery"?details.area:"Pickup",
      delivery_address:details.area,
      items:keys.map(function(k){return {product_id:PRODUCTS[k].id,bottle_option:PRODUCTS[k].option,quantity:qty[k]}})
    };
    var sig=JSON.stringify(payload);
    if(sig===lastSaved.sig&&Date.now()-lastSaved.at<60000)return;
    if(saving[sig])return;
    saving[sig]=true;
    fetch("/api/v1/orders",{method:"POST",headers:{"Content-Type":"application/json",Accept:"application/json"},credentials:"same-origin",keepalive:true,body:sig})
      .then(function(r){
        if(!r.ok)throw new Error("HTTP "+r.status);
        lastSaved.sig=sig;lastSaved.at=Date.now();
      })
      .catch(function(err){if(window.console&&console.warn)console.warn("Could not save order",err)})
      .then(function(){delete saving[sig]});
  }
  document.getElementById("sendOrder").addEventListener("click",function(e){
    e.preventDefault();
    var ks=Object.keys(qty).filter(function(k){return qty[k]>0});
    if(!ks.length){list.innerHTML='<li class="empty">Add at least one item first.</li>';return}
    var details=orderDetails();
    if(!details)return;
    var tot=0,unpriced=false;
    ks.forEach(function(k){
      if(Object.prototype.hasOwnProperty.call(PRICE,k)){tot+=(PRICE[k]||0)*qty[k];}
      else{unpriced=true;}
    });
    var message=t("Hello Aquamor, I would like to order:\n","Habari Aquamor, ningependa kuagiza:\n");
    message+=ks.map(function(k){return "- "+tr(k)+" x "+qty[k]}).join("\n");
    if(tot>0){message+=(sw()?"\nJumla ya makadirio (kutoka bei zilizoorodheshwa): Ksh ":"\nEstimated total (from listed prices): Ksh ")+tot.toLocaleString("en-US")}
    if(unpriced){message+=(sw()?"\nBaadhi ya bidhaa zinauzwa kwa bei ya kuuliza; tafadhali thibitisha bei.":"\nSome items are quoted on request; please confirm their prices.")}
    message+="\n"+t("Name","Jina")+": "+details.name;
    message+="\n"+t("Phone","Simu")+": "+details.phone;
    message+="\n"+tr(mode)+(details.area?": "+details.area:"");
    message+=(sw()?"\nTafadhali nithibitishie bei.":"\nPlease confirm the price.");
    saveLast(ks,details.area);
    refreshReorder();
    saveOrder(ks,details);
    go(message);
  });

  /* Online payment. The backend prices the order again and confirms the total. */
  var payBtn=document.getElementById("payOrder"),payNote=document.getElementById("payNote"),payDlg=document.getElementById("payDlg"),payForm=document.getElementById("payForm"),payError=document.getElementById("payError"),payTotal=document.getElementById("payTotal"),paySubmit=document.getElementById("paySubmit"),payClose=document.getElementById("payClose"),payAmount=0,payBusy=false;
  function pricedOrder(){
    var ks=Object.keys(qty).filter(function(k){return qty[k]>0}),tot=0,unknown=[];
    ks.forEach(function(k){
      if(Object.prototype.hasOwnProperty.call(PRICE,k)){tot+=(PRICE[k]||0)*qty[k]}
      if(!PRODUCTS[k])unknown.push(k);
    });
    return {keys:ks,total:tot,unknown:unknown};
  }
  function setPayNote(text){if(!payNote)return;payNote.hidden=!text;payNote.textContent=text||""}
  function setPayError(text){if(!payError)return;payError.hidden=!text;payError.textContent=text||""}
  function payErrorText(message){
    var map={
      "Online payment is being set up. Please send your order on WhatsApp and we will confirm the total.":t("Online payment is being set up. Please send your order on WhatsApp and we will confirm the total.","Malipo ya mtandaoni yanaandaliwa. Tafadhali tuma oda yako kupitia WhatsApp nasi tutathibitisha jumla."),
      "Could not start the payment. Please try again or send your order on WhatsApp.":t("Could not start the payment. Please try again or send your order on WhatsApp.","Imeshindikana kuanzisha malipo. Tafadhali jaribu tena au tuma oda yako kupitia WhatsApp."),
      "A valid email address is needed for the payment receipt":t("A valid email address is needed for the payment receipt","Barua pepe sahihi inahitajika kwa risiti ya malipo")
    };
    return map[message]||message;
  }
  function renderPayTotal(){if(payTotal)payTotal.textContent=t("Estimated total: Ksh ","Jumla ya makadirio: Ksh ")+payAmount.toLocaleString("en-US")}
  function setPayBusy(busy){payBusy=!!busy;if(!paySubmit)return;paySubmit.disabled=payBusy;paySubmit.textContent=payBusy?t("Starting payment…","Tunanza malipo…"):t("Continue to payment","Endelea kulipa")}
  if(payBtn){
    fetch("/api/v1/payments/config",{headers:{Accept:"application/json"},credentials:"same-origin"})
      .then(function(r){return r.ok?r.json():null})
      .then(function(cfg){if(cfg&&cfg.enabled)payBtn.hidden=false})
      .catch(function(){});
    payBtn.addEventListener("click",function(){
      if(payBusy)return;
      var o=pricedOrder();
      if(!o.keys.length){list.innerHTML='<li class="empty">Add at least one item first.</li>';return}
      if(o.unknown.length){setPayNote(t("Online payment is available for refills and 20L bottles. Please send other items on WhatsApp.","Malipo ya mtandaoni yanapatikana kwa kujaza maji na chupa za 20L. Tafadhali tuma bidhaa nyingine kupitia WhatsApp."));return}
      setPayNote("");
      setOrderError("");
      var area=document.getElementById("area"),payArea=document.getElementById("payArea");
      if(payArea&&area)payArea.value=area.value;
      var payNameEl=document.getElementById("payName"),payPhoneEl=document.getElementById("payPhone");
      if(payNameEl&&!payNameEl.value&&nameInput)payNameEl.value=nameInput.value.trim();
      if(payPhoneEl&&!payPhoneEl.value&&phoneInput)payPhoneEl.value=phoneInput.value.trim();
      payAmount=o.total;
      renderPayTotal();
      setPayError("");
      setPayBusy(false);
      payDlg.showModal();
      var first=document.getElementById("payName");
      if(first&&!first.value)first.focus();
    });
  }
  if(payClose)payClose.addEventListener("click",function(){payDlg.close()});
  if(payDlg)payDlg.addEventListener("click",function(e){if(e.target===payDlg)payDlg.close()});
  if(payForm)payForm.addEventListener("submit",function(e){
    e.preventDefault();
    if(payBusy)return;
    var o=pricedOrder();
    if(!o.keys.length||o.unknown.length){setPayError(t("Please choose refills or 20L bottles to pay online.","Tafadhali chagua kujaza maji au chupa za 20L kulipa mtandaoni."));return}
    var name=(document.getElementById("payName").value||"").trim();
    var phone=(document.getElementById("payPhone").value||"").trim();
    var email=(document.getElementById("payEmail").value||"").trim();
    var area=(document.getElementById("payArea").value||"").trim();
    if(name.length<2){setPayError(t("Please enter your name.","Tafadhali weka jina lako."));return}
    if(phone.replace(/\D/g,"").length<7){setPayError(t("Please enter a valid phone number.","Tafadhali weka namba sahihi ya simu."));return}
    if(!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)){setPayError(t("Please enter a valid email for the receipt.","Tafadhali weka barua pepe sahihi kwa risiti."));return}
    if(area.length<2){setPayError(t("Please enter your delivery area or pickup name.","Tafadhali weka eneo la kuletewa au jina la kuchukua."));return}
    var savedQty={},orderItems=o.keys.map(function(k){
      savedQty[k]=qty[k];
      return {product_id:PRODUCTS[k].id,bottle_option:PRODUCTS[k].option,quantity:qty[k]};
    });
    var savedMode=mode;
    setPayError("");
    setPayBusy(true);
    fetch("/api/v1/payments/checkout",{method:"POST",headers:{"Content-Type":"application/json",Accept:"application/json"},credentials:"same-origin",body:JSON.stringify({customer_name:name,customer_phone:phone,customer_email:email,delivery_estate:savedMode==="Delivery"?area:"Pickup",delivery_address:area,items:orderItems})})
      .then(function(r){return r.json().catch(function(){return {}}).then(function(data){if(!r.ok)throw new Error(data.error||t("Could not start the payment. Please try again.","Imeshindikana kuanzisha malipo. Tafadhali jaribu tena."));return data})})
      .then(function(data){
        if(!data.payment_url)throw new Error(t("The payment page was not returned. Please try again.","Ukurasa wa malipo haukurejeshwa. Tafadhali jaribu tena."));
        saveLast(o.keys,area,savedQty,savedMode);
        refreshReorder();
        window.location.assign(data.payment_url);
      })
      .catch(function(err){setPayBusy(false);setPayError(payErrorText(err.message)||t("Could not start the payment. Please try again.","Imeshindikana kuanzisha malipo. Tafadhali jaribu tena."))});
  });
  document.addEventListener("aquamor:lang",function(){if(paySubmit)paySubmit.textContent=paySubmit.disabled?t("Starting payment…","Tunanza malipo…"):t("Continue to payment","Endelea kulipa");setOrderError("");renderPayTotal()});

  var ev=document.getElementById("ev"),msg=document.getElementById("msg"),t1=document.getElementById("t1"),t2=document.getElementById("t2"),t3=document.getElementById("t3"),lbl=document.getElementById("lbl"),lblRound=document.getElementById("lblRound"),lblOval=document.getElementById("lblOval"),swPick=document.getElementById("swPick"),col="#06476a",style="rect";
  var COLOR_NAMES={"#1a1a1a":"black","#f4c430":"yellow","#8bc34a":"light green","#1d6b43":"dark green","#c62828":"red","#0077be":"ocean blue","#ffffff":"white","#06476a":"blue"};
  var COLOR_NAMES_SW={"#1a1a1a":"nyeusi","#f4c430":"njano","#8bc34a":"kijani kibichi","#1d6b43":"kijani kikolevu","#c62828":"nyekundu","#0077be":"bluu ya bahari","#ffffff":"nyeupe","#06476a":"bluu"};
  function colorName(){
    var n=(sw()?COLOR_NAMES_SW:COLOR_NAMES)[col]||col;
    return n.charAt(0).toUpperCase()+n.slice(1);
  }
  function updateColorPick(){
    if(swPick)swPick.textContent=t("Selected color","Rangi iliyochaguliwa")+": "+colorName();
  }
  var PRESETS={
    wedding:{color:"#c62828",en:["John & Mary","12 December 2026"],sw:["John na Mary","12 Desemba 2026"]},
    church:{color:"#06476a",en:["Grace Church","Sunday service"],sw:["Kanisa la Neema","Ibada ya Jumapili"]},
    graduation:{color:"#1d6b43",en:["Class of 2026","Graduation day"],sw:["Darasa la 2026","Siku ya mahafali"]},
    corporate:{color:"#0077be",en:["Your company","Team event"],sw:["Kampuni yako","Hafla ya timu"]}
  };
  function fit(el,max,base,min){
    el.removeAttribute("textLength");
    var size=base;
    el.setAttribute("font-size",size);
    if(!el.getComputedTextLength)return;
    while(el.getComputedTextLength()>max&&size>min){size-=0.5;el.setAttribute("font-size",size)}
    // Extreme input still gets one final squeeze rather than being clipped.
    if(el.getComputedTextLength()>max){el.setAttribute("textLength",max);el.setAttribute("lengthAdjust","spacingAndGlyphs")}
  }
  function lightColor(c){
    var m=c.replace("#","");if(m.length===3)m=m.replace(/./g,"$&$&");
    return (0.299*parseInt(m.substr(0,2),16)+0.587*parseInt(m.substr(2,2),16)+0.114*parseInt(m.substr(4,2),16))>150;
  }
  function paintLabel(c){
    [lbl,lblRound,lblOval].forEach(function(s){if(!s)return;
      s.setAttribute("fill",c);
      var white=c.toLowerCase()==="#ffffff";
      s.setAttribute("stroke",white?"rgba(6,71,106,.35)":"none");
      s.setAttribute("stroke-width",white?"1":"0");
    });
    var ink=lightColor(c)?"#14212b":"#ffffff";
    [t1,t2,t3].forEach(function(x){if(x)x.setAttribute("fill",ink)});
  }
  function updateLabel(){
    var compact=style!=="rect";
    t1.textContent=ev.value||t("Your event","Hafla yako");
    t2.textContent=msg.value||t("date or message","tarehe au ujumbe");
    fit(t1,compact?48:54,12,4.5);
    fit(t2,compact?44:54,8,4.5);
  }
  function setColor(c){
    col=c; paintLabel(c); updateColorPick();
    document.querySelectorAll("#sw button").forEach(function(x){x.setAttribute("aria-pressed",x.dataset.c===c)});
  }
  function applyStyle(){
    var round=style==="round",oval=style==="oval",compact=round||oval;
    lbl.style.display=compact?"none":"";
    lblRound.style.display=round?"":"none";
    lblOval.style.display=oval?"":"none";
    t1.setAttribute("y",compact?"176":"170");
    t2.setAttribute("y",compact?"192":"190");
    t3.style.display=compact?"none":"";
    updateLabel();
  }
  /* Turn the label around the bottle to give the preview a 3D feel. */
  var svgPrev=document.getElementById("prev"), labelGroup=document.getElementById("labelGroup"), rotInput=document.getElementById("rot"), angle=0;
  function applyRotation(){
    if(!labelGroup)return;
    var a=angle*Math.PI/180, sx=Math.cos(a), dx=60*Math.sin(a);
    labelGroup.setAttribute("transform","translate("+dx.toFixed(2)+" 0) translate(60 0) scale("+Math.max(sx,0.06).toFixed(3)+" 1) translate(-60 0)");
    if(rotInput&&String(rotInput.value)!==String(angle))rotInput.value=angle;
  }
  function setAngle(a){angle=Math.max(-72,Math.min(72,a));applyRotation()}
  if(rotInput)rotInput.addEventListener("input",function(){setAngle(Number(rotInput.value))});
  if(svgPrev){
    var drag=null;
    svgPrev.addEventListener("pointerdown",function(e){
      drag={x:e.clientX,a:angle,id:e.pointerId};
      try{svgPrev.setPointerCapture(e.pointerId)}catch(_){}
    });
    svgPrev.addEventListener("pointermove",function(e){
      if(!drag||e.pointerId!==drag.id)return;
      setAngle(drag.a+(e.clientX-drag.x)/2);
    });
    function endDrag(e){
      if(!drag)return;
      drag=null;
      try{svgPrev.releasePointerCapture(e.pointerId)}catch(_){}
    }
    svgPrev.addEventListener("pointerup",endDrag);
    svgPrev.addEventListener("pointercancel",endDrag);
  }
  ev.addEventListener("input",updateLabel);
  msg.addEventListener("input",updateLabel);
  document.addEventListener("aquamor:lang",function(){drawRows();updateLabel();updateColorPick();if(note)showNote(note)});
  document.querySelectorAll("#sw button").forEach(function(b){b.addEventListener("click",function(){setColor(b.dataset.c)})});
  document.querySelectorAll("#presets button").forEach(function(b){b.addEventListener("click",function(){
    var p=PRESETS[b.dataset.p]; if(!p)return;
    var v=sw()?p.sw:p.en;
    ev.value=v[0]; msg.value=v[1]; setColor(p.color); updateLabel();
  })});
  document.querySelectorAll("#styleSeg button").forEach(function(b){b.addEventListener("click",function(){
    style=b.dataset.st;
    document.querySelectorAll("#styleSeg button").forEach(function(x){x.setAttribute("aria-pressed",x===b)});
    applyStyle();
  })});
  document.getElementById("sendDesign").addEventListener("click",function(e){
    e.preventDefault();
    var color=(sw()?COLOR_NAMES_SW:COLOR_NAMES)[col];
    var stylesEn={rect:"rectangular",round:"round",oval:"oval"},stylesSw={rect:"mstatili",round:"duara",oval:"mviringo"};
    var styleName=sw()?stylesSw[style]:stylesEn[style];
    var notSet=t("(not set)","(hujaweka)");
    go(t("Hello Aquamor, I would like a branded bottle:\nEvent or name: ","Habari Aquamor, ningependa chupa yenye chapa yangu:\nHafla au jina: ")+(ev.value||notSet)+
      t("\nMessage: ","\nUjumbe: ")+(msg.value||notSet)+
      t("\nLabel color: ","\nRangi ya lebo: ")+color+
      t("\nLabel style: ","\nMtindo wa lebo: ")+styleName+
      t("\nBottle size: ","\nUkubwa wa chupa: ")+document.getElementById("sz").value+
      t("\nPlease tell me the quantity options and price.","\nTafadhali niambie idadi zinazopatikana na bei."));
  });
  setColor(col);
  drawRows();
  applyStyle();
  applyRotation();
  if(document.fonts&&document.fonts.ready){document.fonts.ready.then(updateLabel)}
})();
