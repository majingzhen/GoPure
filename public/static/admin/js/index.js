
var element;
// 使用新的layui模块加载方式
layui.config({
    base: '/static/layui/'
}).use(['element', 'layer', 'util'], function(){
    element = layui.element;
    var layer = layui.layer;
    var util = layui.util;
    var $ = layui.jquery;

    // 初始化完成后再获取用户信息
    getLoginUser();

    //头部事件
    util.event('lay-header-event', {
        menuLeft: function(othis){
            layer.msg('展开左侧菜单的操作', {icon: 0});
        }
    });
});

// 获取登录用户信息
function getLoginUser() {
    request.get('/getLoginUser')
        .then(res => {
            $('#userName').text(res.data.userName);
            renderMenuTree(res.data.menus);
        });
}

// 渲染菜单树
function renderMenuTree(menus) {
    var menuHtml = '';
    if (menus && menus.length > 0) {
        menuHtml = buildMenuHtml(menus);
    }
    $('#menuTree').empty().html(menuHtml);
    // 重新渲染导航菜单
    element.render('nav', 'menuTree');
    // 绑定菜单点击事件
    bindMenuClick();
}

// 递归构建菜单HTML
function buildMenuHtml(menus) {
    var html = '';
    menus.forEach(function(menu) {
        // 跳过按钮类型的菜单
        if (menu.menuType === '2') return;

        // 添加菜单项，不再默认添加展开类
        html += '<li class="layui-nav-item">';

        if (menu.menuType === '0') {
            // 目录类型，可以包含子菜单
            html += '<a href="javascript:;" class="menu-dir">' +
                (menu.icon ? '<i class="layui-icon ' + menu.icon + '"></i> ' : '') +
                menu.name + '<span class="layui-nav-more"></span></a>';

            // 递归构建子菜单
            if (menu.children && menu.children.length > 0) {
                html += '<dl class="layui-nav-child">';
                menu.children.forEach(function(child) {
                    // 只有非按钮类型的子菜单才会被构建
                    if (child.menuType !== '2') {
                        html += buildSubMenuHtml(child);
                    }
                });
                html += '</dl>';
            }
        } else {
            // 菜单类型，添加跳转链接
            html += buildMenuLink(menu);
        }

        html += '</li>';
    });
    return html;
}

// 递归构建子菜单HTML
function buildSubMenuHtml(menu) {
    // 跳过按钮类型的菜单
    if (menu.menuType === '2') return '';

    var html = '';
    if (menu.menuType === '0') {
        // 如果是目录类型，创建子菜单组
        html += '<dd>';
        html += '<a href="javascript:;" class="menu-dir">' +
            (menu.icon ? '<i class="layui-icon ' + menu.icon + '"></i> ' : '') +
            menu.name + '<span class="layui-nav-more"></span></a>';

        if (menu.children && menu.children.length > 0) {
            html += '<dl class="layui-nav-child">';
            menu.children.forEach(function(child) {
                // 只有非按钮类型的子菜单才会被构建
                if (child.menuType !== '2') {
                    html += buildSubMenuHtml(child);
                }
            });
            html += '</dl>';
        }
        html += '</dd>';
    } else {
        // 菜单类型，添加跳转链接
        html += '<dd>' + buildMenuLink(menu) + '</dd>';
    }
    return html;
}

// 构建菜单链接
function buildMenuLink(menu) {
    if (menu.target === '1') {
        // 新窗口打开
        return '<a href="' + menu.url + '" target="_blank">' +
            (menu.icon ? '<i class="layui-icon ' + menu.icon + '"></i> ' : '') +
            menu.name + '</a>';
    } else {
        // 当前页面打开
        return '<a href="javascript:;" class="menu-link" data-url="' + menu.url + '" data-title="' + menu.name + '">' +
            (menu.icon ? '<i class="layui-icon ' + menu.icon + '"></i> ' : '') +
            menu.name + '</a>';
    }
}

// 绑定菜单点击事件
function bindMenuClick() {
    // 处理目录菜单的点击
    $('.menu-dir').off('click').on('click', function(e) {
        e.preventDefault();
        e.stopPropagation();
        var $parent = $(this).parent();
        $parent.toggleClass('layui-nav-itemed');
    });

    // 处理内部页面跳转的菜单项点击
    $('.menu-link').off('click').on('click', function(e) {
        e.preventDefault();
        e.stopPropagation();
        var url = $(this).data('url');
        var title = $(this).data('title');
        openTab(url, title);
    });
}

// 打开新标签页
function openTab(url, title) {
    // 如果标签已存在，切换到该标签
    var layId = url.replace(/\//g, '_');
    var exists = $(".layui-tab-title li[lay-id='" + layId + "']").length > 0;
    if (exists) {
        element.tabChange('mainTabs', layId);
        return;
    }

    // 添加新标签页
    element.tabAdd('mainTabs', {
        title: title,
        content: '<iframe src="' + url + '" frameborder="0" class="layui-anim layui-anim-upbit"></iframe>',
        id: layId
    });
    element.tabChange('mainTabs', layId);
}

// 退出登录
function logout() {
    layui.layer.confirm('确定要退出登录吗？', {icon: 3, title: '提示'}, function (index) {
        request.get('/logout')
            .then(() => {
                window.location.href = '/login';
            });
        layui.layer.close(index);
    });
}