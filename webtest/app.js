const baseURL = 'https://13a2-2405-9800-ba10-716d-dc54-380c-bec9-531c.ngrok-free.app/course';

console.log("Fetching from:", baseURL);

fetch(baseURL, {
    headers: {
      'ngrok-skip-browser-warning': 'anyvalue'
    }
  })
  .then(resp => resp.json())
  .then(data => {
    appendData(data); // เรียกใช้ฟังก์ชัน appendData
  })
  .catch(err => {
    document.querySelector("#myData").innerText = "Error loading data.\n" + err.message;
  });

function appendData(data) {
    var mainContainer = document.getElementById("myData");
    for (var i = 0; i < data.length; i++) {
        var div = document.createElement("div");
        div.innerHTML = 'CourseID: ' + data[i].id + ' ' + data[i].name + ' ' + data[i].price + ' ' + data[i].instructor;
        mainContainer.appendChild(div);
    }
}

