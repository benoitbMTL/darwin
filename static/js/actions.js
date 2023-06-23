///////////////////////////////////////////////////////////////////////////////////
// PING                                                                          //
///////////////////////////////////////////////////////////////////////////////////

function performPing(event) {
    event.preventDefault();
    var ipFqdn = document.getElementById('ip-fqdn').value;

    // Create a new EventSource instance for server-sent events
    var eventSource = new EventSource('/ping?ip-fqdn=' + encodeURIComponent(ipFqdn));

    // Define the onmessage handler to update the ping result with the new data
    eventSource.onmessage = function (event) {
        document.getElementById('ping-result').innerText = event.data;
    };

    // Define the onerror handler
    eventSource.onerror = function (event) {
        console.error('Error:', event);
        eventSource.close();  // Close the connection in case of an error
    };
}

// Attach the performPing function to the submit event of the ping form
document.getElementById('ping-form').addEventListener('submit', performPing);

// Function to reset the ping form and clear the ping result
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

// Function to reset the command injection result
function resetCommandInjection() {
    var iframe = document.getElementById('command-injection-result');

    // Remove the iframe element
    iframe.parentNode.removeChild(iframe);

    // Create a new iframe element
    var newIframe = document.createElement('iframe');
    newIframe.id = 'command-injection-result';
    newIframe.className = 'action-result border';
    newIframe.style.width = '100%';
    newIframe.onload = function () {
        resizeIframe(this);
    };

    // Append the new iframe to its parent container
    var parentContainer = document.getElementById('command-injection-container');
    parentContainer.appendChild(newIframe);
}

// Function to adjust the height of the iframe based on its content:
function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 5) + 'px';
}

