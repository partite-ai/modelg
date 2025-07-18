lexer grammar SQLQueryLexer;

options {
	language = Go;
}

fragment LINEWS: [ \t]+;
fragment NL: '\r'? '\n';

fragment STRING_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'';

fragment QUOTED_IDENTIFIER: '"' ('""' | ~ [\u0000"])* '"';

fragment DIGITS: [0-9]+;

fragment NUMBER:
  DIGITS
	| DIGITS '.' DIGITS? ('E' [+-]? DIGITS)?
	| '.' DIGITS ('E' [+-]? DIGITS)?
	| DIGITS 'E' [+-]? DIGITS;

STARTCOMMENT: '--' -> skip, pushMode(CommentStart);

STARTEXPR: ':' -> pushMode(Expr);

EOL: NL;

WS: [ \t\r]+ -> skip;

SQLTEXT:
	STRING_LITERAL
	| QUOTED_IDENTIFIER
	| '-' NUMBER
  | '-'
	| ~[ \t\r\n:\-'"]+;

mode CommentStart;

STARTQUERY: LINEWS* 'name:' LINEWS* -> mode(Query);

STARTWHEN: 'when' -> popMode;

STARTCHOMPWHEN: '<when' -> popMode;

STARTJOINWHEN: '+when' -> popMode;

ENDWHEN: 'endwhen' LINEWS* NL -> popMode;

ENDCOMMENT: NL -> skip, popMode;

RAWCOMMENT: . -> skip;

mode Query;

QUERY_NAME: [a-zA-Z_][a-zA-Z0-9_]*;

ENDQUERY: LINEWS* NL -> skip, popMode;

mode Expr;

IDENT: [a-zA-Z_][a-zA-Z0-9_]*;

IDENT_SEP: '.';

IDENT_LITERAL: '!';

SPACE_EXPR: [ \t] -> skip, popMode;

WS_EXPR: WS -> skip, popMode;

EOL_EXPR: NL -> type(EOL), popMode;

SQLTEXT_EXPR: (
		STRING_LITERAL
		| QUOTED_IDENTIFIER
    | '-'
		| ~[ \t\r\n:\-'"@]
	) -> type(SQLTEXT), popMode;
