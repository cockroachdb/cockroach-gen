// Code generated by "stringer -type=Key"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NamespaceTableWithSchemas-0]
	_ = x[Start20_2-1]
	_ = x[GeospatialType-2]
	_ = x[Enums-3]
	_ = x[RangefeedLeases-4]
	_ = x[AlterColumnTypeGeneral-5]
	_ = x[AlterSystemJobsAddCreatedByColumns-6]
	_ = x[AddScheduledJobsTable-7]
	_ = x[UserDefinedSchemas-8]
	_ = x[NoOriginFKIndexes-9]
	_ = x[NodeMembershipStatus-10]
	_ = x[MinPasswordLength-11]
	_ = x[AbortSpanBytes-12]
	_ = x[AlterSystemJobsAddSqllivenessColumnsAddNewSystemSqllivenessTable-13]
	_ = x[MaterializedViews-14]
	_ = x[Box2DType-15]
	_ = x[UpdateScheduledJobsSchema-16]
	_ = x[CreateLoginPrivilege-17]
	_ = x[HBAForNonTLS-18]
	_ = x[V20_2-19]
	_ = x[Start21_1-20]
	_ = x[EmptyArraysInInvertedIndexes-21]
	_ = x[UniqueWithoutIndexConstraints-22]
	_ = x[VirtualComputedColumns-23]
	_ = x[CPutInline-24]
	_ = x[ReplicaVersions-25]
	_ = x[replacedTruncatedAndRangeAppliedStateMigration-26]
	_ = x[replacedPostTruncatedAndRangeAppliedStateMigration-27]
	_ = x[NewSchemaChanger-28]
	_ = x[LongRunningMigrations-29]
	_ = x[TruncatedAndRangeAppliedStateMigration-30]
	_ = x[PostTruncatedAndRangeAppliedStateMigration-31]
	_ = x[SeparatedIntents-32]
	_ = x[TracingVerbosityIndependentSemantics-33]
	_ = x[SequencesRegclass-34]
	_ = x[ImplicitColumnPartitioning-35]
	_ = x[MultiRegionFeatures-36]
}

const _Key_name = "NamespaceTableWithSchemasStart20_2GeospatialTypeEnumsRangefeedLeasesAlterColumnTypeGeneralAlterSystemJobsAddCreatedByColumnsAddScheduledJobsTableUserDefinedSchemasNoOriginFKIndexesNodeMembershipStatusMinPasswordLengthAbortSpanBytesAlterSystemJobsAddSqllivenessColumnsAddNewSystemSqllivenessTableMaterializedViewsBox2DTypeUpdateScheduledJobsSchemaCreateLoginPrivilegeHBAForNonTLSV20_2Start21_1EmptyArraysInInvertedIndexesUniqueWithoutIndexConstraintsVirtualComputedColumnsCPutInlineReplicaVersionsreplacedTruncatedAndRangeAppliedStateMigrationreplacedPostTruncatedAndRangeAppliedStateMigrationNewSchemaChangerLongRunningMigrationsTruncatedAndRangeAppliedStateMigrationPostTruncatedAndRangeAppliedStateMigrationSeparatedIntentsTracingVerbosityIndependentSemanticsSequencesRegclassImplicitColumnPartitioningMultiRegionFeatures"

var _Key_index = [...]uint16{0, 25, 34, 48, 53, 68, 90, 124, 145, 163, 180, 200, 217, 231, 295, 312, 321, 346, 366, 378, 383, 392, 420, 449, 471, 481, 496, 542, 592, 608, 629, 667, 709, 725, 761, 778, 804, 823}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
