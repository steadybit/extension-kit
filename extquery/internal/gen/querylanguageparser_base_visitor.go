// Code generated from QueryLanguageParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // QueryLanguageParser
import "github.com/antlr4-go/antlr/v4"

type BaseQueryLanguageParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryLanguageParserVisitor) VisitTopLevelQuery(ctx *TopLevelQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitClause(ctx *ClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitIsPresentClause(ctx *IsPresentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitInClause(ctx *InClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitCompareClause(ctx *CompareClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitCountClause(ctx *CountClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitValuelist(ctx *ValuelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageParserVisitor) VisitInt_value(ctx *Int_valueContext) interface{} {
	return v.VisitChildren(ctx)
}
