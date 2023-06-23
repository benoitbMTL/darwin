document.addEventListener("DOMContentLoaded", function () {
    // DARK MODE TOGGLE
    document.getElementById('darkModeSwitch').addEventListener('change', function () {
        var logoElement = document.getElementById('fortiweb-logo');
        if (this.checked) {
            document.body.classList.add('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb-dark.svg"; // Set the dark mode logo
            }
        } else {
            document.body.classList.remove('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb.svg"; // Set the light mode logo
            }
        }
    });
});
