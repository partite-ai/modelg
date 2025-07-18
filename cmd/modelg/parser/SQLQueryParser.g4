parser grammar SQLQueryParser;

options {
  language = Go;
  tokenVocab = SQLQueryLexer;
}

querySet
  : query+ EOF
  ;

query
  : STARTQUERY queryName=QUERY_NAME queryBody=sqlChunk
  ;

sqlChunk
  : (
      sqlLine
      | sqlFragment* whenDirective
      | chompWhenDirective
  )+
  ;

sqlLine
  : sqlFragment* EOL
  ;

sqlFragment
  : (sqlText | expr)
  ;

sqlText
  : textFragments+=SQLTEXT+
  ;

whenDirective
  : STARTWHEN whenCondition=expr EOL body=sqlChunk joinWhens+=joinWhen* ENDWHEN
  ;

chompWhenDirective
  : prefix=sqlFragment+ STARTCHOMPWHEN whenCondition=expr EOL body=sqlChunk joinWhens+=joinWhen* ENDWHEN
  ;

joinWhen
  : prefix=sqlFragment+ STARTJOINWHEN whenCondition=expr EOL body=sqlChunk;


expr
  : STARTEXPR path+=IDENT (IDENT_SEP path+=IDENT)* (literalFlag=IDENT_LITERAL literalMode=IDENT?)?
  ;