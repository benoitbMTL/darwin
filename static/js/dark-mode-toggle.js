document.addEventListener("DOMContentLoaded", function () {
    // DARK MODE TOGGLE
    document.getElementById('darkModeSwitch').addEventListener('change', function () {
        var logoElement = document.getElementById('fortiweb-logo');

        console.log("Dark mode toggle triggered"); // Log to console

        if (this.checked) {
            console.log("Switching to dark mode"); // Log to console
            document.body.classList.add('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb-dark.svg"; // Set the dark mode logo
            }
        } else {
            console.log("Switching to light mode"); // Log to console
            document.body.classList.remove('dark-mode');
            if (logoElement) {
                logoElement.src = "/static/images/icon-fortiweb.svg"; // Set the light mode logo
            }
        }
    });
});
