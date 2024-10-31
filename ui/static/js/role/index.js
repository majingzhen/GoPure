layui.use(['table', 'form', 'layer'], function(){
    var table = layui.table;
    var form = layui.form;
    var layer = layui.layer;
    
    // 渲染表格
    table.render({
        elem: '#roleTable',
        url: '/role/page',
        method: 'get',
        toolbar: '#toolbarTpl',
        defaultToolbar: ['filter', 'exports', 'print'],
        cols: [[
            {type: 'checkbox', fixed: 'left'},
            {field: 'id', title: 'ID', sort: true, width: 80},
            {field: 'name', title: '角色名称'},
            {field: 'code', title: '角色编码'},
            {field: 'description', title: '描述'},
            {field: 'createTime', title: '创建时间', sort: true},
            {title: '操作', toolbar: '#rowToolbarTpl', width: 120}
        ]],
        page: true,
        request: {
            pageName: 'pageNum',
            limitName: 'pageSize'
        },
        response: {
            statusName: 'code',
            statusCode: 0,
            msgName: 'msg',
            countName: 'total',
            dataName: 'rows'
        },
        parseData: function(res){
            return {
                "code": res.code,
                "msg": res.msg,
                "total": res.data.total,
                "rows": res.data.rows
            };
        }
    });

    // 搜索表单提交
    form.on('submit(searchForm)', function(data){
        if(data.field.status === '') {
            delete data.field.status;
        }
        
        table.reload('roleTable', {
            where: data.field,
            page: {curr: 1}
        });
        return false;
    });

    // 头部工具栏事件
    table.on('toolbar(roleTable)', function(obj){
        switch(obj.event){
            case 'add':
                openRoleAddForm();
                break;
            case 'batchDel':
                var checkStatus = table.checkStatus('roleTable');
                var data = checkStatus.data;
                if(data.length === 0){
                    layer.msg('请选择要删除的数据');
                    return;
                }
                var ids = data.map(item => parseInt(item.id));
                layer.confirm('确定删除选中的角色吗？', function(index){
                    deleteRoles(ids);
                    layer.close(index);
                });
                break;
        }
    });

    // 表格行工具事件
    table.on('tool(roleTable)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){
            layer.confirm('确定删除该角色吗？', function(index){
                deleteRoles([parseInt(data.id)]);
                layer.close(index);
            });
        } else if(obj.event === 'edit'){
            openRoleEditForm(parseInt(data.id));
        } else if(obj.event === 'auth'){
            openRoleAuthForm(parseInt(data.id));
        }
    });

    // 状态切换事件
    form.on('switch(statusSwitch)', function(obj){
        var roleId = parseInt(this.value);
        var status = obj.elem.checked ? "0" : "1";
        updateRoleStatus(roleId, status);
    });
});

// 打开角色添加表单
function openRoleAddForm() {
    layer.open({
        type: 2,
        title: '新增角色',
        area: ['600px', '80%'],
        content: '/role/add',
        maxmin: true,
        end: function(){
            layui.table.reload('roleTable');
        }
    });
}

// 打开角色编辑表单
function openRoleEditForm(id) {
    layer.open({
        type: 2,
        title: '编辑角色',
        area: ['600px', '80%'],
        content: '/role/edit?id=' + id,
        maxmin: true,
        end: function(){
            layui.table.reload('roleTable');
        }
    });
}

// 删除角色
function deleteRoles(ids) {
    request.post('/role/delete', {ids: ids})
        .then(() => {
            layer.msg('删除成功');
            layui.table.reload('roleTable');
        });
}

// 更新角色状态
function updateRoleStatus(id, status) {
    request.post('/role/updateStatus', {id: id, status: status})
        .then(() => {
            layer.msg('更新成功');
        })
        .catch(() => {
            layui.table.reload('roleTable');
        });
}

// 添加打开角色授权表单的函数
function openRoleAuthForm(id) {
    layer.open({
        type: 2,
        title: '角色授权',
        area: ['600px', '80%'],
        content: '/role/auth?id=' + id,
        maxmin: true,
        end: function(){
            layui.table.reload('roleTable');
        }
    });
} 