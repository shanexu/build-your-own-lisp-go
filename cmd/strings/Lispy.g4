grammar Lispy;

NUMBER : '-'? [0-9]+;
SYMBOL : [a-zA-Z0-9_+\\\-*/=<>!&]+;
STRING : '"' ('\\"'|.)*? '"';
COMMENT: ';' [^\r\n]*;
WS     : [ \t\n\r]+ -> skip;

number : NUMBER;
symbol : SYMBOL;
string : STRING;
comment: COMMENT;
sexpr  : '(' expr* ')';
qexpr  : '{' expr* '}';
expr   : number | symbol | string | comment | sexpr | qexpr;
lispy  : expr*;