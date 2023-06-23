document.addEventListener("DOMContentLoaded", function () {
    // Fetch configuration values from the server
    fetch('/api/config')
        .then(response => response.json())
        .then(config => {
            // Populate the form fields with the configuration values
            document.getElementById('dvwa-url').value = config.DVWA_URL;
            document.getElementById('dvwa-host').value = config.DVWA_HOST;
            // Populate other fields in the same way
        })
        .catch(error => console.error('Error fetching configuration:', error));
});

function saveConfiguration() {
    // Implement saving the configuration
}

function resetConfiguration() {
    // Implement resetting the configuration to default values
}