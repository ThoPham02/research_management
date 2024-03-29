package logic

import (
	"context"

	"github.com/ThoPham02/research_management/common"
	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"
	"github.com/ThoPham02/research_management/service/progress/api/internal/types"
	"github.com/ThoPham02/research_management/service/progress/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEventsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEventsLogic {
	return &GetEventsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEventsLogic) GetEvents(req *types.GetEventsReq) (resp *types.GetEventsRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("GetEvents", req)

	var eventsModel []model.EventTbl
	var eventModel model.EventTbl
	var events []types.Event
	var event types.Event
	var total int64
	var mapStages = map[int64][]types.Stage{}

	eventsModel, err = l.svcCtx.EventModel.FindEvents(l.ctx, req.IsCurrent)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetEventsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if len(eventsModel) == 0 {
		return &types.GetEventsRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}
	stagesModel, err := l.svcCtx.StageModel.FindStages(l.ctx, 0)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetEventsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if len(stagesModel) == 0 {
		return &types.GetEventsRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}
	for _, stage := range stagesModel {
		mapStages[stage.EventId] = append(mapStages[stage.EventId], types.Stage{
			ID:          stage.Id,
			Name:        stage.Name,
			Description: stage.Description.String,
			Url:         stage.Url.String,
			EventID:     stage.EventId,
			TimeStart:   stage.TimeStart.Int64,
			TimeEnd:     stage.TimeEnd.Int64,
		})
	}

	for _, eventModel = range eventsModel {
		event = types.Event{
			ID:         eventModel.Id,
			Name:       eventModel.Name,
			SchoolYear: eventModel.SchoolYear.String,
			IsCurrent:  eventModel.IsCurrent.Int64,
			Stages:     mapStages[eventModel.Id],
		}
		events = append(events, event)
	}
	total = int64(len(events))

	return &types.GetEventsRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Total:  total,
		Events: events,
	}, nil
}
