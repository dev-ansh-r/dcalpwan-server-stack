package gatewayconfigurationserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.thethings.network/lorawan-stack/v3/pkg/pfconfig/shared"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

func (s *Server) handleGetFrequencyPlan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	frequencyPlanID := mux.Vars(r)["frequency_plan_id"]
	fps, err := s.component.FrequencyPlansStore(ctx)
	if err != nil {
		webhandlers.Error(w, r, err)
		return
	}
	plan, err := fps.GetByID(frequencyPlanID)
	if err != nil {
		webhandlers.Error(w, r, err)
		return
	}
	config, err := shared.BuildSX1301Config(plan)
	if err != nil {
		webhandlers.Error(w, r, err)
		return
	}
	if r.Header.Get("User-Agent") == "TTNGateway" {
		// Filter out fields to reduce response size.
		config.TxLUTConfigs = nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		SX1301Conf *shared.SX1301Config `json:"SX1301_conf"`
	}{
		SX1301Conf: config,
	})
}
