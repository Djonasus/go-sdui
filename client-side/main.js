$(document).ready(function() {
    connectWebSocket();
  });
  
  function connectWebSocket() {
    const socket = new WebSocket("ws://localhost:8080/ws");
  
    socket.onopen = function(event) {
      console.log("WebSocket connected");
    };
  
    socket.onmessage = function(event) {
      const data = JSON.parse(event.data);
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
            case "text":
                elem = $("<p>").text(component.content);
                break;
            case "image":
                elem = $("<img>").attr("src", component.url);
                break;
            case "button":
                elem = $("<button>").text(component.label);
                elem.click(function() {
                    if (component.action === "open_link") {
                        window.open(component.data, "_blank")
                    }
                });
                break;
            case "div":
                elem = $("<div>")
                buildUI(elem, component.children)
            default:
                break;
        }
        if (component.style != null) {
            for (var key in component.style) {
                elem.css(component.style[key]);
            }
        }
        container.append(elem)
    });
  }
  