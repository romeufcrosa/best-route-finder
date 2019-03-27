package v1

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/providers"
)

// AddNode handles the endpoint to insert a node
func AddNode(params Params) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()

		jsonPayload, err := bodyBytes(r)
		if err != nil {
			Error(ctx, w, err)
			return
		}

		node, err := domain.NewNode(jsonPayload)
		if err != nil {
			Error(ctx, w, err)
			return
		}

		interactor, err := providers.GetRoutesProvider()
		if err != nil {
			Error(ctx, w, err)
		}

		result, err := interactor.AddNode(ctx, node)
		if err != nil {
			Error(ctx, w, err)
			return
		}

		Response(ctx, w, result)
	}
}
