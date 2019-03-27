package v1

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/romeufcrosa/best-route-finder/providers"
)

// GetRoute handles a request to the endpoint
func GetRoute(params Params) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()

		interactor, err := providers.GetRoutesProvider()
		if err != nil {
			Error(ctx, w, err)
		}

		origin := originFrom(ps)
		destination := destinationFrom(ps)

		result, err := interactor.GetRoute(ctx, origin, destination)

		if err != nil {
			Error(ctx, w, err)
		}

		Response(ctx, w, result)
	}
}

func originFrom(ps httprouter.Params) int {
	originNodeID, err := strconv.ParseInt(ps.ByName("from"), 10, 64)
	if err != nil {
		return 0
	}

	return int(originNodeID)
}

func destinationFrom(ps httprouter.Params) int {
	destinationNodeID, err := strconv.ParseInt(ps.ByName("to"), 10, 64)
	if err != nil {
		return 0
	}

	return int(destinationNodeID)
}
