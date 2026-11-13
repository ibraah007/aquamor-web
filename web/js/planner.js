/* Refill planner: estimates weekly refills and cost from the number of people.
   Prices are the confirmed refill prices; edit SIZES and PER_PERSON_DAY here. */
(function(){
  var SIZES={"5L":{litres:5,price:50},"10L":{litres:10,price:100},"20L":{litres:20,price:200}};
  var PER_PERSON_DAY=2; /* planning assumption, stated in the note under the result */
  var people=document.getElementById("people");
  if(!people)return;
  var size="20L";
  var litres=document.getElementById("planLitres"), refills=document.getElementById("planRefills"), cost=document.getElementById("planCost"), wa=document.getElementById("planWa");
  function sw(){return window.LANG==="sw"}
  function t(en,swText){return sw()?swText:en}
  function num(n){return n.toLocaleString("en-US")}
  function readPeople(){
    var n=Math.floor(people.valueAsNumber);
    if(!isFinite(n)||n<1)n=1;
    if(n>30)n=30;
    return n;
  }
  function calc(){
    var n=readPeople(), need=n*PER_PERSON_DAY*7, r=Math.ceil(need/SIZES[size].litres);
    return {people:n,need:need,refills:r,cost:r*SIZES[size].price};
  }
  function draw(){
    var c=calc();
    litres.textContent=t("about ","takriban ")+num(c.need)+t(" litres"," lita");
    refills.textContent=c.refills+" × "+size;
    cost.textContent="Ksh "+num(c.cost);
    wa.href="https://wa.me/254712140006?text="+encodeURIComponent(
      t("Hello Aquamor, I used the refill planner:\n","Habari Aquamor, nimetumia kipimo cha kujaza maji:\n")+
      t("People at home: ","Watu nyumbani: ")+c.people+"\n"+
      t("Container size: ","Ukubwa wa chombo: ")+size+"\n"+
      t("Suggested refills each week: ","Kujaza kunashauriwa kila wiki: ")+c.refills+" × "+size+"\n"+
      t("Estimated cost each week: ","Makadirio ya gharama kila wiki: ")+"Ksh "+num(c.cost)+"\n"+
      t("Please confirm.","Tafadhali nithibitishie.")
    );
  }
  function step(d){
    people.value=Math.min(30,Math.max(1,readPeople()+d));
    draw();
  }
  document.getElementById("peopleLess").addEventListener("click",function(){step(-1)});
  document.getElementById("peopleMore").addEventListener("click",function(){step(1)});
  people.addEventListener("input",draw);
  people.addEventListener("change",function(){people.value=readPeople();draw()});
  document.querySelectorAll("#planSize button").forEach(function(b){
    b.addEventListener("click",function(){
      size=b.dataset.s;
      document.querySelectorAll("#planSize button").forEach(function(x){x.setAttribute("aria-pressed",x===b)});
      draw();
    });
  });
  document.addEventListener("aquamor:lang",draw);
  draw();
})();
