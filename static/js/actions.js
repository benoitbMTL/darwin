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

    // Set the iframe source to about:blank after a short delay
    setTimeout(function () {
        iframe.src = 'about:blank';

        // Reset the iframe height after another short delay
        setTimeout(function () {
            iframe.style.height = '35px'; // Set this to the initial height you want
        }, 100);
    }, 100);
}


// Function to adjust the height of the iframe based on its content:
function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 5) + 'px';
}

