// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package delegate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompletions(t *testing.T) {
	tests := []struct {
		stmt                string
		offset              int
		expectedCompletions []string
	}{
		{
			stmt:                "creat",
			expectedCompletions: []string{"CREATE", "CREATEDB", "CREATELOGIN", "CREATEROLE"},
		},
		{
			stmt:                "CREAT",
			expectedCompletions: []string{"CREATE", "CREATEDB", "CREATELOGIN", "CREATEROLE"},
		},
		{
			stmt:                "creat ",
			expectedCompletions: nil,
		},
		{
			stmt:                "SHOW CREAT",
			expectedCompletions: []string{"CREATE", "CREATEDB", "CREATELOGIN", "CREATEROLE"},
		},
		{
			stmt:                "show creat",
			expectedCompletions: []string{"CREATE", "CREATEDB", "CREATELOGIN", "CREATEROLE"},
		},
		{
			stmt: "se",
			expectedCompletions: []string{
				"SEARCH", "SECOND", "SECONDARY", "SECURITY", "SELECT", "SEQUENCE", "SEQUENCES",
				"SERIALIZABLE", "SERVER", "SESSION", "SESSIONS", "SESSION_USER",
				"SET", "SETOF", "SETS", "SETTING", "SETTINGS",
			},
		},
		{
			stmt:                "sel",
			expectedCompletions: []string{"SELECT"},
		},
		{
			stmt:                "create ta",
			expectedCompletions: []string{"TABLE", "TABLES", "TABLESPACE"},
		},
		{
			stmt:                "create ta",
			offset:              3,
			expectedCompletions: []string{"CREATE"},
		},
		{
			stmt:                "select",
			offset:              2,
			expectedCompletions: []string{"SELECT"},
		},
		{
			stmt:                "select ",
			offset:              7,
			expectedCompletions: nil,
		},
		{
			stmt:                "你好，我的名字是鲍勃 SELECT",
			offset:              2,
			expectedCompletions: []string{"你好，我的名字是鲍勃"},
		},
		{
			stmt:                "你好，我的名字是鲍勃 SELECT",
			offset:              11,
			expectedCompletions: nil,
		},
		{
			stmt:                "你好，我的名字是鲍勃 SELECT",
			offset:              12,
			expectedCompletions: []string{"SELECT"},
		},
		{
			stmt:                "😋😋😋 😋😋😋",
			offset:              25,
			expectedCompletions: nil,
		},
		{
			stmt:                "Jalapeño",
			offset:              9,
			expectedCompletions: nil,
		},
		// Test an offset greater than the number of runes in the string, but less
		// than or equal to the number of bytes in the string.
		// 🦹 has 1 rune and 4 bytes.
		{
			stmt:   "🦹",
			offset: 2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.stmt, func(t *testing.T) {
			offset := tc.offset
			if tc.offset == 0 {
				offset = len([]rune(tc.stmt))
			}
			actualCompletions, err := RunShowCompletions(tc.stmt, offset)
			require.NoError(t, err)
			require.Equal(t, tc.expectedCompletions, actualCompletions)
		})
	}
}
