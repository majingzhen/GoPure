var layer, form,$, editRoleSelect ;
layui.use(['form', 'layer', 'upload'], function(){
    form = layui.form;
    layer = layui.layer;
    upload = layui.upload;
    userId = getUrlParam('id');
    // 初始化表单
    initForm();
    // 初始化头像上传
    var uploadInst = upload.render({
        elem: '#avatarUpload',
        url: '/upload',
        accept: 'images',
        acceptMime: 'image/*',
        field: 'file',
        before: function(obj){
            obj.preview(function(index, file, result){
                $('#avatarImage').attr('src', result);
            });
            layer.load();
        },
        done: function(res){
            layer.closeAll('loading');
            if(res.code === 0){
                layer.msg('上传成功');
                // 将返回的图片URL存储到隐藏域
                $("#avatar").val(res.data);
            } else {
                layer.msg(res.msg);
            }
        },
        error: function(){
            layer.closeAll('loading');
            var uploadText = $('#uploadText');
            uploadText.html('<span style="color: #FF5722;">上传失败</span> <a class="layui-btn layui-btn-xs demo-reload">重试</a>');
            uploadText.find('.demo-reload').on('click', function(){
                uploadInst.upload();
            });
        }
    });

    // 监听提交
    form.on('submit(userEditForm)', function(data){
        var formData = data.field;
        formData.id = userId;
        // 获取选中的角色ID
        formData.roleIds = editRoleSelect.getValue().map(obj => obj.id);
        request.post('/user/update', formData).then(res => {
            console.log(res);
            if (res.code === 0) {
                layer.msg('修改成功', {
                    icon: 1,
                    time: 1000
                }, function(){
                    // 关闭当前页面并刷新父页面的表格
                    var index = parent.layer.getFrameIndex(window.name);
                    parent.layer.close(index);
                    parent.layui.table.reload('userTable');
                });
            } else {
                layer.msg(res.msg, {icon: 2});
            }
        });
        return false;
    });
    
    // 初始化表单
    async function initForm() {
        try {
            getUrlParam();
            // 加载性别选项
            loadSexOptions();
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
// 获取用户数据

// 加载角色列表和用户数据
function initRolesAndUserData() {
    // 再获取用户数据
    request.get('/user/get', { id: userId }).then(res => {
        const userData = res.data;
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
                initValue: userData.roleIds
            })
        })

        // 填充表单数据
        form.val('userEditForm', userData);

        // 加载头像
        $('#avatarImage').attr('src', userData.avatar);
        $('#avatar').val(userData.avatar);
    });
}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
}

function cancel() {
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}