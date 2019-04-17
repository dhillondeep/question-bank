let currID = "";

let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");

function readTextFile() {
  const fileToLoad = document.getElementById("fileToLoad").files[0];

  const fileReader = new FileReader();
  fileReader.onload = function(fileLoadedEvent){
    ws.send(JSON.stringify({
      "event": "file_read",
      "data": fileLoadedEvent.target.result
    }))
  };

  fileReader.readAsText(fileToLoad, "UTF-8");
  document.getElementById("fileToLoad").files = null
}

function showRandom() {
  ws.send(JSON.stringify({
    "event": "show_random",
  }))
}

function resetQuestions() {
  ws.send(JSON.stringify({
    "event": "reset_questions",
  }))
}

function checkAnswer() {
  if (currID === "") {
    return
  }

  ws.send(JSON.stringify({
    "event": "check_answer",
    "id": currID,
    "answer": document.getElementById("question_answer").value.toString(),
  }))
}

ws.onmessage = (message) => {
  let obj = JSON.parse(message.data);

  if (obj.event === "show_random") {
    currID = obj.id;
    document.getElementById("question_data").textContent = obj.data ;
    console.log(obj.data)
  } else if (obj.event = "feedback") {
    document.getElementById("feedback").textContent = obj.feedback;
    document.getElementById("attempted_questions").textContent = obj.attempted;
    document.getElementById("correct_questions").textContent = obj.correct;
    document.getElementById("total_questions").textContent = obj.total;
    document.getElementById("used_questions").textContent = obj.used;
  }
};

function setRightBottomHeight() {
  let rightBottomElem = document.getElementsByClassName('inner-bottom')[0];
  let rightTopElem = document.getElementsByClassName('inner-top')[0];

  rightBottomElem.style.height = (window.innerHeight - rightTopElem.clientHeight - 50) + "px";
}

document.addEventListener("DOMContentLoaded", setRightBottomHeight);
window.addEventListener("resize", setRightBottomHeight);
