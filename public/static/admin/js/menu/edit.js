layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var menuId = getUrlParam('id');
    
    // 初始化表单
    initForm();
    
    // 监听菜单类型切换
    form.on('radio(menuType)', function(data){
        if(data.value === '0') {
            $('#routeInfo').hide();
        } else {
            $('#routeInfo').show();
        }
    });
    
    // 监听提交
    form.on('submit(menuEditForm)', function(data){
        var formData = data.field;
        
        // 提交表单
        request.post('/menu/update', formData)
            .then(res => {
                if (res.code !== 0) {
                    layer.msg(res.msg, {icon: 2});
                    return;
                }
                layer.msg('保存成功');
                // 关闭弹窗并刷新父页面
                var index = parent.layer.getFrameIndex(window.name);
                parent.layer.close(index);
            });
        return false;
    });
    
    // 初始化表单
    async function initForm() {
        try {
            // 加载菜单数据
            await loadMenuData();
        } catch (error) {
            console.error('初始化表单失败:', error);
            layer.msg('初始化表单失败');
        }
    }
    
    // 加载菜单数据
    async function loadMenuData() {
        // 获取菜单信息
        const menuRes = await request.get('/menu/get/' + menuId);
        // 填充表单数据
        form.val('menuEditForm', menuRes.data);
        
        // 根据菜单类型显示/隐藏路由信息
        if(menuRes.data.menuType === '0') {
            $('#routeInfo').hide();
        } else {
            $('#routeInfo').show();
        }
    }
    
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