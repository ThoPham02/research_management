package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/topic/api/internal/svc"
	"github.com/ThoPham02/research_management/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicSubcommitteeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicSubcommitteeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicSubcommitteeLogic {
	return &UpdateTopicSubcommitteeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicSubcommitteeLogic) UpdateTopicSubcommittee(req *types.UpdateTopicSubcommitteeReq) (resp *types.UpdateTopicSubcommitteeRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("UpdateTopicSubcommittee", req)

	err = l.svcCtx.TopicModel.UpdateSubcommittee(l.ctx, req.ListTopicID, req.SubcommitteeID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateTopicSubcommitteeRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	return &types.UpdateTopicSubcommitteeRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
