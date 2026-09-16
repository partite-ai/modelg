lexer grammar SQLQueryLexer;

options {
	language = Go;
}

fragment LINEWS: [ \t]+;
fragment NL: '\r'? '\n';

fragment STRING_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'';

fragment QUOTED_IDENTIFIER: '"' ('""' | ~ [\u0000"])* '"';

STARTCOMMENT: '--' -> skip, pushMode(Comment);
STARTQUERY: '--' LINEWS* 'name:' LINEWS* -> pushMode(Query);

STARTWHEN: '--when';
STARTCHOMPWHEN: '--<when';
STARTJOINWHEN: '--+when';
ENDWHEN: '--endwhen' LINEWS* NL;

STARTEXPR: ':' -> pushMode(Expr);

EOL: NL;

WS: [ \t\r]+ -> skip;

SQLTEXT:
	STRING_LITERAL
	| QUOTED_IDENTIFIER
	| (~[ \t\r\n\-:'"] | (':' ~[a-zA-Z_]) | ('-' ~'-'))+;

mode Comment;
RAWCOMMENT: ~[\r\n] -> skip;
ENDCOMMENT: NL -> skip, popMode;

mode Query;
QUERY_NAME: [a-zA-Z_][a-zA-Z0-9_]*;
ENDQUERY: LINEWS* NL -> skip, popMode;

mode Expr;

IDENT: [a-zA-Z_][a-zA-Z0-9_]*;

IDENT_SEP: '.';

IDENT_LITERAL: '!';

WS_EXPR: WS -> skip, popMode;

EOL_EXPR: NL -> type(EOL), popMode;

SQLTEXT_EXPR: (STRING_LITERAL | QUOTED_IDENTIFIER | .) -> type(SQLTEXT), popMode;