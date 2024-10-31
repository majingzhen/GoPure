package service

import (
	"errors"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/utils"
)

var Menu = new(MenuService)

type MenuService struct{}

// List 获取菜单列表
func (s *MenuService) List(req view.MenuListReqVO) ([]model.Menu, error) {
	return dao.Menu.List(req)
}

// GetById 根据ID获取菜单
func (s *MenuService) GetById(id string) (*model.Menu, error) {
	return dao.Menu.GetById(id)
}

// Add 添加菜单
func (s *MenuService) Add(menu *model.Menu) error {
	// 如果是子菜单，检查父菜单是否存在
	if menu.Pid != "-1" {
		parent, err := dao.Menu.GetById(menu.Pid)
		if err != nil {
			return err
		}
		if parent == nil {
			return errors.New("父菜单不存在")
		}
	}

	// 生成ID
	menu.Id = utils.GenUID()
	return dao.Menu.Add(menu)
}

// Update 更新菜单
func (s *MenuService) Update(menu *model.Menu) error {
	// 如果是子菜单，检查父菜单是否存在
	if menu.Pid != "-1" {
		parent, err := dao.Menu.GetById(menu.Pid)
		if err != nil {
			return err
		}
		if parent == nil {
			return errors.New("父菜单不存在")
		}
	}

	return dao.Menu.Update(menu)
}

// Delete 删除菜单
func (s *MenuService) Delete(id string) error {
	// 检查是否有子菜单
	hasChildren, err := dao.Menu.HasChildren(id)
	if err != nil {
		return err
	}
	if hasChildren {
		return errors.New("存在子菜单,不允许删除")
	}

	return dao.Menu.Delete(id)
}

// BuildTree 构建菜单树
func (s *MenuService) BuildTree(menus []model.Menu) []*view.MenuVO {
	// 构建id到菜单的映射
	menuMap := make(map[string]*view.MenuVO)
	for _, m := range menus {
		menu := m // 创建副本
		menuMap[menu.Id] = &view.MenuVO{
			Menu:     &menu,
			Children: make([]*view.MenuVO, 0),
		}
	}

	// 构建树形结构
	tree := make([]*view.MenuVO, 0)
	for _, menuVO := range menuMap {
		if menuVO.Pid == "-1" {
			// 根节点
			tree = append(tree, menuVO)
		} else {
			// 子节点
			if parent, ok := menuMap[menuVO.Pid]; ok {
				parent.Children = append(parent.Children, menuVO)
			}
		}
	}

	return tree
}

// GetRoleMenuTree 获取角色菜单树
func (s *MenuService) GetRoleMenuTree(roleId int) ([]*view.MenuVO, error) {
	menus, err := dao.Menu.GetByRoleId(roleId)
	if err != nil {
		return nil, err
	}
	return s.BuildTree(menus), nil
}

// GetByUserId 获取用户菜单
func (s *MenuService) GetByUserId(userId int, backend string) ([]*view.MenuVO, error) {
	// 先获取用户角色
	roles, err := dao.Role.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	// 检查是否是管理员
	isAdmin := false
	for _, role := range roles {
		if role.Code == model.RoleAdmin {
			isAdmin = true
			break
		}
	}

	var menus []model.Menu
	if isAdmin {
		// 管理员获取所有菜单
		menuVO := view.MenuListReqVO{}
		menuVO.MenuPosition = backend
		menus, err = dao.Menu.List(menuVO)
	} else {
		// 非管理员获取角色对应的菜单
		roleIds := make([]int, len(roles))
		for i, role := range roles {
			roleIds[i] = role.Id
		}
		menus, err = dao.Menu.GetByRoleIds(roleIds)
	}
	if err != nil {
		return nil, err
	}
	// 构建菜单树
	return s.BuildTree(menus), nil
}
