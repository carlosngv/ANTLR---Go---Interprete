package Expresion

import "Parser/Interprete/AST"

type Identificador struct {
	Identificador string
}

func NewIdentificador(identificador string) Identificador {
	return Identificador{Identificador: identificador}
}

func (ide Identificador) ObtenerValor(ent AST.Entorno) AST.RetornoType {

	var encontrado bool = ent.ExisteSimbolo(ide.Identificador)

	if !encontrado {
		return AST.RetornoType{Valor: nil, Tipo: AST.NULL}
	}

	simbo := ent.ObtenerSimbolo(ide.Identificador)

	return AST.RetornoType{Valor: simbo.Valor, Tipo: simbo.Tipo}

}
