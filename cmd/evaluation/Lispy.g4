grammar Lispy;

// 词法规则
NUMBER : '-'? [0-9]+;
OPERATOR : '+' | '-' | '*' | '/';
WS : [ \t\n\r]+ -> skip;

number: NUMBER;
operator: OPERATOR;

expr
    : number
    | '(' operator expr+ ')'
    ;

lispy : '(' operator expr+ ')';