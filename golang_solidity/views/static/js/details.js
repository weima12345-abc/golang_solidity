

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





let query = window.location.search;
var j_query=query.substr(1)
alert(j_query) 

        let web3 = new Web3('ws://localhost:8545');

$.ajax({
url:"http://localhost:8080/dai/oneFunding",
type:"post", 
dataType:"json",
data:{id:j_query},
success:function (data) {
    $("#pro").html(data.data.Title)
    console.log(data.data)
//   window.location.href="http://localhost:8080/d"
let inserString =
          `
           <div class="col-lg-5 col-xl-4">
          <div class="card shadow mb-4">
        
              <div class="card-header d-flex justify-content-between align-items-center">
                  <h6 class="text-primary fw-bold m-0">标题: ${data.data.Title}</h6>
                  <div class="dropdown no-arrow"><button class="btn btn-link btn-sm dropdown-toggle" aria-expanded="false" data-bs-toggle="dropdown" type="button"><i class="fas fa-ellipsis-v text-gray-400"></i></button>
                      <div class="dropdown-menu shadow dropdown-menu-end animated--fade-in">
                          <span class="add"> <a class="dropdown-item" >&nbsp;收藏</a></span>
                         <a class="dropdown-item" href="../views/details.html">&nbsp;捐赠</a>
                          <div class="dropdown-divider"></div>
                         
                      </div>
                  </div>
              </div>
           
                                    <div  class="card shadow mb-4">
                                  
                                        <img  onclick="get()" src="../static/images/img/b.png">
                                       
                                        </div>
                                       
                                         <div class="text-center small mt-4"><span class="me-2">简介: ${data.data.Info}</span></div> 
                                         <div class="text-center small mt-4"> <span class="me-2">目标金额: ${data.data.Goal}</span></div> 
                                         <div class="text-center small mt-4"> <span class="me-2">已筹集到金额: ${data.data.Amount}</span></div> 
                                        
                                      </div>
                                     </div> 
  `
  let inserStrings = decodeURIComponent(inserString);
              $(".add").html(inserStrings+ document.getElementsByClassName("add")[0].innerHTML)
      
},
error:function (data){
    console.log(data)
    $(".alert").alert()
}
})


//用户捐助

$("#tj").click(()=>{
    var val=$("#val").val();

$.ajax({
    url:"http://localhost:8080/dai/contribute",
    type:"post", 
    dataType:"json",
    data:{id:j_query,value:val},
    success:function (data) {
        var options = getOptions();
        options.title = '信息';
        options.description = '捐助成功！';
        options.width = 300;
        options.type = 'success';
        options.closeTimeout = 3000;
        GrowlNotification.notify(options);

        window.location.href="http://localhost:8080/index";
       
  
          
    },
    error:function (data){
        console.log(data)
        $(".alert").alert()
    }
    })
})