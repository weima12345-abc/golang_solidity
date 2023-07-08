// window.onload = () => {
//     const data = {};
//     const $ = window.document.querySelector.bind(window.document);
//   };
//
  
      
  // $(function(){
  //   $('select').searchableSelect();
  //          });
  
  
  
  
  
  
  
        //  function a(){
    
        //      var a=document.getElementById("a");
        //      if(a.style.display="none"){
        //          a.style.display="block"
        //          // alert("显示")
             
        //      }
           
        //  }
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


$.ajax({
    url:"http://localhost:8080/dai/getallFundings",
    type:"post",
    dataType:"json",
    success:function (data){
// console.log(document.getElementsByClassName("row")[2].innerHTML)
        console.log(data.data)

        for(var i=0;i<data.data.length;i++){


            let inserString =
                `
                
                             
                             <tr class="detail">
                             <td>
                             ${data.data[i].Title}</td>
                             
                             <td>　${data.data[i].Info}</td>

                         <td id="${i+1}" onclick="del()"> <a> <button  class="btn btn-primary py-0" type="button" >删除</button></a> </td>
                         
                         </tr>
                             `


            let inserStrings = decodeURIComponent(inserString);
            $.ajax({
                url: 'http://localhost:8080/manage',
                type: 'get',
                datatype: 'json',
                success: function(res) {
                    $(".Add").html(inserStrings+ document.getElementsByClassName("Add")[0].innerHTML)
                }
            })

        }

    },
    error:function (data){
        console.log(data)

    }
})
//管理员新增项目
         function tj(){

          var initiator=document.getElementById("initiator").value;
          
// var id=document.getElementById("id").value;
var info=document.getElementById("info").value;
var goal=document.getElementById("goal").value;
var title=document.getElementById("title").value;

// console.log(img)
// var s_img=img.substr(12,10)     //截取字符
// console.log(img.substr(12,9))


$.ajax({
  url:"http://localhost:8080/dai/newFunding",
  type:"post",
  dataType:"json",
  data:{initiator:initiator,info:info,goal:goal,title:title},
  success:function (data){
// console.log(document.getElementsByClassName("row")[2].innerHTML)
var options = getOptions();
options.title = '信息';
options.description = '新增公益项目成功';
options.width = 300;
options.type = 'success';
options.closeTimeout = 10000;
GrowlNotification.notify(options);

     window.location.href="http://localhost:8080/manage"

         }

  
          })
       
      }

//管理员删除项目

// Array.from( document.getElementsByClassName("del")).forEach(i=>i.onclick=function(){

function del(){
    var arr=document.getElementsByTagName('td');
 
    for(var i=0;i<arr.length;i++){
      arr[i].onclick=function(){
        var b=this.id;
       
        $.ajax({
            url:"http://localhost:8080/dai/delFunding",
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
          
                
               window.location.href="http://localhost:8080/manage"
          
                   }
          
            
                    })
      }
     
	
   
   
    

            }
        }
// })



     
  
     
  

  
       
   
  