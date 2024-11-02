var table, treeTable, layPage, form,$;
layui.config({
    base: '/static/module/',
    debug: true
}).extend({
    tableTree: 'tableTree/tableTree',
    tableEdit: 'tableTree/tableEdit'
}).use(['table', 'form', 'layer', 'tableTree'], function(){
    table = layui.table;
    treeTable = layui.treeTable;
    layPage = layui.laypage;
    form = layui.form;
    $ = layui.jquery;
    initPage();
});
/**
 * 页面初始化事件
 */
function initPage() {
    queryTable();

    // 查询
    $("#queryBtn").click(function () {
        queryTable();
    });

    // 添加
    $("#addBtn").click(function () {
        add();
    });
}

/**
 * 查询表格数据
 */
function queryTable() {
    treeTable.render({
        elem: '#menuTable',
        url: '/menu/list',
        id: 'menuTable',
        treeColIndex: 1,
        treeIdName: 'id',
        treePidName: 'pid',
        treeDefaultClose: false,	//是否默认折叠
        treeLinkage: false,		//父级展开时是否自动展开所有子级
        method: 'get',
        title: '菜单列表',
        totalRow: false,
        where: {
            form: {
                name: $("#name").val()
            }
        },
        parseData: function (res) {
            return {
                "code": res.code,
                "msg": res.msg,
                "count": res.total,
                "data": res.data
            };
        },
        cols: [[
            {field: 'id', title: 'ID', width: 80},
            {field: 'name', minWidth: 160,title: '菜单名'},
            {field: 'url', minWidth: 200,title: '菜单链接'},
            {
                field: 'icon',
                align: 'center',
                title: '图标',
                minWidth: 80,
                templet: "<div><i class='fa {{d.icon}}' aria-hidden='true'></i></div>"
            },
            {field: 'target', minWidth: 120,title: '打开方式', templet: "<div>{{d.target === 0 ? '本页' : '新窗口'}}</div>"},
            {
                field: 'status', minWidth: 100,title: '状态', templet: function (d) {
                    let html;
                    if (d.status === '0') {
                        html = "<input type='checkbox' name='status' lay-filter='status' lay-skin='switch' value='" + d.id + "' lay-text='正常|禁用' checked>";
                    } else {
                        html = "<input type='checkbox' name='status' lay-filter='status' value='" + d.id + "' lay-skin='switch' lay-text='正常|禁用'>";
                    }
                    return html;
                }
            },
            {field: 'seq', minWidth: 70,title: '排序'},
            {
                field: 'createTime',
                title: '创建时间',
                minWidth: 150,
                templet: "<span>{{d.createTime ==null?'':layui.util.toDateString(d.createTime, 'yyyy-MM-dd HH:mm:ss')}}</span>"
            },
            {
                field: 'updateTime',
                title: '更新时间',
                minWidth: 150,
                templet: "<span>{{d.updateTime == null?'':layui.util.toDateString(d.updateTime, 'yyyy-MM-dd HH:mm:ss')}}</span>"
            },
            {
                field: 'id', title: '操作', width: 220,
                templet: function (d) {
                    let html = "<div>"
                    html += "<a class='pear-btn pear-btn-xs pear-btn-primary' onclick='add(\"" + d.id + "\")'>添加子菜单</a> " +
                        "<a class='pear-btn pear-btn-xs pear-btn-primary' onclick='openMenuEditForm(\"" + d.id + "\")'>编辑</a> " +
                        "<a class='pear-btn pear-btn-xs pear-btn-danger' onclick='deleteMenus(\"" + d.id + "\")'>删除</a>" +
                        "</div>";
                    return html;
                }
            },
        ]],
        page: false,
        done: function () {
            layer.closeAll('loading');
        }
    });

    form.on('switch(status)', function (data) {
        const id = this.value;
        const status = this.checked ? '0' : '1';
        updateMenuStatus(id, status);
    });
}

/**
 * 添加
 */
function add(pid = -1) {
    let title = "添加一级菜单";
    if (pid !== '-1') {
        title = "添加子菜单";
    }
    layer.open({
        title: title,
        type: 2,
        area: common.layerArea($("html")[0].clientWidth, 500, 400),
        shadeClose: true,
        anim: 1,
        content: '/menu/add?pid=' + pid
    });
}

// 打开菜单编辑表单
function openMenuEditForm(id) {
    layer.open({
        type: 2,
        title: '编辑菜单',
        type: 2,
        area: common.layerArea($("html")[0].clientWidth, 500, 400),
        shadeClose: true,
        anim: 1,
        content: '/menu/edit?id=' + id
    });
}

// 删除菜单
function deleteMenus(id) {
    request.post('/menu/delete', {id: id})
        .then(() => {
            layer.msg('删除成功');
            queryTable();
        });
}

// 更新菜单状态
function updateMenuStatus(id, status) {
    request.post('/menu/updateStatus', {id: id, status: status})
        .then(() => {
            layer.msg('更新成功');
        })
        .catch(() => {
            layui.tableTree.reload('menuTable');
        });
}