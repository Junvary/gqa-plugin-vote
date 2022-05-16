package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginVoteSysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-vote").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_api 表中vote插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_role_api 表中vote插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{
	// plugin-vote组
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-list", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-get", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-add", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-random-delete", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-handle", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/can-vote-dy", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/can-vote-gl", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-voted", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-list", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/vote-result-chart", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-list", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-add", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/voter-delete", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-list", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-add", ApiMethod: "POST"},
	{RoleCode: "super-admin", ApiGroup: "plugin-vote", ApiPath: "/plugin-vote/candidate-delete", ApiMethod: "POST"},
}
