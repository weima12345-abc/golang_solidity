let query = decodeURI(window.location.search);
console.log(decodeURI(query))
var username=(query.substr(1,6).replace(/%/g,'%25'))
var password=(query.substr(7,6).replace(/%/g,'%25')) 
console.log(username)
console.log(password)


$.ajax({
    url:"http://localhost:8080/dai/getUser",
    type:"post",
    dataType:"json",
    data:{username:username,password,password},
    success:function (data){
        $("#username").html(data.data)
        $("#password").html(data.data1)
      
        $("#judge").html("已登录")
        console.log(data)

  // console.log(document.getElementsByClassName("row")[2].innerHTML)
//   var options = getOptions();
//   options.title = '信息';
//   options.description = '删除成功';
//   options.width = 300;
//   options.type = 'success';
//   options.closeTimeout = 3000;
//   GrowlNotification.notify(options);

        
    //    window.location.href="http://localhost:8080/f"
  
           }
  
    
            })
function change(){
 var username=document.getElementById("username1").value;
 var password=document.getElementById("password1").value;
           $("#username").html(username)
           $("#password").html(password)
        $.ajax({
            url:"http://localhost:8080/dai/changeUser",
            type:"post",
            dataType:"json",
            data:{username:username,password:password},
            success:function (data){
          // console.log(document.getElementsByClassName("row")[2].innerHTML)
          $("#username").html(data.data)
          $("#password").html(data.data1)
        
          $("#judge").html("已登录")
          console.log(data)

          
                   }
          
            
                    })
      }
      function exit(){
        // window.location.href="http://localhost:8080/b"
        $("#judge").html("未登录")
        $("#username").html("")
        $("#password").html("")
      }
     
           