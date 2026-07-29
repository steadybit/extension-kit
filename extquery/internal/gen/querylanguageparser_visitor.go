// Code generated from QueryLanguageParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // QueryLanguageParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by QueryLanguageParser.
type QueryLanguageParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryLanguageParser#topLevelQuery.
	VisitTopLevelQuery(ctx *TopLevelQueryContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#clause.
	VisitClause(ctx *ClauseContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#isPresentClause.
	VisitIsPresentClause(ctx *IsPresentClauseContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#inClause.
	VisitInClause(ctx *InClauseContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#compareClause.
	VisitCompareClause(ctx *CompareClauseContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#countClause.
	VisitCountClause(ctx *CountClauseContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#valuelist.
	VisitValuelist(ctx *ValuelistContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#int_value.
	VisitInt_value(ctx *Int_valueContext) interface{}
}
