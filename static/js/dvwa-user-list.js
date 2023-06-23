document.addEventListener("DOMContentLoaded", function () {
    var userList = [
        { value: "admin", text: "admin" },
        { value: "gordonb", text: "gordonb" },
        { value: "1337", text: "1337" },
        { value: "pablo", text: "pablo" },
        { value: "smithy", text: "smithy" }
    ];

    var selectElement = document.getElementById("username");
    userList.forEach(function (option) {
        var opt = document.createElement("option");
        opt.value = option.value;
        opt.textContent = option.text;
        selectElement.appendChild(opt);
    });
});
