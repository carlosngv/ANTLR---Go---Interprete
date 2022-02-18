package Instrucciones

import (
	"Parser/Interprete/AST"
	Expresion "Parser/Interprete/Expresiones"
	"Parser/Interprete/Interfaces"
	"encoding/json"
	"fmt"

	"github.com/colegno/arraylist"
)

var tipoDef = [5][5]AST.TipoDato{
	{AST.INTEGER, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
	{AST.FLOAT, AST.FLOAT, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.STRING, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.BOOLEAN, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
}

type Declaracion struct {
	ValorInicializacion Interfaces.Expresion
	TipoVariables       AST.TipoDato
	ListaVars           *arraylist.List
}

func NewDeclaracion(listaVars *arraylist.List, tipoVariables AST.TipoDato) *Declaracion {
	return &Declaracion{
		TipoVariables: tipoVariables,
		ListaVars:     listaVars,
	}
}
func NewDeclaracionInicializacion(listaVars *arraylist.List, tipoVariables AST.TipoDato, valInicial Interfaces.Expresion) *Declaracion {
	return &Declaracion{
		TipoVariables:       tipoVariables,
		ListaVars:           listaVars,
		ValorInicializacion: valInicial,
	}
}

func (dec *Declaracion) Ejecutar(ent AST.Entorno) interface{} {

	if dec.esInicializado() {
		if dec.ListaVars.Len() > 1 {
			return nil
		}

		retornoExpresion := dec.ValorInicializacion.ObtenerValor(ent)

		tipoExpresion := retornoExpresion.Tipo
		tipoDeclaracion := dec.TipoVariables

		tipoResultante := tipoDef[tipoDeclaracion][tipoExpresion]

		if tipoResultante == AST.NULL {
			return nil
		}

		for i := 0; i < dec.ListaVars.Len(); i++ {

			varDeclarar := dec.ListaVars.GetValue(i).(Expresion.Identificador)

			if ent.ExisteSimbolo(varDeclarar.Identificador) {
				fmt.Printf("Errror, variable %s ya declarada", varDeclarar.Identificador)
			} else {
				simboloTabala := AST.NewSimboloIdentificadorValor(
					0,
					0,
					varDeclarar.Identificador,
					retornoExpresion.Valor,
					tipoResultante)

				ent.AgregarSimbolo(varDeclarar.Identificador, simboloTabala)
			}

		}

	}

	data, err := json.MarshalIndent(ent, "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)
	fmt.Printf("%v", dec.ListaVars)

	return nil
}

func (dec *Declaracion) esInicializado() bool {
	return dec.ValorInicializacion != nil
}
