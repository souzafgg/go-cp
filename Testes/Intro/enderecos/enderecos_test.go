// teste de unidade
package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {
	enderecoParaTest := "Avenida Paulista"
	tipoEsperado := "Avenida"
	tipoRecebido := TipoEndereco(enderecoParaTest)
	if tipoRecebido != tipoEsperado {
		t.Error("Tipo diferente")
	}
}
