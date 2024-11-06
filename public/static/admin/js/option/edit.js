layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    
    // 获取URL中的ID参数
    var id = getUrlParam('id');
    
    // 加载选项数据
    request.get('/option/get/' + id).then(res => {
        if(res.code === 0){
            form.val('optionEditForm', res.data);
        } else {
            layer.msg(res.msg);
        }
    });
    
    // 监听提交
    form.on('submit(optionEditForm)', function(data){
        event.preventDefault();
        request.post('/option/update', data.field).then(res => {
            if(res.code === 0){
                layer.msg('修改成功', {
                    icon: 1,
                    time: 1000
                }, function(){
                    // 关闭弹窗并刷新父页面
                    var index = parent.layer.getFrameIndex(window.name);
                    parent.layer.close(index);
                    parent.layui.table.reload('optionTable');
                });
            } else {
                layer.msg(res.msg);
            }
        });
        
        return false;
    });
});

// 取消
function cancel(){
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
} 