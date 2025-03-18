grammar Lispy;

// 词法规则
NUMBER : '-'? [0-9]+;
SYMBOL : '+' | '-' | '*' | '/';
WS : [ \t\n\r]+ -> skip;

number: NUMBER;
symbol: SYMBOL;
sexpr:  '(' expr* ')';
expr:   number | symbol | sexpr;
lispy : expr*;