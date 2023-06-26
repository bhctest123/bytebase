package mysql

// Framework code is generated by the generator.

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
)

var (
	_ advisor.Advisor = (*StatementDmlDryRunAdvisor)(nil)
	_ ast.Visitor     = (*statementDmlDryRunChecker)(nil)
)

func init() {
	advisor.Register(db.MySQL, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
	advisor.Register(db.TiDB, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
	advisor.Register(db.MariaDB, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
	advisor.Register(db.OceanBase, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
}

// StatementDmlDryRunAdvisor is the advisor checking for DML dry run.
type StatementDmlDryRunAdvisor struct {
}

// Check checks for DML dry run.
func (*StatementDmlDryRunAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementDmlDryRunChecker{
		level:  level,
		title:  string(ctx.Rule.Type),
		driver: ctx.Driver,
		ctx:    ctx.Context,
	}

	if checker.driver != nil {
		for _, stmt := range stmtList {
			checker.text = stmt.Text()
			checker.line = stmt.OriginTextPosition()
			(stmt).Accept(checker)
		}
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type statementDmlDryRunChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	text       string
	line       int
	driver     *sql.DB
	ctx        context.Context
}

// Enter implements the ast.Visitor interface.
func (checker *statementDmlDryRunChecker) Enter(in ast.Node) (ast.Node, bool) {
	switch node := in.(type) {
	case *ast.InsertStmt, *ast.UpdateStmt, *ast.DeleteStmt:
		if _, err := advisor.Query(checker.ctx, checker.driver, fmt.Sprintf("EXPLAIN %s", node.Text())); err != nil {
			checker.adviceList = append(checker.adviceList, advisor.Advice{
				Status:  checker.level,
				Code:    advisor.StatementDMLDryRunFailed,
				Title:   checker.title,
				Content: fmt.Sprintf("\"%s\" dry runs failed: %s", node.Text(), err.Error()),
				Line:    checker.line,
			})
		}
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*statementDmlDryRunChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
