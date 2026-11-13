(function(){
  "use strict";
  var TOKEN_KEY="aquamor-admin-token";
  function $(id){return document.getElementById(id)}

  var loginPanel=$("loginPanel"),ordersView=$("ordersView"),loginForm=$("loginForm"),loginBtn=$("loginBtn"),
      tokenInput=$("token"),loginStatus=$("loginStatus"),ordersEl=$("orders"),statsEl=$("stats"),
      adminStatus=$("adminStatus"),updatedEl=$("updated"),searchInput=$("search"),
      orderFilter=$("orderFilter"),payFilter=$("payFilter"),refreshBtn=$("refresh"),
      exportBtn=$("export"),logoutBtn=$("logout"),clearBtn=$("clearFilters"),shownEl=$("shownCount");

  var state={token:"",orders:[],products:{},loadedAt:null,busy:false};
  var dateFmt=new Intl.DateTimeFormat("en-GB",{day:"2-digit",month:"short",year:"numeric",timeZone:"Africa/Nairobi"});
  var timeFmt=new Intl.DateTimeFormat("en-GB",{hour:"2-digit",minute:"2-digit",hour12:false,timeZone:"Africa/Nairobi"});
  var optionNames={swap:"Refill / swap",new_purchase:"New bottle",not_applicable:"Standard"};
  var payKinds={paid:"paid",pending:"pending",unpaid:"unpaid",failed:"failed"};
  var orderKinds={completed:"paid",out_for_delivery:"pending",processing:"pending",pending:"unpaid"};

  function el(tag,className,text){
    var node=document.createElement(tag);
    if(className)node.className=className;
    if(text!=null)node.textContent=text;
    return node;
  }
  function money(value){
    var n=Number(value);
    if(!isFinite(n))n=0;
    return "Ksh "+n.toLocaleString("en-KE",{maximumFractionDigits:2});
  }
  function when(value){
    if(!value)return "—";
    var d=new Date(value);
    if(isNaN(d.getTime()))return "—";
    return dateFmt.format(d)+", "+timeFmt.format(d);
  }
  function label(value){
    if(!value)return "—";
    return value.replace(/_/g," ").replace(/\b[a-z]/g,function(c){return c.toUpperCase()});
  }
  function setStatus(node,text,isError){
    node.textContent=text||"";
    if(isError)node.classList.add("error");
    else node.classList.remove("error");
  }
  function productName(item){
    var product=state.products[item.product_id];
    return product?product.name:"Product #"+item.product_id;
  }
  function badge(text,kind){
    return el("span","badge badge-"+(kind||"unpaid"),text);
  }

  function fetchOrders(token){
    return fetch("/api/v1/orders/all",{
      method:"GET",
      headers:{Accept:"application/json",Authorization:"Bearer "+token},
      credentials:"same-origin",
      cache:"no-store"
    }).then(function(response){
      return response.text().then(function(text){
        var data=null;
        try{data=JSON.parse(text)}catch(e){}
        if(!response.ok){
          var message=(data&&data.error)||("The server returned HTTP "+response.status+".");
          var error=new Error(message);
          error.status=response.status;
          throw error;
        }
        if(!Array.isArray(data))throw new Error("The server did not return an order list.");
        return data;
      });
    });
  }

  function loadProducts(){
    return fetch("/api/v1/products",{headers:{Accept:"application/json"},credentials:"same-origin"})
      .then(function(response){
        if(!response.ok)throw new Error("products");
        return response.json();
      })
      .then(function(list){
        var map={};
        if(Array.isArray(list))list.forEach(function(p){map[p.id]=p});
        state.products=map;
      })
      .catch(function(){state.products={}});
  }

  function unlock(token){
    state.token=token;
    loginBtn.disabled=true;
    return loadProducts().then(function(){
      return fetchOrders(token);
    }).then(function(orders){
      state.orders=orders;
      state.loadedAt=new Date();
      try{sessionStorage.setItem(TOKEN_KEY,token)}catch(e){}
      setStatus(loginStatus,"",false);
      loginPanel.hidden=true;
      ordersView.hidden=false;
      tokenInput.value="";
      render();
      return true;
    }).catch(function(error){
      if(error&&error.status===401){
        forgetToken();
        setStatus(loginStatus,"That token was not accepted. Check ADMIN_TOKEN and try again.",true);
        tokenInput.focus();
      }else{
        setStatus(loginStatus,error&&error.message?error.message:"Could not load orders. Check your connection.",true);
      }
      return false;
    }).then(function(ok){
      loginBtn.disabled=false;
      return ok;
    });
  }

  function forgetToken(){
    try{sessionStorage.removeItem(TOKEN_KEY)}catch(e){}
    state.token="";
    state.orders=[];
    state.loadedAt=null;
  }

  function refresh(){
    if(state.busy||!state.token)return;
    state.busy=true;
    refreshBtn.disabled=true;
    setStatus(adminStatus,"Refreshing orders…",false);
    fetchOrders(state.token).then(function(orders){
      state.orders=orders;
      state.loadedAt=new Date();
      render();
      setStatus(adminStatus,"Orders updated.",false);
    }).catch(function(error){
      if(error&&error.status===401){
        forgetToken();
        ordersView.hidden=true;
        loginPanel.hidden=false;
        setStatus(loginStatus,"Your session expired. Enter the admin token again.",true);
        tokenInput.focus();
      }else{
        setStatus(adminStatus,error&&error.message?error.message:"Could not refresh orders. Check your connection.",true);
      }
    }).then(function(){
      state.busy=false;
      refreshBtn.disabled=false;
    });
  }

  function filteredOrders(){
    var query=searchInput.value.trim().toLowerCase();
    var orderStatus=orderFilter.value;
    var payStatus=payFilter.value;
    return state.orders.filter(function(order){
      if(orderStatus&&order.order_status!==orderStatus)return false;
      if(payStatus&&order.payment_status!==payStatus)return false;
      if(!query)return true;
      var haystack=[order.id,order.customer_name,order.customer_phone,order.customer_email,
        order.delivery_estate,order.delivery_address,order.payment_reference,order.mpesa_receipt_number]
        .join(" ").toLowerCase();
      return haystack.indexOf(query)!==-1;
    });
  }

  function stat(labelText,valueText){
    var box=el("div","stat");
    box.appendChild(el("b",null,valueText));
    box.appendChild(el("span",null,labelText));
    return box;
  }

  function renderStats(){
    var paid=0,paidTotal=0,awaiting=0,failed=0;
    state.orders.forEach(function(order){
      var status=(order.payment_status||"").toLowerCase();
      if(status==="paid"){
        paid++;
        paidTotal+=Number(order.total_amount)||0;
      }else if(status==="failed"){
        failed++;
      }else{
        awaiting++;
      }
    });
    statsEl.textContent="";
    statsEl.appendChild(stat("Orders loaded",String(state.orders.length)));
    statsEl.appendChild(stat("Paid",paid+" · "+money(paidTotal)));
    statsEl.appendChild(stat("Awaiting payment",String(awaiting)));
    statsEl.appendChild(stat("Failed payments",String(failed)));
  }

  function appendLine(parent,node){
    parent.appendChild(node);
    parent.appendChild(document.createElement("br"));
  }

  function contactBlock(order){
    var box=el("div");
    box.appendChild(el("h3",null,"Customer"));
    var p=el("p");
    if(order.customer_name)appendLine(p,el("strong",null,order.customer_name));
    if(order.customer_phone){
      var phone=el("a",null,order.customer_phone);
      phone.href="tel:"+order.customer_phone.replace(/[^\d+]/g,"");
      appendLine(p,phone);
    }
    if(order.customer_email){
      var email=el("a",null,order.customer_email);
      email.href="mailto:"+order.customer_email;
      appendLine(p,email);
    }
    if(!p.childNodes.length)p.textContent="—";
    box.appendChild(p);
    return box;
  }

  function deliveryBlock(order){
    var box=el("div");
    box.appendChild(el("h3",null,"Delivery"));
    var p=el("p");
    if(order.delivery_estate)appendLine(p,document.createTextNode(order.delivery_estate));
    if(order.delivery_address)appendLine(p,document.createTextNode(order.delivery_address));
    if(!p.childNodes.length)p.textContent="—";
    box.appendChild(p);
    return box;
  }

  function paymentBlock(order){
    var box=el("div");
    box.appendChild(el("h3",null,"Payment"));
    var p=el("p");
    appendLine(p,document.createTextNode(label(order.payment_status)));
    if(order.payment_provider)appendLine(p,document.createTextNode("via "+label(order.payment_provider)));
    if(order.payment_reference)appendLine(p,document.createTextNode("Ref "+order.payment_reference));
    if(order.mpesa_receipt_number)appendLine(p,document.createTextNode("M-Pesa "+order.mpesa_receipt_number));
    if(order.paid_at)appendLine(p,document.createTextNode("Paid "+when(order.paid_at)));
    box.appendChild(p);
    return box;
  }

  function itemsTable(items){
    if(!items||!items.length)return el("p","note","No line items were saved for this order.");
    var wrap=el("div","items-scroll");
    var table=el("table");
    var head=el("thead");
    var headRow=el("tr");
    [["Item",false],["Option",false],["Qty",true],["Unit",true],["Subtotal",true]].forEach(function(col){
      var th=el("th",col[1]?"num":null,col[0]);
      th.scope="col";
      headRow.appendChild(th);
    });
    head.appendChild(headRow);
    table.appendChild(head);
    var body=el("tbody");
    items.forEach(function(item){
      var row=el("tr");
      var cells=[
        [productName(item),false],
        [optionNames[item.bottle_option]||label(item.bottle_option),false],
        [String(item.quantity),true],
        [money(item.unit_price),true],
        [money(item.subtotal),true]
      ];
      cells.forEach(function(cell){
        row.appendChild(el("td",cell[1]?"num":null,cell[0]));
      });
      body.appendChild(row);
    });
    table.appendChild(body);
    wrap.appendChild(table);
    return wrap;
  }

  function orderCard(order){
    var card=el("article","order");
    var head=el("header","order-head");

    var left=el("div");
    left.appendChild(el("h2","order-id","Order #"+order.id));
    var meta=el("p","note",when(order.created_at)+" (Nairobi)");
    left.appendChild(meta);
    head.appendChild(left);

    var right=el("div","order-head-right");
    right.appendChild(badge(label(order.payment_status),payKinds[order.payment_status]||"unpaid"));
    right.appendChild(badge(label(order.order_status),orderKinds[order.order_status]||"unpaid"));
    right.appendChild(el("span","order-total",money(order.total_amount)));
    head.appendChild(right);

    card.appendChild(head);

    var grid=el("div","order-grid");
    grid.appendChild(contactBlock(order));
    grid.appendChild(deliveryBlock(order));
    grid.appendChild(paymentBlock(order));
    card.appendChild(grid);

    card.appendChild(itemsTable(order.items));
    return card;
  }

  function render(){
    var orders=filteredOrders();
    renderStats();
    ordersEl.textContent="";

    if(!state.orders.length){
      ordersEl.appendChild(el("p","empty","No orders have been saved yet."));
    }else if(!orders.length){
      ordersEl.appendChild(el("p","empty","No orders match these filters."));
    }else{
      var fragment=document.createDocumentFragment();
      orders.forEach(function(order){fragment.appendChild(orderCard(order))});
      ordersEl.appendChild(fragment);
    }

    shownEl.textContent="Showing "+orders.length+" of "+state.orders.length+" orders.";
    updatedEl.textContent=state.loadedAt?"Last updated "+when(state.loadedAt)+" (Nairobi).":"Not loaded yet.";
  }

  function csvCell(value){
    var text=value==null?"":String(value);
    if(/^[\s]*[=+\-@]/.test(text))text="'"+text;
    if(/[",\r\n]/.test(text))text='"'+text.replace(/"/g,'""')+'"';
    return text;
  }

  function exportCsv(){
    var orders=filteredOrders();
    if(!orders.length){
      setStatus(adminStatus,"There are no orders to export with the current filters.",true);
      return;
    }
    var header=["Order","Created (Nairobi)","Customer","Phone","Email","Area","Address",
      "Order status","Payment status","Provider","Reference","M-Pesa receipt","Paid at",
      "Total","Items"];
    var lines=[header.map(csvCell).join(",")];
    orders.forEach(function(order){
      var items=(order.items||[]).map(function(item){
        return item.quantity+" x "+productName(item)+" @ "+item.unit_price+" = "+item.subtotal;
      }).join("; ");
      lines.push([order.id,when(order.created_at),order.customer_name,order.customer_phone,
        order.customer_email,order.delivery_estate,order.delivery_address,order.order_status,
        order.payment_status,order.payment_provider,order.payment_reference,
        order.mpesa_receipt_number,order.paid_at?when(order.paid_at):"",order.total_amount,items]
        .map(csvCell).join(","));
    });
    var blob=new Blob(["\ufeff"+lines.join("\r\n")],{type:"text/csv;charset=utf-8"});
    var url=URL.createObjectURL(blob);
    var link=document.createElement("a");
    link.href=url;
    link.download="aquamor-orders-"+new Date().toISOString().slice(0,10)+".csv";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    setTimeout(function(){URL.revokeObjectURL(url)},1000);
    setStatus(adminStatus,"Exported "+orders.length+" order(s) to CSV.",false);
  }

  function signOut(){
    forgetToken();
    ordersView.hidden=true;
    loginPanel.hidden=false;
    loginForm.reset();
    ordersEl.textContent="";
    statsEl.textContent="";
    shownEl.textContent="";
    setStatus(adminStatus,"",false);
    setStatus(loginStatus,"Signed out of this tab.",false);
    tokenInput.focus();
  }

  loginForm.addEventListener("submit",function(event){
    event.preventDefault();
    var token=tokenInput.value.trim();
    if(!token){
      setStatus(loginStatus,"Enter the admin token to continue.",true);
      tokenInput.focus();
      return;
    }
    setStatus(loginStatus,"Checking token…",false);
    unlock(token);
  });

  refreshBtn.addEventListener("click",refresh);
  exportBtn.addEventListener("click",exportCsv);
  logoutBtn.addEventListener("click",signOut);
  clearBtn.addEventListener("click",function(){
    searchInput.value="";
    orderFilter.value="";
    payFilter.value="";
    render();
    searchInput.focus();
  });
  searchInput.addEventListener("input",render);
  orderFilter.addEventListener("change",render);
  payFilter.addEventListener("change",render);

  var saved="";
  try{saved=sessionStorage.getItem(TOKEN_KEY)||""}catch(e){}
  if(saved){
    setStatus(loginStatus,"Checking saved token…",false);
    unlock(saved);
  }else{
    loginPanel.hidden=false;
  }
})();
