// Inspired by the open-source lucene parser.
// Copied at: 2022-07-16
// License at time of writing: MIT
// Link to source file: https://github.com/antlr/grammars-v4/blob/a771ddc7e659d7b31f74b44f2034b23816d90738/lucene/LuceneParser.g4

parser grammar QueryLanguageParser;

options { tokenVocab=QueryLanguageLexer; }

topLevelQuery
 : query? EOF
 ;

query
 : clause
 | LPAREN query RPAREN
 | NOT query
 | query AND query
 | query OR query
 ;

clause
 : isPresentClause
 | inClause
 | compareClause
 | countClause
 ;

isPresentClause
 : fieldName IS NOT? PRESENT
 ;

inClause
 : fieldName NOT? IN LPAREN valuelist RPAREN
 ;

compareClause
 :  fieldName op=(OP_TILDE | OP_NOT_TILDE | OP_EQUAL | OP_NOT_EQUAL | OP_TILDE_IGNORE_CASE | OP_NOT_TILDE_IGNORE_CASE | OP_EQUAL_IGNORE_CASE | OP_NOT_EQUAL_IGNORE_CASE) value
 ;

countClause
 :  COUNT LPAREN fieldName RPAREN op=(OP_EQUAL | OP_NOT_EQUAL | OP_GREATER_THAN | OP_GREATER_THAN_EQUAL | OP_LESS_THAN | OP_LESS_THAN_EQUAL) int_value
 ;

valuelist
 : value (COMMA value)*
 ;

value
 : QUOTED | TERM | INTEGER | VARIABLE | PLACEHOLDER
 ;

fieldName
 : QUOTED | TERM
 ;

int_value
 : INTEGER
 ;