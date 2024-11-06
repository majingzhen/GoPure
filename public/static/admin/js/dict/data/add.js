
var dictType = getUrlParam('dictType');
layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    // 监听提交
    form.on('submit(dictDataAddForm)', function(data){
        // 阻止表单默认提交行为，避免页面刷新
        event.preventDefault();
        data.field.dictType = dictType;
        // 将seq字段转换为整数类型
        if(data.field.seq) {
            data.field.seq = parseInt(data.field.seq);
        }
        request.post('/dict/data/add', data.field).then(res => {
            if(res.code === 0){
                layer.msg('添加成功', {
                    icon: 1,
                    time: 1000
                }, function(){
                    // 关闭弹窗并刷新父页面
                    var index = parent.layer.getFrameIndex(window.name);
                    parent.layer.close(index);
                    parent.layui.table.reload('dictDataTable');
                });
            } else {
                layer.msg(res.msg, {icon: 2});
            }
        });
        
        return false;
    });
});

// 取消
function cancel(){
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
}