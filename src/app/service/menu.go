package service

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
	"sort"
)

var Menu = new(MenuService)

type MenuService struct{}

// GetMenuById 根据id获取菜单
func (s *MenuService) GetMenuById(id string) (*model.Menu, error) {
	return dao.Menu.GetMenuById(global.GormDao, id)
}

// GetMenusByPid 根据父ID获取子菜单列表
func (s *MenuService) GetMenusByPid(pid string) ([]*model.Menu, error) {
	return dao.Menu.GetMenusByPid(global.GormDao, pid)
}

// Page 获取菜单分页列表
func (s *MenuService) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	return dao.Menu.Page(pageNum, pageSize, query)
}

// GetMenuTree 获取菜单树结构
func (s *MenuService) GetMenuTree(menuType int) ([]*model.Menu, error) {
	// 先获取所有根菜单，并按类型过滤
	rootMenus, err := dao.Menu.GetMenusByPidAndType(global.GormDao, "-1", menuType)
	if err != nil {
		return nil, err
	}

	// 递归获取子菜单
	for _, menu := range rootMenus {
		if err := s.buildMenuTreeWithType(menu, menuType); err != nil {
			return nil, err
		}
	}
	return rootMenus, nil
}

// buildMenuTreeWithType 递归构建指定类型的菜单树
func (s *MenuService) buildMenuTreeWithType(menu *model.Menu, menuType int) error {
	children, err := dao.Menu.GetMenusByPidAndType(global.GormDao, menu.Id, menuType)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		menu.Children = children
		for _, child := range children {
			if err := s.buildMenuTreeWithType(child, menuType); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetMenusByRoleId 根据角色ID获取菜单列表
func (s *MenuService) GetMenusByRoleId(roleId int) ([]*model.Menu, error) {
	return dao.RoleMenu.GetMenusByRoleId(global.GormDao, roleId)
}

// GetMenuIdsByRoleId 根据角色ID获取菜单ID列表
func (s *MenuService) GetMenuIdsByRoleId(roleId int) ([]string, error) {
	return dao.RoleMenu.GetMenuIdsByRoleId(global.GormDao, roleId)
}

// SaveRoleMenus 保存角色菜单关联
func (s *MenuService) SaveRoleMenus(roleId int, menuIds []string) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		return dao.RoleMenu.BatchSave(tx, roleId, menuIds)
	})
}

// GetByUserId 根据用户ID获取菜单列表
func (s *MenuService) GetByUserId(id int, menuType string) ([]*model.Menu, error) {
	// 1. 先获取用户的所有角色
	roles, err := Role.GetByUserId(id)
	if err != nil {
		return nil, err
	}

	// 2. 检查是否包含管理员角色
	isAdmin := false
	for _, role := range roles {
		if role.Code == model.RoleAdmin {
			isAdmin = true
			break
		}
	}

	var allMenus []*model.Menu
	if isAdmin {
		// 如果是管理员，获取所有启用的菜单
		allMenus, err = dao.Menu.GetAllEnabledMenusByType(global.GormDao, menuType)
	} else {
		// 如果不是管理员，获取角色对应的菜单
		menuMap := make(map[string]*model.Menu)
		for _, role := range roles {
			menus, err := dao.Menu.GetMenusByRoleIdAndType(global.GormDao, role.Id, menuType)
			if err != nil {
				return nil, err
			}
			for _, menu := range menus {
				menuMap[menu.Id] = menu
			}
		}
		// 将map转换为切片
		for _, menu := range menuMap {
			allMenus = append(allMenus, menu)
		}
	}

	if err != nil {
		return nil, err
	}

	// 3. 构建菜单树
	menuMap := make(map[string]*model.Menu)
	var rootMenus []*model.Menu

	// 先将所有菜单放入map中
	for _, menu := range allMenus {
		menuMap[menu.Id] = menu
	}

	// 构建树形结构
	for _, menu := range allMenus {
		if menu.Pid == "-1" {
			// 这是根菜单
			rootMenus = append(rootMenus, menu)
		} else {
			// 找到父菜单
			if parent, exists := menuMap[menu.Pid]; exists {
				if parent.Children == nil {
					parent.Children = make([]*model.Menu, 0)
				}
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	// 对所有层级的菜单进行排序
	for _, menu := range menuMap {
		if menu.Children != nil && len(menu.Children) > 0 {
			s.sortMenus(menu.Children)
		}
	}
	s.sortMenus(rootMenus)

	return rootMenus, nil
}

// sortMenus 对菜单列表进行排序
func (s *MenuService) sortMenus(menus []*model.Menu) {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Seq < menus[j].Seq
	})
}
