document.addEventListener("DOMContentLoaded", function () {
    var userList = [
        { value: "admin", text: "admin" },
        { value: "gordonb", text: "gordonb" },
        { value: "1337", text: "1337" },
        { value: "pablo", text: "pablo" },
        { value: "smithy", text: "smithy" }
    ];

    var stolenCredentials = [
        { value: "pklangdon4@msn.com", text: "pklangdon4@msn.com" },
        { value: "muldersstan@gmail.com", text: "muldersstan@gmail.com" },
        { value: "forsternp2@aol.com", text: "forsternp2@aol.com" },
        { value: "cragsy@msn.com", text: "cragsy@msn.com" },
        { value: "bjrehdorf@hotmail.com", text: "bjrehdorf@hotmail.com" },
        { value: "baz2709@icloud.com", text: "baz2709@icloud.com" },
        { value: "amysiura@ymail.com", text: "amysiura@ymail.com" },
        { value: "jond714@gmail.com", text: "jond714@gmail.com" },
        { value: "josefahorenstein87@hotmail.com", text: "josefahorenstein87@hotmail.com" },
        { value: "bizotic6@gmail.com", text: "bizotic6@gmail.com" }
    ];

    var countryList = [
        { value: "United Kingdom", text: "United Kingdom" },
        { value: "France", text: "France" },
        { value: "Germany", text: "Germany" },
        { value: "Italy", text: "Italy" },
        { value: "Spain", text: "Spain" },
        { value: "Canada", text: "Canada" },
        { value: "United States", text: "United States" },
        { value: "Russia", text: "Russia" },
        { value: "Brazil", text: "Brazil" },
        { value: "Japan", text: "Japan" },
        { value: "Australia", text: "Australia" },
        { value: "Mexico", text: "Mexico" },
        { value: "Ukraine", text: "Ukraine" },
        { value: "China", text: "China" },
    ];

    var userSelectElement = document.getElementById("username");
    userList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        userSelectElement.appendChild(opt);
    });

    var stolenSelectElement = document.getElementById("stolen-credential");
    stolenCredentials.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        stolenSelectElement.appendChild(opt);
    });

    var countrySelectElement = document.getElementById("country");
    countryList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        countrySelectElement.appendChild(opt);
    });
});