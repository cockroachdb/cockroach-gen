// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by generate_visitor.go. DO NOT EDIT.

package scop

import "context"

// ValidationOp is an operation which can be visited by ValidationVisitor.
type ValidationOp interface {
	Op
	Visit(context.Context, ValidationVisitor) error
}

// ValidationVisitor is a visitor for ValidationOp operations.
type ValidationVisitor interface {
	ValidateIndex(context.Context, ValidateIndex) error
	ValidateConstraint(context.Context, ValidateConstraint) error
}

// Visit is part of the ValidationOp interface.
func (op ValidateIndex) Visit(ctx context.Context, v ValidationVisitor) error {
	return v.ValidateIndex(ctx, op)
}

// Visit is part of the ValidationOp interface.
func (op ValidateConstraint) Visit(ctx context.Context, v ValidationVisitor) error {
	return v.ValidateConstraint(ctx, op)
}
