layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    
    // 监听提交
    form.on('submit(optionAddForm)', function(data){
        event.preventDefault();
        
        request.post('/option/add', data.field).then(res => {
            if(res.code === 0){
                layer.msg('添加成功', {
                    icon: 1,
                    time: 1000
                }, function(){
                    // 关闭弹窗并刷新父页面
                    var index = parent.layer.getFrameIndex(window.name);
                    parent.layer.close(index);
                    parent.layui.table.reload('optionTable');
                });
            } else {
                layer.msg(res.msg, {icon: 2});
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