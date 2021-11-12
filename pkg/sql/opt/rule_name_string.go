// Code generated by "stringer"; DO NOT EDIT.

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
	_ = x[TryDecorrelateLimit-67]
	_ = x[TryDecorrelateProjectSet-68]
	_ = x[TryDecorrelateWindow-69]
	_ = x[TryDecorrelateMax1Row-70]
	_ = x[HoistSelectExists-71]
	_ = x[HoistSelectNotExists-72]
	_ = x[HoistSelectSubquery-73]
	_ = x[HoistProjectSubquery-74]
	_ = x[HoistJoinSubquery-75]
	_ = x[HoistValuesSubquery-76]
	_ = x[HoistProjectSetSubquery-77]
	_ = x[NormalizeSelectAnyFilter-78]
	_ = x[NormalizeJoinAnyFilter-79]
	_ = x[NormalizeSelectNotAnyFilter-80]
	_ = x[NormalizeJoinNotAnyFilter-81]
	_ = x[FoldNullCast-82]
	_ = x[FoldNullUnary-83]
	_ = x[FoldNullBinaryLeft-84]
	_ = x[FoldNullBinaryRight-85]
	_ = x[FoldNullInNonEmpty-86]
	_ = x[FoldInEmpty-87]
	_ = x[FoldNotInEmpty-88]
	_ = x[FoldArray-89]
	_ = x[FoldBinary-90]
	_ = x[FoldUnary-91]
	_ = x[FoldComparison-92]
	_ = x[FoldCast-93]
	_ = x[FoldIndirection-94]
	_ = x[FoldColumnAccess-95]
	_ = x[FoldFunctionWithNullArg-96]
	_ = x[FoldFunction-97]
	_ = x[FoldEqualsAnyNull-98]
	_ = x[ConvertGroupByToDistinct-99]
	_ = x[EliminateGroupByProject-100]
	_ = x[EliminateJoinUnderGroupByLeft-101]
	_ = x[EliminateJoinUnderGroupByRight-102]
	_ = x[EliminateDistinct-103]
	_ = x[ReduceGroupingCols-104]
	_ = x[ReduceNotNullGroupingCols-105]
	_ = x[EliminateAggDistinctForKeys-106]
	_ = x[EliminateAggFilteredDistinctForKeys-107]
	_ = x[EliminateDistinctNoColumns-108]
	_ = x[EliminateEnsureDistinctNoColumns-109]
	_ = x[EliminateDistinctOnValues-110]
	_ = x[PushAggDistinctIntoGroupBy-111]
	_ = x[PushAggFilterIntoScalarGroupBy-112]
	_ = x[ConvertCountToCountRows-113]
	_ = x[ConvertRegressionCountToCount-114]
	_ = x[FoldGroupingOperators-115]
	_ = x[InlineConstVar-116]
	_ = x[InlineProjectConstants-117]
	_ = x[InlineSelectConstants-118]
	_ = x[InlineJoinConstantsLeft-119]
	_ = x[InlineJoinConstantsRight-120]
	_ = x[PushSelectIntoInlinableProject-121]
	_ = x[InlineSelectVirtualColumns-122]
	_ = x[InlineProjectInProject-123]
	_ = x[CommuteRightJoin-124]
	_ = x[SimplifyJoinFilters-125]
	_ = x[DetectJoinContradiction-126]
	_ = x[PushFilterIntoJoinLeftAndRight-127]
	_ = x[MapFilterIntoJoinLeft-128]
	_ = x[MapFilterIntoJoinRight-129]
	_ = x[MapEqualityIntoJoinLeftAndRight-130]
	_ = x[PushFilterIntoJoinLeft-131]
	_ = x[PushFilterIntoJoinRight-132]
	_ = x[SimplifyLeftJoin-133]
	_ = x[SimplifyRightJoin-134]
	_ = x[EliminateSemiJoin-135]
	_ = x[SimplifyZeroCardinalitySemiJoin-136]
	_ = x[EliminateAntiJoin-137]
	_ = x[SimplifyZeroCardinalityAntiJoin-138]
	_ = x[EliminateJoinNoColsLeft-139]
	_ = x[EliminateJoinNoColsRight-140]
	_ = x[HoistJoinProjectRight-141]
	_ = x[HoistJoinProjectLeft-142]
	_ = x[SimplifyJoinNotNullEquality-143]
	_ = x[ExtractJoinEqualities-144]
	_ = x[SortFiltersInJoin-145]
	_ = x[LeftAssociateJoinsLeft-146]
	_ = x[LeftAssociateJoinsRight-147]
	_ = x[RightAssociateJoinsLeft-148]
	_ = x[RightAssociateJoinsRight-149]
	_ = x[RemoveJoinNotNullCondition-150]
	_ = x[ProjectInnerJoinValues-151]
	_ = x[EliminateLimit-152]
	_ = x[EliminateOffset-153]
	_ = x[PushLimitIntoProject-154]
	_ = x[PushOffsetIntoProject-155]
	_ = x[PushLimitIntoOffset-156]
	_ = x[PushLimitIntoOrdinality-157]
	_ = x[PushLimitIntoJoinLeft-158]
	_ = x[PushLimitIntoJoinRight-159]
	_ = x[FoldLimits-160]
	_ = x[AssociateLimitJoinsLeft-161]
	_ = x[AssociateLimitJoinsRight-162]
	_ = x[EliminateMax1Row-163]
	_ = x[SimplifyPartialIndexProjections-164]
	_ = x[FoldPlusZero-165]
	_ = x[FoldZeroPlus-166]
	_ = x[FoldMinusZero-167]
	_ = x[FoldMultOne-168]
	_ = x[FoldOneMult-169]
	_ = x[FoldDivOne-170]
	_ = x[InvertMinus-171]
	_ = x[EliminateUnaryMinus-172]
	_ = x[SimplifyLimitOrdering-173]
	_ = x[SimplifyOffsetOrdering-174]
	_ = x[SimplifyGroupByOrdering-175]
	_ = x[SimplifyOrdinalityOrdering-176]
	_ = x[SimplifyExplainOrdering-177]
	_ = x[SimplifyWithBindingOrdering-178]
	_ = x[EliminateJoinUnderProjectLeft-179]
	_ = x[EliminateJoinUnderProjectRight-180]
	_ = x[EliminateProject-181]
	_ = x[MergeProjects-182]
	_ = x[MergeProjectWithValues-183]
	_ = x[PushColumnRemappingIntoValues-184]
	_ = x[FoldTupleAccessIntoValues-185]
	_ = x[FoldJSONAccessIntoValues-186]
	_ = x[ConvertZipArraysToValues-187]
	_ = x[PruneProjectCols-188]
	_ = x[PruneScanCols-189]
	_ = x[PruneSelectCols-190]
	_ = x[PruneLimitCols-191]
	_ = x[PruneOffsetCols-192]
	_ = x[PruneJoinLeftCols-193]
	_ = x[PruneJoinRightCols-194]
	_ = x[PruneSemiAntiJoinRightCols-195]
	_ = x[PruneAggCols-196]
	_ = x[PruneGroupByCols-197]
	_ = x[PruneValuesCols-198]
	_ = x[PruneOrdinalityCols-199]
	_ = x[PruneExplainCols-200]
	_ = x[PruneProjectSetCols-201]
	_ = x[PruneWindowOutputCols-202]
	_ = x[PruneWindowInputCols-203]
	_ = x[PruneMutationFetchCols-204]
	_ = x[PruneMutationInputCols-205]
	_ = x[PruneMutationReturnCols-206]
	_ = x[PruneWithScanCols-207]
	_ = x[PruneWithCols-208]
	_ = x[PruneUnionAllCols-209]
	_ = x[RejectNullsLeftJoin-210]
	_ = x[RejectNullsRightJoin-211]
	_ = x[RejectNullsGroupBy-212]
	_ = x[RejectNullsUnderJoinLeft-213]
	_ = x[RejectNullsUnderJoinRight-214]
	_ = x[RejectNullsProject-215]
	_ = x[CommuteVar-216]
	_ = x[CommuteConst-217]
	_ = x[EliminateCoalesce-218]
	_ = x[SimplifyCoalesce-219]
	_ = x[EliminateCast-220]
	_ = x[NormalizeInConst-221]
	_ = x[FoldInNull-222]
	_ = x[SimplifyInSingleElement-223]
	_ = x[SimplifyNotInSingleElement-224]
	_ = x[UnifyComparisonTypes-225]
	_ = x[EliminateExistsZeroRows-226]
	_ = x[EliminateExistsProject-227]
	_ = x[EliminateExistsGroupBy-228]
	_ = x[InlineExistsSelectTuple-229]
	_ = x[IntroduceExistsLimit-230]
	_ = x[EliminateExistsLimit-231]
	_ = x[SimplifyCaseWhenConstValue-232]
	_ = x[InlineAnyValuesSingleCol-233]
	_ = x[InlineAnyValuesMultiCol-234]
	_ = x[SimplifyEqualsAnyTuple-235]
	_ = x[SimplifyAnyScalarArray-236]
	_ = x[FoldCollate-237]
	_ = x[NormalizeArrayFlattenToAgg-238]
	_ = x[SimplifySameVarEqualities-239]
	_ = x[SimplifySameVarInequalities-240]
	_ = x[SimplifyNotDisjoint-241]
	_ = x[SimplifySelectFilters-242]
	_ = x[ConsolidateSelectFilters-243]
	_ = x[DeduplicateSelectFilters-244]
	_ = x[EliminateSelect-245]
	_ = x[MergeSelects-246]
	_ = x[PushSelectIntoProject-247]
	_ = x[MergeSelectInnerJoin-248]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-249]
	_ = x[PushSelectIntoJoinLeft-250]
	_ = x[PushSelectIntoGroupBy-251]
	_ = x[RemoveNotNullCondition-252]
	_ = x[PushSelectIntoProjectSet-253]
	_ = x[PushFilterIntoSetOp-254]
	_ = x[EliminateSetLeft-255]
	_ = x[EliminateSetRight-256]
	_ = x[EliminateDistinctSetLeft-257]
	_ = x[EliminateDistinctSetRight-258]
	_ = x[SimplifyExcept-259]
	_ = x[SimplifyIntersectLeft-260]
	_ = x[SimplifyIntersectRight-261]
	_ = x[ConvertUnionToDistinctUnionAll-262]
	_ = x[EliminateWindow-263]
	_ = x[ReduceWindowPartitionCols-264]
	_ = x[SimplifyWindowOrdering-265]
	_ = x[PushSelectIntoWindow-266]
	_ = x[PushLimitIntoWindow-267]
	_ = x[InlineWith-268]
	_ = x[startExploreRule-269]
	_ = x[ReplaceScalarMinMaxWithScalarSubqueries-270]
	_ = x[ReplaceFilteredScalarMinMaxWithSubqueries-271]
	_ = x[ReplaceScalarMinMaxWithLimit-272]
	_ = x[ReplaceMinWithLimit-273]
	_ = x[ReplaceMaxWithLimit-274]
	_ = x[GenerateStreamingGroupBy-275]
	_ = x[SplitGroupByScanIntoUnionScans-276]
	_ = x[SplitGroupByFilteredScanIntoUnionScans-277]
	_ = x[EliminateIndexJoinOrProjectInsideGroupBy-278]
	_ = x[GenerateLimitedGroupByScans-279]
	_ = x[ReorderJoins-280]
	_ = x[CommuteLeftJoin-281]
	_ = x[CommuteSemiJoin-282]
	_ = x[ConvertSemiToInnerJoin-283]
	_ = x[GenerateMergeJoins-284]
	_ = x[GenerateLookupJoins-285]
	_ = x[GenerateInvertedJoins-286]
	_ = x[GenerateInvertedJoinsFromSelect-287]
	_ = x[GenerateLookupJoinsWithFilter-288]
	_ = x[GenerateLookupJoinsWithVirtualCols-289]
	_ = x[GenerateLookupJoinsWithVirtualColsAndFilter-290]
	_ = x[PushJoinIntoIndexJoin-291]
	_ = x[HoistProjectFromInnerJoin-292]
	_ = x[HoistProjectFromLeftJoin-293]
	_ = x[GenerateLocalityOptimizedAntiJoin-294]
	_ = x[GenerateLocalityOptimizedLookupJoin-295]
	_ = x[GenerateLimitedScans-296]
	_ = x[PushLimitIntoFilteredScan-297]
	_ = x[PushLimitIntoIndexJoin-298]
	_ = x[SplitLimitedScanIntoUnionScans-299]
	_ = x[GenerateTopK-300]
	_ = x[EliminateIndexJoinInsideProject-301]
	_ = x[GenerateIndexScans-302]
	_ = x[GenerateLocalityOptimizedScan-303]
	_ = x[GeneratePartialIndexScans-304]
	_ = x[GenerateConstrainedScans-305]
	_ = x[GenerateInvertedIndexScans-306]
	_ = x[GenerateZigzagJoins-307]
	_ = x[GenerateInvertedIndexZigzagJoins-308]
	_ = x[SplitDisjunction-309]
	_ = x[SplitDisjunctionAddKey-310]
	_ = x[GenerateStreamingSetOp-311]
	_ = x[NumRuleNames-312]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightFoldEqTrueFoldEqFalseFoldNeTrueFoldNeFalseNormCycleTestRelTrueToFalseNormCycleTestRelFalseToTrueDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateLimitTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionWithNullArgFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsConvertRegressionCountToCountFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineSelectVirtualColumnsInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionProjectInnerJoinValuesEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowSimplifyPartialIndexProjectionsFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingSimplifyWithBindingOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullSimplifyInSingleElementSimplifyNotInSingleElementUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifyNotDisjointSimplifySelectFiltersConsolidateSelectFiltersDeduplicateSelectFiltersEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateSetLeftEliminateSetRightEliminateDistinctSetLeftEliminateDistinctSetRightSimplifyExceptSimplifyIntersectLeftSimplifyIntersectRightConvertUnionToDistinctUnionAllEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithScalarSubqueriesReplaceFilteredScalarMinMaxWithSubqueriesReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupBySplitGroupByScanIntoUnionScansSplitGroupByFilteredScanIntoUnionScansEliminateIndexJoinOrProjectInsideGroupByGenerateLimitedGroupByScansReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateLookupJoinsWithFilterGenerateLookupJoinsWithVirtualColsGenerateLookupJoinsWithVirtualColsAndFilterPushJoinIntoIndexJoinHoistProjectFromInnerJoinHoistProjectFromLeftJoinGenerateLocalityOptimizedAntiJoinGenerateLocalityOptimizedLookupJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitLimitedScanIntoUnionScansGenerateTopKEliminateIndexJoinInsideProjectGenerateIndexScansGenerateLocalityOptimizedScanGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsSplitDisjunctionSplitDisjunctionAddKeyGenerateStreamingSetOpNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 917, 928, 938, 949, 976, 1003, 1018, 1039, 1059, 1080, 1107, 1137, 1160, 1187, 1208, 1235, 1257, 1279, 1298, 1322, 1342, 1363, 1380, 1400, 1419, 1439, 1456, 1475, 1498, 1522, 1544, 1571, 1596, 1608, 1621, 1639, 1658, 1676, 1687, 1701, 1710, 1720, 1729, 1743, 1751, 1766, 1782, 1805, 1817, 1834, 1858, 1881, 1910, 1940, 1957, 1975, 2000, 2027, 2062, 2088, 2120, 2145, 2171, 2201, 2224, 2253, 2274, 2288, 2310, 2331, 2354, 2378, 2408, 2434, 2456, 2472, 2491, 2514, 2544, 2565, 2587, 2618, 2640, 2663, 2679, 2696, 2713, 2744, 2761, 2792, 2815, 2839, 2860, 2880, 2907, 2928, 2945, 2967, 2990, 3013, 3037, 3063, 3085, 3099, 3114, 3134, 3155, 3174, 3197, 3218, 3240, 3250, 3273, 3297, 3313, 3344, 3356, 3368, 3381, 3392, 3403, 3413, 3424, 3443, 3464, 3486, 3509, 3535, 3558, 3585, 3614, 3644, 3660, 3673, 3695, 3724, 3749, 3773, 3797, 3813, 3826, 3841, 3855, 3870, 3887, 3905, 3931, 3943, 3959, 3974, 3993, 4009, 4028, 4049, 4069, 4091, 4113, 4136, 4153, 4166, 4183, 4202, 4222, 4240, 4264, 4289, 4307, 4317, 4329, 4346, 4362, 4375, 4391, 4401, 4424, 4450, 4470, 4493, 4515, 4537, 4560, 4580, 4600, 4626, 4650, 4673, 4695, 4717, 4728, 4754, 4779, 4806, 4825, 4846, 4870, 4894, 4909, 4921, 4942, 4962, 5000, 5022, 5043, 5065, 5089, 5108, 5124, 5141, 5165, 5190, 5204, 5225, 5247, 5277, 5292, 5317, 5339, 5359, 5378, 5388, 5404, 5443, 5484, 5512, 5531, 5550, 5574, 5604, 5642, 5682, 5709, 5721, 5736, 5751, 5773, 5791, 5810, 5831, 5862, 5891, 5925, 5968, 5989, 6014, 6038, 6071, 6106, 6126, 6151, 6173, 6203, 6215, 6246, 6264, 6293, 6318, 6342, 6368, 6387, 6419, 6435, 6457, 6479, 6491}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
