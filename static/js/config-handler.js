document.addEventListener("DOMContentLoaded", function () {
    // Fetch configuration values from the server
    fetch('/api/config')
        .then(response => response.json())
        .then(config => {
            // Populate the form fields with the configuration values
            document.getElementById('dvwa-url').value = config.dvwa_url;
            document.getElementById('dvwa-host').value = config.dvwa_host;
            document.getElementById('shop-url').value = config.shop_url;
            document.getElementById('fwb-url').value = config.fwb_url;
            document.getElementById('speedtest-url').value = config.speedtest_url;
            document.getElementById('kali-url').value = config.kali_url;
            document.getElementById('username-api').value = config.username_api;
            document.getElementById('password-api').value = config.password_api;
            document.getElementById('vdom-api').value = config.vdom_api;
            document.getElementById('token').value = config.token;
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
        shop_url: document.getElementById('shop-url').value,
        fwb_url: document.getElementById('fwb-url').value,
        speedtest_url: document.getElementById('speedtest-url').value,
        kali_url: document.getElementById('kali-url').value,
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
        })
        .catch(error => {
            console.error('Error saving configuration:', error);
        });
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
            document.getElementById('shop-url').value = config.shop_url;
            document.getElementById('fwb-url').value = config.fwb_url;
            document.getElementById('speedtest-url').value = config.speedtest_url;
            document.getElementById('kali-url').value = config.kali_url;
            document.getElementById('token').value = config.token;
            document.getElementById('fwb-mgt-ip').value = config.fwb_mgt_ip;
            document.getElementById('policy').value = config.policy;
            document.getElementById('user-agent').value = config.user_agent;
        })
        .catch(error => {
            console.error('Error resetting configuration:', error);
        });
}


