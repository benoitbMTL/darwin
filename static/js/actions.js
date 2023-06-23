///////////////////////////////////////////////////////////////////////////////////
// PING                                                                          //
///////////////////////////////////////////////////////////////////////////////////

function performPing(event) {
    event.preventDefault();
    var ipFqdn = document.getElementById('ip-fqdn').value;

    fetch('/ping', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'ip-fqdn=' + encodeURIComponent(ipFqdn)
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('ping-result').innerText = result;
        })
        .catch(error => console.error('Error:', error));
}

function resetPingForm() {
    document.getElementById('ip-fqdn').value = '';
    document.getElementById('ping-result').innerText = '';
}


///////////////////////////////////////////////////////////////////////////////////
// COMMAND INJECTION                                                             //
///////////////////////////////////////////////////////////////////////////////////

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
        .then(htmlContent => {
            document.getElementById('command-injection-result').srcdoc = htmlContent;
            document.getElementById('command-injection-result').style.height = '0px';
        })
        .catch(error => console.error('Error:', error));
}

function resetCommandInjection() {
    var iframe = document.getElementById('command-injection-result');
    iframe.parentNode.removeChild(iframe);
    var newIframe = document.createElement('iframe');
    newIframe.id = 'command-injection-result';
    newIframe.className = 'action-result border';
    newIframe.style.width = '100%';
    newIframe.onload = function () {
        resizeIframe(this);
    };
    var parentContainer = document.getElementById('command-injection-container');
    parentContainer.appendChild(newIframe);
}


///////////////////////////////////////////////////////////////////////////////////
// SQL INJECTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

function performSQLInjection() {
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

    fetch('/sql-injection', { // Changed the endpoint here
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username) + '&password=' + encodeURIComponent(password)
    })
        .then(response => response.text())
        .then(htmlContent => {
            document.getElementById('sql-injection-result').srcdoc = htmlContent;
            document.getElementById('sql-injection-result').style.height = '0px';
        })
        .catch(error => console.error('Error:', error));
}

function resetSQLInjection() {
    var iframe = document.getElementById('sql-injection-result');
    iframe.parentNode.removeChild(iframe);
    var newIframe = document.createElement('iframe');
    newIframe.id = 'sql-injection-result';
    newIframe.className = 'action-result border';
    newIframe.style.width = '100%';
    newIframe.onload = function () {
        resizeIframe(this);
    };
    var parentContainer = document.getElementById('sql-injection-container');
    parentContainer.appendChild(newIframe);
}


///////////////////////////////////////////////////////////////////////////////////
// Function to adjust the height of the iframe based on its content              //
///////////////////////////////////////////////////////////////////////////////////

// Function to adjust the height of the iframe based on its content:
function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 1) + 'px';
}
