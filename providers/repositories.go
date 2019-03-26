package providers

import (
	"github.com/romeufcrosa/best-route-finder/domain/services"
	"github.com/romeufcrosa/best-route-finder/repositories/mysql"
)

// Provider list...
var (
	Routes = Provider("repositories/mysql/routes")
)

// RegisterRepositoryProviders ...
func RegisterRepositoryProviders() {
	Register(Routes, func() (provider interface{}, err error) { // nolint:errcheck
		return mysql.NewRoutes(manufacturer.params.pool), nil
	})
}

// GetRoutesProvider returns the provider of the routes repo
func GetRoutesProvider() (services.Routes, error) {
	provider, err := Get(Routes)
	if err != nil {
		return services.Routes{}, err
	}

	routesStorage := provider.(mysql.Routes)

	return services.NewRoutesInteractor(routesStorage), nil
}
