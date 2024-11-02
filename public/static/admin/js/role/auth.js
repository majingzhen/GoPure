
var form,layer,roleId,authSelect,menuData;
layui.use(['form', 'layer'], function(){
    form = layui.form;
    layer = layui.layer;
    roleId = getUrlParam('id');
    // 初始化表单
    initForm();
    // 监听提交
    form.on('submit(roleAuthForm)', function(data){
        // 获取选中的菜单ID
        var menuIds = authSelect.getValue().map(obj => obj.id);
        // 提交表单
        request.post('/role/authRole', {
            id: parseInt(roleId),
            menuIds: menuIds
        }).then(() => {
            layer.msg('保存成功');
            // 关闭弹窗并刷新父页面
            var index = parent.layer.getFrameIndex(window.name);
            parent.layer.close(index);
        });
        return false;
    });
});
// 初始化表单
function initForm() {
    try {
        // 加载角色信息和菜单树
        loadRoleAndMenus();
    } catch (error) {
        console.error('初始化表单失败:', error);
        layer.msg('初始化表单失败');
    }
}

// 取消
function cancel() {
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}

// 加载角色信息和菜单树
function loadRoleAndMenus() {
    request.get('/menu/list').then(res => {
        menuData = res.data;
    })
    // 获取角色信息
    request.get('/role/get', { id: roleId }).then(res => {
        // 填充表单数据
        form.val('roleAuthForm', res.data);
        console.log(menuData)
        authSelect = xmSelect.render({
            el: "#menuTree",
            name: "menuTree",
            initValue: res.data.menuIds,
            data: menuData,
            tree: {
                show:true,
                expandedKeys:res.data.menuIds,
                showLine: true,
                indent: 20,
            },
            prop: {
                value: 'id',
                name: 'name'
            },
            toolbar: {show:true,list:["ALL","CLEAR","REVERSE"]},
        });
    })

}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
}