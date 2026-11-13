/* Confirms a hosted-checkout payment after the provider redirects back. */
(function(){
  var isSw=false;
  try{isSw=localStorage.getItem("aquamor-lang")==="sw"}catch(e){}
  document.documentElement.lang=isSw?"sw":"en";
  function t(en,sw){return isSw?sw:en}

  var card=document.getElementById("payCard"),
      title=document.getElementById("payTitle"),
      msg=document.getElementById("payMsg"),
      spinner=document.getElementById("paySpinner"),
      actions=document.getElementById("payActions"),
      wa=document.getElementById("payWhatsApp"),
      back=document.getElementById("backOrder"),
      again=document.getElementById("payOrderAgain"),
      helpNote=document.getElementById("payHelpNote");

  if(wa)wa.href="https://wa.me/254712140006?text="+encodeURIComponent(t(
    "Hello Aquamor, I need help with an online payment.",
    "Habari Aquamor, ninahitaji msaada kuhusu malipo ya mtandaoni."
  ));
  if(back)back.textContent=t("Back to order","Rudi kwenye oda");
  if(again)again.textContent=t("Order again","Agiza tena");
  if(helpNote)helpNote.textContent=t(
    "If you were charged and this page does not confirm it, message us on WhatsApp with your receipt.",
    "Ikiwa ulitozwa na ukurasa huu hauthibitishi malipo, tutumie ujumbe WhatsApp na risiti yako."
  );

  function show(state,heading,text){
    if(card)card.dataset.state=state;
    title.textContent=heading;
    msg.textContent=text;
    if(spinner)spinner.hidden=true;
    if(actions)actions.hidden=false;
  }

  var q=new URLSearchParams(location.search);
  var ref=q.get("tx_ref")||q.get("reference")||"";
  var txid=q.get("transaction_id")||q.get("transactionId")||"";

  if(!ref&&!txid){
    show("info",t("No payment to check","Hakuna malipo ya kukagua"),t(
      "This page confirms an Aquamor online payment after you return from the provider. Go back to the order page to start a payment.",
      "Ukurasa huu huthibitisha malipo ya mtandaoni ya Aquamor baada ya kurejea kutoka kwa mtoa huduma. Rudi kwenye ukurasa wa oda kuanza malipo."
    ));
    return;
  }
  if(!ref||!txid){
    show("info",t("Payment details are incomplete","Maelezo ya malipo hayakamiliki"),t(
      "We could not find the transaction reference. If you were charged, contact us on WhatsApp with your receipt.",
      "Hatukuweza kupata kumbukumbu ya malipo. Ikiwa ulitozwa, wasiliana nasi kupitia WhatsApp na risiti yako."
    ));
    return;
  }

  fetch("/api/v1/payments/verify?reference="+encodeURIComponent(ref)+"&transaction_id="+encodeURIComponent(txid),{headers:{Accept:"application/json"},credentials:"same-origin"})
    .then(function(r){return r.json().catch(function(){return {}}).then(function(data){return {ok:r.ok,data:data}})})
    .then(function(res){
      if(!res.ok)throw new Error(res.data.error||t(
        "We could not confirm the payment. Please try again.",
        "Hatukuweza kuthibitisha malipo. Tafadhali jaribu tena."
      ));
      if(res.data.status==="paid"){
        var order=res.data.order_id?" "+t("Order","Oda")+" #"+res.data.order_id+".":"";
        show("paid",t("Payment received","Malipo yamepokelewa"),t(
          "Thank you! Your order is confirmed and Aquamor will contact you to arrange delivery.",
          "Asante! Oda yako imethibitishwa na Aquamor atawasiliana nawe kupanga uletaji."
        )+order);
      }else if(res.data.status==="failed"){
        show("failed",t("Payment not completed","Malipo hayakukamilika"),t(
          "The payment was not completed. You can try again or send your order on WhatsApp.",
          "Malipo hayakukamilika. Unaweza kujaribu tena au kutuma oda yako kupitia WhatsApp."
        ));
      }else{
        show("pending",t("Payment is still being confirmed","Malipo bado yanathibitishwa"),t(
          "We have not received the final confirmation yet. Please check again in a moment, or contact us on WhatsApp with your receipt.",
          "Bado hatujapokea uthibitisho wa mwisho. Tafadhali angalia tena baada ya muda mfupi, au wasiliana nasi kupitia WhatsApp na risiti yako."
        ));
      }
    })
    .catch(function(err){
      show("error",t("We could not confirm the payment","Hatukuweza kuthibitisha malipo"),err.message||t(
        "Please try again or contact us on WhatsApp with your receipt.",
        "Tafadhali jaribu tena au wasiliana nasi kupitia WhatsApp na risiti yako."
      ));
    });
})();
