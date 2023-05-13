// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/ThoPham02/research_management/service/progress/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/reports",
				Handler: GetTopicReportsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/report",
				Handler: CreateTopicReportHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/event",
				Handler: CreateEventHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/events",
				Handler: GetEventsHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/event-current/:eventID",
				Handler: UpdateCurrentEventHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/stage",
				Handler: CreateStageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/stages",
				Handler: GetStagesHandler(serverCtx),
			},
		},
	)
}
