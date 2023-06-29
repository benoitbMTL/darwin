///////////////////////////////////////////////////////////////////////////////////
// Function to adjust the height of the iframe based on its content              //
///////////////////////////////////////////////////////////////////////////////////

function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 5) + 'px';
}

///////////////////////////////////////////////////////////////////////////////////
// COMMAND INJECTION                                                             //
///////////////////////////////////////////////////////////////////////////////////

function performCommandInjection() {
    var username = document.getElementById('username').value;

    fetch('/command-injection', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username)
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
    iframe.srcdoc = '';
    iframe.style.display = 'none'; // Hide the iframe
}

///////////////////////////////////////////////////////////////////////////////////
// SQL INJECTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

function performSQLInjection() {
    var username = document.getElementById('username').value;

    fetch('/sql-injection', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username)
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
// CROSS SITE SCRIPTING (XSS)                                                    //
///////////////////////////////////////////////////////////////////////////////////

function performCrossSiteScripting() {
    var username = document.getElementById('username').value;

    fetch('/cross-site-scripting', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username)
    })
        .then(response => response.text())
        .then(htmlContent => {
            var iframe = document.getElementById('cross-site-scripting-result');
            iframe.srcdoc = htmlContent;
            iframe.style.height = '0px';
            iframe.style.display = 'block'; // Show the iframe
        })
        .catch(error => console.error('Error:', error));
}

function resetCrossSiteScripting() {
    var iframe = document.getElementById('cross-site-scripting-result');
    iframe.srcdoc = '';
    iframe.style.display = 'none'; // Hide the iframe
}

///////////////////////////////////////////////////////////////////////////////////
// COOKIE SECURITY                                                                 //
///////////////////////////////////////////////////////////////////////////////////

function performCookieSecurity() {
    var username = document.getElementById('username').value;
    fetch('/cookie-security', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username)
    })
        .then(response => response.json())
        .then(data => {

            document.getElementById('initial-cookie-additional-text').innerText = "You are now authenticated. Your cookie security level is set to low.";
            let initialCookieHtml = '<html><body><pre>' + data.initialCookie.replace(/low/g, '<span style="color: red;">low</span>') + '</pre></body></html>';
            var iframe1 = document.getElementById('initial-cookie');
            iframe1.srcdoc = initialCookieHtml;
            iframe1.style.display = 'block';
            iframe1.onload = function () {
                iframe1.style.height = (iframe1.contentWindow.document.body.scrollHeight + 30) + 'px';
            }

            document.getElementById('modified-cookie-additional-text').innerText = "Let's change the cookie security level to medium.";
            let modifiedCookieHtml = '<html><body><pre>' + data.modifiedCookie.replace(/medium/g, '<span style="color: red;">medium</span>') + '</pre></body></html>';
            var iframe2 = document.getElementById('modified-cookie');
            iframe2.srcdoc = modifiedCookieHtml;
            iframe2.style.display = 'block';
            iframe2.onload = function () {
                iframe2.style.height = (iframe2.contentWindow.document.body.scrollHeight + 30) + 'px';
            }

            document.getElementById('web-page-iframe-additional-text').innerText = "Let's connect again to the web app with the new crafted cookie.";
            var iframe3 = document.getElementById('web-page-iframe');
            iframe3.srcdoc = data.webPageHTML;
            iframe3.style.height = '0px';
            iframe3.style.display = 'block';

        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function resetCookieSecurity() {
    document.getElementById('initial-cookie-additional-text').innerText = '';
    document.getElementById('modified-cookie-additional-text').innerText = '';
    document.getElementById('web-page-iframe-additional-text').innerText = '';


    var iframe1 = document.getElementById('initial-cookie');
    iframe1.srcdoc = '';
    iframe1.style.display = 'none'; // Hide the iframe

    var iframe2 = document.getElementById('modified-cookie');
    iframe2.srcdoc = '';
    iframe2.style.display = 'none'; // Hide the iframe

    var iframe3 = document.getElementById('web-page-iframe');
    iframe3.srcdoc = '';
    iframe3.style.display = 'none'; // Hide the iframe
}

///////////////////////////////////////////////////////////////////////////////////
// CREDENTIAL STUFFING                                                           //
///////////////////////////////////////////////////////////////////////////////////

function performCredentialStuffing() {
    var username = document.getElementById('stolen-credential').value;

    fetch('/credential-stuffing', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'username=' + encodeURIComponent(username)
    })
        .then(response => response.text())
        .then(htmlContent => {
            var iframe = document.getElementById('credential-stuffing-result');
            iframe.srcdoc = htmlContent;
            iframe.style.height = '0px';
            iframe.style.display = 'block'; // Show the iframe
        })
        .catch(error => console.error('Error:', error));
}

function resetCredentialStuffing() {
    var iframe = document.getElementById('credential-stuffing-result');
    iframe.srcdoc = '';
    iframe.style.display = 'none'; // Hide the iframe
}

///////////////////////////////////////////////////////////////////////////////////
// WEB SCAN                                                                      //
///////////////////////////////////////////////////////////////////////////////////

function performWebScan() {
    document.getElementById('web-scan-spinner').style.display = 'inline-block';

    var country = document.getElementById('country').value;

    fetch('/web-scan', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: 'country=' + encodeURIComponent(country)
    })
        .then(response => response.text())
        .then(result => {
            var scanResult = document.getElementById('web-scan-result');
            scanResult.innerText = result;
            scanResult.style.display = 'block';
        })
        .catch(error => {
            console.error('Error:', error);
        })
        .finally(() => {
            document.getElementById('web-scan-spinner').style.display = 'none';
        });
}

function resetWebScan() {
    document.getElementById('country').value = 'All';
    var scanResult = document.getElementById('web-scan-result');
    scanResult.innerText = '';
    scanResult.style.display = 'none';
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
// REST API CREATE POLICY                                                        //
///////////////////////////////////////////////////////////////////////////////////

function performOnboardNewApplicationPolicy() {
    var spinner = document.getElementById('create-spinner');
    spinner.style.display = 'inline-block';

    fetch('/create-policy', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
    })
        .then(response => response.json())
        .then(data => {
            // Iterate over the array of statuses and update the task status for each one
            data.forEach(status => {
                updateTaskStatus(status.taskId, status.status, status.description);
            });
            // Hide the spinner when all tasks are done
            spinner.style.display = 'none';
        })
        .catch(error => {
            console.error('Error:', error);
            // Hide the spinner in case of an error
            spinner.style.display = 'none';
        });
}

function performDeleteApplicationPolicy() {
    var spinner = document.getElementById('delete-spinner');
    spinner.style.display = 'inline-block';

    fetch('/delete-policy', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
    })
        .then(response => response.json())
        .then(data => {
            // Iterate over the array of statuses and update the task status for each one
            data.forEach(status => {
                updateTaskStatus(status.taskId, status.status, status.description);
            });
            // Hide the spinner when all tasks are done
            spinner.style.display = 'none';
        })
        .catch(error => {
            console.error('Error:', error);
            // Hide the spinner in case of an error
            spinner.style.display = 'none';
        });
}

function updateTaskStatus(taskId, status, description) {
    var taskElement = document.getElementById(taskId);
    var badgeElement = taskElement.querySelector('.badge');
    var descriptionElement = taskElement.querySelector('.task-description');

    if (status === 'success') {
        badgeElement.textContent = 'Complete';
        badgeElement.classList.remove('bg-primary');
        badgeElement.classList.add('bg-success');
        descriptionElement.textContent = description;
    } else {
        badgeElement.textContent = 'Failed';
        badgeElement.classList.remove('bg-primary');
        badgeElement.classList.add('bg-danger');
        descriptionElement.textContent = description;
    }
}

function resetOnboardNewApplicationPolicy() {
    var tasks = [
        { id: 'createNewVirtualIP', description: 'Create New Virtual IP' },
        { id: 'createNewServerPool', description: 'Create New Server Pool' },
        { id: 'createNewMemberPool', description: 'Create New Member Pool' },
        { id: 'createNewVirtualServer', description: 'Create New Virtual Server' },
        { id: 'assignVirtualIP', description: 'Assign Virtual IP to Virtual Server' },
        { id: 'cloneSignatureProtection', description: 'Clone Signature Standard Protection' },
        { id: 'cloneInlineProtection', description: 'Clone Inline Standard Protection' },
        { id: 'configureProtectionProfile', description: 'Configure Protection Profile' },
        { id: 'createNewPolicy', description: 'Create New Policy' }
    ];
    tasks.forEach(task => {
        var taskElement = document.getElementById(task.id);
        var badgeElement = taskElement.querySelector('.badge');
        var descriptionElement = taskElement.querySelector('.task-description');

        // Reset the badge to 'Incomplete'
        badgeElement.textContent = 'Incomplete';
        badgeElement.classList.remove('bg-success', 'bg-danger');
        badgeElement.classList.add('bg-secondary');

        // Reset the task description
        descriptionElement.textContent = task.description;
    });
}

function resetDeleteApplicationPolicy() {
    var tasks = [
        { id: 'deletePolicy', description: 'Delete Policy' },
        { id: 'deleteProtectionProfile', description: 'Delete Protection Profile' },
        { id: 'deleteSignatureProtection', description: 'Delete Signature Protection' },
        { id: 'deleteVirtualServer', description: 'Delete Virtual Server' },
        { id: 'deleteServerPool', description: 'Delete Server Pool' },
        { id: 'deleteVirtualIP', description: 'Delete Virtual IP' }
    ];
    tasks.forEach(task => {
        var taskElement = document.getElementById(task.id);
        var badgeElement = taskElement.querySelector('.badge');
        var descriptionElement = taskElement.querySelector('.task-description');

        // Reset the badge to 'Incomplete'
        badgeElement.textContent = 'Incomplete';
        badgeElement.classList.remove('bg-success', 'bg-danger');
        badgeElement.classList.add('bg-secondary');

        // Reset the task description
        descriptionElement.textContent = task.description;
    });
}

///////////////////////////////////////////////////////////////////////////////////
// HEALTH CHECK                                                                  //
///////////////////////////////////////////////////////////////////////////////////

function performHealthCheck() {
    document.getElementById('health-check-spinner').style.display = 'inline-block';

    fetch('/health-check')
        .then(response => response.text())
        .then(result => {
            var healthCheckResult = document.getElementById('health-check-result');
            healthCheckResult.innerHTML = result;
            healthCheckResult.style.display = 'block';
        })
        .catch(error => {
            console.error('Error:', error);
        })
        .finally(() => {
            document.getElementById('health-check-spinner').style.display = 'none';
        });
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

    document.getElementById('spinner').style.display = 'inline-block';

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
        .catch(error => {
            console.error('Error:', error);
        })
        .finally(() => {
            document.getElementById('spinner').style.display = 'none';
        });
}

function resetPingForm() {
    document.getElementById('ip-fqdn').value = '';
    var pingResult = document.getElementById('ping-result');
    pingResult.innerText = '';
    pingResult.style.display = 'none';
}

