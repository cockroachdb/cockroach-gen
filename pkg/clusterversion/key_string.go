// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_1-0]
	_ = x[Start21_1PLUS-1]
	_ = x[Start21_2-2]
	_ = x[SpanConfigurationsTable-3]
	_ = x[BoundedStaleness-4]
	_ = x[DateAndIntervalStyle-5]
	_ = x[TenantUsageSingleConsumptionColumn-6]
	_ = x[SQLStatsTables-7]
	_ = x[SQLStatsCompactionScheduledJob-8]
	_ = x[V21_2-9]
	_ = x[Start22_1-10]
	_ = x[AvoidDrainingNames-11]
	_ = x[DrainingNamesMigration-12]
	_ = x[TraceIDDoesntImplyStructuredRecording-13]
	_ = x[AlterSystemTableStatisticsAddAvgSizeCol-14]
	_ = x[AlterSystemStmtDiagReqs-15]
	_ = x[MVCCAddSSTable-16]
	_ = x[InsertPublicSchemaNamespaceEntryOnRestore-17]
	_ = x[UnsplitRangesInAsyncGCJobs-18]
	_ = x[ValidateGrantOption-19]
	_ = x[PebbleFormatBlockPropertyCollector-20]
	_ = x[ProbeRequest-21]
	_ = x[SelectRPCsTakeTracingInfoInband-22]
	_ = x[PreSeedTenantSpanConfigs-23]
	_ = x[SeedTenantSpanConfigs-24]
	_ = x[PublicSchemasWithDescriptors-25]
}

const _Key_name = "V21_1Start21_1PLUSStart21_2SpanConfigurationsTableBoundedStalenessDateAndIntervalStyleTenantUsageSingleConsumptionColumnSQLStatsTablesSQLStatsCompactionScheduledJobV21_2Start22_1AvoidDrainingNamesDrainingNamesMigrationTraceIDDoesntImplyStructuredRecordingAlterSystemTableStatisticsAddAvgSizeColAlterSystemStmtDiagReqsMVCCAddSSTableInsertPublicSchemaNamespaceEntryOnRestoreUnsplitRangesInAsyncGCJobsValidateGrantOptionPebbleFormatBlockPropertyCollectorProbeRequestSelectRPCsTakeTracingInfoInbandPreSeedTenantSpanConfigsSeedTenantSpanConfigsPublicSchemasWithDescriptors"

var _Key_index = [...]uint16{0, 5, 18, 27, 50, 66, 86, 120, 134, 164, 169, 178, 196, 218, 255, 294, 317, 331, 372, 398, 417, 451, 463, 494, 518, 539, 567}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
