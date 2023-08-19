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
      buildUI(data.components);
    };
  
    socket.onclose = function(event) {
      console.log("WebSocket closed");
    };
  }
  
  function buildUI(components) {
    const container = $("#ui-container");
    container.empty();
  
    components.forEach(function(component) {
      if (component.type === "text") {
        const textElement = $("<p>").text(component.content);
        container.append(textElement);
      } else if (component.type === "image") {
        const imageElement = $("<img>").attr("src", component.url);
        container.append(imageElement);
      } else if (component.type === "button") {
        const buttonElement = $("<button>").text(component.label);
        buttonElement.click(function() {
          if (component.action === "open_link") {
            window.open(component.data, "_blank");
          }
        });
        container.append(buttonElement);
      }
    });
  }
  