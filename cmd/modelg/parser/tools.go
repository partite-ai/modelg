package parser

//go:generate antlr -Dlanguage=Go -o generated -package generated -visitor SQLQueryLexer.g4 SQLQueryParser.g4
