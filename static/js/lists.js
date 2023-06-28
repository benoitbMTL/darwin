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

    var ipCountry = [
        { value: "United Kingdom", text: "2.24.0.1" },
        { value: "France", text: "2.0.0.1" },
        { value: "Germany", text: "2.160.0.1" },
        { value: "Italy", text: "2.224.0.1" },
        { value: "Spain", text: "2.136.0.1" },
        { value: "Canada", text: "24.0.0.1" },
        { value: "United States", text: "3.0.0.1" },
        { value: "Russia", text: "5.128.0.1" },
        { value: "Brazil", text: "131.0.0.1" },
        { value: "Japan", text: "1.0.16.1" },
        { value: "Australia", text: "1.0.0.1" },
        { value: "Mexico", text: "24.224.0.1" },
        { value: "Ukraine", text: "5.34.0.1" },
        { value: "China", text: "1.1.4.1" },
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
        stolenSelectElement.appendChild(opt);
    });
});
