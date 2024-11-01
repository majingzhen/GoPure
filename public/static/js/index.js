// 菜单管理对话框
function showMenuDialog(title, menu) {
    $('#menuDialog').dialog({
        closed: false,
        modal: true,
        title: title,
        buttons: [{
            text: '保存',
            iconCls: 'icon-ok',
            handler: function() {
                saveMenu();
            }
        }, {
            text: '取消',
            iconCls: 'icon-cancel',
            handler: function() {
                $('#menuDialog').dialog('close');
            }
        }],
        btn: null,
        cancel: function(index, layero){
            layer.close(index);
            return false;
        }
    });
}

// 用户管理对话框
function showUserDialog(title, user) {
    $('#userDialog').dialog({
        closed: false,
        modal: true,
        title: title,
        buttons: [{
            text: '保存',
            iconCls: 'icon-ok',
            handler: function() {
                saveUser();
            }
        }, {
            text: '取消',
            iconCls: 'icon-cancel',
            handler: function() {
                $('#userDialog').dialog('close');
            }
        }],
        btn: null,
        cancel: function(index, layero){
            layer.close(index);
            return false;
        }
    });
}

// 角色管理对话框
function showRoleDialog(title, role) {
    $('#roleDialog').dialog({
        closed: false,
        modal: true,
        title: title,
        buttons: [{
            text: '保存',
            iconCls: 'icon-ok',
            handler: function() {
                saveRole();
            }
        }, {
            text: '取消',
            iconCls: 'icon-cancel',
            handler: function() {
                $('#roleDialog').dialog('close');
            }
        }],
        btn: null,
        cancel: function(index, layero){
            layer.close(index);
            return false;
        }
    });
} 