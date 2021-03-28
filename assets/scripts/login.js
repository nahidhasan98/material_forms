$(document).ready(function () {
    // mdc ripple effect
    mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-username'));
    mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-password'));
    mdc.ripple.MDCRipple.attachTo(document.querySelector('.cancel'));
    mdc.ripple.MDCRipple.attachTo(document.querySelector('.next'));

    $('#loginForm').submit(function () {
        // checking if any field is empty(only space) or not
        let flag = 1;
        if ($('#username').val().trim() == "") {
            $('#username').val("");
            alert("Please give a username!");
            flag = 0;
        }

        if (flag) {
            //sending ajax post request
            let request = $.ajax({
                async: true,
                type: "POST",
                url: "/login",
                data: $('#loginForm').serialize(),
                // error: function (err, statusCode) {
                //     alert(err, statusCode);
                // }
            });
            request.done(function (response) {
                if (response.trim() == "Login Done") {
                    //alert("Login successful!");
                    window.location.href = "/dashboard";
                } else {
                    console.log(response)
                    alert(response);
                }
                // //resetting form field
                // $('#username').val("");
                // $('#password').val("");
            });
            request.fail(function (response) {
                alert(response);
            });
        }
        return false;
    });
});