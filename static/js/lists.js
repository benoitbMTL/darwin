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
        { value: "findByStatus?status=xx& var1=l var2=s;$var1$var2", text: "Status with Zero-Day" },
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
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "available" },
            text: "Add new pet FortiPet [id:999]"
        },
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "/bin/ls" },
            text: "New Pet with Command Injection"
        },
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "<script>alert(123)</script>" },
            text: "New Pet with Cross-Site-Scripting"
        },
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "xx& var1=l var2=s;$var1$var2" },
            text: "New Pet with Zero-Day"
        },
        {
            value: "eyAiaWQiOiA5OTksICJuYW1lIjogIkZvcnRpUGV0IiwgImNhdGVnb3J5IjogeyAiaWQiOiAxLCAibmFtZSI6ICJEb2dzIiB9LCAicGhvdG9VcmxzIjogWyJmb3J0aXBldC5wbmciXSwgInRhZ3MiOiBbeyAiaWQiOiAwLCAibmFtZSI6ICJzbyBjdXRlIiB9XSwgInN0YXR1cyI6ICJYNU8hUCVAQVBbNFxcUFpYNTQoUF4pN0NDKTd9JEVJQ0FSLVNUQU5EQVJELUFOVElWSVJVUy1URVNULUZJTEUhJEgrSCoiIH0=",
            text: "New Pet with Malware"
        },
    ];

    var newPetSelectElement = document.getElementById("new-pet");
    postNewPetList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = JSON.stringify(option.value);
        opt.textContent = option.text;
        newPetSelectElement.appendChild(opt);
    });

    // API PUT Pet
    var putPetList = [
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "sold" },
            text: "Modify FortiPet [id:999]"
        },
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "ls;;cmd.exe" },
            text: "Modify FortiPet with Command Injection"
        },
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "<script>alert(123)</script>" },
            text: "Modify FortiPet with Cross-Site-Scripting"
        },
        {
            value: { "id": 999, "name": "FortiPet", "category": { "id": 1, "name": "Dogs" }, "photoUrls": ["fortipet.png"], "tags": [{ "id": 0, "name": "so cute" }], "status": "xx& var1=l var2=s;$var1$var2" },
            text: "Modify FortiPet with Zero-Day"
        },
        {
            value: "eyAiaWQiOiA5OTksICJuYW1lIjogIkZvcnRpUGV0IiwgImNhdGVnb3J5IjogeyAiaWQiOiAxLCAibmFtZSI6ICJEb2dzIiB9LCAicGhvdG9VcmxzIjogWyJmb3J0aXBldC5wbmciXSwgInRhZ3MiOiBbeyAiaWQiOiAwLCAibmFtZSI6ICJzbyBjdXRlIiB9XSwgInN0YXR1cyI6ICJYNU8hUCVAQVBbNFxcUFpYNTQoUF4pN0NDKTd9JEVJQ0FSLVNUQU5EQVJELUFOVElWSVJVUy1URVNULUZJTEUhJEgrSCoiIH0=",
            text: "Modify FortiPet with Malware"
        },
    ];

    var userSelectElement = document.getElementById("modify-pet");
    putPetList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = JSON.stringify(option.value);
        opt.textContent = option.text;
        userSelectElement.appendChild(opt);
    });

    // API DELETE Pet
    var deletePetList = [
        { value: 999, text: "Delete FortiPet [id:999]" },
        { value: "?=cat /etc/passwd", text: "Delete FortiPet with Command Injection" },
        { value: "xx& var1=l var2=s;$var1$var2", text: "Delete FortiPet with Zero-Day" },
    ];

    var userSelectElement = document.getElementById("pet-id");
    deletePetList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = String(option.value);
        opt.textContent = option.text;
        userSelectElement.appendChild(opt);
    });


});
