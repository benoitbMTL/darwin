document.addEventListener("DOMContentLoaded", function () {

    // USERNAME
    var userList = [
        { value: "admin", text: "admin" },
        { value: "gordonb", text: "gordonb" },
        { value: "1337", text: "1337" },
        { value: "pablo", text: "pablo" },
        { value: "smithy", text: "smithy" }
    ];

    var userSelectElement = document.getElementById("username");
    userList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        userSelectElement.appendChild(opt);
    });



    // CREDENTIAL STUFFING
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

    var stolenSelectElement = document.getElementById("stolen-credential");
    stolenCredentials.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        stolenSelectElement.appendChild(opt);
    });



    // COUNTRY LIST
    var countryList = [
        { value: "All", text: "All" },
        { value: "Argentina", text: "Argentina" },
        { value: "Australia", text: "Australia" },
        { value: "Brazil", text: "Brazil" },
        { value: "Canada", text: "Canada" },
        { value: "Chile", text: "Chile" },
        { value: "France", text: "France" },
        { value: "Germany", text: "Germany" },
        { value: "Italy", text: "Italy" },
        { value: "Japan", text: "Japan" },
        { value: "Mexico", text: "Mexico" },
        { value: "Norway", text: "Norway" },
        { value: "Spain", text: "Spain" },
        { value: "Ukraine", text: "Ukraine" },
        { value: "United Kingdom", text: "United Kingdom" },
        { value: "United States", text: "United States" },
    ];

    var countrySelectElement = document.getElementById("country");
    countryList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        countrySelectElement.appendChild(opt);
    });


    // API GET findByStatus
    var findByStatusList = [
        { value: "findByStatus?status=available", text: "Status Available" },
        { value: "findByStatus?status=sold", text: "Status Sold" },
        { value: "findByStatus?status=pending", text: "Status Pending" },
        { value: "findByStatus?", text: "Empty Status" },
        { value: "findByStatus?status=ABCDEFGHIJKL", text: "Very Long Status" },
        { value: "findByStatus?status=A", text: "Very Short Status" },
        { value: "findByStatus?status=;cmd.exe", text: "Status with Command Injection" },
        { value: "findByStatus?status=sold&status=pending", text: "Duplicate Status" },
    ];

    var statusSelectElement = document.getElementById("status");
    findByStatusList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        statusSelectElement.appendChild(opt);
    });


    // API POST newPet
    var postNewPetList = [
        { value: "ls;;cmd.exe", text: "New Pet with Command Injection" },
        { value: "xx& var1=l var2=s;$var1$var2", text: "New Pet with Zero-Day" },
        { value: "<script>alert(123)</script>", text: "New Pet with Cross-Site-Scripting" },
    ];

    var userSelectElement = document.getElementById("new-pet");
    postNewPetList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        userSelectElement.appendChild(opt);
    });


    // API PUT Pet
    var putPetList = [
        { value: "aa", text: "aa" },
        { value: "bb", text: "bb" },
        { value: "aa", text: "cc" },
    ];

    var userSelectElement = document.getElementById("modify-pet");
    putPetList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        userSelectElement.appendChild(opt);
    });

});
