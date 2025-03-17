grammar Lispy;

// 词法规则
NUMBER : '-'? [0-9]+;
OPERATOR : '+' | '-' | '*' | '/';
WS : [ \t\n\r]+ -> skip;

expr
    : NUMBER
    | '(' OPERATOR expr+ ')'
    ;

program : OPERATOR expr+ EOF;