package privaterouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/api/privateapi"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	voteRouter := privateGroup.Use(middleware.LogOperationHandler())
	{
		//随机和固定投票人列表
		voteRouter.POST("voter-random-list", privateapi.VoterRandomList)
		//获取随机投票人
		voteRouter.POST("voter-random-get", privateapi.VoterRandomGet)
		//新增随机和固定投票人
		voteRouter.POST("voter-random-add", privateapi.VoterRandomAdd)
		//删除随机和固定投票人
		voteRouter.POST("voter-random-delete", privateapi.VoterRandomDelete)
		//检查是否可以投票-党员
		voteRouter.POST("can-vote-dy", privateapi.CanVoteDy)
		//检查是否可以投票-管理
		voteRouter.POST("can-vote-gl", privateapi.CanVoteGl)
		//投票结果查看已经投票人
		voteRouter.POST("vote-result-voted", privateapi.VoteResultVoted)
		//投票结果列表
		voteRouter.POST("vote-result-list", privateapi.VoteResultList)
		//投票结果图形
		voteRouter.POST("vote-result-chart", privateapi.VoteResultChart)
		//投票
		voteRouter.POST("vote-handle", privateapi.VoteHandle)
		//固定投票人列表
		voteRouter.POST("voter-list", privateapi.VoterList)
		//新增固定投票人
		voteRouter.POST("voter-add", privateapi.VoterAdd)
		//删除固定投票人
		voteRouter.POST("voter-delete", privateapi.VoterDelete)
		//候选人列表
		voteRouter.POST("candidate-list", privateapi.CandidateList)
		//新增候选人
		voteRouter.POST("candidate-add", privateapi.CandidateAdd)
		//删除候选人
		voteRouter.POST("candidate-delete", privateapi.CandidateDelete)
	}
}
