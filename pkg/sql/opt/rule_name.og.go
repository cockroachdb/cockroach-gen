// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateAggDistinct
	NormalizeNestedAnds
	SimplifyTrueAnd
	SimplifyAndTrue
	SimplifyFalseAnd
	SimplifyAndFalse
	SimplifyTrueOr
	SimplifyOrTrue
	SimplifyFalseOr
	SimplifyOrFalse
	SimplifyRange
	FoldNullAndOr
	FoldNotTrue
	FoldNotFalse
	FoldNotNull
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	ExtractRedundantConjunct
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	FoldIsNull
	FoldNonNullIsNull
	FoldNullTupleIsTupleNull
	FoldNonNullTupleIsTupleNull
	FoldIsNotNull
	FoldNonNullIsNotNull
	FoldNonNullTupleIsTupleNotNull
	FoldNullTupleIsTupleNotNull
	CommuteNullIs
	NormalizeCmpTimeZoneFunction
	NormalizeCmpTimeZoneFunctionTZ
	FoldEqZeroSTDistance
	FoldCmpSTDistanceLeft
	FoldCmpSTDistanceRight
	FoldCmpSTMaxDistanceLeft
	FoldCmpSTMaxDistanceRight
	FoldEqTrue
	FoldEqFalse
	FoldNeTrue
	FoldNeFalse
	NormCycleTestRelTrueToFalse
	NormCycleTestRelFalseToTrue
	DecorrelateJoin
	DecorrelateProjectSet
	TryDecorrelateSelect
	TryDecorrelateProject
	TryDecorrelateProjectSelect
	TryDecorrelateProjectInnerJoin
	TryDecorrelateInnerJoin
	TryDecorrelateInnerLeftJoin
	TryDecorrelateGroupBy
	TryDecorrelateScalarGroupBy
	TryDecorrelateSemiJoin
	TryDecorrelateLimitOne
	TryDecorrelateLimit
	TryDecorrelateProjectSet
	TryDecorrelateWindow
	TryDecorrelateMax1Row
	HoistSelectExists
	HoistSelectNotExists
	HoistSelectSubquery
	HoistProjectSubquery
	HoistJoinSubquery
	HoistValuesSubquery
	HoistProjectSetSubquery
	NormalizeSelectAnyFilter
	NormalizeJoinAnyFilter
	NormalizeSelectNotAnyFilter
	NormalizeJoinNotAnyFilter
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldInEmpty
	FoldNotInEmpty
	FoldArray
	FoldBinary
	FoldUnary
	FoldComparison
	FoldCast
	FoldAssignmentCast
	FoldIndirection
	FoldColumnAccess
	FoldFunctionWithNullArg
	FoldFunction
	FoldEqualsAnyNull
	ConvertGroupByToDistinct
	EliminateGroupByProject
	EliminateJoinUnderGroupByLeft
	EliminateJoinUnderGroupByRight
	EliminateDistinct
	ReduceGroupingCols
	ReduceNotNullGroupingCols
	EliminateAggDistinctForKeys
	EliminateAggFilteredDistinctForKeys
	EliminateDistinctNoColumns
	EliminateEnsureDistinctNoColumns
	EliminateDistinctOnValues
	PushAggDistinctIntoGroupBy
	PushAggFilterIntoScalarGroupBy
	ConvertCountToCountRows
	ConvertRegressionCountToCount
	FoldGroupingOperators
	InlineConstVar
	InlineProjectConstants
	InlineSelectConstants
	InlineJoinConstantsLeft
	InlineJoinConstantsRight
	PushSelectIntoInlinableProject
	InlineSelectVirtualColumns
	InlineProjectInProject
	CommuteRightJoin
	SimplifyJoinFilters
	DetectJoinContradiction
	PushFilterIntoJoinLeftAndRight
	MapFilterIntoJoinLeft
	MapFilterIntoJoinRight
	MapEqualityIntoJoinLeftAndRight
	PushFilterIntoJoinLeft
	PushFilterIntoJoinRight
	SimplifyLeftJoin
	SimplifyRightJoin
	EliminateSemiJoin
	SimplifyZeroCardinalitySemiJoin
	EliminateAntiJoin
	SimplifyZeroCardinalityAntiJoin
	EliminateJoinNoColsLeft
	EliminateJoinNoColsRight
	HoistJoinProjectRight
	HoistJoinProjectLeft
	SimplifyJoinNotNullEquality
	ExtractJoinEqualities
	SortFiltersInJoin
	LeftAssociateJoinsLeft
	LeftAssociateJoinsRight
	RightAssociateJoinsLeft
	RightAssociateJoinsRight
	RemoveJoinNotNullCondition
	ProjectInnerJoinValues
	EliminateLimit
	EliminateOffset
	PushLimitIntoProject
	PushOffsetIntoProject
	PushLimitIntoOffset
	PushLimitIntoOrdinality
	PushLimitIntoJoinLeft
	PushLimitIntoJoinRight
	FoldLimits
	AssociateLimitJoinsLeft
	AssociateLimitJoinsRight
	EliminateMax1Row
	SimplifyPartialIndexProjections
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	SimplifyLimitOrdering
	SimplifyOffsetOrdering
	SimplifyGroupByOrdering
	SimplifyOrdinalityOrdering
	SimplifyExplainOrdering
	SimplifyWithBindingOrdering
	EliminateJoinUnderProjectLeft
	EliminateJoinUnderProjectRight
	EliminateProject
	MergeProjects
	MergeProjectWithValues
	PushColumnRemappingIntoValues
	PushAssignmentCastsIntoValues
	FoldTupleAccessIntoValues
	FoldJSONAccessIntoValues
	ConvertZipArraysToValues
	PruneProjectCols
	PruneScanCols
	PruneSelectCols
	PruneLimitCols
	PruneOffsetCols
	PruneJoinLeftCols
	PruneJoinRightCols
	PruneSemiAntiJoinRightCols
	PruneAggCols
	PruneGroupByCols
	PruneValuesCols
	PruneOrdinalityCols
	PruneExplainCols
	PruneProjectSetCols
	PruneWindowOutputCols
	PruneWindowInputCols
	PruneMutationFetchCols
	PruneMutationInputCols
	PruneMutationReturnCols
	PruneWithScanCols
	PruneWithCols
	PruneUnionAllCols
	RejectNullsLeftJoin
	RejectNullsRightJoin
	RejectNullsGroupBy
	RejectNullsUnderJoinLeft
	RejectNullsUnderJoinRight
	RejectNullsProject
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	NormalizeInConst
	FoldInNull
	SimplifyInSingleElement
	SimplifyNotInSingleElement
	UnifyComparisonTypes
	EliminateExistsZeroRows
	EliminateExistsProject
	EliminateExistsGroupBy
	InlineExistsSelectTuple
	IntroduceExistsLimit
	EliminateExistsLimit
	SimplifyCaseWhenConstValue
	InlineAnyValuesSingleCol
	InlineAnyValuesMultiCol
	SimplifyEqualsAnyTuple
	SimplifyAnyScalarArray
	FoldCollate
	NormalizeArrayFlattenToAgg
	SimplifySameVarEqualities
	SimplifySameVarInequalities
	SimplifyNotDisjoint
	SimplifySelectFilters
	ConsolidateSelectFilters
	DeduplicateSelectFilters
	EliminateSelect
	MergeSelects
	PushSelectIntoProject
	MergeSelectInnerJoin
	PushSelectCondLeftIntoJoinLeftAndRight
	PushSelectIntoJoinLeft
	PushSelectIntoGroupBy
	RemoveNotNullCondition
	SimplifyIsNullCondition
	PushSelectIntoProjectSet
	PushFilterIntoSetOp
	EliminateSetLeft
	EliminateSetRight
	EliminateDistinctSetLeft
	EliminateDistinctSetRight
	SimplifyExcept
	SimplifyIntersectLeft
	SimplifyIntersectRight
	ConvertUnionToDistinctUnionAll
	EliminateWindow
	ReduceWindowPartitionCols
	SimplifyWindowOrdering
	PushSelectIntoWindow
	PushLimitIntoWindow
	InlineWith

	// startExploreRule tracks the number of normalization rules;
	// all rules greater than this value are exploration rules.
	startExploreRule

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	MemoCycleTestRelRule
	ReplaceScalarMinMaxWithScalarSubqueries
	ReplaceFilteredScalarMinMaxWithSubqueries
	ReplaceScalarMinMaxWithLimit
	ReplaceMinWithLimit
	ReplaceMaxWithLimit
	GenerateStreamingGroupBy
	SplitGroupByScanIntoUnionScans
	SplitGroupByFilteredScanIntoUnionScans
	EliminateIndexJoinOrProjectInsideGroupBy
	GenerateLimitedGroupByScans
	ReorderJoins
	CommuteLeftJoin
	CommuteSemiJoin
	ConvertSemiToInnerJoin
	SplitDisjunctionOfJoinTerms
	SplitDisjunctionOfAntiJoinTerms
	GenerateMergeJoins
	GenerateLookupJoins
	GenerateInvertedJoins
	GenerateInvertedJoinsFromSelect
	GenerateLookupJoinsWithFilter
	GenerateLookupJoinsWithVirtualCols
	GenerateLookupJoinsWithVirtualColsAndFilter
	PushJoinIntoIndexJoin
	HoistProjectFromInnerJoin
	HoistProjectFromLeftJoin
	GenerateLocalityOptimizedAntiJoin
	GenerateLocalityOptimizedLookupJoin
	GenerateLimitedScans
	PushLimitIntoFilteredScan
	PushLimitIntoIndexJoin
	SplitLimitedScanIntoUnionScans
	SplitLimitedSelectIntoUnionSelects
	GenerateTopK
	GenerateLimitedTopKScans
	GeneratePartialOrderTopK
	EliminateIndexJoinInsideProject
	GenerateIndexScans
	GenerateLocalityOptimizedScan
	GeneratePartialIndexScans
	GenerateConstrainedScans
	GenerateInvertedIndexScans
	GenerateZigzagJoins
	GenerateInvertedIndexZigzagJoins
	SplitDisjunction
	SplitDisjunctionAddKey
	GenerateStreamingSetOp

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
