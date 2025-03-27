grammar Lispy;

NUMBER : '-'? [0-9]+;
SYMBOL : [a-zA-Z0-9_+\\\-*/=<>!&]+;
WS     : [ \t\n\r]+ -> skip;

number : NUMBER;
symbol : SYMBOL;
sexpr  : '(' expr* ')';
qexpr  : '{' expr* '}';
expr   : number | symbol | sexpr | qexpr;
lispy  : expr*;