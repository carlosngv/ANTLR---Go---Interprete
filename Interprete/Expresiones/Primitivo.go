package Expresion

import "Parser/Interprete/AST"

type Primitivo struct {
	Valor interface{}
	Tipo  AST.TipoDato
}

func (p Primitivo) ObtenerValor(ent AST.Entorno) AST.RetornoType {

	return AST.RetornoType{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NewPrimitivo(val interface{}, tipo AST.TipoDato) Primitivo {
	e := Primitivo{val, tipo}
	return e
}
