///////////////////////////////////////////////////////////////////////////////////
// Function to adjust the height of the iframe based on its content              //
///////////////////////////////////////////////////////////////////////////////////

function resizeIframe(iframe) {
    iframe.style.height = (iframe.contentWindow.document.body.scrollHeight + 7) + 'px';
}

///////////////////////////////////////////////////////////////////////////////////
// COMMAND INJECTION                                                             //
///////////////////////////////////////////////////////////////////////////////////

function performCommandInjection() {
    var username = document.getElementById('username').value;
    var spinner = document.getElementById('spinner');

    // Show the spinner
    //spinner.style.display = 'inline-block';

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
            var errorDiv = document.getElementById('command-injection-error-result');

            // Checks if the HTML content contains an error message.
            if (htmlContent.includes('The Virtual Server is not reachable')) {
                // Insert the HTML content into the errorDiv and display it.
                errorDiv.innerHTML = htmlContent;
                errorDiv.style.display = 'block'; // Show the div
                iframe.style.display = 'none'; // Hide the iframe
            } else {
                // Display the HTML content in the iframe as usual.
                iframe.srcdoc = htmlContent;
                iframe.style.display = 'block'; // Show the iframe
                errorDiv.style.display = 'none'; // Hide the div
            }

            // Hide the spinner
            spinner.style.display = 'none';
        })
        .catch(error => {
            console.error('Error:', error);

            // Hide the spinner in case of an error
            spinner.style.display = 'none';
        });
}

function resetCommandInjection() {
    var iframe = document.getElementById('command-injection-result');
    var errorDiv = document.getElementById('command-injection-error-result');

    // Reset the iframe
    iframe.srcdoc = '';
    iframe.style.display = 'none'; // Hide the iframe

    // Reset the div
    errorDiv.innerHTML = '';
    errorDiv.style.display = 'none'; // Hide the div
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
            var iframe1 = document.getElementById('initial-cookie');
            iframe1.innerHTML = data.initialCookie;
            iframe1.style.display = 'block';

            document.getElementById('modified-cookie-additional-text').innerText = "Let's change the cookie security level to medium.";
            var iframe2 = document.getElementById('modified-cookie');
            iframe2.innerHTML = data.modifiedCookie;
            iframe2.style.display = 'block';

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
// BOT THRESHOLDS                                                                //
///////////////////////////////////////////////////////////////////////////////////

function performBotCrawler() {
    fetch('/bot-threshold', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
        .then(response => response.text())
        .then(htmlContent => {
            document.getElementById('bot-threshold-additional-text').innerText = "Crawler Detection: detects tools that browse your web site for indexing purposes. Monitors the frequency of 403 and 404 response codes.";
            var iframeResult = document.getElementById('bot-threshold-result');
            iframeResult.srcdoc = htmlContent;
            iframeResult.style.display = 'block';
            document.getElementById('bot-threshold-text-result').style.display = 'none';
        })
        .catch(error => console.error('Error:', error));
}

function resetBotThreshold() {
    // Clear the additional text
    document.getElementById('bot-threshold-additional-text').innerText = '';

    // Reset the iframe
    var iframe = document.getElementById('bot-threshold-result');
    iframe.srcdoc = '';
    iframe.style.display = 'none';
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
            document.getElementById('bot-deception-additional-text').innerText = "We can see a hidden link on the login page (display:none).";
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
            document.getElementById('bot-deception-additional-text').innerText = "We simulate a malicious bot by following the hidden link.";
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
// PETSTORE API PROTECTION                                                        //
///////////////////////////////////////////////////////////////////////////////////

function resetPetstoreResult() {
    // Reset the result display area for response body
    var petstoreResultText = document.getElementById('petstore-result-text');
    var petstoreResultHtml = document.getElementById('petstore-result-html');
    var petstoreResultTitle = document.getElementById('petstore-result-title');

    petstoreResultText.innerText = '';
    petstoreResultHtml.srcdoc = '';
    petstoreResultTitle.innerText = '';
    petstoreResultText.style.display = 'none';
    petstoreResultHtml.style.display = 'none';
    petstoreResultTitle.style.display = 'none';

    // Reset the API URL spans
    var apiGetSpan = document.getElementById('api-get');
    var apiPostSpan = document.getElementById('api-post');
    var apiPutSpan = document.getElementById('api-put');
    var apiDeleteSpan = document.getElementById('api-delete');

    apiGetSpan.innerText = '';
    apiPostSpan.innerText = '';
    apiPutSpan.innerText = '';
    apiDeleteSpan.innerText = '';

    // Reset the CURL command display area
    var petstoreCurlText = document.getElementById('petstore-curl-text');
    var petstoreCurlTitle = document.getElementById('petstore-curl-title');

    petstoreCurlText.innerText = '';
    petstoreCurlTitle.innerText = '';
    petstoreCurlText.style.display = 'none';
    petstoreCurlTitle.style.display = 'none';
}

///////////////////////////////////////////////////////////////////////////////////
// PETSTORE GET                                                                  //
///////////////////////////////////////////////////////////////////////////////////

function performPetstoreGETfindByStatus() {
    resetPetstoreResult();
    var selectedOption = document.getElementById('status').value;

    // Fetch the config
    fetch('/config')
        .then(response => response.json())
        .then(config => {
            // Then perform the pet-get request
            fetch('/petstore-pet-get', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: 'status=' + encodeURIComponent(selectedOption)
            })
                .then(response => {
                    var contentType = response.headers.get("content-type");
                                console.log("Response Content-Type:", contentType); // Log the content type

                    if (contentType.includes("application/json")) {
                        return response.json();
                    } else if (contentType.includes("text/plain")) {
                        return response.text();
                    } else if (contentType.includes("text/html")) {
                        return response.text();
                    } else {
                        throw new Error("Unsupported content type: " + contentType);
                    }
                })
                .then(result => {
                    // Display URL from the response
                    document.getElementById('api-get').innerText = result.url;

                    // Display the Curl Title
                    var petstoreCurlTitle = document.getElementById('petstore-curl-title');
                    petstoreCurlTitle.innerText = "Curl";
                    petstoreCurlTitle.style.display = 'block';
                    petstoreCurlTitle.style.fontWeight = "bold";

                    // Display the Curl command
                    var petstoreCurlText = document.getElementById('petstore-curl-text');
                    petstoreCurlText.style.display = 'block';
                    petstoreCurlText.innerText = result.curlCommand;

                    // Display the Response Body Title
                    var petstoreResultTitle = document.getElementById('petstore-result-title');
                    petstoreResultTitle.innerText = "Response Body";
                    petstoreResultTitle.style.display = 'block';
                    petstoreResultTitle.style.fontWeight = "bold";

                    // Display the Response Body
                    var petstoreResultText = document.getElementById('petstore-result-text');
                    var petstoreResultHtml = document.getElementById('petstore-result-html');
                    document.getElementById('petstore-result-title').innerText = "Response Body";
                    document.getElementById('petstore-result-title').style.fontWeight = "bold";

                    if (typeof result.data === 'object') {
                        // JSON data
                        petstoreResultHtml.style.display = 'none';
                        petstoreResultText.style.display = 'block';
                        petstoreResultText.innerText = JSON.stringify(result.data, null, 2);
                    } else if (typeof result.data === 'string') {
                        // Check if the string is HTML
                        if (result.data.indexOf('<') > -1 && result.data.indexOf('>') > -1) {
                            // HTML data
                            petstoreResultText.style.display = 'none';
                            petstoreResultHtml.style.display = 'block';
                            petstoreResultHtml.srcdoc = result.data;
                        } else {
                            // Plain text data
                            petstoreResultHtml.style.display = 'none';
                            petstoreResultText.style.display = 'block';
                            petstoreResultText.innerText = result.data;
                        }
                    }
                })
                .catch((error) => {
                    console.error('Error during fetch operation:', error);
                });
        });
}

///////////////////////////////////////////////////////////////////////////////////
// PETSTORE POST                                                                 //
///////////////////////////////////////////////////////////////////////////////////

function performPetstorePOSTNewPet() {
    resetPetstoreResult();
    var selectedOptionValue = document.getElementById('new-pet').value;

    try {
        var selectedOptionObject = JSON.parse(selectedOptionValue);
    } catch (e) {
        console.error("Error parsing selected option value:", e);
        return; // Exit the function if parsing fails
    }

    // Fetch the config
    fetch('/config')
        .then(response => response.json())
        .then(config => {
            var PETSTORE_URL = config.PETSTORE_URL;

            // Send the POST request to the specified endpoint
            fetch('/petstore-pet-post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(selectedOptionObject),
            })
                .then(response => {
                    var contentType = response.headers.get("content-type");
                    if (contentType.includes("application/json")) {
                        return response.json();
                    } else if (contentType.includes("text/plain")) {
                        return response.text();
                    } else if (contentType.includes("text/html")) {
                        return response.text(); // Treat HTML as text
                    } else {
                        throw new Error("Unsupported content type: " + contentType);
                    }
                })
                .then(result => {
                    // Display URL from the response
                    document.getElementById('api-post').innerText = result.url;

                    // Display the Curl Title
                    var petstoreCurlTitle = document.getElementById('petstore-curl-title');
                    petstoreCurlTitle.innerText = "Curl";
                    petstoreCurlTitle.style.display = 'block';
                    petstoreCurlTitle.style.fontWeight = "bold";

                    // Display the Curl command
                    var petstoreCurlText = document.getElementById('petstore-curl-text');
                    petstoreCurlText.style.display = 'block';
                    petstoreCurlText.innerText = result.curlCommand;

                    // Display the Response Body Title
                    var petstoreResultTitle = document.getElementById('petstore-result-title');
                    petstoreResultTitle.innerText = "Response Body";
                    petstoreResultTitle.style.display = 'block';
                    petstoreResultTitle.style.fontWeight = "bold";

                    // Display the Response Body
                    var petstoreResultText = document.getElementById('petstore-result-text');
                    var petstoreResultHtml = document.getElementById('petstore-result-html');
                    document.getElementById('petstore-result-title').innerText = "Response Body";
                    document.getElementById('petstore-result-title').style.fontWeight = "bold";

                    if (typeof result.data === 'object') {
                        // JSON data
                        petstoreResultHtml.style.display = 'none';
                        petstoreResultText.style.display = 'block';
                        petstoreResultText.innerText = JSON.stringify(result.data, null, 2);
                    } else if (typeof result.data === 'string') {
                        // Check if the string is HTML
                        if (result.data.indexOf('<') > -1 && result.data.indexOf('>') > -1) {
                            // HTML data
                            petstoreResultText.style.display = 'none';
                            petstoreResultHtml.style.display = 'block';
                            petstoreResultHtml.srcdoc = result.data;
                        } else {
                            // Plain text data
                            petstoreResultHtml.style.display = 'none';
                            petstoreResultText.style.display = 'block';
                            petstoreResultText.innerText = result.data;
                        }
                    }
                })
                .catch((error) => {
                    console.error('Error during fetch operation:', error);
                });
        });
}

///////////////////////////////////////////////////////////////////////////////////
// PETSTORE PUT                                                                  //
///////////////////////////////////////////////////////////////////////////////////

function performPetstorePUTPet() {
    resetPetstoreResult();
    var selectedOptionValue = document.getElementById('modify-pet').value;
    // console.log("Selected option object RAW:", selectedOptionValue);
    try {
        var selectedOptionObject = JSON.parse(selectedOptionValue);
        // console.log("Selected option object JSON:", selectedOptionObject);
    } catch (e) {
        console.error("Error parsing selected option value:", e);
        return; // Exit the function if parsing fails
    }

    // Fetch the config
    fetch('/config')
        .then(response => response.json())
        .then(config => {
            // Extract the PETSTORE_URL from the config
            var PETSTORE_URL = config.PETSTORE_URL;

            // Send the POST request to the specified endpoint
            fetch('/petstore-pet-put', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(selectedOptionObject),
            })
                .then(response => {
                    // console.log('Response received:', response);
                    var contentType = response.headers.get("content-type");
                    // console.log('Content-Type:', contentType);
                    if (contentType.includes("application/json")) {
                        return response.json();
                    } else if (contentType.includes("text/plain")) {
                        return response.text();
                    } else if (contentType.includes("text/html")) {
                        return response.text(); // treat HTML as text
                    } else {
                        throw new Error("Unsupported content type: " + contentType);
                    }
                })
                .then(result => {
                    // Display URL from the response
                    document.getElementById('api-put').innerText = result.url;

                    // Display the Curl Title
                    var petstoreCurlTitle = document.getElementById('petstore-curl-title');
                    petstoreCurlTitle.innerText = "Curl";
                    petstoreCurlTitle.style.display = 'block';
                    petstoreCurlTitle.style.fontWeight = "bold";

                    // Display the Curl command
                    var petstoreCurlText = document.getElementById('petstore-curl-text');
                    petstoreCurlText.style.display = 'block';
                    petstoreCurlText.innerText = result.curlCommand;

                    // Display the Response Body Title
                    var petstoreResultTitle = document.getElementById('petstore-result-title');
                    petstoreResultTitle.innerText = "Response Body";
                    petstoreResultTitle.style.display = 'block';
                    petstoreResultTitle.style.fontWeight = "bold";

                    // Display the Response Body
                    var petstoreResultText = document.getElementById('petstore-result-text');
                    var petstoreResultHtml = document.getElementById('petstore-result-html');
                    document.getElementById('petstore-result-title').innerText = "Response Body";
                    document.getElementById('petstore-result-title').style.fontWeight = "bold";

                    if (typeof result.data === 'object') {
                        // JSON data
                        petstoreResultHtml.style.display = 'none';
                        petstoreResultText.style.display = 'block';
                        petstoreResultText.innerText = JSON.stringify(result.data, null, 2);
                    } else if (typeof result.data === 'string') {
                        // Check if the string is HTML
                        if (result.data.indexOf('<') > -1 && result.data.indexOf('>') > -1) {
                            // HTML data
                            petstoreResultText.style.display = 'none';
                            petstoreResultHtml.style.display = 'block';
                            petstoreResultHtml.srcdoc = result.data;
                        } else {
                            // Plain text data
                            petstoreResultHtml.style.display = 'none';
                            petstoreResultText.style.display = 'block';
                            petstoreResultText.innerText = result.data;
                        }
                    }
                })
                .catch((error) => {
                    console.error('Error during fetch operation:', error);
                });
        });
}

///////////////////////////////////////////////////////////////////////////////////
// PETSTORE DELETE                                                               //
///////////////////////////////////////////////////////////////////////////////////

function performPetstoreDELETEPet() {
    resetPetstoreResult();
    var selectedPetId = document.getElementById('pet-id').value;
    // console.log("Selected pet ID:", selectedPetId);
    // console.log("Type of selected pet ID:", typeof selectedPetId);

    if (!selectedPetId) {
        console.error("No pet ID provided");
        return; // Exit the function if no pet ID is provided
    }

    // Fetch the config
    fetch('/config')
        .then(response => response.json())
        .then(config => {
            // Extract the PETSTORE_URL from the config
            var PETSTORE_URL = config.PETSTORE_URL;

            // Send the DELETE request to the specified endpoint
            fetch('/petstore-pet-delete', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: 'petId=' + encodeURIComponent(selectedPetId)
            })
                .then(response => {
                    // console.log('Response received:', response);
                    var contentType = response.headers.get("content-type");
                    // console.log('Content-Type:', contentType);
                    if (contentType.includes("application/json")) {
                        return response.json();
                    } else if (contentType.includes("text/plain")) {
                        return response.text();
                    } else if (contentType.includes("text/html")) {
                        return response.text(); // treat HTML as text
                    } else {
                        throw new Error("Unsupported content type: " + contentType);
                    }
                })
                .then(result => {
                    // Display URL from the response
                    document.getElementById('api-delete').innerText = result.url;

                    // Display the Curl Title
                    var petstoreCurlTitle = document.getElementById('petstore-curl-title');
                    petstoreCurlTitle.innerText = "Curl";
                    petstoreCurlTitle.style.display = 'block';
                    petstoreCurlTitle.style.fontWeight = "bold";

                    // Display the Curl command
                    var petstoreCurlText = document.getElementById('petstore-curl-text');
                    petstoreCurlText.style.display = 'block';
                    petstoreCurlText.innerText = result.curlCommand;

                    // Display the Response Body Title
                    var petstoreResultTitle = document.getElementById('petstore-result-title');
                    petstoreResultTitle.innerText = "Response Body";
                    petstoreResultTitle.style.display = 'block';
                    petstoreResultTitle.style.fontWeight = "bold";

                    // Display the Response Body
                    var petstoreResultText = document.getElementById('petstore-result-text');
                    var petstoreResultHtml = document.getElementById('petstore-result-html');
                    document.getElementById('petstore-result-title').innerText = "Response Body";
                    document.getElementById('petstore-result-title').style.fontWeight = "bold";

                    if (typeof result.data === 'object') {
                        // JSON data
                        petstoreResultHtml.style.display = 'none';
                        petstoreResultText.style.display = 'block';
                        petstoreResultText.innerText = JSON.stringify(result.data, null, 2);
                    } else if (typeof result.data === 'string') {
                        // Check if the string is HTML
                        if (result.data.indexOf('<') > -1 && result.data.indexOf('>') > -1) {
                            // HTML data
                            petstoreResultText.style.display = 'none';
                            petstoreResultHtml.style.display = 'block';
                            petstoreResultHtml.srcdoc = result.data;
                        } else {
                            // Plain text data
                            petstoreResultHtml.style.display = 'none';
                            petstoreResultText.style.display = 'block';
                            petstoreResultText.innerText = result.data;
                        }
                    }
                })
                .catch((error) => {
                    console.error('Error during fetch operation:', error);
                });
        });
}

///////////////////////////////////////////////////////////////////////////////////
// API TRAFFIC GENERATOR                                                         //
///////////////////////////////////////////////////////////////////////////////////

function generateAPITraffic() {
    // console.log("Starting API Traffic Generation...");

    var spinner = document.getElementById('api-spinner');
    // console.log("Displaying spinner...");
    spinner.style.display = 'inline-block';

    fetch('/api-traffic', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
        .then(response => {
            // console.log("Received response from server...");
            return response.text();
        })
        .then(textContent => {
            // console.log("Processing text content...");

            var resultElement = document.getElementById('api-traffic-result');
            resultElement.textContent = textContent; // Set text content directly
            // console.log("Content written to pre element, hiding spinner...");
            resultElement.style.display = 'block';
            spinner.style.display = 'none';
        })
        .catch(error => {
            console.error('Error:', error);
            spinner.style.display = 'none';
        });

    // console.log("API traffic generation request sent.");
}

function resetAPITraffic() {
    // Clear the content of the api-traffic-result element
    var resultElement = document.getElementById('api-traffic-result');
    resultElement.textContent = '';
    resultElement.style.display = 'none'; // Hide the element

    // Hide the spinner, if it's visible
    var spinner = document.getElementById('api-spinner');
    spinner.style.display = 'none';

    // console.log("API Traffic has been reset."); // Debug message
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
        badgeElement.textContent = 'Done';
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
        { id: 'createNewVirtualIP', description: 'Create new Virtual IP' },
        { id: 'createNewServerPool', description: 'Create new Server Pool' },
        { id: 'createNewMemberPool', description: 'Create new Member Pool' },
        { id: 'createNewVirtualServer', description: 'Create new Virtual Server' },
        { id: 'assignVIPToVirtualServer', description: 'Assign Virtual IP to Virtual Server' },
        { id: 'cloneSignatureProtection', description: 'Clone Signature Protection' },
        { id: 'cloneInlineProtection', description: 'Clone Inline Protection' },
        { id: 'createNewXForwardedForRule', description: 'Create new X-Forwarded-For Rule' },
        { id: 'configureProtectionProfile', description: 'Configure Protection Profile' },
        { id: 'createNewPolicy', description: 'Create new Policy' }
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
        { id: 'deleteInlineProtection', description: 'Delete Inline Protection Profile' },
        { id: 'deleteXForwardedForRule', description: 'Delete X-Forwarded-For Rule' },
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