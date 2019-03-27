package v1

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/providers"
)

// AddEdge handles the endpoint to add an edge
func AddEdge(params Params) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()

		jsonPayload, err := bodyBytes(r)
		if err != nil {
			Error(ctx, w, err)
			return
		}

		edge, err := domain.NewEdge(jsonPayload)
		if err != nil {
			Error(ctx, w, err)
			return
		}

		interactor, err := providers.GetRoutesProvider()
		if err != nil {
			Error(ctx, w, err)
		}

		result, err := interactor.AddEdge(ctx, edge)
		if err != nil {
			Error(ctx, w, err)
			return
		}

		Response(ctx, w, result)
	}
}

func bodyBytes(r *http.Request) ([]byte, error) {
	var bodyBytes []byte

	if r.Body == nil {
		return nil, ErrCartNoContent
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}
