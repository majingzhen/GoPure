layui.extend({
    dtree: '/static/dtree/dtree'
}).use(['form', 'layer', 'dtree'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var dtree = layui.dtree;
    
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
    
    // 初始化上级菜单树
    dtree.render({
        elem: "#parentTree",
        url: "/menu/tree",
        method: 'get',
        initLevel: 1,
        type: "all",
        dataStyle: "layuiStyle",
        response: {
            statusName: "code",
            statusCode: 0,
            message: "msg",
            rootName: "data",
            treeId: "id",
            parentId: "pid",
            title: "name"
        },
        done: function(){
            // 绑定节点点击事件
            dtree.on("node('parentTree')", function(obj){
                $('input[name="pid"]').val(obj.param.nodeId);
                $('input[name="parentName"]').val(obj.param.context);
            });
        }
    });
}); 