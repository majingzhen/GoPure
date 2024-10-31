layui.extend({
    dtree: '/static/dtree/dtree'
}).use(['form', 'layer', 'dtree'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var dtree = layui.dtree;
    var roleId = getUrlParam('id');
    var menuTree;
    
    // 初始化表单
    initForm();
    
    // 监听提交
    form.on('submit(roleAuthForm)', function(data){
        // 获取选中的菜单ID
        var menuIds = menuTree.getCheckbarNodesParam();
        var ids = menuIds.map(item => parseInt(item.nodeId));
        
        // 提交表单
        request.post('/role/auth', {
            id: roleId,
            menuIds: ids
        }).then(() => {
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
            // 加载角色信息和菜单树
            await loadRoleAndMenus();
        } catch (error) {
            console.error('初始化表单失败:', error);
            layer.msg('初始化表单失败');
        }
    }
    
    // 加载角色信息和菜单树
    async function loadRoleAndMenus() {
        // 获取角色信息
        const roleRes = await request.get('/role/get', { id: roleId });
        // 填充表单数据
        form.val('roleAuthForm', roleRes.data);
        
        // 初始化菜单树
        menuTree = dtree.render({
            elem: "#menuTree",
            url: "/menu/tree",
            method: 'get',
            checkbar: true,
            checkbarType: "all",
            done: function(){
                // 获取角色菜单
                request.get('/role/menus', { id: roleId })
                    .then(res => {
                        // 设置选中节点
                        menuTree.checkNodeByIds(res.data);
                    });
            }
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