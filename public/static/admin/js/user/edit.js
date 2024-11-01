
var layer, form,$, editRoleSelect;
layui.use(['form', 'layer'], function(){
    form = layui.form;
    layer = layui.layer;
    userId = getUrlParam('id');
    // 初始化表单
    initForm();
    // 监听提交
    form.on('submit(userEditForm)', function(data){
        var formData = data.field;
        // 获取选中的角色ID
        formData.roleIds = editRoleSelect.getValue().map(obj => obj.id);
        // 提交表单
        request.post('/user/update', formData)
            .then(() => {
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
            getUrlParam();
            // 加载性别选项
            await loadSexOptions();
            // 加载角色列表和用户数据
            initRolesAndUserData();
        } catch (error) {
            console.error('初始化表单失败:', error);
            layer.msg('初始化表单失败');
        }
    }
    

});

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

// 加载角色列表和用户数据
function initRolesAndUserData() {
    // 再获取用户数据
    const userRes = request.get('/user/get', { id: userId });
    var options = {
        el: '#editRoleSelect',
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
    editRoleSelect = xmSelect.render(options);
    request.get('/role/list').then(res => {
        editRoleSelect.update({
            data: res.data,
            initValue: userRes.data.roleIds
        })
    })

    // 填充表单数据
    form.val('userEditForm', userRes.data);
}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
}