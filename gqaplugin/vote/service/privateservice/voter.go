package privateservice

import (
	"errors"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"gorm.io/gorm"
)

func VoterList(getVoterList model.RequestVoterList, username string) (err error, voter []model.GqaPluginVoteVoter, total int64) {
	pageSize := getVoterList.PageSize
	offset := getVoterList.PageSize * (getVoterList.Page - 1)
	var voterList []model.GqaPluginVoteVoter
	var db *gorm.DB
	db = gqaGlobal.GqaDb.Model(&model.GqaPluginVoteVoter{})
	//配置搜索
	if getVoterList.Voter != "" {
		db = db.Where("title like ?", "%"+getVoterList.Voter+"%")
	}
	if getVoterList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoterList.VoteType+"%")
	}
	if getVoterList.VoteRatio != "" {
		db = db.Where("vote_ratio like ?", "%"+getVoterList.VoteRatio+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn("voter", false)).Preload("VoterByUser").Find(&voterList).Error
	return err, voterList, total
}

func VoterAdd(toAddVoter *model.RequestAddVoter, username string) (err error) {
	var addVoter []model.GqaPluginVoteVoter
	var voterList []model.GqaPluginVoteVoter
	//if err = gqaGlobal.GqaDb.
	//	Where("vote_type = ? and vote_ratio != ? and vote_ratio != ?", toAddVoter.VoteType, "dy_zhenggong", "dy_jijian").
	//	Find(&voterList).Error; err != nil {
	//	return errors.New("获取投票人列表失败！")
	//}
	if err = gqaGlobal.GqaDb.
		Where("vote_type = ? ", toAddVoter.VoteType).
		Find(&voterList).Error; err != nil {
		return errors.New("获取投票人列表失败！")
	}
	for _, v := range toAddVoter.Voter {
		//避免同一个投票人出现多个投票权重
		for _, vl := range voterList {
			if vl.Voter == v {
				return errors.New("添加的固定投票人中包含已存在的投票权重类别！")
			}
		}
	}
	for _, v := range toAddVoter.Voter {
		av := model.GqaPluginVoteVoter{
			Voter:       v,
			VoteType:    toAddVoter.VoteType,
			VoterByUser: gqaModel.SysUser{Username: v},
			VoteRatio:   toAddVoter.VoteRatio,
			CreatedBy:   username,
		}
		addVoter = append(addVoter, av)
	}
	if len(addVoter) != 0 {
		err = gqaGlobal.GqaDb.Model(&model.GqaPluginVoteVoter{}).Save(&addVoter).Error
		return err
	} else {
		return errors.New("本次操作没有影响！")
	}
}

func VoterDelete(id uint, username string) (err error) {
	db := gqaGlobal.GqaDb.Model(&model.GqaPluginVoteVoter{})
	var candidate model.GqaPluginVoteVoter
	if err = db.Where("id = ?", id).First(&candidate).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&candidate).Error
	return err
}
