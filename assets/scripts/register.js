$(document).ready(function () {
    // mdc ripple effect
    mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-username'));
    mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-email'));
    mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-password'));
    mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-confirm-password'));
    mdc.ripple.MDCRipple.attachTo(document.querySelector('.next'));


    $('#regForm').submit(function () {
        // checking if any field is empty(only space) or not
        let flag = 0;
        if ($('#username').val().trim() == "") {
            $('#username').val("");
            alert("Please give a username!");
            flag = 1;
        }
        if (flag == 0 && ($('#password').val() != $('#confirmPassword').val())) {
            alert("Password didn't match!");
        }

        //sending ajax post request
        let request = $.ajax({
            async: true,
            type: "POST",
            url: "/register",
            data: $('#regForm').serialize(),
            // error: function (err, statusCode) {
            //     alert(err, statusCode);
            // }
        });
        request.done(function (response) {
            if (response.trim() == "Registration Done") {
                alert("Registration successful. Email verification link was sent to your provided email.");
                window.location.href = "/login";
            } else {
                alert("Registration unsuccessful. Something went wrong. Please try again!");
            }
            //resetting form field
            $('#username').val("");
            $('#email').val("");
            $('#password').val("");
            $('#confirmPassword').val("");
        });
        request.fail(function (response) {
            alert(response);
        });

        return false;
    });
});