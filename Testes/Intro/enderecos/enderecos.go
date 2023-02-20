package enderecos

import "strings"

// TipoEndereco de endereço se o endereço tem um tipo válido
func TipoEndereco(e string) string {
	tipoValid := []string{"Rua", "Avenida", "Estrada"}
	letraMinuscula := strings.ToLower(e)
	primeiraPalavra := strings.Split(letraMinuscula, " ")[0]

	eTipoValido := false

	for _, tipo := range tipoValid {
		if tipo == primeiraPalavra {
			eTipoValido = true
		}
	}
	if eTipoValido {
		return strings.ToTitle(primeiraPalavra)
	}
	return "Tipo Inválido"
}
