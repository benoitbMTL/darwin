function performCommandInjection() {
    var username = document.getElementById('username').value;

    // Map usernames to passwords
    var userPassMap = {
        "admin": "password",
        "gordonb": "abc123",
        "1337": "charley",
        "pablo": "letmein",
        "smithy": "password"
    };

    var password = userPassMap[username];

    fetch('/command-injection', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username) + '&password=' + encodeURIComponent(password)
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('command-injection-result').innerText = result;
        })
        .catch(error => console.error('Error:', error));
}
