package Routes

import (
	"net/http"

	"github.com/DonShanilka/user-service/internal/Handler"
)

func RegisterAdminRoutes(mux *http.ServeMux, handler *Handler.AdminHandler) {
	mux.HandleFunc("/api/admin/creatAdmin", handler.CreateAdmin)
	mux.HandleFunc("/api/admin/getAllAdmins", handler.GetAllAdmins)
	mux.HandleFunc("/api/admin/updateAdmin", handler.UpdateAdmin)
	mux.HandleFunc("/api/admin/deleteAdmin", handler.DeleteAdmin)
}
