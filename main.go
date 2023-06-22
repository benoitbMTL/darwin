<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Darwin Project</title>
</head>

<body>
    <h1>Darwin Project</h1>
    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi.</p>

    <form id="ping-form" action="/ping" method="post">
        <label for="ip-fqdn">Enter IP Address or FQDN:</label>
        <input type="text" id="ip-fqdn" name="ip-fqdn" required>
        <button type="submit">Ping</button>
    </form>

    <div id="ping-result"></div>

    <script>
        document.getElementById('ping-form').addEventListener('submit', function(event) {
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
    </script>
</body>

</html>
