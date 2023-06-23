// Define the userPassMap globally
var userPassMap = {
    "admin": "password",
    "gordonb": "abc123",
    "1337": "charley",
    "pablo": "letmein",
    "smithy": "password"
};

///////////////////////////////////////////////////////////////////////////////////
// COMMAND INJECTION                                                             //
///////////////////////////////////////////////////////////////////////////////////

function performCommandInjection() {
    var username = document.getElementById('username').value;
    var password = userPassMap[username]; // Use the global userPassMap

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
    var password = userPassMap[username]; // Use the global userPassMap

    fetch('/sql-injection', {
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
// BOT DECEPTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

function viewPageSource() {
    // Clear previous results
    document.getElementById('bot-deception-result').innerHTML = '';
    document.getElementById('bot-deception-additional-text').innerHTML = '';

    // Fetch the page source
    fetch('/viewPageSource', {
        method: 'GET'
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('bot-deception-result').innerText = result;
            document.getElementById('bot-deception-additional-text').innerText = 'We can see a hidden link on the login page (display:none):';
        })
        .catch(error => console.error('Error:', error));
}

function performBotDeception() {
    // Clear previous results
    document.getElementById('bot-deception-result').innerHTML = '';
    document.getElementById('bot-deception-additional-text').innerHTML = '';

    // Fetch the fake URL
    fetch('/botDeception', {
        method: 'GET'
    })
        .then(response => response.text())
        .then(htmlContent => {
            document.getElementById('bot-deception-result').innerHTML = htmlContent;
            document.getElementById('bot-deception-additional-text').innerText = 'We simulate a malicious bot by following the hidden link:';
        })
        .catch(error => console.error('Error:', error));
}

function resetBotDeception() {
    document.getElementById('bot-deception-result').innerHTML = '';
    document.getElementById('bot-deception-additional-text').innerHTML = '';
}

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
// Function to adjust the height of the iframe based on its content              //
///////////////////////////////////////////////////////////////////////////////////

// Function to adjust the height of the iframe based on its content:
function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 1) + 'px';
}
