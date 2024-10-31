package service

import (
	"errors"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
)

var Role = new(RoleService)

type RoleService struct{}

// Page 获取角色分页
func (s *RoleService) Page(req view.RoleReqPageVO) (*common.PageInfo, error) {
	return dao.Role.Page(req)
}

// List 获取角色列表
func (s *RoleService) List() ([]model.Role, error) {
	return dao.Role.List()
}

// GetById 根据ID获取角色
func (s *RoleService) GetById(id int) (*model.Role, error) {
	return dao.Role.GetById(id)
}

// Add 添加角色
func (s *RoleService) Add(role *model.Role) error {
	// 检查角色名称是否存在
	exist, err := dao.Role.CheckNameExist(role.Name)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("角色名称已存在")
	}

	// 检查角色编码是否存在
	exist, err = dao.Role.CheckCodeExist(role.Code)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("角色编码已存在")
	}

	// 添加角色
	return dao.Role.Add(role)
}

// Update 更新角色
func (s *RoleService) Update(role *model.Role) error {
	// 检查角色名称是否存在
	exist, err := dao.Role.CheckNameExist(role.Name, role.Id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("角色名称已存在")
	}

	// 检查角色编码是否存在
	exist, err = dao.Role.CheckCodeExist(role.Code, role.Id)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("角色编码已存在")
	}

	// 不允许修改管理员角色
	oldRole, err := dao.Role.GetById(role.Id)
	if err != nil {
		return err
	}
	if oldRole.Code == model.RoleAdmin {
		return errors.New("不允许修改管理员角色")
	}

	// 更新角色
	return dao.Role.Update(role)
}

// Delete 删除角色
func (s *RoleService) Delete(ids []int) error {
	// 检查是否包含管理员角色
	for _, id := range ids {
		role, err := dao.Role.GetById(id)
		if err != nil {
			return err
		}
		if role.Code == model.RoleAdmin {
			return errors.New("不允许删除管理员角色")
		}
	}

	// 删除角色
	return dao.Role.Delete(ids)
}

// UpdateStatus 更新角色状态
func (s *RoleService) UpdateStatus(id int, status string) error {
	// 不允许禁用管理员角色
	role, err := dao.Role.GetById(id)
	if err != nil {
		return err
	}
	if role.Code == model.RoleAdmin && status == "1" {
		return errors.New("不允许禁用管理员角色")
	}

	return dao.Role.UpdateStatus(id, status)
}

// GetByCode 根据角色编码获取角色
func (s *RoleService) GetByCode(code string) (*model.Role, error) {
	return dao.Role.GetByCode(code)
}

func (s *RoleService) GetByUserId(id int) ([]*model.Role, error) {
	return dao.Role.GetByUserId(id)
}

// GetRoleMenus 获取角色菜单ID列表
func (s *RoleService) GetRoleMenus(roleId int) ([]int, error) {
	return dao.Role.GetRoleMenus(roleId)
}

// AuthRole 角色授权
func (s *RoleService) AuthRole(roleId int, menuIds []int) error {
	return dao.Role.AuthRole(roleId, menuIds)
}
