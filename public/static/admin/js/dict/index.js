// 表格工具栏模板
const toolbarTpl = `
    <div class="layui-btn-container">
        <button class="layui-btn layui-btn-sm" lay-event="edit">
            <i class="layui-icon layui-icon-edit"></i>
        </button>
        <button class="layui-btn layui-btn-danger layui-btn-sm" lay-event="del">
            <i class="layui-icon layui-icon-delete"></i>
        </button>
    </div>
`;

// 状态列模板
const statusTpl = `
    <input type="checkbox" name="status" value={{d.id}} lay-skin="switch" lay-text="正常|禁用" 
           lay-filter="statusSwitch" {{d.status === '0' ? 'checked' : ''}}>
`;

var layer,table,form,$;
layui.use(['table', 'form', 'layer'], function() {
    table = layui.table;
    form = layui.form;
    layer = layui.layer;
    $ = layui.jquery;
    initPage();
});

function initPage() {
    loadTable();
    
    // 搜索表单提交
    form.on('submit(searchForm)', function(data){
        console.log('Search form submitted:', data.field);
        
        if(data.field.status === '') {
            delete data.field.status;
        }

        table.reload('dictTable', {
            where: data.field,
            page: {curr: 1}
        });
        return false;
    });

    // 重置按钮点击事件 
    $('button[type="reset"]').click(function() {
        console.log('Reset button clicked');
        
        // 重置表单
        $('input[name="dictName"]').val('');
        $('input[name="dictType"]').val('');
        $('select[name="status"]').val('');
        form.render('select');
        
        // 重新加载表格数据
        table.reload('dictTable', {
            where: {},
            page: {curr: 1}
        });
    });
    // 新增点击事件
    $('button[type="add"]').click(function() {
        openDictAddForm();
    });

    // 监听状态开关切换
    form.on('switch(statusSwitch)', function(data) {
        var id = parseInt(data.value)
        var status = data.elem.checked ? '0' : '1';
        request.post('/dict/editStatus', {id: id, status: status})
            .then(res => {
                if(res.code === 0) {
                    layer.msg('更新成功');
                    table.reload('dictTable');
                } else {
                    layer.msg(res.msg);
                    return false;
                }
            });
    });

    // 监听工具条点击事件
    table.on('tool(dictTable)', function(obj) {
        var data = obj.data;
        if(obj.event === 'del') {
            layer.confirm('确定删除该字典吗？', function(index) {
                deleteDict(parseInt(data.id));
                layer.close(index);
            });
        } else if(obj.event === 'edit') {
            openDictEditForm(parseInt(data.id));
        }
    });
}

function loadTable() {
    console.log('Loading table...');
    
    table.render({
        elem: '#dictTable',
        url: '/dict/page',
        method: 'get',
        toolbar: '#toolbar',
        defaultToolbar: ['filter', 'exports', 'print', {
            title: '提示',
            layEvent: 'LAYTABLE_TIPS',
            icon: 'layui-icon-tips'
        }],
        cols: [[
            {field: 'id', title: 'ID', sort: true, width: 80},
            {field: 'dictName', title: '字典名称'},
            {
                field: 'dictType',
                title: '字典类型',
                templet: function(d) {
                    return '<a href="javascript:;" lay-event="dictData" class="layui-table-link">' + d.dictType + '</a>';
                }
            },
            {
                field: 'status',
                title: '状态',
                templet: statusTpl,
                width: 100
            },
            {
                field: 'createTime',
                title: '创建时间',
                width: 180,
                sort: true,
                templet: "<span>{{d.createTime ==null?'':layui.util.toDateString(d.createTime, 'yyyy-MM-dd HH:mm:ss')}}</span>"
            },
            {
                field: 'updateTime',
                title: '更新时间',
                width: 180,
                sort: true,
                templet: "<span>{{d.updateTime ==null?'':layui.util.toDateString(d.updateTime, 'yyyy-MM-dd HH:mm:ss')}}</span>"
            },
            {title: '操作', toolbar: toolbarTpl, width: 180}
        ]],
        page: true,
        request: {
            pageName: 'pageNum',
            limitName: 'pageSize'
        },
        parseData: function (res) {
            console.log('Parsing response:', res);
            return {
                "code": res.code,
                "msg": res.msg,
                "count": res.data.total,
                "data": res.data.rows
            }
        },
        done: function(res) {
            console.log('Table render complete:', res);
        }
    });
}

// 新增字典
function openDictAddForm() {
    layer.open({
        type: 2,
        title: '新增字典',
        area: common.layerArea($("html")[0].clientWidth, 500, 400),
        shadeClose: true,
        anim: 1,
        content: '/dict/add'
    });
}

// 编辑字典
function openDictEditForm(id) {
    layer.open({
        type: 2,
        title: '编辑字典',
        area: common.layerArea($("html")[0].clientWidth, 500, 400),
        shadeClose: true,
        anim: 1,
        content: '/dict/edit?id=' + id
    });
}

function deleteDict(id) {
    console.log('Deleting dict:', id);
    request.delete("/dict/delete/" + id, null, ).then(res => {
        if(res.code === 0) {
            layer.msg('删除成功');
            table.reload('dictTable');
        } else {
            layer.msg(res.msg);
            return false;
        }
    });
}
