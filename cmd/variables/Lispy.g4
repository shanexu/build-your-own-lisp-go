grammar Lispy;

// 词法规则
NUMBER : '-'? [0-9]+;
SYMBOL : 'list' | 'head' | 'tail' | 'eval'
       | 'join' | '+' | '-' | '*' | '/';
WS : [ \t\n\r]+ -> skip;

number: NUMBER;
symbol: SYMBOL;
sexpr:  '(' expr* ')';
qexpr:  '{' expr* '}';
expr:   number | symbol | sexpr | qexpr;
lispy : expr*;