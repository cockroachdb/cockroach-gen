// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package lang

// The Optgen language compiler uses itself to generate its AST expressions.
// This is a form of compiler bootstrapping:
//   https://en.wikipedia.org/wiki/Bootstrapping_(compilers)
//
// In order to regenerate expr.og.go or operator.og.go, perform these steps
// from the optgen directory (and make sure that $GOROOT/bin is in your path):
//
//   cd cmd/langgen
//   go build && go install
//   cd ../../lang
//   go generate

//go:generate langgen -out expr.og.go exprs lang.opt
//go:generate langgen -out operator.og.go ops lang.opt
//go:generate stringer -type=Operator operator.og.go
