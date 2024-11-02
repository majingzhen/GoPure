
var layer, form,$, addRoleSelect;
layui.use(['form', 'layer'], function(){
    form = layui.form;
    layer = layui.layer;
    addRoleSelect; // 保存xmSelect实例

    // 初始化表单
    initForm();
    // 监听提交
    form.on('submit(userAddForm)', function(data){
        var formData = data.field;
        // 获取选中的角色ID
        formData.roleIds = addRoleSelect.getValue().map(obj => obj.id);
        console.log(formData)
        // 提交表单
        request.post('/user/add', formData)
            .then(() => {
                layer.msg('添加成功');
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
        // 加载性别选项
        loadSexOptions();
        // 加载角色列表
        loadRoles();

    } catch (error) {
        console.error('初始化表单失败:', error);
        layer.msg('初始化表单失败');
    }
}

// 加载角色列表
function loadRoles() {
    var options = {
        el: '#addRoleSelect',
        name: 'roleIds',
        layVerify: 'required',
        toolbar: {
            show: true,
            list: [
                'ALL',
                'CLEAR',
                'REVERSE'
            ]
        },
        data:[],
        filterable: true,
        autoRow: true,
        prop: {
            value: 'id',
            name: 'name'
        }
    }
    addRoleSelect = xmSelect.render(options);
    request.get('/role/list').then(res => {
        addRoleSelect.update({
            data: res.data
        })
    })
}


// 加载性别选项
function loadSexOptions() {
    return request.get('/dict/data/list', { dictType: 'sys_user_sex' })
        .then(res => {
            var html = '';
            res.data.forEach(function(dict) {
                html += `<input type="radio" name="sex" value="${dict.dictValue}" 
                                  title="${dict.dictLabel}">`;
            });
            $('#sexRadios').html(html);
            form.render('radio');
        });
}

function cancel() {
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}