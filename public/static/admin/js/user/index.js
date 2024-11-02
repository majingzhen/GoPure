// 表格工具栏模板
const toolbarTpl = `
    <div class="layui-btn-container">
        <button class="layui-btn layui-btn-sm" lay-event="edit">
            <i class="layui-icon layui-icon-edit"></i>
        </button>
        <button class="layui-btn layui-btn-warm layui-btn-sm" lay-event="resetPwd">
            <i class="layui-icon layui-icon-key"></i>
        </button>
        <button class="layui-btn layui-btn-danger layui-btn-sm" lay-event="del">
            <i class="layui-icon layui-icon-delete"></i>
        </button>
    </div>
`;

// 重置密码弹窗模板
const resetPwdTpl = `
    <form class="layui-form" style="padding: 20px;" lay-filter="resetPwdForm">
        <div class="layui-form-item">
            <label class="layui-form-label">新密码</label>
            <div class="layui-input-block">
                <input type="password" name="password" required lay-verify="required" 
                       placeholder="请输入新密码" autocomplete="off" class="layui-input">
            </div>
        </div>
        <div class="layui-form-item">
            <label class="layui-form-label">确认密码</label>
            <div class="layui-input-block">
                <input type="password" name="confirmPassword" required lay-verify="required|confirmPwd" 
                       placeholder="请再次输入新密码" autocomplete="off" class="layui-input">
            </div>
        </div>
    </form>
`;

// 状态列模板
const statusTpl = `
    <input type="checkbox" name="status" value={{d.id}} lay-skin="switch" lay-text="正常|禁用" 
           lay-filter="statusSwitch" {{d.status === '0' ? 'checked' : ''}}>
`;

// 性别列模板
const sexTpl = `
    {{#  if(d.sex === '1'){ }}
        <span>男</span>
    {{#  } else if(d.sex === '2'){ }}
        <span>女</span>
    {{#  } else { }}
        <span>未知</span>
    {{#  } }}
`;

layui.use(['table', 'form', 'layer'], function(){
    var table = layui.table;
    var form = layui.form;
    var layer = layui.layer;
    
    // 渲染表格
    table.render({
        elem: '#userTable',
        url: '/user/page',
        method: 'get',
        toolbar: '#toolbarTpl',
        defaultToolbar: ['filter', 'exports', 'print'],
        cols: [[
            {type: 'checkbox', fixed: 'left'},
            {field: 'id', title: 'ID', sort: true, width: 80},
            {field: 'account', title: '账号'},
            {field: 'userName', title: '用户名'},
            {field: 'sex', title: '性别', templet: sexTpl, width: 80},
            {field: 'mobile', title: '手机号'},
            {field: 'email', title: '邮箱'},
            {field: 'status', title: '状态', templet: statusTpl, width: 100},
            {field: 'createTime', title: '创建时间', sort: true},
            {title: '操作', toolbar: toolbarTpl, width: 180}
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
        
        table.reload('userTable', {
            where: data.field,
            page: {curr: 1}
        });
        return false;
    });

    // 头部工具栏事件
    table.on('toolbar(userTable)', function(obj){
        switch(obj.event){
            case 'add':
                openUserAddForm();
                break;
            case 'batchDel':
                var checkStatus = table.checkStatus('userTable');
                var data = checkStatus.data;
                if(data.length === 0){
                    layer.msg('请选择要删除的数据');
                    return;
                }
                var ids = data.map(item => parseInt(item.id));
                layer.confirm('确定删除选中的用户吗？', function(index){
                    deleteUsers(ids);
                    layer.close(index);
                });
                break;
        }
    });

    // 表格行工具事件
    table.on('tool(userTable)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){
            layer.confirm('确定删除该用户吗？', function(index){
                deleteUsers([parseInt(data.id)]);
                layer.close(index);
            });
        } else if(obj.event === 'edit'){
            openUserEditForm(parseInt(data.id));
        } else if(obj.event === 'resetPwd'){
            resetPassword(parseInt(data.id));
        }
    });

    // 状态切换事件
    form.on('switch(statusSwitch)', function(obj){
        var userId = parseInt(this.value);
        var status = obj.elem.checked ? "0" : "1";
        updateUserStatus(userId, status);
    });
});

// 打开用户添加表单
function openUserAddForm() {
    layer.open({
        type: 2,
        title: '新增用户',
        area: ['600px', '80%'],
        content: '/user/add',
        maxmin: true,
        end: function(){
            layui.table.reload('userTable');
        }
    });
}

// 打开用户编辑表单
function openUserEditForm(id) {
    layer.open({
        type: 2,
        title: '编辑用户',
        area: ['600px', '80%'],
        content: '/user/edit?id=' + id,
        maxmin: true,
        end: function(){
            layui.table.reload('userTable');
        }
    });
}

// 删除用户
function deleteUsers(ids) {
    request.post('/user/delete', {ids: ids})
        .then(() => {
            layer.msg('删除成功');
            layui.table.reload('userTable');
        });
}

// 更新用户状态
function updateUserStatus(id, status) {
    request.post('/user/updateStatus', {id: id, status: status})
        .then(() => {
            layer.msg('更新成功');
        })
        .catch(() => {
            layui.table.reload('userTable');
        });
}

// 重置密码
function resetPassword(userId) {
    layer.open({
        type: 1,
        title: '重置密码',
        area: ['500px', '300px'],
        content: resetPwdTpl,
        btn: ['确定', '取消'],
        success: function(layero, index) {
            layui.form.render(null, 'resetPwdForm');
            // 添加确认密码验证
            layui.form.verify({
                confirmPwd: function(value) {
                    var pwd = layero.find('input[name=password]').val();
                    if(pwd !== value) {
                        return '两次输入的密码不一致';
                    }
                }
            });
        },
        yes: function(index, layero) {
            var pwd = layero.find('input[name=password]').val();
            if(!pwd) {
                layer.msg('请输入新密码');
                return;
            }
            // 提交重置密码请求
            request.post('/user/resetPassword', {
                id: userId,
                password: pwd
            }).then(() => {
                layer.msg('密码重置成功');
                layer.close(index);
            });
        }
    });
}

