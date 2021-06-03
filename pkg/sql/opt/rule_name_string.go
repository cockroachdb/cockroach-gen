// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidRuleName-0]
	_ = x[SimplifyRootOrdering-1]
	_ = x[PruneRootCols-2]
	_ = x[SimplifyZeroCardinalityGroup-3]
	_ = x[NumManualRuleNames-4]
	_ = x[startAutoRule-4]
	_ = x[EliminateAggDistinct-5]
	_ = x[NormalizeNestedAnds-6]
	_ = x[SimplifyTrueAnd-7]
	_ = x[SimplifyAndTrue-8]
	_ = x[SimplifyFalseAnd-9]
	_ = x[SimplifyAndFalse-10]
	_ = x[SimplifyTrueOr-11]
	_ = x[SimplifyOrTrue-12]
	_ = x[SimplifyFalseOr-13]
	_ = x[SimplifyOrFalse-14]
	_ = x[SimplifyRange-15]
	_ = x[FoldNullAndOr-16]
	_ = x[FoldNotTrue-17]
	_ = x[FoldNotFalse-18]
	_ = x[FoldNotNull-19]
	_ = x[NegateComparison-20]
	_ = x[EliminateNot-21]
	_ = x[NegateAnd-22]
	_ = x[NegateOr-23]
	_ = x[ExtractRedundantConjunct-24]
	_ = x[CommuteVarInequality-25]
	_ = x[CommuteConstInequality-26]
	_ = x[NormalizeCmpPlusConst-27]
	_ = x[NormalizeCmpMinusConst-28]
	_ = x[NormalizeCmpConstMinus-29]
	_ = x[NormalizeTupleEquality-30]
	_ = x[FoldNullComparisonLeft-31]
	_ = x[FoldNullComparisonRight-32]
	_ = x[FoldIsNull-33]
	_ = x[FoldNonNullIsNull-34]
	_ = x[FoldNullTupleIsTupleNull-35]
	_ = x[FoldNonNullTupleIsTupleNull-36]
	_ = x[FoldIsNotNull-37]
	_ = x[FoldNonNullIsNotNull-38]
	_ = x[FoldNonNullTupleIsTupleNotNull-39]
	_ = x[FoldNullTupleIsTupleNotNull-40]
	_ = x[CommuteNullIs-41]
	_ = x[NormalizeCmpTimeZoneFunction-42]
	_ = x[NormalizeCmpTimeZoneFunctionTZ-43]
	_ = x[FoldEqZeroSTDistance-44]
	_ = x[FoldCmpSTDistanceLeft-45]
	_ = x[FoldCmpSTDistanceRight-46]
	_ = x[FoldCmpSTMaxDistanceLeft-47]
	_ = x[FoldCmpSTMaxDistanceRight-48]
	_ = x[FoldEqTrue-49]
	_ = x[FoldEqFalse-50]
	_ = x[FoldNeTrue-51]
	_ = x[FoldNeFalse-52]
	_ = x[NormCycleTestRelTrueToFalse-53]
	_ = x[NormCycleTestRelFalseToTrue-54]
	_ = x[DecorrelateJoin-55]
	_ = x[DecorrelateProjectSet-56]
	_ = x[TryDecorrelateSelect-57]
	_ = x[TryDecorrelateProject-58]
	_ = x[TryDecorrelateProjectSelect-59]
	_ = x[TryDecorrelateProjectInnerJoin-60]
	_ = x[TryDecorrelateInnerJoin-61]
	_ = x[TryDecorrelateInnerLeftJoin-62]
	_ = x[TryDecorrelateGroupBy-63]
	_ = x[TryDecorrelateScalarGroupBy-64]
	_ = x[TryDecorrelateSemiJoin-65]
	_ = x[TryDecorrelateLimitOne-66]
	_ = x[TryDecorrelateProjectSet-67]
	_ = x[TryDecorrelateWindow-68]
	_ = x[TryDecorrelateMax1Row-69]
	_ = x[HoistSelectExists-70]
	_ = x[HoistSelectNotExists-71]
	_ = x[HoistSelectSubquery-72]
	_ = x[HoistProjectSubquery-73]
	_ = x[HoistJoinSubquery-74]
	_ = x[HoistValuesSubquery-75]
	_ = x[HoistProjectSetSubquery-76]
	_ = x[NormalizeSelectAnyFilter-77]
	_ = x[NormalizeJoinAnyFilter-78]
	_ = x[NormalizeSelectNotAnyFilter-79]
	_ = x[NormalizeJoinNotAnyFilter-80]
	_ = x[FoldNullCast-81]
	_ = x[FoldNullUnary-82]
	_ = x[FoldNullBinaryLeft-83]
	_ = x[FoldNullBinaryRight-84]
	_ = x[FoldNullInNonEmpty-85]
	_ = x[FoldInEmpty-86]
	_ = x[FoldNotInEmpty-87]
	_ = x[FoldArray-88]
	_ = x[FoldBinary-89]
	_ = x[FoldUnary-90]
	_ = x[FoldComparison-91]
	_ = x[FoldCast-92]
	_ = x[FoldIndirection-93]
	_ = x[FoldColumnAccess-94]
	_ = x[FoldFunctionWithNullArg-95]
	_ = x[FoldFunction-96]
	_ = x[FoldEqualsAnyNull-97]
	_ = x[ConvertGroupByToDistinct-98]
	_ = x[EliminateGroupByProject-99]
	_ = x[EliminateJoinUnderGroupByLeft-100]
	_ = x[EliminateJoinUnderGroupByRight-101]
	_ = x[EliminateDistinct-102]
	_ = x[ReduceGroupingCols-103]
	_ = x[ReduceNotNullGroupingCols-104]
	_ = x[EliminateAggDistinctForKeys-105]
	_ = x[EliminateAggFilteredDistinctForKeys-106]
	_ = x[EliminateDistinctNoColumns-107]
	_ = x[EliminateEnsureDistinctNoColumns-108]
	_ = x[EliminateDistinctOnValues-109]
	_ = x[PushAggDistinctIntoGroupBy-110]
	_ = x[PushAggFilterIntoScalarGroupBy-111]
	_ = x[ConvertCountToCountRows-112]
	_ = x[ConvertRegressionCountToCount-113]
	_ = x[FoldGroupingOperators-114]
	_ = x[InlineConstVar-115]
	_ = x[InlineProjectConstants-116]
	_ = x[InlineSelectConstants-117]
	_ = x[InlineJoinConstantsLeft-118]
	_ = x[InlineJoinConstantsRight-119]
	_ = x[PushSelectIntoInlinableProject-120]
	_ = x[InlineSelectVirtualColumns-121]
	_ = x[InlineProjectInProject-122]
	_ = x[CommuteRightJoin-123]
	_ = x[SimplifyJoinFilters-124]
	_ = x[DetectJoinContradiction-125]
	_ = x[PushFilterIntoJoinLeftAndRight-126]
	_ = x[MapFilterIntoJoinLeft-127]
	_ = x[MapFilterIntoJoinRight-128]
	_ = x[MapEqualityIntoJoinLeftAndRight-129]
	_ = x[PushFilterIntoJoinLeft-130]
	_ = x[PushFilterIntoJoinRight-131]
	_ = x[SimplifyLeftJoin-132]
	_ = x[SimplifyRightJoin-133]
	_ = x[EliminateSemiJoin-134]
	_ = x[SimplifyZeroCardinalitySemiJoin-135]
	_ = x[EliminateAntiJoin-136]
	_ = x[SimplifyZeroCardinalityAntiJoin-137]
	_ = x[EliminateJoinNoColsLeft-138]
	_ = x[EliminateJoinNoColsRight-139]
	_ = x[HoistJoinProjectRight-140]
	_ = x[HoistJoinProjectLeft-141]
	_ = x[SimplifyJoinNotNullEquality-142]
	_ = x[ExtractJoinEqualities-143]
	_ = x[SortFiltersInJoin-144]
	_ = x[LeftAssociateJoinsLeft-145]
	_ = x[LeftAssociateJoinsRight-146]
	_ = x[RightAssociateJoinsLeft-147]
	_ = x[RightAssociateJoinsRight-148]
	_ = x[RemoveJoinNotNullCondition-149]
	_ = x[ProjectInnerJoinValues-150]
	_ = x[EliminateLimit-151]
	_ = x[EliminateOffset-152]
	_ = x[PushLimitIntoProject-153]
	_ = x[PushOffsetIntoProject-154]
	_ = x[PushLimitIntoOffset-155]
	_ = x[PushLimitIntoOrdinality-156]
	_ = x[PushLimitIntoJoinLeft-157]
	_ = x[PushLimitIntoJoinRight-158]
	_ = x[FoldLimits-159]
	_ = x[AssociateLimitJoinsLeft-160]
	_ = x[AssociateLimitJoinsRight-161]
	_ = x[EliminateMax1Row-162]
	_ = x[SimplifyPartialIndexProjections-163]
	_ = x[FoldPlusZero-164]
	_ = x[FoldZeroPlus-165]
	_ = x[FoldMinusZero-166]
	_ = x[FoldMultOne-167]
	_ = x[FoldOneMult-168]
	_ = x[FoldDivOne-169]
	_ = x[InvertMinus-170]
	_ = x[EliminateUnaryMinus-171]
	_ = x[SimplifyLimitOrdering-172]
	_ = x[SimplifyOffsetOrdering-173]
	_ = x[SimplifyGroupByOrdering-174]
	_ = x[SimplifyOrdinalityOrdering-175]
	_ = x[SimplifyExplainOrdering-176]
	_ = x[EliminateJoinUnderProjectLeft-177]
	_ = x[EliminateJoinUnderProjectRight-178]
	_ = x[EliminateProject-179]
	_ = x[MergeProjects-180]
	_ = x[MergeProjectWithValues-181]
	_ = x[PushColumnRemappingIntoValues-182]
	_ = x[FoldTupleAccessIntoValues-183]
	_ = x[FoldJSONAccessIntoValues-184]
	_ = x[ConvertZipArraysToValues-185]
	_ = x[PruneProjectCols-186]
	_ = x[PruneScanCols-187]
	_ = x[PruneSelectCols-188]
	_ = x[PruneLimitCols-189]
	_ = x[PruneOffsetCols-190]
	_ = x[PruneJoinLeftCols-191]
	_ = x[PruneJoinRightCols-192]
	_ = x[PruneSemiAntiJoinRightCols-193]
	_ = x[PruneAggCols-194]
	_ = x[PruneGroupByCols-195]
	_ = x[PruneValuesCols-196]
	_ = x[PruneOrdinalityCols-197]
	_ = x[PruneExplainCols-198]
	_ = x[PruneProjectSetCols-199]
	_ = x[PruneWindowOutputCols-200]
	_ = x[PruneWindowInputCols-201]
	_ = x[PruneMutationFetchCols-202]
	_ = x[PruneMutationInputCols-203]
	_ = x[PruneMutationReturnCols-204]
	_ = x[PruneWithScanCols-205]
	_ = x[PruneWithCols-206]
	_ = x[PruneUnionAllCols-207]
	_ = x[RejectNullsLeftJoin-208]
	_ = x[RejectNullsRightJoin-209]
	_ = x[RejectNullsGroupBy-210]
	_ = x[RejectNullsUnderJoinLeft-211]
	_ = x[RejectNullsUnderJoinRight-212]
	_ = x[RejectNullsProject-213]
	_ = x[CommuteVar-214]
	_ = x[CommuteConst-215]
	_ = x[EliminateCoalesce-216]
	_ = x[SimplifyCoalesce-217]
	_ = x[EliminateCast-218]
	_ = x[NormalizeInConst-219]
	_ = x[FoldInNull-220]
	_ = x[SimplifyInSingleElement-221]
	_ = x[SimplifyNotInSingleElement-222]
	_ = x[UnifyComparisonTypes-223]
	_ = x[EliminateExistsZeroRows-224]
	_ = x[EliminateExistsProject-225]
	_ = x[EliminateExistsGroupBy-226]
	_ = x[InlineExistsSelectTuple-227]
	_ = x[IntroduceExistsLimit-228]
	_ = x[EliminateExistsLimit-229]
	_ = x[SimplifyCaseWhenConstValue-230]
	_ = x[InlineAnyValuesSingleCol-231]
	_ = x[InlineAnyValuesMultiCol-232]
	_ = x[SimplifyEqualsAnyTuple-233]
	_ = x[SimplifyAnyScalarArray-234]
	_ = x[FoldCollate-235]
	_ = x[NormalizeArrayFlattenToAgg-236]
	_ = x[SimplifySameVarEqualities-237]
	_ = x[SimplifySameVarInequalities-238]
	_ = x[SimplifyNotDisjoint-239]
	_ = x[SimplifySelectFilters-240]
	_ = x[ConsolidateSelectFilters-241]
	_ = x[EliminateSelect-242]
	_ = x[MergeSelects-243]
	_ = x[PushSelectIntoProject-244]
	_ = x[MergeSelectInnerJoin-245]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-246]
	_ = x[PushSelectIntoJoinLeft-247]
	_ = x[PushSelectIntoGroupBy-248]
	_ = x[RemoveNotNullCondition-249]
	_ = x[PushSelectIntoProjectSet-250]
	_ = x[PushFilterIntoSetOp-251]
	_ = x[EliminateUnionAllLeft-252]
	_ = x[EliminateUnionAllRight-253]
	_ = x[EliminateUnionLeft-254]
	_ = x[EliminateUnionRight-255]
	_ = x[ConvertUnionToDistinctUnionAll-256]
	_ = x[EliminateWindow-257]
	_ = x[ReduceWindowPartitionCols-258]
	_ = x[SimplifyWindowOrdering-259]
	_ = x[PushSelectIntoWindow-260]
	_ = x[PushLimitIntoWindow-261]
	_ = x[InlineWith-262]
	_ = x[startExploreRule-263]
	_ = x[ReplaceScalarMinMaxWithLimit-264]
	_ = x[ReplaceMinWithLimit-265]
	_ = x[ReplaceMaxWithLimit-266]
	_ = x[GenerateStreamingGroupBy-267]
	_ = x[SplitGroupByScanIntoUnionScans-268]
	_ = x[SplitGroupByFilteredScanIntoUnionScans-269]
	_ = x[ReorderJoins-270]
	_ = x[CommuteLeftJoin-271]
	_ = x[CommuteSemiJoin-272]
	_ = x[ConvertSemiToInnerJoin-273]
	_ = x[GenerateMergeJoins-274]
	_ = x[GenerateLookupJoins-275]
	_ = x[GenerateInvertedJoins-276]
	_ = x[GenerateInvertedJoinsFromSelect-277]
	_ = x[GenerateLookupJoinsWithFilter-278]
	_ = x[PushJoinIntoIndexJoin-279]
	_ = x[HoistProjectFromInnerJoin-280]
	_ = x[HoistProjectFromLeftJoin-281]
	_ = x[GenerateLocalityOptimizedAntiJoin-282]
	_ = x[GenerateLocalityOptimizedLookupJoin-283]
	_ = x[GenerateLimitedScans-284]
	_ = x[PushLimitIntoFilteredScan-285]
	_ = x[PushLimitIntoIndexJoin-286]
	_ = x[SplitLimitedScanIntoUnionScans-287]
	_ = x[EliminateIndexJoinInsideProject-288]
	_ = x[GenerateIndexScans-289]
	_ = x[GenerateLocalityOptimizedScan-290]
	_ = x[GeneratePartialIndexScans-291]
	_ = x[GenerateConstrainedScans-292]
	_ = x[GenerateInvertedIndexScans-293]
	_ = x[GenerateZigzagJoins-294]
	_ = x[GenerateInvertedIndexZigzagJoins-295]
	_ = x[SplitDisjunction-296]
	_ = x[SplitDisjunctionAddKey-297]
	_ = x[NumRuleNames-298]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightFoldEqTrueFoldEqFalseFoldNeTrueFoldNeFalseNormCycleTestRelTrueToFalseNormCycleTestRelFalseToTrueDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionWithNullArgFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsConvertRegressionCountToCountFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineSelectVirtualColumnsInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionProjectInnerJoinValuesEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowSimplifyPartialIndexProjectionsFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullSimplifyInSingleElementSimplifyNotInSingleElementUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifyNotDisjointSimplifySelectFiltersConsolidateSelectFiltersEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateUnionLeftEliminateUnionRightConvertUnionToDistinctUnionAllEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupBySplitGroupByScanIntoUnionScansSplitGroupByFilteredScanIntoUnionScansReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateLookupJoinsWithFilterPushJoinIntoIndexJoinHoistProjectFromInnerJoinHoistProjectFromLeftJoinGenerateLocalityOptimizedAntiJoinGenerateLocalityOptimizedLookupJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitLimitedScanIntoUnionScansEliminateIndexJoinInsideProjectGenerateIndexScansGenerateLocalityOptimizedScanGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 917, 928, 938, 949, 976, 1003, 1018, 1039, 1059, 1080, 1107, 1137, 1160, 1187, 1208, 1235, 1257, 1279, 1303, 1323, 1344, 1361, 1381, 1400, 1420, 1437, 1456, 1479, 1503, 1525, 1552, 1577, 1589, 1602, 1620, 1639, 1657, 1668, 1682, 1691, 1701, 1710, 1724, 1732, 1747, 1763, 1786, 1798, 1815, 1839, 1862, 1891, 1921, 1938, 1956, 1981, 2008, 2043, 2069, 2101, 2126, 2152, 2182, 2205, 2234, 2255, 2269, 2291, 2312, 2335, 2359, 2389, 2415, 2437, 2453, 2472, 2495, 2525, 2546, 2568, 2599, 2621, 2644, 2660, 2677, 2694, 2725, 2742, 2773, 2796, 2820, 2841, 2861, 2888, 2909, 2926, 2948, 2971, 2994, 3018, 3044, 3066, 3080, 3095, 3115, 3136, 3155, 3178, 3199, 3221, 3231, 3254, 3278, 3294, 3325, 3337, 3349, 3362, 3373, 3384, 3394, 3405, 3424, 3445, 3467, 3490, 3516, 3539, 3568, 3598, 3614, 3627, 3649, 3678, 3703, 3727, 3751, 3767, 3780, 3795, 3809, 3824, 3841, 3859, 3885, 3897, 3913, 3928, 3947, 3963, 3982, 4003, 4023, 4045, 4067, 4090, 4107, 4120, 4137, 4156, 4176, 4194, 4218, 4243, 4261, 4271, 4283, 4300, 4316, 4329, 4345, 4355, 4378, 4404, 4424, 4447, 4469, 4491, 4514, 4534, 4554, 4580, 4604, 4627, 4649, 4671, 4682, 4708, 4733, 4760, 4779, 4800, 4824, 4839, 4851, 4872, 4892, 4930, 4952, 4973, 4995, 5019, 5038, 5059, 5081, 5099, 5118, 5148, 5163, 5188, 5210, 5230, 5249, 5259, 5275, 5303, 5322, 5341, 5365, 5395, 5433, 5445, 5460, 5475, 5497, 5515, 5534, 5555, 5586, 5615, 5636, 5661, 5685, 5718, 5753, 5773, 5798, 5820, 5850, 5881, 5899, 5928, 5953, 5977, 6003, 6022, 6054, 6070, 6092, 6104}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
