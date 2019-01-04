

function regexpCheck(expration, testedString) {
  return expration.test(testedString);
}

$(document).ready(function() { 
    
    let regMail = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    let regName = /^[a-zA-Z0-9_-]{3,30}$/;
    let regPass = /^[a-zA-Z0-9_-]{3,60}$/;
    let userIP = "";


    $.getJSON('https://api.ipify.org?format=jsonp&callback=?', function(data) {
      userIP = JSON.stringify(data.ip, null, 2).substr(1).slice(0, -1);
      $("#ip").val(userIP);
      console.log(userIP);//test ip
    });


    /*проверка email*/
    $('#email, #username, #password').on('input', function() {
      validEmail = regexpCheck(regMail, $('#email').val());
      validName = regexpCheck(regName, $('#username').val());
      validPass = regexpCheck(regPass, $('#password').val());
      
      if(validEmail && validName && validPass) {
        $('#btn-submit').prop('disabled', false);
      }
      else {
        $('#btn-submit').prop('disabled', true);
      }
    });
    
});
