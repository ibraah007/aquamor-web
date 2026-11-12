/* English/Swahili switch, live open/closed status, Paybill copy buttons, testimonials.
   Edit SW for translations and TESTIMONIALS for customer quotes. */
(function(){
  var SW = {
  "Order": "Agiza",
  "Branded bottles": "Chupa zenye chapa",
  "Dispenser cleaning": "Kusafisha dispenser",
  "Gallery": "Picha",
  "About": "Kuhusu sisi",
  "Contact": "Wasiliana",
  "Order now": "Agiza sasa",
  "Pure water, delivered the way you need it.": "Maji safi, yanayoletwa jinsi unavyoyahitaji.",
  "Purified drinking water for homes, businesses, churches and events, from Musindi Road, Ngara, Nairobi.": "Maji safi ya kunywa kwa nyumba, biashara, makanisa na hafla, kutoka Barabara ya Musindi, Ngara, Nairobi.",
  "Build your order": "Panga oda yako",
  "Design an event bottle": "Buni chupa ya hafla",
  "Serving Nairobi since 2019": "Tunahudumia Nairobi tangu 2019",
  "Refills from Ksh 50": "Kujaza maji kuanzia Ksh 50",
  "Branded bottles for events": "Chupa zenye chapa kwa hafla",
  "Free delivery": "Huduma ya kuleta bila malipo",
  "Choose your sizes and how many. We will confirm the price and delivery on WhatsApp.": "Chagua ukubwa na idadi. Tutathibitisha bei na uletaji kupitia WhatsApp.",
  "Water refills": "Kujaza maji",
  "Bottled water": "Maji ya chupa",
  "New 20L bottles": "Chupa mpya za 20L",
  "5L refill": "Kujaza 5L",
  "10L refill": "Kujaza 10L",
  "20L refill": "Kujaza 20L",
  "Ksh 50 · bring your container": "Ksh 50 · leta chombo chako",
  "Ksh 100 · bring your container": "Ksh 100 · leta chombo chako",
  "Ksh 200 · bring your container": "Ksh 200 · leta chombo chako",
  "300ml bottle": "Chupa ya 300ml",
  "500ml bottle": "Chupa ya 500ml",
  "1L bottle": "Chupa ya 1L",
  "1.5L bottle": "Chupa ya 1.5L",
  "Bottled drinking water": "Maji ya kunywa ya chupa",
  "20L hard bottle + water": "Chupa ngumu ya 20L + maji",
  "Ksh 1,400 · bottle and water": "Ksh 1,400 · chupa na maji",
  "20L regular bottle": "Chupa ya kawaida ya 20L",
  "Your order": "Oda yako",
  "Nothing added yet.": "Bado hujaongeza chochote.",
  "Add at least one item first.": "Ongeza angalau bidhaa moja kwanza.",
  "Estimated total": "Jumla ya makadirio",
  "Delivery": "Kuletewa",
  "Pickup": "Kuchukua",
  "Delivery area": "Eneo la kuletewa",
  "Name for pickup": "Jina la kuchukua",
  "Send order on WhatsApp": "Tuma oda kupitia WhatsApp",
  "Refills: 5L Ksh 50, 10L Ksh 100, 20L Ksh 200. New 20L bottles: hard bottle with water Ksh 1,400, regular bottle Ksh 450. We offer free delivery: tell us your area and we will confirm. Small bottled water prices are confirmed on WhatsApp.": "Kujaza maji: 5L Ksh 50, 10L Ksh 100, 20L Ksh 200. Chupa mpya za 20L: chupa ngumu na maji Ksh 1,400, chupa ya kawaida Ksh 450. Tunaleta bila malipo: tuambie eneo lako nasi tutathibitisha. Bei za maji ya chupa ndogo zinathibitishwa kupitia WhatsApp.",
  "Your name on every bottle": "Jina lako kwenye kila chupa",
  "Design a label for your wedding, church service, graduation or company event and see it on the bottle straight away.": "Buni lebo kwa ajili ya harusi, ibada ya kanisa, mahafali au hafla ya kampuni yako na uione kwenye chupa papo hapo.",
  "Event or name": "Hafla au jina",
  "Date or short message": "Tarehe au ujumbe mfupi",
  "Label color": "Rangi ya lebo",
  "Bottle size": "Ukubwa wa chupa",
  "Request this design on WhatsApp": "Omba muundo huu kupitia WhatsApp",
  "Aquamor will confirm the final design, quantity and price with you.": "Aquamor itathibitisha muundo wa mwisho, idadi na bei nawe.",
  "Dispenser cleaning and sanitizing": "Kusafisha na kuua vijidudu kwenye dispenser",
  "A clean dispenser keeps your drinking water fresh. Our trained team cleans and sanitizes water dispensers for homes and workplaces.": "Dispenser safi huweka maji yako ya kunywa yakiwa safi. Timu yetu iliyofunzwa husafisha na kuua vijidudu kwenye dispenser za maji kwa nyumba na sehemu za kazi.",
  "Deep cleaning": "Usafi wa kina",
  "Removes dirt and stains from your dispenser.": "Huondoa uchafu na madoa kwenye dispenser yako.",
  "Complete sanitizing": "Kuua vijidudu kikamilifu",
  "The tank, taps and every part that touches the water.": "Tanki, mabomba na kila sehemu inayogusa maji.",
  "Safe and hygienic": "Salama na safi",
  "Helps keep your drinking water fresh.": "Husaidia maji yako ya kunywa kubaki safi.",
  "Trained team": "Timu iliyofunzwa",
  "Done by the experienced Aquamor team.": "Hufanywa na timu yenye uzoefu ya Aquamor.",
  "Before": "Kabla",
  "Cleaning": "Kusafisha",
  "After": "Baada",
  "Ask us for a quote on WhatsApp.": "Tuulize bei kupitia WhatsApp.",
  "Book a cleaning on WhatsApp": "Weka miadi ya kusafisha kupitia WhatsApp",
  "Call 0113 952 100": "Piga 0113 952 100",
  "Our work": "Kazi zetu",
  "Real photos from our shop, our events and our custom labels. Tap a photo to enlarge it.": "Picha halisi kutoka dukani kwetu, hafla zetu na lebo zetu maalum. Gusa picha ili kuikuza.",
  "All": "Zote",
  "Bottles": "Chupa",
  "Events": "Hafla",
  "Refills": "Kujaza",
  "Team": "Timu",
  "A row of our glass bottles in the shop window": "Safu ya chupa zetu za glasi dirishani dukani",
  "Our refill station on Musindi Road": "Kituo chetu cha kujaza maji Barabara ya Musindi",
  "Custom round labels for a church event, 2026": "Lebo za duara maalum kwa hafla ya kanisa, 2026",
  "A custom-branded glass bottle": "Chupa ya glasi yenye chapa maalum",
  "Custom wedding label designs in two colors": "Miundo ya lebo za harusi kwa rangi mbili",
  "Glass bottles with custom labels": "Chupa za glasi zenye lebo maalum",
  "Bottled water stacked and ready": "Maji ya chupa yaliyopangwa tayari",
  "Our team at work in the shop": "Timu yetu ikifanya kazi dukani",
  "Pumps and dispensers at our shop": "Pampu na dispenser dukani kwetu",
  "Our team, with a pack of custom-labelled bottles": "Timu yetu, na kifurushi cha chupa zenye lebo maalum",
  "A team member wiping down a 20L container": "Mfanyakazi akifuta chombo cha 20L",
  "Inside the shop, packs of bottles ready for orders": "Ndani ya duka, vifurushi vya chupa tayari kwa oda",
  "Portable automatic pumps for 20L containers": "Pampu otomatiki zinazobebeka kwa vyombo vya 20L",
  "Custom labels celebrating 50 years of God’s faithfulness": "Lebo maalum zinazosherehekea miaka 50 ya uaminifu wa Mungu",
  "Bottles ready for a church event": "Chupa zilizo tayari kwa hafla ya kanisa",
  "A hall set for a wedding reception": "Ukumbi ulioandaliwa kwa sherehe ya harusi",
  "Bottled water on the shelves in our shop": "Maji ya chupa rafu dukani kwetu",
  "20 litre containers lined up at the counter": "Vyombo vya lita 20 vilivyopangwa kaunta",
  "A customer carrying a 20 litre refill": "Mteja akibeba chombo cha lita 20",
  "A team member refilling a 20 litre container": "Mfanyakazi akijaza chombo cha lita 20",
  "A trusted drinking water partner in Ngara": "Mshirika wa kuaminika wa maji ya kunywa Ngara",
  "Aquamor Purified Drinking Water was established in 2019 and is located along Musindi Road, Ngara, Nairobi. We are committed to quality purified drinking water and reliable customer service.": "Aquamor Purified Drinking Water ilianzishwa mwaka 2019 na iko Barabara ya Musindi, Ngara, Nairobi. Tumejitolea kutoa maji safi ya kunywa yenye ubora na huduma ya kuaminika kwa wateja.",
  "We serve households, businesses, churches and organizations with convenient and dependable water solutions. Whether you need water for your home, your business or a special event, Aquamor is here to serve you.": "Tunahudumia kaya, biashara, makanisa na mashirika kwa suluhisho za maji zilizo rahisi na za kutegemewa. Iwe unahitaji maji kwa nyumba yako, biashara yako au hafla maalum, Aquamor iko tayari kukuhudumia.",
  "What customers say": "Wateja wetu wanasema nini",
  "Questions people ask": "Maswali ambayo watu huuliza",
  "Talk to us": "Ongea nasi",
  "Visit us": "Tutembelee",
  "Along Musindi Road, opposite Ngara Girls Secondary School, next to Ngara Heights building, Nairobi": "Barabara ya Musindi, mkabala na Shule ya Upili ya Wasichana ya Ngara, karibu na jengo la Ngara Heights, Nairobi",
  "Opening hours": "Saa za kufungua",
  "Monday to Saturday: 8:30am to 8:30pm": "Jumatatu hadi Jumamosi: 8:30am hadi 8:30pm",
  "Sunday: 12:00pm to 7:00pm": "Jumapili: 12:00pm hadi 7:00pm",
  "Phone calls": "Simu",
  "Email": "Barua pepe",
  "Pay by Paybill (Equity)": "Lipa kwa Paybill (Equity)",
  "Business no.": "Namba ya biashara",
  "Account no.": "Namba ya akaunti",
  "Copy": "Nakili",
  "Copied": "Imenakiliwa",
  "Chat on WhatsApp": "Ongea nasi kwenye WhatsApp",
  "Call us": "Tupigie",
  "Get directions": "Pata maelekezo",
  "WhatsApp us": "Tuandikie WhatsApp",
  "© 2026 Aquamor Purified Drinking Water. Musindi Road, Ngara, Nairobi.": "© 2026 Aquamor Purified Drinking Water. Barabara ya Musindi, Ngara, Nairobi.",
  "What sizes do you sell?": "Mnauza ukubwa gani?",
  "We refill 5L, 10L and 20L containers, and we sell bottled drinking water in 300ml, 500ml, 1L and 1.5L bottles.": "Tunajaza vyombo vya 5L, 10L na 20L, na tunauza maji ya kunywa ya chupa za 300ml, 500ml, 1L na 1.5L.",
  "How much does a refill cost?": "Kujaza maji kunagharimu kiasi gani?",
  "A 5L refill is Ksh 50, a 10L refill is Ksh 100 and a 20L refill is Ksh 200.": "Kujaza 5L ni Ksh 50, 10L ni Ksh 100 na 20L ni Ksh 200.",
  "Do you deliver?": "Mnaleta?",
  "Yes, we offer free delivery. Tell us your area when you order and we will confirm.": "Ndiyo, tunaleta bila malipo. Tuambie eneo lako unapoagiza nasi tutathibitisha.",
  "Can I order bottles with my own label?": "Naweza kuagiza chupa zenye lebo yangu mwenyewe?",
  "Yes. We make custom-branded bottled water for churches, weddings, graduations, corporate functions and other events. Try the label designer on this page, or message us on WhatsApp.": "Ndiyo. Tunatengeneza maji ya chupa yenye chapa maalum kwa makanisa, harusi, mahafali, hafla za kampuni na matukio mengine. Jaribu kibuni lebo kwenye ukurasa huu, au tutumie ujumbe kwenye WhatsApp.",
  "When are you open?": "Mnafunguliwa saa ngapi?",
  "Monday to Saturday from 8:30am to 8:30pm, and Sunday from 12:00pm to 7:00pm.": "Jumatatu hadi Jumamosi kuanzia 8:30am hadi 8:30pm, na Jumapili kuanzia 12:00pm hadi 7:00pm.",
  "How do I pay?": "Nalipaje?",
  "You can pay with Equity Paybill 247247, account number 395210. Confirm your total with us on WhatsApp first.": "Unaweza kulipa kwa Equity Paybill 247247, namba ya akaunti 395210. Thibitisha jumla yako nasi kupitia WhatsApp kwanza.",
  "Do you clean water dispensers?": "Mnasafisha dispenser za maji?",
  "Yes. We clean and sanitize dispensers, including the tank and taps. Ask us for a quote on WhatsApp.": "Ndiyo. Tunasafisha na kuua vijidudu kwenye dispenser, ikiwa ni pamoja na tanki na mabomba. Tuulize bei kupitia WhatsApp.",
  "Where are you located?": "Mko wapi?",
  "Along Musindi Road, opposite Ngara Girls Secondary School, next to Ngara Heights building, Nairobi.": "Barabara ya Musindi, mkabala na Shule ya Upili ya Wasichana ya Ngara, karibu na jengo la Ngara Heights, Nairobi.",
  "Estimated total (listed items)": "Jumla ya makadirio (bidhaa zilizo na bei)",
  "Bottled water prices are confirmed on WhatsApp.": "Bei za maji ya chupa zinathibitishwa kupitia WhatsApp.",
  "Bottled water prices will be confirmed on WhatsApp.": "Bei za maji ya chupa zitathibitishwa kupitia WhatsApp.",
  "Skip to content": "Nenda kwenye maudhui",
  "Aquamor purified water": "Maji safi ya Aquamor",
  "Menu": "Menyu",
  "Product type": "Aina ya bidhaa",
  "Delivery or pickup": "Kuletewa au kuchukua",
  "Preview of your custom bottle": "Onyesho la chupa yako maalum",
  "Deep blue": "Bluu iliyokolea",
  "Burgundy": "Rangi ya divai",
  "Green": "Kijani",
  "Purple": "Zambarau",
  "Black": "Nyeusi",
  "A dispenser before, during and after cleaning": "Dispenser kabla, wakati na baada ya kusafishwa",
  "A dirty dispenser basin": "Bakuli la dispenser lililochafuliwa",
  "A team member cleaning dispenser parts at a sink": "Mfanyakazi akisafisha sehemu za dispenser kwenye sinki",
  "A clean white dispenser basin": "Bakuli jeupe la dispenser lililosafishwa",
  "Filter photos": "Chuja picha",
  "A team member wiping down a 20 litre container": "Mfanyakazi akifuta chombo cha lita 20",
  "Inside the shop, with packs of bottles ready for orders": "Ndani ya duka, pakiti za chupa ziko tayari kwa oda",
  "A team member showing a portable automatic water pump for 20 litre containers": "Mfanyakazi akionyesha pampu ndogo ya maji ya kiotomatiki kwa chombo cha lita 20",
  "Close": "Funga",
  "An Aquamor team member smiling and holding a pack of custom-labelled bottles in front of the refill price poster": "Mfanyakazi wa Aquamor akitabasamu akiwa ameshika pakiti ya chupa zenye lebo maalum mbele ya bango la bei za kujaza maji"
};
  var PH = {
  "e.g. Ngara, Pangani, Eastleigh": "mf. Ngara, Pangani, Eastleigh",
  "e.g. John & Mary": "mf. John & Mary",
  "e.g. 12 December 2026": "mf. 12 Desemba 2026",
  "e.g. Your name": "mf. Jina lako"
};
  window.__sw=SW; window.LANG="en";
  var TESTIMONIALS=[]; // add real quotes: {quote:"...", name:"...", place:"..."}
  var orig=new WeakMap();
  function textNodes(root){var w=document.createTreeWalker(root,NodeFilter.SHOW_TEXT),n,a=[];while(n=w.nextNode()){var t=n.parentNode&&n.parentNode.nodeName;if(t!=="SCRIPT"&&t!=="STYLE")a.push(n)}return a}
  function apply(root){textNodes(root).forEach(function(n){
    var o=orig.has(n)?orig.get(n):n.nodeValue,k=o.replace(/\s+/g," ").trim(); if(!k)return;
    if(window.LANG==="sw"&&SW[k]){if(!orig.has(n))orig.set(n,o);n.nodeValue=o.match(/^\s*/)[0]+SW[k]+o.match(/\s*$/)[0]}
    else if(orig.has(n)&&window.LANG!=="sw"){n.nodeValue=orig.get(n)}
  })}
  var H={Mon:[510,1230],Tue:[510,1230],Wed:[510,1230],Thu:[510,1230],Fri:[510,1230],Sat:[510,1230],Sun:[720,1140]};
  var order=["Mon","Tue","Wed","Thu","Fri","Sat","Sun"];
  function fmt(m){var h=Math.floor(m/60),mm=m%60,ap=h>=12?"pm":"am";h=h%12||12;return h+(mm?":"+(mm<10?"0":"")+mm:":00")+ap}
  function status(){
    var isSw=window.LANG==="sw",o={};
    new Intl.DateTimeFormat("en-GB",{timeZone:"Africa/Nairobi",weekday:"short",hour:"numeric",minute:"numeric",hourCycle:"h23"}).formatToParts(new Date()).forEach(function(x){o[x.type]=x.value});
    var d=o.weekday,m=(+o.hour)*60+(+o.minute),h=H[d]; if(!h)return "";
    if(m>=h[0]&&m<h[1])return (isSw?"Wazi sasa hadi ":"Open now until ")+fmt(h[1]);
    if(m<h[0])return (isSw?"Imefungwa sasa, inafunguliwa leo saa ":"Closed now, opens today at ")+fmt(h[0]);
    return (isSw?"Imefungwa sasa, inafunguliwa kesho saa ":"Closed now, opens tomorrow at ")+fmt(H[order[(order.indexOf(d)+1)%7]][0]);
  }
  function refreshStatus(){
    var text=status();
    document.querySelectorAll(".js-open").forEach(function(e){e.textContent=text});
  }
  function wa(msg){return "https://wa.me/254712140006?text="+encodeURIComponent(msg)}
  function setLang(l){
    window.LANG=l; document.documentElement.lang=l; apply(document.body);
    document.querySelectorAll("[placeholder]").forEach(function(e){if(!e.dataset.ph0)e.dataset.ph0=e.placeholder;e.placeholder=(l==="sw"&&PH[e.dataset.ph0])?PH[e.dataset.ph0]:e.dataset.ph0});
    // Accessible names use the same dictionary, so screen readers get Swahili too.
    document.querySelectorAll("[aria-label]").forEach(function(e){
      if(!e.dataset.al0)e.dataset.al0=e.getAttribute("aria-label");
      var v=e.dataset.al0;
      e.setAttribute("aria-label",(l==="sw"&&SW[v])?SW[v]:v);
    });
    refreshStatus();
    var isSw=l==="sw";
    var book=document.getElementById("bookClean");
    if(book)book.href=wa(isSw?"Habari Aquamor, ningependa kuweka miadi ya kusafisha dispenser.":"Hello Aquamor, I would like to book a dispenser cleaning and sanitizing service.");
    var chat=document.getElementById("chatWa");
    if(chat)chat.href=wa(isSw?"Habari Aquamor, ningependa kuuliza kuhusu maji yenu.":"Hello Aquamor, I would like to ask about your water.");
    var langBtn=document.getElementById("langBtn");
    if(langBtn)langBtn.textContent=isSw?"EN":"SW";
    try{localStorage.setItem("aquamor-lang",l)}catch(e){}
    document.dispatchEvent(new CustomEvent("aquamor:lang",{detail:{lang:l}}));
  }
  window.__setLang=setLang;
  new MutationObserver(function(ms){ms.forEach(function(m){m.addedNodes.forEach(function(a){
    if(a.nodeType===1)apply(a);else if(a.nodeType===3&&a.parentNode)apply(a.parentNode)})})}).observe(document.body,{childList:true,subtree:true});
  var langBtn=document.getElementById("langBtn");
  if(langBtn)langBtn.addEventListener("click",function(){setLang(window.LANG==="sw"?"en":"sw")});
  document.querySelectorAll(".cp").forEach(function(b){b.addEventListener("click",function(){
    var v=b.dataset.v;function ok(){b.textContent="Copied";setTimeout(function(){b.textContent="Copy"},1600)}
    try{navigator.clipboard.writeText(v).then(ok,function(){b.textContent=v})}catch(e){b.textContent=v}
  })});
  if(TESTIMONIALS.length){
    var box=document.getElementById("revList");
    TESTIMONIALS.forEach(function(r){var f=document.createElement("figure");f.className="rev";
      var p=document.createElement("p");p.textContent="\u201c"+r.quote+"\u201d";var c=document.createElement("footer");c.textContent=r.name+(r.place?", "+r.place:"");
      f.appendChild(p);f.appendChild(c);box.appendChild(f)});
    document.getElementById("reviews").hidden=false;
  }
  var saved=null;try{saved=localStorage.getItem("aquamor-lang")}catch(e){}
  var start=saved||((navigator.language||"").toLowerCase().indexOf("sw")===0?"sw":"en");
  setLang(start);
  setInterval(refreshStatus,60000);
})();
