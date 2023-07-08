

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
//ajax调用合约 返回所有公益项目的具体详情

$.ajax({
  url:"http://localhost:8080/dai/getallFundings",
  type:"post",
  dataType:"json",
  success:function (data){
// console.log(document.getElementsByClassName("row")[2].innerHTML)
      console.log(data)
    
      for(var i=0;i<data.data.length;i++){


          let inserString =
              `
               <div class="col-lg-5 col-xl-4">
              <div class="card shadow mb-4">
            
                  <div class="card-header d-flex justify-content-between align-items-center">
                      <h6 class="text-primary fw-bold m-0">标题: ${data.data[i].Title}</h6>
                      <div class="dropdown no-arrow"><button class="btn btn-link btn-sm dropdown-toggle" aria-expanded="false" data-bs-toggle="dropdown" type="button"><i class="fas fa-ellipsis-v text-gray-400"></i></button>
                          <div class="dropdown-menu shadow dropdown-menu-end animated--fade-in">
                              <span class="add" > <a   class="dropdown-item"  ><li id="${data.data[i].Id}" onclick="ins()"> &nbsp;收藏</li></a></span>
                           
                              <div class="dropdown-divider"></div>
                             
                          </div>
                      </div>
                  </div>
               
                                        <div  class="card shadow mb-4">
                                      
                                            <img id="${i+1}" onclick="get()" src="../static/images/img/b.png">
                                           
                                            </div> 
                                           
                                             <div class="text-center small mt-4"><span class="me-2">简介: ${data.data[i].Info}</span></div> 
                                             <div class="text-center small mt-4"> <span class="me-2">目标金额: ${data.data[i].Goal}</span></div> 
                                           
                                             <div class="text-center small mt-4"> <span class="me-2">已筹集到金额: ${data.data[i].Amount}</span></div>
                                          </div>
                                         </div> 
      `
          let inserStrings = decodeURIComponent(inserString);
          $.ajax({
              url: 'http://localhost:8080/index',
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


//ajax调用合约 返回某个公益项目的具体详情

var id;
function get(){
 
    var arr=document.getElementsByTagName('img');
   for(var i=0;i<arr.length;i++){
     arr[i].onclick=function(){
       id=this.id;
     
      window.location.href="/d/?"+id;
     
     }
    
   }
 

     
    

   
  }

//项目加入收藏
  function ins(){
var a;
    var arr=document.getElementsByTagName('li');
 
   for(var i=0;i<arr.length;i++){
     arr[i].onclick=function(){
       a=this.id;
       console.log(a)
       $.ajax({
        url:"http://localhost:8080/dai/insFunding",
        type:"post",
        dataType:"json",
        data:{id:a},
        success:function (data){
          console.log(data.data)
          var options = getOptions();
          options.title = '信息';
          options.description = '项目加入收藏成功，您可以在收藏页面查看您收藏的公益';
          options.width = 300;
          options.type = 'success';
          options.closeTimeout = 3000;
          GrowlNotification.notify(options);
        

        }
      })
     }
    
     
    
   }
 

     
    

   
  }






// window.onload = () => {
//   const data = {};
//   const $ = window.document.querySelector.bind(window.document);
// };

// $("#myselect").change(function(){
//   var opt=$("#myselect").val();
//   if(opt=="教育助学"){
//     $(".Add").html("")
//            var inserStrings=`
//             <tr class="detail">
//                                             <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                             <!-- <td>爱心午餐</td> -->
//                                             <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                            
//                                            <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                            
//                                         </tr>
//                                         `
//             $.ajax({
//               url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//               type: 'get',
//               datatype: 'json',
//               success: function(res) {
//                   $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//               }
//           })
//  }else if(opt=="全部项目"){
//   $(".Add").html("")
//   for (var i=1;i<12;i++){

 
//          var inserStrings=`
//           <tr class="detail">
//                                           <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                           <!-- <td>爱心午餐</td> -->
//                                           <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                          
//                                          <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                          
//                                       </tr>
//                                       `
//           $.ajax({
//             url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//             type: 'get',
//             datatype: 'json',
//             success: function(res) {
//                 $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//             }
//         })
//       }}
//  else if(opt=="疾病救助"){
//   $(".Add").html("")
//   for (var i=1;i<4;i++){

 
//          var inserStrings=`
//           <tr class="detail">
//                                           <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                           <!-- <td>爱心午餐</td> -->
//                                           <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                          
//                                          <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                          
//                                       </tr>
//                                       `
//           $.ajax({
//             url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//             type: 'get',
//             datatype: 'json',
//             success: function(res) {
//                 $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//             }
//         })
//       }
// }
// else if(opt=="疾病救助"){
//   $(".Add").html("")
//   for (var i=1;i<4;i++){

 
//          var inserStrings=`
//           <tr class="detail">
//                                           <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                           <!-- <td>爱心午餐</td> -->
//                                           <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                          
//                                          <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                          
//                                       </tr>
//                                       `
//           $.ajax({
//             url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//             type: 'get',
//             datatype: 'json',
//             success: function(res) {
//                 $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//             }
//         })
//       }}
//       else if(opt=="济困救灾"){
//         $(".Add").html("")
//         for (var i=1;i<6;i++){
      
       
//                var inserStrings=`
//                 <tr class="detail">
//                                                 <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                                 <!-- <td>爱心午餐</td> -->
//                                                 <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                                
//                                                <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                                
//                                             </tr>
//                                             `
//                 $.ajax({
//                   url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//                   type: 'get',
//                   datatype: 'json',
//                   success: function(res) {
//                       $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//                   }
//               })
//             }}else{
//               $(".Add").html("")
//               for (var i=1;i<3;i++){
            
             
//                      var inserStrings=`
//                       <tr class="detail">
//                                                       <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                                       <!-- <td>爱心午餐</td> -->
//                                                       <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                                      
//                                                      <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                                      
//                                                   </tr>
//                                                   `
//                       $.ajax({
//                         url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//                         type: 'get',
//                         datatype: 'json',
//                         success: function(res) {
//                             $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//                         }
//                     })
//             }}
// });
    
// // $(function(){
// //   $('select').searchableSelect();
// //          });







//        function a(){
  
//            var a=document.getElementById("a");
//            if(a.style.display="none"){
//                a.style.display="block"
//                // alert("显示")
           
//            }
         
//        }

//        function q(){
//            $(".Add").html("")
//           var inserStrings=`
//            <tr class="detail">
//                                            <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                            <!-- <td>爱心午餐</td> -->
//                                            <td>　“楼道更干净了，邻里关系更和谐温馨了，大家参加‘邻舍+’志愿活动的热情很高涨。我们既是参与者，也是受益者。”嵊州市仙景花苑居民丁伟东说。嵊州市委宣传部工作人员冯波表示，当地通过“邻舍+”志愿服务品牌“4+计划”项目化运作，以“小家”的和睦带动“大家”的文明，实现了“孩子和乐、家庭和美、邻舍和睦、社会和谐”的“四和”愿景。<br></td>
                                           
//                                           <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                           
//                                        </tr>
//                                        `
//            $.ajax({
//              url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//              type: 'get',
//              datatype: 'json',
//              success: function(res) {
//                  $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
//              }
//          })
//        }

//        function a(){
  
//   var a=document.getElementById("a");
//   if(a.style.display="none"){
//       a.style.display="block"
//       // alert("显示")
  
//   }

// }

// function w(){
  
//   $(".Add").html("")
//   for(i=1;i<=2;i++){
//  var inserStrings=`
//   <tr class="detail">
//                                   <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                   <!-- <td>爱心午餐</td> -->
//                                   <td>　“邻舍+”是“邻居”的嵊州本土方言，主要通过开展“邻舍+”宝贝星、“邻舍+”万家和等亲子志愿服务携手共建文明。“邻舍+”志愿服务队以“爱相邻，和为贵”为宗旨，秉承“机制化保障、项目化运作、品牌化经营”的思路，由嵊州市文明办牵头主导，结合新时代文明实践阵地，在各社区层面设立“邻舍+”志愿服务点，为“邻舍+”志愿服务提供活动场地，实行纵向扁平化管理模式。统一制作“邻舍+”标识牌设置于志愿家庭门口，亮明“身份”；<br></td>
                                  
//                                  <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                  
//                               </tr>
//                               `
//   $.ajax({
//     url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//     type: 'get',
//     datatype: 'json',
//     success: function(res) {
        
//         $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
        
//     }
// })
// }
// }

// function e(){
//   $(".Add").html("")
//   for(i=1;i<=3;i++){
//  var inserStrings=`
//   <tr class="detail">
//                                   <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                   <!-- <td>爱心午餐</td> -->
//                                   <td>　“邻舍+”是“邻居”的嵊州本土方言，主要通过开展“邻舍+”宝贝星、“邻舍+”万家和等亲子志愿服务携手共建文明。“邻舍+”志愿服务队以“爱相邻，和为贵”为宗旨，秉承“机制化保障、项目化运作、品牌化经营”的思路，由嵊州市文明办牵头主导，结合新时代文明实践阵地，在各社区层面设立“邻舍+”志愿服务点，为“邻舍+”志愿服务提供活动场地，实行纵向扁平化管理模式。统一制作“邻舍+”标识牌设置于志愿家庭门口，亮明“身份”；<br></td>
                                  
//                                  <td> <a href="../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                  
//                               </tr>
//                               `
//   $.ajax({
//     url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//     type: 'get',
//     datatype: 'json',
//     success: function(res) {
        
//         $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
        
//     }
// })
// }
// }

// function r(){
//   $(".Add").html("")
//   for(i=1;i<=10;i++){
//  var inserStrings=`
//   <tr class="detail">
//                                   <td><img class="rounded-circle me-2" width="100" height="100" src="../asset1/asset1/img/avatars/avatar1.jpeg">爱心午餐1</td>
//                                   <!-- <td>爱心午餐</td> -->
//                                   <td>　“邻舍+”是“邻居”的嵊州本土方言，主要通过开展“邻舍+”宝贝星、“邻舍+”万家和等亲子志愿服务携手共建文明。“邻舍+”志愿服务队以“爱相邻，和为贵”为宗旨，秉承“机制化保障、项目化运作、品牌化经营”的思路，由嵊州市文明办牵头主导，结合新时代文明实践阵地，在各社区层面设立“邻舍+”志愿服务点，为“邻舍+”志愿服务提供活动场地，实行纵向扁平化管理模式。统一制作“邻舍+”标识牌设置于志愿家庭门口，亮明“身份”；<br></td>
                                  
//                                  <td> <a href=""../views/details.html"> <button class="btn btn-primary py-0" type="button" >详情</button></a> </td>
                                  
//                               </tr>
//                               `
//   $.ajax({
//     url: 'http://127.0.0.1:5501/static/views/project_classify.html',
//     type: 'get',
//     datatype: 'json',
//     success: function(res) {
        
//         $(".Add").html(inserStrings + document.getElementsByClassName("Add")[0].innerHTML)
        
//     }
// })
// }
// }

// Array.from(document.getElementsByClassName("detail")).forEach(i =>i.onclick=function(){
  
//    window.location.href="http://127.0.0.1:5501/static/views/details.html"
// })




     
 
