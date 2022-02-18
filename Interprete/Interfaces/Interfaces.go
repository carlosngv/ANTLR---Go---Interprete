package Interfaces

import (
	"Parser/Interprete/AST"
)

type Expresion interface {
	ObtenerValor(entorno AST.Entorno) AST.RetornoType
}

type Instruccion interface {
	Ejecutar(entorno AST.Entorno) interface{}
}
