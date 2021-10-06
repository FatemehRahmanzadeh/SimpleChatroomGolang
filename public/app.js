window.addEventListener('DOMContentLoaded', (_) => {
    var room = document.getElementById("chat-text");
    var chatform = document.getElementById("input-form");

    var allChats = [];
    if (localStorage.getItem("chats")) {
        allChats = JSON.parse(localStorage.getItem("chats"));
        console.log(allChats)
        allChats.forEach(element => {
            let chatContent = document.createElement('p');
            chatContent.innerHTML = `${element}` + "\n"
            room.append(chatContent);

        });
    } else {
        allChats = [];
    }
    var loginForm = document.getElementById("loginForm");

    if (!sessionStorage.getItem('user')) {
        room.style.display = "none"
        chatform.style.display = "none"
        loginForm.style.display = "block"
    } else {
        room.style.display = "block"
        chatform.style.display = "block"
        loginForm.style.display = "none"
    }

    loginForm.addEventListener('submit', function(e) {
        e.preventDefault()
        var usr = document.getElementById("username").value
        var pass = document.getElementById("password").value
        fetch("http://localhost:8080/api/login", {
                method: "POST",
                headers: new Headers({
                    'content-type': 'application/json',
                    'Accept': 'application/json'
                }),
                body: JSON.stringify({
                    username: usr,
                    password: pass,
                    UserSessionStorage: sessionStorage.getItem("user"),
                }),
            })
            .then((result) => {
                result.text().then(function(data) {
                    console.log(data)
                    res = JSON.parse(data)
                    $.cookie('token', res.token)
                    $.cookie('username', res.username)
                    sessionStorage.setItem("user", "loggedin")
                    room.style.display = "block"
                    chatform.style.display = "block"
                    loginForm.style.display = "none"
                    connectToSocket(res.username)
                })

            })
    });

    function connectToSocket(username) {

        let websocket = new WebSocket("ws://" + window.location.host + "/websocket" + "?bearer=" + $.cookie('token'));

        websocket.addEventListener("message", function(e) {
            let data = JSON.parse(e.data);
            allChats.push(data.username + ":" + data.text)
            localStorage.setItem("chats", JSON.stringify(allChats))
            let chatContent = document.createElement('p');
            chatContent.innerHTML = `${data.username}: ${data.text}`
            room.append(chatContent);
        });

        let form = document.getElementById("input-form");
        form.addEventListener("submit", function(event) {
            event.preventDefault();
            let text = document.getElementById("input-text");
            websocket.send(
                JSON.stringify({
                    username: username,
                    text: text.value,
                })
            );
            text.value = "";
        });
        websocket.onclose = function() {
            localStorage.removeItem("chats")
            sessionStorage.removeItem("user")
        }
    }
})