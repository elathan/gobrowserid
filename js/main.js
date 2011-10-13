function update_login_info(res) {
  var gravatar = 'http://www.gravatar.com/avatar/' +
    Crypto.MD5($.trim(res).toLowerCase()) +
    "?s=32";
    
    document.getElementById("avatar").innerHTML = "<img src='" + gravatar + "'/>";
    document.getElementById("user").innerHTML = res + " | <a href='/logout'>Logout</a>";
    document.getElementById("login").innerHTML = "";
    
}

function login() {
  navigator.id.getVerifiedEmail(function(assertion) {
      if (assertion) {
        // alert(assertion)
        $.ajax({
          type: 'POST',
          url: '/verify',
          data: 'assertion=' + assertion,
          success: function(res, status, xhr) {
            update_login_info(res);
          }
          
        });
      } else {
          // something went wrong!  the user isn't logged in.
      }
  });  
}
