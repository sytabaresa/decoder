package main

import (
	"fmt"
	"time"

	"github.com/sytabaresa/decoder"
)

const (
	frameBegin  = 0x3C
	frameEnding = 0x3E
	frameLength = 8
)

//go:generate stringer -type=Dato
type Dato int

const (
	Nulo Dato = iota
	TensionLineaR
	TensionLineaS
	TensionLineaT
	CorrienteLineaR
	CorrienteLineaS
	CorrienteLineaT
	CorrienteN
	AnguloS
	AnguloT
	FrecuenciaR
	FrecuenciaS
	FrecuenciaT
	DesfaseRS
	DesfaseST
	DesfaseTR
	//CorrienteN
	//Temperatura
	//Presion
	//Iluminancia
	//Par2_5
	//Par10
	//Humedad
	//Fase
)

type DatoType struct {
	Nodo     int
	TipoDato Dato
	Fecha    time.Time
	Medicion float32
}

const (
	Measure = "Potencia"
)

func (e DatoType) String() string {
	return fmt.Sprintf("%s,sensor=%v %s=%v %v", Measure, e.Nodo, e.TipoDato, e.Medicion, e.Fecha.Unix())
}

func ParseToken(token []byte) (e DatoType, err error) {
	if len(token) != frameLength {
		return DatoType{}, fmt.Errorf("tamaño invalido de datagrama 0x%X", token)
	}
	e.Nodo = int(token[0])
	dat := int(token[1])
	if dat > 15 {
		return DatoType{}, fmt.Errorf("Tipo de Dato no válido")
	}
	e.TipoDato = Dato(dat)

	if token[2] != 0 {
		return DatoType{}, fmt.Errorf("No se encuentra el byte 0x00 en el medio")
	}

	var digit float32 = 100.0
	var measure float32 = 0.0
	for i := 3; i < 8; i++ {
		d, err := decoder.ToBCD(token[i])
		if err != nil {
			return DatoType{}, fmt.Errorf("%e : en el byte %v", err, i)
		}
		measure += float32(d) * digit
		digit /= 10
	}
	e.Medicion = measure
	e.Fecha = time.Now()
	return e, nil
}

func ReadToken(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var start int
	// Busca el inicio del token
	for start = 0; start < len(data); start++ {
		if data[start] == frameBegin {
			break
		}
	}
	// Busca el final del token
	var end = start + 1
	// if end >= len(data)-1 {
	// 	return start, nil, nil
	// }
	// if data[end] == frameEnding {
	// 	start = end
	// 	end++
	// }
	for ; end < len(data); end++ {
		if data[end] == frameEnding {
			return end + 1, data[start+1 : end], nil
		}
	}

	//Requiere mas datos
	return start, nil, nil
}
