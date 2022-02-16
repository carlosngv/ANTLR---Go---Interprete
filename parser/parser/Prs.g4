parser grammar Prs;

options {
  tokenVocab = PrsLexer;
}

// Imports
@header{
    import arrayList "github.com/colegno/arraylist"
}

// Aux structs
@members{

}

start returns [*arrayList.List lista]
  : L_LLAVE instrucciones R_LLAVE {$lista = $instrucciones.l}
;


instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instruccion*  {
      listInt := localctx.(*InstruccionesContext).GetE()
      		for _, e := range listInt {
            $l.Add(e.GetInstr())
          }
      fmt.Printf("tipo %T",localctx.(*InstruccionesContext).GetE())
  }
;

instruccion returns [interfaces.Instruccion instr]
  : SYSTEM '.' OUT '.' PRINTLN LP expression RP ';' {$instr = funbasica.NewImprimir($expression.p)}
;

expression returns[interfaces.Expresion p]
    : expr_rel    {$p = $expr_rel.p}
    | expr_arit    {$p = $expr_arit.p}
;

expr_rel returns[interfaces.Expresion p]
    : opIz = expr_rel op=( MAYORIGUAL | MENORIGUAL | MENOR | MAYOR ) opDe = expr_rel {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | expr_arit  {$p = $expr_arit.p}
;

expr_arit returns[interfaces.Expresion p]
    : opIz = expr_arit op=('*'|'/') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | opIz = expr_arit op=('+'|'-') opDe = expr_arit {$p = expresion.NewOperacion($opIz.p,$op.text,$opDe.p,false)}
    | primitivo {$p = $primitivo.p}
    | LP expression RP {$p = $expression.p}
;

primitivo returns[interfaces.Expresion p]
    :NUMBER {
            	num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expresion.NewPrimitivo (num,interfaces.INTEGER)
       }
    | STRING {
      str:= $STRING.text[1:len($STRING.line)-1]
      $p = expresion.NewPrimitivo(str,interfaces.STRING)}
;
