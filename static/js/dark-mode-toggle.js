document.addEventListener("DOMContentLoaded", function () {
    // DARK MODE TOGGLE
    document.getElementById('darkModeSwitch').addEventListener('change', function () {
        var logoElement = document.getElementById('fortiweb-logo');
        var tiles = document.getElementsByClassName('tile');
        if (this.checked) {
            document.body.classList.add('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb-dark.svg"; // Set the dark mode logo
            }
            // Add the dark-tile class to all tiles
            for (var i = 0; i < tiles.length; i++) {
                tiles[i].classList.add('dark-tile');
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
        }
    });
});
