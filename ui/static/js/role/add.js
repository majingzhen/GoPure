layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    
    // 监听提交
    form.on('submit(roleAddForm)', function(data){
        var formData = data.field;
        
        // 提交表单
        request.post('/role/add', formData)
            .then(() => {
                layer.msg('添加成功');
                // 关闭弹窗并刷新父页面
                var index = parent.layer.getFrameIndex(window.name);
                parent.layer.close(index);
            });
        return false;
    });
}); 