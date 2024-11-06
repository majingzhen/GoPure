layui.use(['table', 'layer', 'form'], function(){
    var table = layui.table;
    var layer = layui.layer;
    var form = layui.form;
    
    // 初始化表格
    table.render({
        elem: '#optionTable',
        method: 'get',
        url: '/option/page',
        cols: [[
            {field: 'id', title: 'ID', width: 80, sort: true},
            {field: 'key', title: 'Key', width: 150},
            {field: 'value', title: 'Value'},
            {field: 'title', title: '标题', width: 150},
            {field: 'identification', title: '标识', width: 150},
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
            {title: '操作', width: 150, toolbar: '#optionBar', fixed: 'right'}
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
        table.reload('optionTable', {
            where: data.field,
            page: {curr: 1}
        });
        return false;
    });
    
    // 监听工具条
    table.on('tool(optionTable)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){
            layer.confirm('确认删除吗？', function(index){
                request.delete('/option/delete/' + data.id).then(res => {
                    if(res.code === 0){
                        layer.msg('删除成功');
                        obj.del();
                    } else {
                        layer.msg(res.msg);
                    }
                    layer.close(index);
                });
            });
        } else if(obj.event === 'edit'){
            layer.open({
                type: 2,
                title: '编辑选项',
                content: '/option/edit?id=' + data.id,
                area: ['500px', '500px']
            });
        }
    });
});

// 添加选项
function add(){
    layer.open({
        type: 2,
        title: '添加选项',
        content: '/option/add',
        area: ['500px', '500px']
    });
} 