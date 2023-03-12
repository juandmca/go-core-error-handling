package middleware

import (
	"fmt"
	"net/http"

	"github.com/magiconair/properties"
	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type RubikMiddleware struct{}

// Funcion que valida la estructura de las cabeceras asegurandose que vengan los valores
// requeridos para una correcta trazabilidad de la peticion
func (r *RubikMiddleware) HeaderChecker() web.Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			p := properties.MustLoadFile("../properties/header_structure.properties", properties.UTF8).Map()
			for indice, property := range p {
				fmt.Println("Indice: ", indice, "Nombre: ", property)
			}
		}
	}
}
