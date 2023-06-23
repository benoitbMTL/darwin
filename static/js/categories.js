function showCategory(categoryId) {
    // Hide all categories
    var categories = document.getElementsByClassName('category');
    for (var i = 0; i < categories.length; i++) {
        categories[i].style.display = 'none';
    }

    // Show the selected category
    document.getElementById(categoryId).style.display = 'block';

    // Get all buttons
    var buttons = document.querySelectorAll('.category-buttons .btn');

    // Loop through all buttons
    for (var i = 0; i < buttons.length; i++) {
        var button = buttons[i];

        // If the button's onclick attribute matches the selected category, make it active
        if (button.getAttribute('onclick') === "showCategory('" + categoryId + "')") {
            button.classList.add('active');
            button.setAttribute('aria-pressed', 'true');
        }
        // Otherwise, make it not active
        else {
            button.classList.remove('active');
            button.setAttribute('aria-pressed', 'false');
        }
    }
}
