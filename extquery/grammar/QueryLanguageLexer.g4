// Inspired by the open-source lucene lexer.
// Copied at: 2022-07-16
// License at time of writing: MIT
// Link to source file: https://github.com/antlr/grammars-v4/blob/a771ddc7e659d7b31f74b44f2034b23816d90738/lucene/LuceneLexer.g4

lexer grammar QueryLanguageLexer;

AND : 'AND' | 'and' | '&&';

OR : 'OR' | 'or' | '||';

NOT : 'NOT' | 'not' | '!';

IS : 'IS' | 'is';

IN : 'IN' | 'in';

PRESENT: 'PRESENT' | 'present';

COUNT: 'COUNT' | 'count';

LPAREN : '(';

RPAREN : ')';

COMMA : ',';

OP_EQUAL : '=';

OP_NOT_EQUAL : '!=';

OP_EQUAL_IGNORE_CASE : '=*';

OP_NOT_EQUAL_IGNORE_CASE : '!=*';

OP_TILDE : '~';

OP_NOT_TILDE : '!~';

OP_TILDE_IGNORE_CASE : '~*';

OP_NOT_TILDE_IGNORE_CASE : '!~*';

OP_GREATER_THAN : '>';

OP_GREATER_THAN_EQUAL : '>=';

OP_LESS_THAN : '<';

OP_LESS_THAN_EQUAL : '<=';

INTEGER : '0' | ([1-9][0-9]*);

QUOTED : '"' QUOTED_CHAR* '"';

// Variable ({{key}}) and template-placeholder ([[key]]) markers as bare value tokens, so a query
// can reference them unquoted, e.g. `k8s.deployment IN ({{deployment}})`. The quoted form
// (`IN ("{{deployment}}")`) still lexes as QUOTED, so this is additive. The key charset mirrors the
// variable/placeholder key pattern used elsewhere ([\w-_.]+).
VARIABLE : '{{' MARKER_CHAR+ '}}';

PLACEHOLDER : '[[' MARKER_CHAR+ ']]';

TERM : TERM_START_CHAR TERM_CHAR*;

WHITESPACE : [ \t\n\r\u3000]+ -> channel(HIDDEN);

fragment QUOTED_CHAR : ~["\\] | ESCAPED_CHAR;

fragment TERM_START_CHAR
 : ~[ \t\n\r\u3000+\-!():^@<>=[\]"{}~\\/,] | ESCAPED_CHAR
 ;
fragment TERM_CHAR : ( TERM_START_CHAR | [\-+] );

fragment MARKER_CHAR : [a-zA-Z0-9_.\-];

fragment ESCAPED_CHAR : '\\' .;
