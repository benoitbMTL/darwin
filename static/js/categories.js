function showCategory(categoryId) {
    // Hide all categories
    var categories = document.getElementsByClassName('category');
    for (var i = 0; i < categories.length; i++) {
        categories[i].style.display = 'none';
    }

    // Show the selected category
    document.getElementById(categoryId).style.display = 'block';
}
