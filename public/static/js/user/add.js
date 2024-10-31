layui.use(['form', 'layer'], function(){
    var form = layui.form;
    var layer = layui.layer;
    var addRoleSelect; // 保存xmSelect实例


    // 初始化表单
    initForm();
    
    // 监听提交
    form.on('submit(userAddForm)', function(data){
        var formData = data.field;
        
        // 获取选中的角色ID
        var roleIds = addRoleSelect.getValue().map(item => parseInt(item.value));
        formData.roleIds = roleIds;
        
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
    
    // 添加确认密码验证
    form.verify({
        confirmPwd: function(value) {
            var pwd = $('input[name=password]').val();
            if(pwd !== value) {
                return '两次输入的密码不一致';
            }
        }
    });
    
    // 初始化表单
    async function initForm() {
        try {
            // 加载性别选项
            await loadSexOptions();
            // 加载角色列表
            await loadRoles();
        } catch (error) {
            console.error('初始化表单失败:', error);
            layer.msg('初始化表单失败');
        }
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
    
    // 加载角色列表
    async function loadRoles() {
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
}); 