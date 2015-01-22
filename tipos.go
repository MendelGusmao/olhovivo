package olhovivo

const (
	TPTS Sentido = 1
	TSTP Sentido = 2
)

type Sentido byte

type Linha struct {
	CodigoLinha     int64
	Circular        bool
	Letreiro        string
	Sentido         Sentido
	Tipo            byte
	DenominacaoTPTS string
	DenominacaoTSTP string
	Informacoes     *string
}

type Parada struct {
	CodigoParada int64
	Nome         string
	Endereco     string
	Latitude     float64
	Longitude    float64
}

type Corredor struct {
	CodigoCorredor int64 `json:"CodCorredor"`
	Nome           string
}

type Veiculo struct {
	Prefixo    string  `json:"p"`
	Articulado bool    `json:"a"`
	Latitude   float64 `json:"py"`
	Longitude  float64 `json:"px"`
	Hora       string  `json:"t"`
}

type Posicao struct {
	Hora     string    `json:"hr"`
	Veiculos []Veiculo `json:"vs"`
}

// GET /Previsao?codigoParada={codigoParada}&codigoLinha={codigoLinha}
// GET /Previsao/Parada?codigoParada={codigoParada}
type Previsao struct {
	Hora   string           `json:"hr"`
	Parada *ParadaComLinhas `json:"p"`
}

type ParadaComLinhas struct {
	CodigoParada int64           `json:"cp"`
	Nome         string          `json:"np"`
	Latitude     float64         `json:"py"`
	Longitude    float64         `json:"px"`
	Linhas       []LinhaPrevisao `json:"l"`
}

type LinhaPrevisao struct {
	Letreiro    string    `json:"c"`
	CodigoLinha int64     `json:"cl"`
	SL          int64     `json:"sl"` // TODO sl???
	Letreiro0   string    `json:"lt0"`
	Letreiro1   string    `json:"lt1"`
	QV          int64     `json:"qv"` // TODO qv???
	Veiculos    []Veiculo `json:"vs"`
}

// GET /Previsao/Linha?codigoLinha={codigoLinha}
type PrevisaoLinha struct {
	Hora    string              `json:"hr"`
	Paradas []ParadaComVeiculos `json:"ps"`
}

type ParadaComVeiculos struct {
	CodigoParada int64     `json:"cp"`
	Nome         string    `json:"np"`
	Latitude     float64   `json:"py"`
	Longitude    float64   `json:"px"`
	Veiculos     []Veiculo `json:"vs"`
}
