package Routes

import (
	"backend/payment-service/internal/Handler"
	"net/http"
)

func RegisterPlanRoutes(mux *http.ServeMux, handler *Handler.PlanHandler) {
	mux.HandleFunc("/api/plan/creatPlan", handler.CreatePlane)
	mux.HandleFunc("/api/plan/getAllPlan", handler.GetAllPlan)
	mux.HandleFunc("/api/plan/updatePlan", handler.UpdatePlan)
	mux.HandleFunc("/api/plan/deletePlan", handler.DeletePlan)
}
