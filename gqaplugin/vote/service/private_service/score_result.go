package private_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"gorm.io/gorm"
)

func VoteResultVoted(getVoteScoreVoted model.RequestScoreResultVoted) (err error, score []model.GqaPluginVoteScoreResult) {
	var voteScoreList []model.GqaPluginVoteScoreResult
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteScoreResult{})
	//配置搜索
	if getVoteScoreVoted.VoteMonth != "" {
		db = db.Where("vote_month like ?", "%"+getVoteScoreVoted.VoteMonth+"%")
	}
	if getVoteScoreVoted.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoteScoreVoted.VoteType+"%")
	}
	if err != nil {
		return
	}
	err = db.Distinct("vote_from").Order(global.OrderByColumn("vote_from", true)).
		Preload("VoteFromByUser").Find(&voteScoreList).Error
	return err, voteScoreList
}

func VoteResultList(getVoteScoreList model.RequestScoreResultList, username string) (err error, score []model.GqaPluginVoteScoreResult, total int64) {
	pageSize := getVoteScoreList.PageSize
	offset := getVoteScoreList.PageSize * (getVoteScoreList.Page - 1)
	var voteScoreList []model.GqaPluginVoteScoreResult
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteScoreResult{})
	//配置搜索
	if getVoteScoreList.VoteMonth != "" {
		db = db.Where("vote_month like ?", "%"+getVoteScoreList.VoteMonth+"%")
	}
	if getVoteScoreList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoteScoreList.VoteType+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn("vote_month", true)).
		Order(global.OrderByColumn("candidate", false)).
		Preload("CandidateByUser").Preload("VoteFromByUser").Find(&voteScoreList).Error
	return err, voteScoreList, total
}

func VoteResultChart(getVoteScoreList model.RequestScoreResultList, username string) (err error, score []model.GqaPluginVoteScoreResult, total int64) {
	//pageSize := getVoteScoreList.PageSize
	//offset := getVoteScoreList.PageSize * (getVoteScoreList.Page - 1)
	var voteScoreList []model.GqaPluginVoteScoreResult
	var db *gorm.DB
	db = global.GqaDb.Model(&model.GqaPluginVoteScoreResult{})
	//配置搜索
	if getVoteScoreList.VoteMonth != "" {
		db = db.Where("vote_month like ?", "%"+getVoteScoreList.VoteMonth+"%")
	}
	if getVoteScoreList.VoteType != "" {
		db = db.Where("vote_type like ?", "%"+getVoteScoreList.VoteType+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//db = db.Select("id, candidate, vote_type, vote_type_detail, AVG(vote_score) as vote_score, vote_ratio, vote_from, vote_month, created_at").
	//	Group("candidate").Group("vote_type_detail").Group("vote_ratio").Find(&voteScoreList)
	//err = db.
	//	Select("id, candidate, vote_type, vote_type_detail, SUM(vote_score * ratio / 100) as vote_score, vote_ratio, vote_from, vote_month, created_at").
	//	Group("candidate").Group("vote_type_detail").Find(&voteScoreList).Error
	err = global.GqaDb.Table("(?) as u",
		db.Model(&model.GqaPluginVoteScoreResult{}).
			Select("candidate, vote_type, vote_type_detail, AVG(vote_score) as vote_score, vote_ratio, ratio, vote_from, vote_month").
			Group("candidate").Group("vote_type_detail").Group("vote_ratio")).
		Select("candidate, vote_type, vote_type_detail, SUM(vote_score * ratio / 100) as vote_score, vote_ratio, vote_from, vote_month").
		Group("candidate").Group("vote_type_detail").
		Preload("CandidateByUser").Preload("VoteFromByUser").
		Find(&voteScoreList).Error
	return err, voteScoreList, total
}
