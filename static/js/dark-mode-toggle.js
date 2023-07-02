document.addEventListener("DOMContentLoaded", function () {
    // DARK MODE TOGGLE
    document.getElementById('darkModeSwitch').addEventListener('change', function () {
        var logoElement = document.getElementById('fortiweb-logo');
        var tiles = document.getElementsByClassName('tile');
        var tileHeaders = document.getElementsByClassName('tile-header'); // Get all tile headers
        if (this.checked) {
            document.body.classList.add('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb-dark.svg"; // Set the dark mode logo
            }
            // Add the dark-tile class to all tiles
            for (var i = 0; i < tiles.length; i++) {
                tiles[i].classList.add('dark-tile');
            }
            // Add the dark-tile-header class to all tile headers
            for (var i = 0; i < tileHeaders.length; i++) {
                tileHeaders[i].classList.add('dark-tile-header');
            }
        } else {
            document.body.classList.remove('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb.svg"; // Set the light mode logo
            }
            // Remove the dark-tile class from all tiles
            for (var i = 0; i < tiles.length; i++) {
                tiles[i].classList.remove('dark-tile');
            }
            // Remove the dark-tile-header class from all tile headers
            for (var i = 0; i < tileHeaders.length; i++) {
                tileHeaders[i].classList.remove('dark-tile-header');
            }
        }
    });
});
