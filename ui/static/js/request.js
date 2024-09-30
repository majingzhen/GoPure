$(document).ready(function() {
    // 401 未登录
    $.ajaxSetup({
        complete: function(xhr, status) {
          if (xhr.status === 401) {
            // 处理未授权响应
            window.location.href = '/login';
          }
        }
    });
})