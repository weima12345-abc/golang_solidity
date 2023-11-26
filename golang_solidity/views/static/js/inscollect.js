    
    // var daiaccounts=[];
    // let web3 = new Web3('ws://localhost:8546');
    //   web3.eth.getAccounts().then(function(accounts) {
    //     daiaccounts = accounts;
    //     $("#show").html(daiaccounts[0]);
    //     console.log(daiaccounts[0])
    // });
    
    // ethereum.on('accountsChanged', function(account) {
    //     // Time to reload your interface with accounts[0]!
    
    
    //     $("#show").html(account);
    
    //     daiaccounts[0] = account[0];
    
    // });
  
    function getOptions() {
        var position = $('#notification-position').val();
        var closeTimeout = $('#close-timeout').val();
        var animation = $('#animation').val();
        var showButtons = $('#show-buttons').get(0).checked;
        var showProgressBar = $('#show-progress-bar').get(0).checked;
        var animationOptions = {
            open: animation + '-in',
            close: animation + '-out'
        };
    
        if (animation === 'none') {
            animationOptions = {
                open: false,
                close: false
            };
        }
    
        return options = {
            description: 'I am a success notification',
            position: position,
            closeTimeout: closeTimeout,
            closeWith: ['click'],
            animation: animationOptions,
            showButtons: showButtons,
            buttons: {
                action: {
                    callback: function(notification) {
                        console.log('action button');
                    }
                }
            },
            showProgress: showProgressBar
        };
    }  
    var a="0xE13BE5b9d25f42fC0a9f391d58962cB947F27A81"
$.ajax({
    url:"http://localhost:8080/dai/getallinsFunding",
    type:"post",
    dataType:"json",
    data:{initiator:a},
    success:function (data){
// console.log(document.getElementsByClassName("row")[2].innerHTML)
        console.log(data.data)

        for(var i=0;i<data.data.length;i++){


            let inserString =
                `
                
                <tr>
                <td>
                
                ${data.data[i].Title}</td>
               
                <td>2022/06/09</td>
                <td>${data.data[i].Goal}</td>
               <td> ${data.data[i].Info} </td>
               <td id="${data.data[i].Id}" onclick="del()" > <a> <button class="btn btn-primary py-0" type="button" >删除</button></a> </td>
            </tr>
                             `


            let inserStrings = decodeURIComponent(inserString);
            $.ajax({
                url: 'http://localhost:8080/f',
                type: 'get',
                datatype: 'json',
                success: function(res) {
                    $(".add").html(inserStrings+ document.getElementsByClassName("add")[0].innerHTML)
                }
            })

        }

    },
    error:function (data){
        console.log(data)

    }
})

function del(){
    var arr=document.getElementsByTagName('td');
 
    for(var i=0;i<arr.length;i++){
      arr[i].onclick=function(){
        var b=this.id;
        
        $.ajax({
            url:"http://localhost:8080/dai/delInscollect",
            type:"post",
            dataType:"json",
            data:{id:b},
            success:function (data){
          // console.log(document.getElementsByClassName("row")[2].innerHTML)
          var options = getOptions();
          options.title = '信息';
          options.description = '删除成功';
          options.width = 300;
          options.type = 'success';
          options.closeTimeout = 3000;
          GrowlNotification.notify(options);
       
                
               window.location.href="http://localhost:8080/f"
          
                   }
          
            
                    })
      }
     
            }
        }