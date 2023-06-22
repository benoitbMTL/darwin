document.getElementById('ping-form').addEventListener('submit', function (event) {
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
});

function resetForm() {
    document.getElementById('ip-fqdn').value = '';
    document.getElementById('ping-result').innerText = '';
}
