package olhovivo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/lafikl/fluent"
)

const (
	servidor                 string = "http://api.olhovivo.sptrans.com.br/v0"
	loginAutenticar          string = servidor + "/Login/Autenticar?token=%s"
	buscarLinha              string = servidor + "/Linha/Buscar?termosBusca=%s"
	carregarDetalhesLinha    string = servidor + "/Linha/CarregarDetalhes?codigoLinha=%d"
	buscarParada             string = servidor + "/Linha/Buscar?termosBusca=%s"
	buscarParadasPorLinha    string = servidor + "/Parada/BuscarParadasPorLinha?codigoLinha=%d"
	buscarParadasPorCorredor string = servidor + "/Parada/BuscarParadasPorCorredor?codigoCorredor=%d"
	corredor                 string = servidor + "/Corredor"
	posicao                  string = servidor + "/Posicao?codigoLinha=%d"
	previsao                 string = servidor + "/Previsao?codigoParada=%d&codigoLinha=%d"
	previsaoLinha            string = servidor + "/Previsao/Linha?codigoLinha=%d"
	previsaoParada           string = servidor + "/Previsao/Parada?codigoParada=%d"
)

type Cliente struct {
	cookie string
}

func (c *Cliente) Autenticar(chave string) (bool, error) {
	request := fluent.New().Post(fmt.Sprintf(loginAutenticar, url.QueryEscape(chave)))
	response, err := request.Send()

	if err != nil {
		return false, fmt.Errorf("olhovivo.Cliente.Autenticar: %v", err)
	}

	ok := false
	content, err := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(content, &ok); err != nil {
		return ok, fmt.Errorf("olhovivo.Cliente.Autenticar: %v", err)
	}

	if !ok {
		return ok, fmt.Errorf("olhovivo.Cliente.Autenticar: falhou")
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == "apiCredentials" {
			c.cookie = fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
		}
	}

	return ok, nil
}

func (c *Cliente) BuscarLinha(termos string) ([]Linha, error) {
	url := fmt.Sprintf(buscarParada, termos)
	linhas := []Linha{}

	if err := c.request(url, &linhas); err != nil {
		return linhas, fmt.Errorf("olhovivo.Cliente.BuscarLinha: %v", err)
	}

	return linhas, nil
}

func (c *Cliente) CarregarDetalhesLinha(codigoLinha int64) ([]Linha, error) {
	url := fmt.Sprintf(carregarDetalhesLinha, codigoLinha)
	linhas := []Linha{}

	if err := c.request(url, &linhas); err != nil {
		return linhas, fmt.Errorf("olhovivo.Cliente.CarregarDetalhesLinha: %v", err)
	}

	return linhas, nil
}

func (c *Cliente) BuscarParada(termos string) ([]Parada, error) {
	url := fmt.Sprintf(buscarParada, termos)
	paradas := []Parada{}

	if err := c.request(url, &paradas); err != nil {
		return paradas, fmt.Errorf("olhovivo.Cliente.BuscarParada: %v", err)
	}

	return paradas, nil
}

func (c *Cliente) BuscarParadasPorLinha(codigoLinha int64) ([]Parada, error) {
	url := fmt.Sprintf(buscarParadasPorLinha, codigoLinha)
	paradas := []Parada{}

	if err := c.request(url, &paradas); err != nil {
		return paradas, fmt.Errorf("olhovivo.Cliente.BuscarParadasPorLinha: %v", err)
	}

	return paradas, nil
}

func (c *Cliente) BuscarParadasPorCorredor(codigoCorredor int) ([]Parada, error) {
	url := fmt.Sprintf(buscarParadasPorCorredor, codigoCorredor)
	paradas := []Parada{}

	if err := c.request(url, &paradas); err != nil {
		return paradas, fmt.Errorf("olhovivo.Cliente.BuscarParadasPorCorredor: %v", err)
	}

	return paradas, nil
}

func (c *Cliente) Corredor() ([]Corredor, error) {
	url := corredor
	corredor := []Corredor{}

	if err := c.request(url, &corredor); err != nil {
		return corredor, fmt.Errorf("olhovivo.Cliente.Corredor: %v", err)
	}

	return corredor, nil
}

func (c *Cliente) Posicao(codigoLinha int64) (Posicao, error) {
	url := posicao
	posicao := Posicao{}

	if err := c.request(url, &posicao); err != nil {
		return posicao, fmt.Errorf("olhovivo.Cliente.Posicao: %v", err)
	}

	return posicao, nil
}

func (c *Cliente) Previsao(codigoParada int64, codigoLinha int64) (Previsao, error) {
	url := fmt.Sprintf(previsao, codigoParada, codigoLinha)
	previsao := Previsao{}

	if err := c.request(url, &previsao); err != nil {
		return previsao, fmt.Errorf("olhovivo.Cliente.PrevisaoParada: %v", err)
	}

	return previsao, nil
}

func (c *Cliente) PrevisaoLinha(codigoLinha int64) (PrevisaoLinha, error) {
	url := fmt.Sprintf(previsaoLinha, codigoLinha)
	previsao := PrevisaoLinha{}

	if err := c.request(url, &previsao); err != nil {
		return previsao, fmt.Errorf("olhovivo.Cliente.PrevisaoLinha: %v", err)
	}

	return previsao, nil
}

func (c *Cliente) PrevisaoParada(codigoParada int64) (Previsao, error) {
	url := fmt.Sprintf(previsaoParada, codigoParada)
	previsao := Previsao{}

	if err := c.request(url, &previsao); err != nil {
		return previsao, fmt.Errorf("olhovivo.Cliente.PrevisaoParada: %v", err)
	}

	return previsao, nil
}

func (c *Cliente) request(url string, object interface{}) error {
	request := fluent.New().Get(url)
	request.SetHeader("Cookie", c.cookie)

	response, err := request.Send()

	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &object); err != nil {
		return err
	}

	return nil
}
