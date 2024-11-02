layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var pid = getUrlParam('pid');
    $('#pid').val(pid);
    // 监听菜单类型切换
    form.on('radio(menuType)', function(data){
        if(data.value === '0') {
            $('#routeInfo').hide();
        } else {
            $('#routeInfo').show();
        }
    });
    
    // 监听提交
    form.on('submit(menuAddForm)', function(data){
        var formData = data.field;
        
        // 提交表单
        request.post('/menu/add', formData)
            .then(() => {
                layer.msg('添加成功');
                // 关闭弹窗并刷新父页面
                var index = parent.layer.getFrameIndex(window.name);
                parent.layer.close(index);
            });
        return false;
    });

    // 获取URL参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
        var r = window.location.search.substr(1).match(reg);
        if (r != null) return decodeURI(r[2]);
        return null;
    }
});

function cancel() {
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}