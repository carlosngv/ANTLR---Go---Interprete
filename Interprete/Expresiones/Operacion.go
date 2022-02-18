package Expresion

import (
	"Parser/Interprete/AST"
	"Parser/Interprete/Interfaces"
	"fmt"
	"reflect"
	"strconv"
)

var suma_dominante = [5][5]AST.TipoDato{
	{AST.INTEGER, AST.FLOAT, AST.STRING, AST.NULL, AST.NULL},
	{AST.FLOAT, AST.FLOAT, AST.STRING, AST.NULL, AST.NULL},
	{AST.STRING, AST.STRING, AST.STRING, AST.STRING, AST.NULL},
	{AST.NULL, AST.NULL, AST.STRING, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
}

var multi_division_dominante = [5][5]AST.TipoDato{
	{AST.INTEGER, AST.FLOAT, AST.NULL, AST.NULL, AST.NULL},
	{AST.FLOAT, AST.FLOAT, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
}
var resta_dominante = [5][5]AST.TipoDato{
	{AST.INTEGER, AST.FLOAT, AST.NULL, AST.NULL, AST.NULL},
	{AST.FLOAT, AST.FLOAT, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
	{AST.NULL, AST.NULL, AST.NULL, AST.NULL, AST.NULL},
}

type Operacion struct {
	Op1      Interfaces.Expresion
	Operador string
	Op2      Interfaces.Expresion
	Unario   bool
}

func NewOperacion(Op1 Interfaces.Expresion, Operador string, Op2 Interfaces.Expresion, unario bool) Operacion {

	e := Operacion{Op1, Operador, Op2, unario}
	return e
}

func (p Operacion) ObtenerValor(ent AST.Entorno) AST.RetornoType {

	var retornoIzq AST.RetornoType
	var retornoDer AST.RetornoType

	if p.Unario {
		retornoIzq = p.Op1.ObtenerValor(ent)
	} else {

		if reflect.TypeOf(p.Op1).Name() == "Identificador" {
			existeIzquierdo := ent.ExisteSimbolo(p.Op1.(Identificador).Identificador)
			if !existeIzquierdo {
				return AST.RetornoType{Tipo: AST.NULL, Valor: nil}
			}
		}
		if reflect.TypeOf(p.Op2).Name() == "Identificador" {
			existeDerecho := ent.ExisteSimbolo(p.Op2.(Identificador).Identificador)
			if !existeDerecho {
				return AST.RetornoType{Tipo: AST.NULL, Valor: nil}
			}
		}

		retornoIzq = p.Op1.ObtenerValor(ent)
		retornoDer = p.Op2.ObtenerValor(ent)
	}

	var dominante AST.TipoDato

	switch p.Operador {
	case "+":
		{

			dominante = suma_dominante[retornoIzq.Tipo][retornoDer.Tipo]

			if dominante == AST.INTEGER {

				fmt.Println(retornoIzq.Tipo)
				fmt.Println(retornoDer.Tipo)

				/*

					nuevaVariable :=   variable.(instrucciones.Imprimir)

				*/

				return AST.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(int) + retornoDer.Valor.(int)}

			} else if dominante == AST.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoDer.Valor), 64)
				return AST.RetornoType{Tipo: dominante, Valor: val1 + val2}

			} else if dominante == AST.STRING {

				r1 := fmt.Sprintf("%v", retornoIzq.Valor)
				r2 := fmt.Sprintf("%v", retornoDer.Valor)

				return AST.RetornoType{Tipo: dominante, Valor: r1 + r2}
			}

		}

	case "*":
		{
			dominante = multi_division_dominante[retornoIzq.Tipo][retornoDer.Tipo]

			if dominante == AST.INTEGER {
				return AST.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(int) * retornoDer.Valor.(int)}

			} else if dominante == AST.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoDer.Valor), 64)
				return AST.RetornoType{Tipo: dominante, Valor: val1 * val2}

			} else if dominante == AST.NULL {
				return AST.RetornoType{Tipo: dominante, Valor: nil}
			}

		}
	case "-":
		{
			if p.Unario {

				if retornoIzq.Tipo != AST.INTEGER && retornoIzq.Tipo != AST.FLOAT {
					return AST.RetornoType{Tipo: AST.NULL, Valor: nil}
				}

				if retornoIzq.Tipo == AST.INTEGER {
					return AST.RetornoType{Tipo: retornoIzq.Tipo, Valor: -1 * retornoIzq.Valor.(int)}
				} else if retornoIzq.Tipo == AST.FLOAT {
					return AST.RetornoType{Tipo: retornoIzq.Tipo, Valor: -1 * retornoIzq.Valor.(float64)}
				}

			} else {
				dominante = resta_dominante[retornoIzq.Tipo][retornoDer.Tipo]

				if dominante == AST.INTEGER {

					fmt.Println(retornoIzq.Tipo)
					fmt.Println(retornoDer.Tipo)

					return AST.RetornoType{Tipo: dominante, Valor: retornoIzq.Valor.(int) - retornoDer.Valor.(int)}

				} else if dominante == AST.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoIzq.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retornoDer.Valor), 64)
					return AST.RetornoType{Tipo: dominante, Valor: val1 - val2}

				} else if dominante == AST.NULL {
					return AST.RetornoType{Tipo: dominante, Valor: nil}
				}
			}
		}
	}

	return AST.RetornoType{Tipo: AST.NULL, Valor: nil}
}
