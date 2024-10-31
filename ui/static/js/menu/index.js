layui.extend({
    dtree: '/static/dtree/dtree'
}).use(['table', 'form', 'layer', 'dtree'], function(){
    var table = layui.table;
    var form = layui.form;
    var layer = layui.layer;
    var dtree = layui.dtree;
    
    // 状态列模板
    var statusTpl = `
        <input type="checkbox" name="status" value="{{d.id}}" lay-skin="switch" lay-text="正常|停用" 
               lay-filter="statusSwitch" {{d.status === '0' ? 'checked' : ''}}>
    `;

    // 类型列模板
    var typeTpl = `
        {{#  if(d.menuType === '0'){ }}
            <span class="layui-badge layui-bg-blue">目录</span>
        {{#  } else if(d.menuType === '1'){ }}
            <span class="layui-badge layui-bg-green">菜单</span>
        {{#  } else if(d.menuType === '2'){ }}
            <span class="layui-badge layui-bg-gray">按钮</span>
        {{#  } }}
    `;
    
    // 渲染表格
    table.render({
        elem: '#menuTable',
        url: '/menu/list',
        method: 'get',
        toolbar: '#toolbarTpl',
        defaultToolbar: ['filter', 'exports', 'print'],
        cols: [[
            {type: 'checkbox', fixed: 'left'},
            {field: 'name', title: '菜单名称'},
            {field: 'menuType', title: '类型', templet: typeTpl, width: 80},
            {field: 'icon', title: '图标', templet: '<div><i class="layui-icon {{d.icon}}"></i></div>', width: 80},
            {field: 'url', title: '菜单链接'},
            {field: 'seq', title: '排序', width: 80},
            {field: 'menuPosition', title: '菜单位置', templet: menuPositionTpl, width: 100},
            {field: 'target', title: '打开方式', templet: targetTpl, width: 100},
            {field: 'status', title: '状态', templet: statusTpl, width: 100},
            {title: '操作', toolbar: '#rowToolbarTpl', width: 120}
        ]],
        response: {
            statusName: 'code',
            statusCode: 0,
            msgName: 'msg',
            dataName: 'data'
        }
    });

    // 搜索表单提交
    form.on('submit(searchForm)', function(data){
        if(data.field.status === '') {
            delete data.field.status;
        }
        
        table.reload('menuTable', {
            where: data.field
        });
        return false;
    });

    // 头部工具栏事件
    table.on('toolbar(menuTable)', function(obj){
        switch(obj.event){
            case 'add':
                openMenuAddForm();
                break;
            case 'batchDel':
                var checkStatus = table.checkStatus('menuTable');
                var data = checkStatus.data;
                if(data.length === 0){
                    layer.msg('请选择要删除的数据');
                    return;
                }
                var ids = data.map(item => parseInt(item.id));
                layer.confirm('确定删除选中的菜单吗？', function(index){
                    deleteMenus(ids);
                    layer.close(index);
                });
                break;
        }
    });

    // 表格行工具事件
    table.on('tool(menuTable)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){
            layer.confirm('确定删除该菜单吗？', function(index){
                deleteMenus([parseInt(data.id)]);
                layer.close(index);
            });
        } else if(obj.event === 'edit'){
            openMenuEditForm(parseInt(data.id));
        }
    });

    // 状态切换事件
    form.on('switch(statusSwitch)', function(obj){
        var menuId = parseInt(this.value);
        var status = obj.elem.checked ? "0" : "1";
        updateMenuStatus(menuId, status);
    });
});

// 打开菜单添加表单
function openMenuAddForm() {
    layer.open({
        type: 2,
        title: '新增菜单',
        area: ['600px', '80%'],
        content: '/menu/add',
        maxmin: true,
        end: function(){
            layui.table.reload('menuTable');
        }
    });
}

// 打开菜单编辑表单
function openMenuEditForm(id) {
    layer.open({
        type: 2,
        title: '编辑菜单',
        area: ['600px', '80%'],
        content: '/menu/edit?id=' + id,
        maxmin: true,
        end: function(){
            layui.table.reload('menuTable');
        }
    });
}

// 删除菜单
function deleteMenus(ids) {
    request.post('/menu/delete', {ids: ids})
        .then(() => {
            layer.msg('删除成功');
            layui.table.reload('menuTable');
        });
}

// 更新菜单状态
function updateMenuStatus(id, status) {
    request.post('/menu/updateStatus', {id: id, status: status})
        .then(() => {
            layer.msg('更新成功');
        })
        .catch(() => {
            layui.table.reload('menuTable');
        });
}

// 菜单位置模板
var menuPositionTpl = `
    {{#  if(d.menuPosition === '0'){ }}
        <span class="layui-badge layui-bg-blue">前台</span>
    {{#  } else if(d.menuPosition === '1'){ }}
        <span class="layui-badge layui-bg-green">后台</span>
    {{#  } }}
`;

// 打开方式模板
var targetTpl = `
    {{#  if(d.target === '0'){ }}
        <span>本页</span>
    {{#  } else if(d.target === '1'){ }}
        <span>新窗口</span>
    {{#  } }}
`;

// 修改类型模板
var typeTpl = `
    {{#  if(d.menuType === '0'){ }}
        <span class="layui-badge layui-bg-blue">目录</span>
    {{#  } else if(d.menuType === '1'){ }}
        <span class="layui-badge layui-bg-green">菜单</span>
    {{#  } else if(d.menuType === '2'){ }}
        <span class="layui-badge layui-bg-gray">按钮</span>
    {{#  } }}
`;