// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package exec

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/execgen"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

// Use execgen package to remove unused import warning.
var _ interface{} = execgen.UNSAFEGET

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	rows sqlbase.EncDatumRows,
	vec coldata.Vec,
	columnIdx int,
	columnType *types.T,
	alloc *sqlbase.DatumAlloc,
) error {

	switch columnType.Family() {
	case types.DateFamily:

		col := vec.Int64()
		datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}

				castV := v.(int64)
				col[i] = castV
			}
		}
	case types.OidFamily:

		col := vec.Int64()
		datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}

				castV := v.(int64)
				col[i] = castV
			}
		}
	case types.BoolFamily:

		col := vec.Bool()
		datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}

				castV := v.(bool)
				col[i] = castV
			}
		}
	case types.BytesFamily:

		col := vec.Bytes()
		datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}

				castV := v.([]byte)
				col.Set(i, castV)
			}
		}
	case types.IntFamily:
		switch columnType.Width() {
		case 16:

			col := vec.Int16()
			datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}

					castV := v.(int16)
					col[i] = castV
				}
			}
		case 32:

			col := vec.Int32()
			datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}

					castV := v.(int32)
					col[i] = castV
				}
			}
		case 64:

			col := vec.Int64()
			datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}

					castV := v.(int64)
					col[i] = castV
				}
			}
		default:
			execerror.VectorizedInternalPanic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width(), columnType.String()))
		}
	case types.DecimalFamily:

		col := vec.Decimal()
		datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}

				castV := v.(apd.Decimal)
				col[i] = castV
			}
		}
	case types.StringFamily:

		col := vec.Bytes()
		datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
		for i := range rows {
			row := rows[i]
			if row[columnIdx].Datum == nil {
				if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := row[columnIdx].Datum
			if datum == tree.DNull {
				vec.Nulls().SetNull(uint16(i))
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}

				castV := v.([]byte)
				col.Set(i, castV)
			}
		}
	case types.FloatFamily:
		switch columnType.Width() {
		case 64:

			col := vec.Float64()
			datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}

					castV := v.(float64)
					col[i] = castV
				}
			}
		case 32:

			col := vec.Float64()
			datumToPhysicalFn := typeconv.GetDatumToPhysicalFn(columnType)
			for i := range rows {
				row := rows[i]
				if row[columnIdx].Datum == nil {
					if err := row[columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := row[columnIdx].Datum
				if datum == tree.DNull {
					vec.Nulls().SetNull(uint16(i))
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}

					castV := v.(float64)
					col[i] = castV
				}
			}
		default:
			execerror.VectorizedInternalPanic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width(), columnType.String()))
		}
	default:
		execerror.VectorizedInternalPanic(fmt.Sprintf("unsupported column type %s", columnType.String()))
	}
	return nil
}
