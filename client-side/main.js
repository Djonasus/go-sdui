$(document).ready(function() {
    connectWebSocket("gettodos");
  });
  
  function connectWebSocket(link) {
    const socket = new WebSocket("ws://localhost:8080/"+link);
  
    socket.onopen = function(event) {
      console.log("WebSocket connected");
    };
  
    socket.onmessage = function(event) {
      console.log(event.data)
      const data = JSON.parse(event.data);
      //console.log(data.components)
      buildUI($("#ui-container"), data.components);
    };
  
    socket.onclose = function(event) {
      console.log("WebSocket closed");
    };
  }
  
  function buildUI(container, components) {
    //const container = $("#ui-container");
    container.empty();
  
    components.forEach(function(component) {
        var elem;

        switch (component.type) {
          case "title":
            elem = $("<h1>").text(component.content);
            break;
          case "element":
            elem = $("<div>");
            elem.append("<b>"+component.title+"</b>");
            elem.append("<p>"+component.description+"</p>");
            if (component.checked == "0") {
              var che = "";
            } else {
              var che = "checked";
            }
            elem.append("<input type='checkbox' disabled "+che+">");
            elem.append("<button onclick='query(`check`,`"+component.id+"`)' >Выполнить</button><br>");
            elem.append("<button onClick='query(`remove`,`"+component.id+"`)'>Удалить</button>");
            elem.append("<br><br>");
            break;
          case "footer":
            elem = $("<h3>").text("SDUI by Djonasus 2023");
            break;
          case "button":
            elem = $("<button>").text(component.content)
            elem.click(function () {
              connectWebSocket(component.link);
            })
            break;
          case "text":
            elem = $("<p>").text(component.content)
            break;
          case "input":
            elem = $("<input>").attr('name',component.name)
            break;
          case "submit":
            elem = $("<button>").text(component.content)
            elem.click(function () {
              var names = "";
              for (var key in component.inputs) {
                names+=$("input[name="+component.inputs[key]+"]").val()+"/";
              }
              query("create",names)
            })
        }

        if (component.style != null) {
            for (var key in component.style) {
                elem.css(component.style[key]);
            }
        }
        container.append(elem)
    });
  }
  
  function query(action, arg) {
    const socket = new WebSocket("ws://localhost:8080/"+action+"/"+arg);

    socket.onopen = function(event) {
      console.log("ActionSocket connected");
    };

    socket.onmessage = function(event) {
      if (event.data === "ok"){
        connectWebSocket("gettodos");
      }
    };

    socket.onclose = function(event) {
      console.log("ActionSocket closed");
    };
  }