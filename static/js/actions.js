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
            var iframe = document.getElementById('command-injection-result');
            iframe.srcdoc = htmlContent;
            iframe.style.height = '0px';
            iframe.style.display = 'block'; // Show the iframe
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
    newIframe.style.display = 'none'; // Hide the iframe
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
            var iframe = document.getElementById('sql-injection-result');
            iframe.srcdoc = htmlContent;
            iframe.style.height = '0px';
            iframe.style.display = 'block'; // Show the iframe
        })
        .catch(error => console.error('Error:', error));
}

function resetSQLInjection() {
    var iframe = document.getElementById('sql-injection-result');
    iframe.srcdoc = '';
    iframe.style.display = 'none'; // Hide the iframe
}

///////////////////////////////////////////////////////////////////////////////////
// BOT DECEPTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

function viewPageSource() {
    fetch('/view-page-source', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('bot-deception-additional-text').innerText = "We can see a hidden link on the login page (display:none)";
            document.getElementById('bot-deception-result').style.display = 'none';
            var textResult = document.getElementById('bot-deception-text-result');
            textResult.innerText = result;
            textResult.style.display = 'block';
        })
        .catch(error => console.error('Error:', error));
}

function performBotDeception() {
    fetch('/bot-deception', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
        .then(response => response.text())
        .then(htmlContent => {
            document.getElementById('bot-deception-additional-text').innerText = "We simulate a malicious bot by following the hidden link";
            var iframeResult = document.getElementById('bot-deception-result');
            iframeResult.srcdoc = htmlContent;
            iframeResult.style.display = 'block';
            document.getElementById('bot-deception-text-result').style.display = 'none';
        })
        .catch(error => console.error('Error:', error));
}

function resetBotDeception() {
    // Reset the <pre> element for view page source
    var preElement = document.getElementById('bot-deception-text-result');
    preElement.innerText = '';
    preElement.style.display = 'none'; // Hide the <pre> element

    // Clear the additional text
    document.getElementById('bot-deception-additional-text').innerText = '';

    // Reset the iframe for bot deception
    var iframe = document.getElementById('bot-deception-result');
    iframe.srcdoc = '';
    iframe.style.display = 'none'; // Make the iframe invisible
}

///////////////////////////////////////////////////////////////////////////////////
// HEALTH CHECK                                                                  //
///////////////////////////////////////////////////////////////////////////////////

function performHealthCheck() {
    fetch('/health-check')
        .then(response => response.text())
        .then(result => {
            var healthCheckResult = document.getElementById('health-check-result');
            healthCheckResult.innerHTML = result;
            healthCheckResult.style.display = 'block';
        })
        .catch(error => console.error('Error:', error));
}

function resetHealthCheck() {
    var healthCheckResult = document.getElementById('health-check-result');
    healthCheckResult.innerHTML = '';
    healthCheckResult.style.display = 'none';
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
            var pingResult = document.getElementById('ping-result');
            pingResult.innerText = result;
            pingResult.style.display = 'block';
        })
        .catch(error => console.error('Error:', error));
}

function resetPingForm() {
    document.getElementById('ip-fqdn').value = '';
    var pingResult = document.getElementById('ping-result');
    pingResult.innerText = '';
    pingResult.style.display = 'none';
}

///////////////////////////////////////////////////////////////////////////////////
// Function to adjust the height of the iframe based on its content              //
///////////////////////////////////////////////////////////////////////////////////

// Function to adjust the height of the iframe based on its content:
function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 1) + 'px';
}
