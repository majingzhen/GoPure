layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var roleId = getUrlParam('id');
    
    // 初始化表单
    loadRoleData();
    
    // 监听提交
    form.on('submit(roleEditForm)', function(data){
        var formData = data.field;
        
        // 提交表单
        request.post('/role/update', formData)
            .then(() => {
                layer.msg('保存成功');
                // 关闭弹窗并刷新父页面
                var index = parent.layer.getFrameIndex(window.name);
                parent.layer.close(index);
            });
        return false;
    });
    
    // 加载角色数据
    function loadRoleData() {
        request.get('/role/get', { id: roleId })
            .then(res => {
                // 填充表单数据
                form.val('roleEditForm', res.data);
            });
    }
    
    // 获取URL参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
        var r = window.location.search.substr(1).match(reg);
        if (r != null) return decodeURI(r[2]);
        return null;
    }
}); 