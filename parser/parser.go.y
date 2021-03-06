%{
package parser
%}
%type<expr> prog
%type<expr> stats
%type<expr> prog_stat
%type<expr> stat
%type<expr> declarator
%type<expr> declarator_list
%type<expr> ident_list
%type<expr> expr_list
%type<expr> expr_list_paren
%type<expr> expr_assign_list
%type<expr> expr
%type<expr> postfix_incdec
%type<expr> _postfix_assign
%type<expr> prefix_expr
%type<expr> assign_stat
%type<expr> for_stat
%type<expr> if_stat
%type<expr> elseif_stat
%type<expr> jmp_stat
%type<expr> flow_stat
%type<expr> func
%type<expr> func_stat
%type<expr> func_params_list

%union {
  token Token
  expr  Node
}

/* Reserved words */
%token<token> TDo TLocal TElseIf TThen TEnd TBreak TElse TFor TWhile TFunc TIf TReturn TReturnVoid TRepeat TUntil TNot TLabel TGoto TIn 

/* Literals */
%token<token> TOr TAnd TEqeq TNeq TLte TGte TIdent TNumber TString 
%token<token> '{' '[' '(' '=' '>' '<' '+' '-' '*' '/' '%' '^' '#' '.' '&' TIDiv
%token<token> TAddEq TSubEq TMulEq TDivEq TModEq
%token<token> TSquare TDotDot 

/* Operators */
%right 'T'
%right TElse
%left ASSIGN
%right FUNC
%left TOr
%left TAnd
%left '>' '<' TGte TLte TEqeq TNeq
%left TDotDot
%left '+' '-' '^'
%left '*' '/' '%' TIDiv
%right UNARY /* not # -(unary) */

%% 

prog: 
        {
            $$ = __chain()
            if l, ok := yylex.(*Lexer); ok {
                l.Stmts = $$
            }
        } |
        prog prog_stat {
       $$ = $1.append($2)
            if l, ok := yylex.(*Lexer); ok {
                l.Stmts = $$
            }
        }

stats: 
        {
            $$ = __chain()
        } |
        stats stat {
       $$ = $1.append($2)
        }

prog_stat:
        jmp_stat       { $$ = $1 } |
        flow_stat      { $$ = $1 } |
        assign_stat    { $$ = $1 } |
   func_stat      { $$ = $1 } |
        TDo prog TEnd  { $$ = __do($2) } |
        ';'            { $$ = emptyNode }

stat:
        jmp_stat       { $$ = $1 } |
        flow_stat      { $$ = $1 } |
        assign_stat    { $$ = $1 } |
        TDo stats TEnd { $$ = __do($2) } |
        ';'            { $$ = emptyNode }

flow_stat:
        for_stat       { $$ = $1 } |
        if_stat        { $$ = $1 }

_postfix_assign:
   TAddEq         { $$ = NewSymbol(AAdd).SetPos($1.Pos) } |
   TSubEq         { $$ = NewSymbol(ASub).SetPos($1.Pos) } |
   TMulEq         { $$ = NewSymbol(AMul).SetPos($1.Pos) } |
   TDivEq         { $$ = NewSymbol(ADiv).SetPos($1.Pos) } |
   TModEq         { $$ = NewSymbol(AMod).SetPos($1.Pos) }

assign_stat:
        prefix_expr {
            if $1.isCallStat() {
                // Single call statement, clear env.V to avoid side effects
                $$ = __chain($1, popvClearNode)
            } else {
                $$ = $1
            }
        } | 
        postfix_incdec {
            $$ = $1
        } |
        TLocal ident_list {
            $$ = __chain()
            for _, v := range $2.Nodes {
                $$ = $$.append(__set(v, NewSymbol(ANil)).SetPos($1.Pos))
            }
        } |
        TLocal ident_list '=' expr_list {
            $$ = __local($2.Nodes, $4.Nodes, $1.Pos)
        } |
        declarator_list '=' expr_list {
            $$ = __moveMulti($1.Nodes, $3.Nodes, $2.Pos)
        } 

postfix_incdec:
        TIdent _postfix_assign expr %prec ASSIGN  {
       if $2.SymbolValue() == AAdd && $3.IsNumber() {
              $$ = __inc(NewSymbolFromToken($1), $3).SetPos($2.Pos())
       } else if $2.SymbolValue() == ASub && $3.IsNumber() {
      f, i, isInt := $3.NumberValue()
      if isInt {
             $$ = __inc(NewSymbolFromToken($1), NewNumberFromInt(-i)).SetPos($2.Pos())
      } else {
             $$ = __inc(NewSymbolFromToken($1), NewNumberFromFloat(-f)).SetPos($2.Pos())
      }
       } else {
              $$ = __move(NewSymbolFromToken($1), NewComplex($2, NewSymbolFromToken($1), $3)).SetPos($2.Pos())
       }
        } |
        prefix_expr '[' expr ']' _postfix_assign expr %prec ASSIGN {
            $$ = __store($1, $3, NewComplex($5, __load($1, $3), $6).SetPos($5.Pos()))
        } |
        prefix_expr '.' TIdent _postfix_assign expr %prec ASSIGN {
       i := NewString($3.Str) 
            $$ = __store($1, i, NewComplex($4, __load($1, i), $5).SetPos($4.Pos()))
        }

for_stat:
        TWhile expr TDo stats TEnd {
            $$ = __loop(__if($2, $4, breakNode).SetPos($1.Pos)).SetPos($1.Pos)
        } |
        TRepeat stats TUntil expr {
            $$ = __loop(
                __chain(
                    $2,
                    __if($4, breakNode, emptyNode).SetPos($1.Pos),
                ).SetPos($1.Pos),
            ).SetPos($1.Pos)
        } |
        TFor TIdent '=' expr ',' expr TDo stats TEnd {
            forVar, forEnd := NewSymbolFromToken($2), randomVarname()
            $$ = __do(
                    __set(forVar, $4).SetPos($1.Pos),
                    __set(forEnd, $6).SetPos($1.Pos),
                    __loop(
                        __if(
                            __lessEq(forVar, forEnd),
                            __chain($8, __inc(forVar, oneNode).SetPos($1.Pos)),
                            breakNode,
                        ).SetPos($1.Pos),
                    ).SetPos($1.Pos),
                )
        } |
        TFor ident_list TIn expr_list TDo stats TEnd {
            f := randomVarname()
            bsDDD, bs := randomDDDVarname()
            if len($4.Nodes) == 1 {
                // for b1, ..., bn in expr do
                vsDDD, vs := randomDDDVarname()
                $$ = __do(
                    // local f, ...vs = expr
                    __local([]Node{f, vsDDD}, $4.Nodes, $1.Pos),
                    // local b1, ..., bn, ...bs = f(vs)
                    __local(append($2.Nodes, bsDDD), []Node{__call(f, NewComplex(vs)).SetPos($1.Pos)}, $1.Pos))
            } else {
                // for b1, ..., bn in expr, v1, ..., vn do
                $$ = __do(
                    // local f = expr
                    __set(f, $4.Nodes[0]).SetPos($1.Pos),
                    // local b1, ..., bn, ...bs = f(v1, ..., vn)
                    __local(append($2.Nodes, bsDDD), []Node{__call(f, NewComplex($4.Nodes[1:]...)).SetPos($1.Pos)}, $1.Pos))
            }
            $$ = $$.append(
                __loop(
                    __chain(
                        __if(
                            NewComplex(NewSymbol(AEq), NewSymbol("nil"), $2.Nodes[0]).SetPos($1.Pos),
                            breakNode,
                            __chain($6),
                        ).SetPos($1.Pos),
                        // b1, ..., bn, ...bs = f(b1, ..., bn, bs)
                        __moveMulti(append($2.Nodes, bsDDD), []Node{
                            __call(f, NewComplex(append($2.DuplicateNodes(), bs)...)).SetPos($1.Pos),
                        }, $1.Pos),
                    ),
                ).SetPos($1.Pos),
            )
        } |
        TFor TIdent '=' expr ',' expr ',' expr TDo stats TEnd {
            forVar, forEnd := NewSymbolFromToken($2), randomVarname()
            if $8.IsNumber() { // step is a static number, easy case
                var cond Node
                if $8.IsNegativeNumber() {
                    cond = __lessEq(forEnd, forVar)
                } else {
                    cond = __lessEq(forVar, forEnd)
                }
                $$ = __do(
                    __set(forVar, $4).SetPos($1.Pos),
                    __set(forEnd, $6).SetPos($1.Pos),
                    __loop(
                        __chain(
                            __if(
                                cond,
                                __chain($10, __inc(forVar, $8)),
                                breakNode,
                            ).SetPos($1.Pos),
                        ),
                    ).SetPos($1.Pos),
                )
            } else { 
                forStep := randomVarname()
                $$ = __do(
                    __set(forVar, $4).SetPos($1.Pos),
                    __set(forEnd, $6).SetPos($1.Pos),
                    __set(forStep, $8).SetPos($1.Pos),
                    __loop(
                        __chain(
                            __if(
                                __less(zeroNode, forStep).SetPos($1.Pos),
                                // +step
                                __if(__less(forEnd, forVar), breakNode, emptyNode).SetPos($1.Pos),
                                // -step
                                __if(__less(forVar, forEnd), breakNode, emptyNode).SetPos($1.Pos),
                            ).SetPos($1.Pos),
                            $10,
                            __inc(forVar, forStep),
                        ),
                    ).SetPos($1.Pos),
                )
            }
            
        } 

if_stat:
        TIf expr TThen stats elseif_stat TEnd %prec 'T' {
            $$ = __if($2, $4, $5).SetPos($1.Pos)
        }

elseif_stat:
        {
            $$ = NewComplex()
        } |
        TElse stats {
            $$ = $2
        } |
        TElseIf expr TThen stats elseif_stat {
            $$ = __if($2, $4, $5).SetPos($1.Pos)
        }

func:
        TFunc        { $$ = NewSymbol(AMove).SetPos($1.Pos) } |
        TLocal TFunc { $$ = NewSymbol(ASet).SetPos($1.Pos) }

func_stat:
        func TIdent func_params_list stats TEnd         { $$ = __func($1, $2, $3, "", $4) } | 
        func TIdent func_params_list TString stats TEnd { $$ = __func($1, $2, $3, $4.Str, $5) }

func_params_list:
        '(' ')'                           { $$ = NewComplex() } |
        '(' ident_list ')'                { $$ = $2 }

jmp_stat:
        TBreak                            { $$ = NewComplex(NewSymbol(ABreak)).SetPos($1.Pos) } |
        TGoto TIdent                      { $$ = NewComplex(NewSymbol(AGoto), NewSymbolFromToken($2)).SetPos($1.Pos) } |
        TLabel TIdent TLabel              { $$ = NewComplex(NewSymbol(ALabel), NewSymbolFromToken($2)) } |
        TReturnVoid                       { $$ = NewComplex(NewSymbol(AReturn), emptyNode).SetPos($1.Pos) } |
        TReturn expr_list                 {
            if len($2.Nodes) == 1 {
                x := $2.Nodes[0]
                if len(x.Nodes) == 3 && x.Nodes[0].SymbolValue() == ACall { 
                  x.Nodes[0].strSym = ATailCall
                }
            }
            $$ = NewComplex(NewSymbol(AReturn), $2).SetPos($1.Pos) 
        }

declarator:
        TIdent                            { $$ = NewSymbolFromToken($1) } |
        prefix_expr '[' expr ']'          { $$ = __load($1, $3).SetPos($2.Pos) } |
        prefix_expr '[' expr ':' expr ']' { $$ = NewComplex(NewSymbol(ASlice), $1, $3, $5).SetPos($2.Pos) } |
        prefix_expr '.' TIdent            { $$ = __load($1, NewString($3.Str)).SetPos($2.Pos) }

declarator_list:
        declarator                        { $$ = NewComplex($1) } |
        declarator_list ',' declarator    { $$ = $1.append($3) }

ident_list:
        TIdent                            { $$ = NewComplex(NewSymbolFromToken($1)) } | 
        ident_list ',' TIdent             { $$ = $1.append(NewSymbolFromToken($3)) }

expr:
        TNumber                           { $$ = NewNumberFromString($1.Str) } |
        TString                           { $$ = NewString($1.Str) } |
        prefix_expr                       { $$ = $1 } |
        expr TOr expr                     { $$ = NewComplex(NewSymbol(AOr), $1,$3).SetPos($2.Pos) } |
        expr TAnd expr                    { $$ = NewComplex(NewSymbol(AAnd), $1,$3).SetPos($2.Pos) } |
        expr '>' expr                     { $$ = NewComplex(NewSymbol(ALess), $3,$1).SetPos($2.Pos) } |
        expr '<' expr                     { $$ = NewComplex(NewSymbol(ALess), $1,$3).SetPos($2.Pos) } |
        expr TGte expr                    { $$ = NewComplex(NewSymbol(ALessEq), $3,$1).SetPos($2.Pos) } |
        expr TLte expr                    { $$ = NewComplex(NewSymbol(ALessEq), $1,$3).SetPos($2.Pos) } |
        expr TEqeq expr                   { $$ = NewComplex(NewSymbol(AEq), $1,$3).SetPos($2.Pos) } |
        expr TNeq expr                    { $$ = NewComplex(NewSymbol(ANeq), $1,$3).SetPos($2.Pos) } |
        expr '+' expr                     { $$ = NewComplex(NewSymbol(AAdd), $1,$3).SetPos($2.Pos) } |
        expr TDotDot expr                 { $$ = NewComplex(NewSymbol(AConcat), $1,$3).SetPos($2.Pos) } |
        expr '-' expr                     { $$ = NewComplex(NewSymbol(ASub), $1,$3).SetPos($2.Pos) } |
        expr '*' expr                     { $$ = NewComplex(NewSymbol(AMul), $1,$3).SetPos($2.Pos) } |
        expr '/' expr                     { $$ = NewComplex(NewSymbol(ADiv), $1,$3).SetPos($2.Pos) } |
        expr TIDiv expr                   { $$ = NewComplex(NewSymbol(AIDiv), $1,$3).SetPos($2.Pos) } |
        expr '%' expr                     { $$ = NewComplex(NewSymbol(AMod), $1,$3).SetPos($2.Pos) } |
        expr '^' expr                     { $$ = NewComplex(NewSymbol(APow), $1,$3).SetPos($2.Pos) } |
        TNot expr %prec UNARY             { $$ = NewComplex(NewSymbol(ANot), $2).SetPos($1.Pos) } |
        '-' expr %prec UNARY              { $$ = NewComplex(NewSymbol(ASub), zeroNode, $2).SetPos($1.Pos) } |
        '#' expr %prec UNARY              { $$ = NewComplex(NewSymbol(ALen), $2).SetPos($1.Pos) }

prefix_expr:
        declarator                        { $$ = $1 } |
        prefix_expr TString               { $$ = __call($1, NewComplex(NewString($2.Str))).SetPos($1.Pos()) } |
        prefix_expr expr_list_paren       { $$ = __call($1, $2).SetPos($1.Pos()) } |
        '(' expr ')'                      { $$ = __chain($2) } | // shift/reduce conflict
        prefix_expr '(' expr_assign_list ')'                { $$ = __callMap($1, emptyNode, $3).SetPos($1.Pos()) } |
        prefix_expr '(' expr_assign_list',' ')'             { $$ = __callMap($1, emptyNode, $3).SetPos($1.Pos()) } |
        prefix_expr '(' expr_list','expr_assign_list ')'    { $$ = __callMap($1, $3, $5).SetPos($1.Pos()) } |
        prefix_expr '(' expr_list','expr_assign_list',' ')' { $$ = __callMap($1, $3, $5).SetPos($1.Pos()) }

expr_list:
        expr                              { $$ = NewComplex($1) } |
        expr_list ',' expr                { $$ = $1.append($3) }

expr_list_paren:
        '(' ')'                           { $$ = NewComplex() } |
        '(' expr_list ')'                 { $$ = $2 }

expr_assign_list:
        TIdent '=' expr                            { $$ = NewComplex(NewString($1.Str), $3) } |
        '[' expr ']' '=' expr                      { $$ = NewComplex($2, $5) } |
        expr_assign_list ',' TIdent '=' expr       { $$ = $1.append(NewString($3.Str)).append($5) } |
        expr_assign_list ',' '[' expr ']' '=' expr { $$ = $1.append($4).append($7) }

%%

