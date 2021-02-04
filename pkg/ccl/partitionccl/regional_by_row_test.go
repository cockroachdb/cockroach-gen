// Copyright 2021 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package partitionccl_test

import (
	"context"
	gosql "database/sql"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/jobs"
	"github.com/cockroachdb/cockroach/pkg/keys"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/catalogkv"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
	"github.com/cockroachdb/cockroach/pkg/sql/sqltestutils"
	"github.com/cockroachdb/cockroach/pkg/sql/tests"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
)

// REGIONAL BY ROW tests are defined in partitionccl as REGIONAL BY ROW
// requires CCL to operate.

// TestAlterToRegionalByRowChangeWithCancel tests an alteration to REGIONAL
// BY ROW AS <col> which gets canceled.
func TestAlterToRegionalByRowAsChangeWithCancel(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	// Decrease the adopt loop interval so that retries happen quickly.
	defer sqltestutils.SetTestJobsAdoptInterval()()

	var chunkSize int64 = 100
	var maxValue = 4000
	if util.RaceEnabled {
		// Race builds are a lot slower, so use a smaller number of rows.
		maxValue = 200
		chunkSize = 5
	}

	ctx := context.Background()
	var db *gosql.DB
	shouldCancel := false
	params, _ := tests.CreateTestServerParams()
	params.Locality.Tiers = []roachpb.Tier{
		{Key: "region", Value: "ajstorm-1"},
	}
	params.Knobs = base.TestingKnobs{
		SQLSchemaChanger: &sql.SchemaChangerTestingKnobs{
			BackfillChunkSize: chunkSize,
		},
		DistSQL: &execinfra.TestingKnobs{
			RunBeforeBackfillChunk: func(sp roachpb.Span) error {
				if !shouldCancel {
					return nil
				}
				if _, err := db.Exec(`CANCEL JOB (
					SELECT job_id FROM [SHOW JOBS]
					WHERE
						job_type = 'SCHEMA CHANGE' AND
						status = $1 AND
						description NOT LIKE 'ROLL BACK%'
				)`, jobs.StatusRunning); err != nil {
					t.Error(err)
				}
				return nil
			},
		},
	}
	s, sqlDB, kvDB := serverutils.StartServer(t, params)
	db = sqlDB
	defer s.Stopper().Stop(ctx)

	// Disable strict GC TTL enforcement because we're going to shove a zero-value
	// TTL into the system with AddImmediateGCZoneConfig.
	defer sqltestutils.DisableGCTTLStrictEnforcement(t, sqlDB)()

	if _, err := sqlDB.Exec(`
CREATE DATABASE t PRIMARY REGION "ajstorm-1";
CREATE TABLE t.test (k INT NOT NULL, v INT) LOCALITY GLOBAL;
`); err != nil {
		t.Fatal(err)
	}

	if err := sqltestutils.BulkInsertIntoTable(sqlDB, maxValue); err != nil {
		t.Fatal(err)
	}

	if _, err := sqlDB.Exec(`
		ALTER TABLE t.test ADD COLUMN cr t.crdb_internal_region
		NOT NULL
		DEFAULT gateway_region()::t.crdb_internal_region
	`); err != nil {
		t.Fatal(err)
	}
	// This will fail, so we don't want to check the error.
	shouldCancel = true
	_, _ = sqlDB.Exec(`ALTER TABLE t.test SET LOCALITY REGIONAL BY ROW AS cr`)

	// Ensure that the mutations corresponding to the primary key change are cleaned up and
	// that the job did not succeed even though it was canceled.
	testutils.SucceedsSoon(t, func() error {
		tableDesc := catalogkv.TestingGetTableDescriptor(kvDB, keys.SystemSQLCodec, "t", "test")
		if len(tableDesc.GetMutations()) != 0 {
			return errors.Errorf("expected 0 mutations after cancellation, found %d", len(tableDesc.GetMutations()))
		}
		if tableDesc.GetPrimaryIndex().NumColumns() != 1 || tableDesc.GetPrimaryIndex().GetColumnName(0) != "rowid" {
			return errors.Errorf("expected primary key change to not succeed after cancellation")
		}
		if !tableDesc.IsLocalityGlobal() {
			return errors.Errorf("expected locality to be global")
		}

		zoneConfigRow := sqlDB.QueryRow("SHOW ZONE CONFIGURATION FROM TABLE t.test")
		var target string
		var rawZoneSQL string
		require.NoError(t, zoneConfigRow.Scan(&target, &rawZoneSQL))
		const expectedRawZoneSQL = `ALTER TABLE t.public.test CONFIGURE ZONE USING
	range_min_bytes = 134217728,
	range_max_bytes = 536870912,
	gc.ttlseconds = 90000,
	global_reads = true,
	num_replicas = 3,
	num_voters = 3,
	constraints = '{+region=ajstorm-1: 1}',
	voter_constraints = '[+region=ajstorm-1]',
	lease_preferences = '[[+region=ajstorm-1]]'`
		if target != "DATABASE t" && rawZoneSQL != expectedRawZoneSQL {
			return errors.Errorf(
				"expected zone configuration to not have changed, got %s, sql %s",
				target,
				rawZoneSQL,
			)
		}
		return nil
	})

	tableDesc := catalogkv.TestingGetTableDescriptor(kvDB, keys.SystemSQLCodec, "t", "test")
	if _, err := sqltestutils.AddImmediateGCZoneConfig(db, tableDesc.GetID()); err != nil {
		t.Fatal(err)
	}
	// Ensure that the writes from the partial new indexes are cleaned up.
	testutils.SucceedsSoon(t, func() error {
		return sqltestutils.CheckTableKeyCount(ctx, kvDB, 1, maxValue)
	})
}
