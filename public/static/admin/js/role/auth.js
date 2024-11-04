
var form,layer,roleId,authSelect,menuData;
layui.use(['form', 'layer'], function(){
    form = layui.form;
    layer = layui.layer;
    roleId = getUrlParam('id');
    // 初始化表单
    initForm();
    // 监听提交
    form.on('submit(roleAuthForm)', function(data){
        // 获取选中的菜单ID
        var menuIds = authSelect.getValue().map(obj => obj.id);
        // 提交表单
        request.post('/role/authRole', {
            id: parseInt(roleId),
            menuIds: menuIds
        }).then(() => {
            layer.msg('保存成功');
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
        // 加载角色信息和菜单树
        loadRoleAndMenus();
    } catch (error) {
        console.error('初始化表单失败:', error);
        layer.msg('初始化表单失败');
    }
}

// 取消
function cancel() {
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
}

// 解决xmSelect tree 严格遵守父子模式时获取不到父节点的值
// 1. 首先定义一个函数来转换数据结构
// 1. 修正后的数据转换函数
function transformTreeData(data) {
    return data.map(item => {
        // 创建新的节点对象，保留原有属性
        let newItem = {...item};

        // 只有包含子节点的节点才需要添加"仅选择"选项
        if (item.children && item.children.length) {
            // 保留原有的子节点
            let originalChildren = [...item.children];
            console.log(item.name)
            // 添加"仅选择"选项作为第一个子节点
            newItem.children = [
                {
                    id: item.id + '_self',
                    name: item.name,
                    value: item.id,
                    isSelf: true
                },
                // 递归处理原有的子节点
                ...originalChildren.map(child => transformTreeData([child])[0])
            ];
            console.log(newItem)
        }

        return newItem;
    });
}

// 加载角色信息和菜单树
function loadRoleAndMenus() {
    request.get('/menu/list').then(res => {
        menuData = transformTreeData(res.data);
        // 获取角色信息
        request.get('/role/get', { id: roleId }).then(res => {
            // 填充表单数据
            form.val('roleAuthForm', res.data);
            authSelect = xmSelect.render({
                el: "#menuTree",
                name: "menuTree",
                initValue: res.data.menuIds,
                data: menuData,
                tree: {
                    show:true,
                    expandedKeys:[-1],
                    showLine: true,
                    indent: 20,
                    strict: true,  // 关闭严格模式
                },
                prop: {
                    value: 'id',
                    name: 'name'
                },
                toolbar: {show:true,list:["ALL","CLEAR","REVERSE"]},
                // 修正后的模板函数
                // 修正后的模板配置
                template({ item, sels, name, value }) {
                    // 如果是父节点的自身选择节点
                    if (item.isSelf) {
                        return item.name + ' (仅选择)';
                    }
                    return item.name;
                },
                // 处理选中值的回调
                on: function(data) {
                    if (data.isAdd && data.change && data.change.length > 0) {
                        let changed = data.change[0];
                        // 处理选中值，去掉 _self 后缀
                        if (changed.id && changed.id.toString().includes('_self')) {
                            changed.id = changed.id.toString().replace('_self', '');
                        }
                    }
                    return true;
                }
            });
        })
    })
}

// 获取URL参数
function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
}