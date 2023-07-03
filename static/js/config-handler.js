document.addEventListener("DOMContentLoaded", function () {
    // Fetch configuration values from the server
    fetch('/api/config')
        .then(response => response.json())
        .then(config => {
            // Populate the form fields with the configuration values
            document.getElementById('dvwa-url').value = config.dvwa_url;
            document.getElementById('dvwa-host').value = config.dvwa_host;
            document.getElementById('juiceshop-url').value = config.juiceshop_url;
            document.getElementById('fwb-url').value = config.fwb_url;
            document.getElementById('speedtest-url').value = config.speedtest_url;
            document.getElementById('petstore-url').value = config.petstore_url;
            document.getElementById('username-api').value = config.username_api;
            document.getElementById('password-api').value = config.password_api;
            document.getElementById('vdom-api').value = config.vdom_api;
            document.getElementById('token').value = config.token;
            document.getElementById('fwb-mgt-ip').value = config.fwb_mgt_ip;
            document.getElementById('policy').value = config.policy;
            document.getElementById('user-agent').value = config.user_agent;
        })
        .catch(error => console.error('Error fetching configuration:', error));
});

function saveConfiguration() {
    // Gather the values from the form fields
    var config = {
        dvwa_url: document.getElementById('dvwa-url').value,
        dvwa_host: document.getElementById('dvwa-host').value,
        juiceshop_url: document.getElementById('juiceshop-url').value,
        fwb_url: document.getElementById('fwb-url').value,
        speedtest_url: document.getElementById('speedtest-url').value,
        petstore_url: document.getElementById('petstore-url').value,
        username_api: document.getElementById('username-api').value,
        password_api: document.getElementById('password-api').value,
        vdom_api: document.getElementById('vdom-api').value,
        fwb_mgt_ip: document.getElementById('fwb-mgt-ip').value,
        policy: document.getElementById('policy').value,
        user_agent: document.getElementById('user-agent').value
    };

    // Send the configuration to the server
    fetch('/api/config', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(config)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Configuration saved successfully:', data);
            // Fetch the new configuration from the server
            fetch('/api/config')
                .then(response => response.json())
                .then(config => {
                    // Update the token field with the new token
                    document.getElementById('token').value = config.token;
                })
                .catch(error => console.error('Error fetching new configuration:', error));
        })
        .catch(error => {
            console.error('Error saving configuration:', error);
        });

    // Show success message
    var messageDiv = document.getElementById('message');
    messageDiv.innerHTML = 'Saved!';
}

function resetConfiguration() {
    // Fetch the default configuration from the server
    fetch('/api/config/default')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(config => {
            // Populate the form fields with the default configuration
            document.getElementById('dvwa-url').value = config.dvwa_url;
            document.getElementById('dvwa-host').value = config.dvwa_host;
            document.getElementById('juiceshop-url').value = config.juiceshop_url;
            document.getElementById('fwb-url').value = config.fwb_url;
            document.getElementById('speedtest-url').value = config.speedtest_url;
            document.getElementById('petstore-url').value = config.petstore_url;
            document.getElementById('username-api').value = config.username_api;
            document.getElementById('password-api').value = config.password_api;
            document.getElementById('vdom-api').value = config.vdom_api;
            document.getElementById('token').value = config.token;
            document.getElementById('fwb-mgt-ip').value = config.fwb_mgt_ip;
            document.getElementById('policy').value = config.policy;
            document.getElementById('user-agent').value = config.user_agent;
        })
        .catch(error => {
            console.error('Error resetting configuration:', error);
        });

    // Show success message
    var messageDiv = document.getElementById('message');
    messageDiv.innerHTML = 'Configuration reset to default!';
}


