// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package schemaexpr

import (
	"bytes"
	"context"
	"fmt"
	"sort"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

// CheckConstraintBuilder creates sqlbase.TableDescriptor_CheckConstraints from
// tree.CheckConstraintTableDefs. See Build for more details.
type CheckConstraintBuilder struct {
	ctx       context.Context
	tableName tree.TableName
	desc      *sqlbase.MutableTableDescriptor
	semaCtx   *tree.SemaContext

	// inUseNames keeps track of names that have been generated previously
	// by generateName or marked as in-use by MarkNameInUse. It is used to
	// prevent the CheckConstraintBuilder from generating constraints with
	// non-unique names.
	inUseNames map[string]struct{}
}

// NewCheckConstraintBuilder returns a CheckConstraintBuilder struct that can
// be used to build sqlbase.TableDescriptor_CheckConstraints. See Build for more
// details.
func NewCheckConstraintBuilder(
	ctx context.Context,
	tableName tree.TableName,
	desc *sqlbase.MutableTableDescriptor,
	semaCtx *tree.SemaContext,
) CheckConstraintBuilder {
	return CheckConstraintBuilder{
		ctx:        ctx,
		tableName:  tableName,
		desc:       desc,
		semaCtx:    semaCtx,
		inUseNames: make(map[string]struct{}),
	}
}

// MarkNameInUse marks a name as in-use, preventing subsequent calls to Build
// from generating a check constraint with the same name.
func (b *CheckConstraintBuilder) MarkNameInUse(name string) {
	b.inUseNames[name] = struct{}{}
}

// Build validates the input tree.CheckConstraintTableDef and, if valid, returns
// a sqlbase.TableDescriptor_CheckConstraint. If the input constraint does not
// have a name, Build generates a name based on the variables referenced in the
// check expression.
//
// A check constraint expression is valid if the following are true:
//
//   - It results in a boolean.
//   - It refers only to columns in the table.
//   - It does not include subqueries.
//   - It does not include aggregate, window, or set returning functions.
//
// Note that mutable functions are currently allowed, but using them can lead
// to unexpected behavior.
// TODO(mgartner): Add unit tests for Build. Some cases should test
// MarkNameInUse, too.
func (b *CheckConstraintBuilder) Build(
	c *tree.CheckConstraintTableDef,
) (*sqlbase.TableDescriptor_CheckConstraint, error) {
	name := string(c.Name)

	if name == "" {
		var err error
		name, err = b.generateUniqueName(c.Expr)
		if err != nil {
			return nil, err
		}
	}

	// Replace the column variables with dummyColumns so that they can be
	// type-checked.
	replacedExpr, colIDsUsed, err := replaceVars(b.desc, c.Expr)
	if err != nil {
		return nil, err
	}

	// Verify that the expression results in a boolean and does not use
	// invalid functions.
	_, err = sqlbase.SanitizeVarFreeExpr(
		b.ctx,
		replacedExpr,
		types.Bool,
		"CHECK",
		b.semaCtx,
		true, /* allowImpure */
	)
	if err != nil {
		return nil, err
	}

	// Collect and sort the column IDs referenced in the check expression.
	colIDs := make(sqlbase.ColumnIDs, 0, len(colIDsUsed))
	for colID := range colIDsUsed {
		colIDs = append(colIDs, colID)
	}
	sort.Sort(colIDs)

	source := sqlbase.NewSourceInfoForSingleTable(
		b.tableName, sqlbase.ResultColumnsFromColDescs(
			b.desc.GetID(),
			b.desc.TableDesc().AllNonDropColumns(),
		),
	)

	// Strip the database and table names from any qualified column names.
	expr, err := DequalifyColumnRefs(b.ctx, source, c.Expr)
	if err != nil {
		return nil, err
	}

	return &sqlbase.TableDescriptor_CheckConstraint{
		Expr:      tree.Serialize(expr),
		Name:      name,
		ColumnIDs: colIDs,
		Hidden:    c.Hidden,
	}, nil
}

// nameInUse returns true if the given name is already inuse.
func (b *CheckConstraintBuilder) nameInUse(name string) bool {
	_, ok := b.inUseNames[name]
	return ok
}

// generateUniqueName returns a name for the given check constraint expression
// that is unique among other names generated by the CheckConstraintBuilder. It
// is also unique from any constraint names marked as in-use via MarkNameInUse.
func (b *CheckConstraintBuilder) generateUniqueName(expr tree.Expr) (string, error) {
	name, err := b.DefaultName(expr)
	if err != nil {
		return "", err
	}

	// If the generated name isn't unique, add a number to the end. Increment
	// the number until a unique name is found.
	if b.nameInUse(name) {
		i := 1
		for {
			numberedName := fmt.Sprintf("%s%d", name, i)
			if !b.nameInUse(numberedName) {
				name = numberedName
				break
			}
			i++
		}
	}

	b.MarkNameInUse(name)

	return name, nil
}

// DefaultName creates a check constraint name based on the columns referenced
// in the check expression. The format is "check_col1_col2...". If columns are
// duplicated in the expression, they will be duplicated in the name.
//
// For example:
//
//   CHECK (a < 0) => check_a
//   CHECK (a < 0 AND b = 'foo') => check_a_b
//   CHECK (a < 0 AND b = 'foo' AND a < 10) => check_a_b_a
//
// Note that the generated name is not guaranteed to be unique among the other
// constraints of the table.
// TODO(mgartner): Add unit tests for DefaultName.
func (b *CheckConstraintBuilder) DefaultName(expr tree.Expr) (string, error) {
	var nameBuf bytes.Buffer
	nameBuf.WriteString("check")

	err := iterColDescriptors(b.desc, expr, func(c *sqlbase.ColumnDescriptor) error {
		nameBuf.WriteByte('_')
		nameBuf.WriteString(c.Name)
		return nil
	})
	if err != nil {
		return "", err
	}

	return nameBuf.String(), nil
}
