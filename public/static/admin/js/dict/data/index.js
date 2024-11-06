
// 状态列模板
const statusTpl = `
    <input type="checkbox" name="status" value={{d.id}} lay-skin="switch" lay-text="正常|禁用" 
           lay-filter="statusSwitch" {{d.status === '0' ? 'checked' : ''}}>
`;
var dictType = getUrlParam('type');
var table, form, layer, $;
layui.use(['table', 'form', 'layer'], function(){
    table = layui.table;
    form = layui.form;
    layer = layui.layer;
    loadTable();
    // 监听搜索表单提交
    form.on('submit(searchForm)', function(data){
        // 将type数据添加到查询条件中
        data.field.dictType = dictType;
        table.reload('dictDataTable', {
            where: data.field,
            page: {curr: 1}
        });
        return false;
    });

    // 重置按钮点击
    $('button[type="reset"]').click(function() {
        $('input[name="dictLabel"]').val('');
        $('select[name="status"]').val('');
        form.render('select');
        loadTable();
    });

    // 新增点击事件
    $('button[type="add"]').click(function() {
        openDictDataAddForm();
    });

    // 关闭点击事件
    $('button[type="close"]').click(function() {
       window.location.href = '/dict';
    });
    
    // 监听工具条
    table.on('tool(dictDataTable)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){
            layer.confirm('确认删除吗？', function(index){
                deleteDictData(data.id);
                layer.close(index);
            });
        } else if(obj.event === 'edit'){
            openDictDataEditForm(data.id);
        }
    });
    
    // 监听状态切换
    form.on('switch(statusSwitch)', function(obj){
        var data = {
            id: parseInt(this.value),
            status: obj.elem.checked ? '0' : '1'
        };
        request.post('/dict/data/editStatus', data).then(res => {
            if(res.code !== 0){
                layer.msg(res.msg);
                $(obj.elem).prop('checked', !obj.elem.checked);
                form.render('checkbox');
            }
        });
    });
});


// 加载表格
function loadTable() {
    // 初始化表格
    table.render({
        elem: '#dictDataTable',
        url: '/dict/data/page',
        where: {dictType: dictType},
        method: 'get',
        page: true,
        cols: [[
            {field: 'id', title: 'ID', sort: true},
            {field: 'dictType', title: '字典类型'},
            {field: 'dictLabel', title: '字典标签'},
            {field: 'dictValue', title: '字典键值'},
            {field: 'dictExtendValue', title: '扩展值'},
            {field: 'seq', title: '排序', sort: true},
            {field: 'status', title: '状态', width: 100, templet: statusTpl},
            {
                field: 'createTime',
                title: '创建时间',
                width: 180,
                sort: true,
                templet: "<span>{{d.createTime ==null?'':layui.util.toDateString(d.createTime, 'yyyy-MM-dd HH:mm:ss')}}</span>"
            },
            {title: '操作', width: 150, toolbar: '#dictDataBar', fixed: 'right'}
        ]],
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

// 删除字典数据
function deleteDictData(id) {
    request.delete('/dict/data/delete/' + id).then(res => {
        if(res.code === 0){
            layer.msg('删除成功');
            table.reload('dictDataTable');
        } else {
            layer.msg(res.msg);
            return false;
        }
    });
}

// 添加字典数据
function openDictDataAddForm() {
    layer.open({
        type: 2,
        title: '新增字典数据',
        area: common.layerArea($("html")[0].clientWidth, 500, 400),
        shadeClose: true,
        anim: 1,
        content: '/dict/data/add?dictType=' + dictType
    });
}
// 编辑字典数据
function openDictDataEditForm(id) {
    layer.open({
        type: 2,
        title: '编辑字典数据',
        area: common.layerArea($("html")[0].clientWidth, 500, 400),
        shadeClose: true,
        anim: 1,
        content: '/dict/data/edit?id=' + id
    });
}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
} 
