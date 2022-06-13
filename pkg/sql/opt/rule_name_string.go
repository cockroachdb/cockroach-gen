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
	_ = x[FoldAssignmentCast-94]
	_ = x[FoldIndirection-95]
	_ = x[FoldColumnAccess-96]
	_ = x[FoldFunctionWithNullArg-97]
	_ = x[FoldFunction-98]
	_ = x[FoldEqualsAnyNull-99]
	_ = x[ConvertGroupByToDistinct-100]
	_ = x[EliminateGroupByProject-101]
	_ = x[EliminateJoinUnderGroupByLeft-102]
	_ = x[EliminateJoinUnderGroupByRight-103]
	_ = x[EliminateDistinct-104]
	_ = x[ReduceGroupingCols-105]
	_ = x[ReduceNotNullGroupingCols-106]
	_ = x[EliminateAggDistinctForKeys-107]
	_ = x[EliminateAggFilteredDistinctForKeys-108]
	_ = x[EliminateDistinctNoColumns-109]
	_ = x[EliminateEnsureDistinctNoColumns-110]
	_ = x[EliminateDistinctOnValues-111]
	_ = x[PushAggDistinctIntoGroupBy-112]
	_ = x[PushAggFilterIntoScalarGroupBy-113]
	_ = x[ConvertCountToCountRows-114]
	_ = x[ConvertRegressionCountToCount-115]
	_ = x[FoldGroupingOperators-116]
	_ = x[InlineConstVar-117]
	_ = x[InlineProjectConstants-118]
	_ = x[InlineSelectConstants-119]
	_ = x[InlineJoinConstantsLeft-120]
	_ = x[InlineJoinConstantsRight-121]
	_ = x[PushSelectIntoInlinableProject-122]
	_ = x[InlineSelectVirtualColumns-123]
	_ = x[InlineProjectInProject-124]
	_ = x[CommuteRightJoin-125]
	_ = x[SimplifyJoinFilters-126]
	_ = x[DetectJoinContradiction-127]
	_ = x[PushFilterIntoJoinLeftAndRight-128]
	_ = x[MapFilterIntoJoinLeft-129]
	_ = x[MapFilterIntoJoinRight-130]
	_ = x[MapEqualityIntoJoinLeftAndRight-131]
	_ = x[PushFilterIntoJoinLeft-132]
	_ = x[PushFilterIntoJoinRight-133]
	_ = x[SimplifyLeftJoin-134]
	_ = x[SimplifyRightJoin-135]
	_ = x[EliminateSemiJoin-136]
	_ = x[SimplifyZeroCardinalitySemiJoin-137]
	_ = x[EliminateAntiJoin-138]
	_ = x[SimplifyZeroCardinalityAntiJoin-139]
	_ = x[EliminateJoinNoColsLeft-140]
	_ = x[EliminateJoinNoColsRight-141]
	_ = x[HoistJoinProjectRight-142]
	_ = x[HoistJoinProjectLeft-143]
	_ = x[SimplifyJoinNotNullEquality-144]
	_ = x[ExtractJoinEqualities-145]
	_ = x[SortFiltersInJoin-146]
	_ = x[LeftAssociateJoinsLeft-147]
	_ = x[LeftAssociateJoinsRight-148]
	_ = x[RightAssociateJoinsLeft-149]
	_ = x[RightAssociateJoinsRight-150]
	_ = x[RemoveJoinNotNullCondition-151]
	_ = x[ProjectInnerJoinValues-152]
	_ = x[EliminateLimit-153]
	_ = x[EliminateOffset-154]
	_ = x[PushLimitIntoProject-155]
	_ = x[PushOffsetIntoProject-156]
	_ = x[PushLimitIntoOffset-157]
	_ = x[PushLimitIntoOrdinality-158]
	_ = x[PushLimitIntoJoinLeft-159]
	_ = x[PushLimitIntoJoinRight-160]
	_ = x[FoldLimits-161]
	_ = x[AssociateLimitJoinsLeft-162]
	_ = x[AssociateLimitJoinsRight-163]
	_ = x[EliminateMax1Row-164]
	_ = x[SimplifyPartialIndexProjections-165]
	_ = x[FoldPlusZero-166]
	_ = x[FoldZeroPlus-167]
	_ = x[FoldMinusZero-168]
	_ = x[FoldMultOne-169]
	_ = x[FoldOneMult-170]
	_ = x[FoldDivOne-171]
	_ = x[InvertMinus-172]
	_ = x[EliminateUnaryMinus-173]
	_ = x[SimplifyLimitOrdering-174]
	_ = x[SimplifyOffsetOrdering-175]
	_ = x[SimplifyGroupByOrdering-176]
	_ = x[SimplifyOrdinalityOrdering-177]
	_ = x[SimplifyExplainOrdering-178]
	_ = x[SimplifyWithBindingOrdering-179]
	_ = x[EliminateJoinUnderProjectLeft-180]
	_ = x[EliminateJoinUnderProjectRight-181]
	_ = x[EliminateProject-182]
	_ = x[MergeProjects-183]
	_ = x[MergeProjectWithValues-184]
	_ = x[PushColumnRemappingIntoValues-185]
	_ = x[PushAssignmentCastsIntoValues-186]
	_ = x[FoldTupleAccessIntoValues-187]
	_ = x[FoldJSONAccessIntoValues-188]
	_ = x[ConvertZipArraysToValues-189]
	_ = x[PruneProjectCols-190]
	_ = x[PruneScanCols-191]
	_ = x[PruneSelectCols-192]
	_ = x[PruneLimitCols-193]
	_ = x[PruneOffsetCols-194]
	_ = x[PruneJoinLeftCols-195]
	_ = x[PruneJoinRightCols-196]
	_ = x[PruneSemiAntiJoinRightCols-197]
	_ = x[PruneAggCols-198]
	_ = x[PruneGroupByCols-199]
	_ = x[PruneValuesCols-200]
	_ = x[PruneOrdinalityCols-201]
	_ = x[PruneExplainCols-202]
	_ = x[PruneProjectSetCols-203]
	_ = x[PruneWindowOutputCols-204]
	_ = x[PruneWindowInputCols-205]
	_ = x[PruneMutationFetchCols-206]
	_ = x[PruneMutationInputCols-207]
	_ = x[PruneMutationReturnCols-208]
	_ = x[PruneWithScanCols-209]
	_ = x[PruneWithCols-210]
	_ = x[PruneUnionAllCols-211]
	_ = x[RejectNullsLeftJoin-212]
	_ = x[RejectNullsRightJoin-213]
	_ = x[RejectNullsGroupBy-214]
	_ = x[RejectNullsUnderJoinLeft-215]
	_ = x[RejectNullsUnderJoinRight-216]
	_ = x[RejectNullsProject-217]
	_ = x[CommuteVar-218]
	_ = x[CommuteConst-219]
	_ = x[EliminateCoalesce-220]
	_ = x[SimplifyCoalesce-221]
	_ = x[EliminateCast-222]
	_ = x[NormalizeInConst-223]
	_ = x[FoldInNull-224]
	_ = x[SimplifyInSingleElement-225]
	_ = x[SimplifyNotInSingleElement-226]
	_ = x[UnifyComparisonTypes-227]
	_ = x[EliminateExistsZeroRows-228]
	_ = x[EliminateExistsProject-229]
	_ = x[EliminateExistsGroupBy-230]
	_ = x[InlineExistsSelectTuple-231]
	_ = x[IntroduceExistsLimit-232]
	_ = x[EliminateExistsLimit-233]
	_ = x[SimplifyCaseWhenConstValue-234]
	_ = x[InlineAnyValuesSingleCol-235]
	_ = x[InlineAnyValuesMultiCol-236]
	_ = x[SimplifyEqualsAnyTuple-237]
	_ = x[SimplifyAnyScalarArray-238]
	_ = x[FoldCollate-239]
	_ = x[NormalizeArrayFlattenToAgg-240]
	_ = x[SimplifySameVarEqualities-241]
	_ = x[SimplifySameVarInequalities-242]
	_ = x[SimplifyNotDisjoint-243]
	_ = x[SimplifySelectFilters-244]
	_ = x[ConsolidateSelectFilters-245]
	_ = x[DeduplicateSelectFilters-246]
	_ = x[EliminateSelect-247]
	_ = x[MergeSelects-248]
	_ = x[PushSelectIntoProject-249]
	_ = x[MergeSelectInnerJoin-250]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-251]
	_ = x[PushSelectIntoJoinLeft-252]
	_ = x[PushSelectIntoGroupBy-253]
	_ = x[RemoveNotNullCondition-254]
	_ = x[PushSelectIntoProjectSet-255]
	_ = x[PushFilterIntoSetOp-256]
	_ = x[EliminateSetLeft-257]
	_ = x[EliminateSetRight-258]
	_ = x[EliminateDistinctSetLeft-259]
	_ = x[EliminateDistinctSetRight-260]
	_ = x[SimplifyExcept-261]
	_ = x[SimplifyIntersectLeft-262]
	_ = x[SimplifyIntersectRight-263]
	_ = x[ConvertUnionToDistinctUnionAll-264]
	_ = x[EliminateWindow-265]
	_ = x[ReduceWindowPartitionCols-266]
	_ = x[SimplifyWindowOrdering-267]
	_ = x[PushSelectIntoWindow-268]
	_ = x[PushLimitIntoWindow-269]
	_ = x[InlineWith-270]
	_ = x[startExploreRule-271]
	_ = x[MemoCycleTestRelRule-272]
	_ = x[ReplaceScalarMinMaxWithScalarSubqueries-273]
	_ = x[ReplaceFilteredScalarMinMaxWithSubqueries-274]
	_ = x[ReplaceScalarMinMaxWithLimit-275]
	_ = x[ReplaceMinWithLimit-276]
	_ = x[ReplaceMaxWithLimit-277]
	_ = x[GenerateStreamingGroupBy-278]
	_ = x[SplitGroupByScanIntoUnionScans-279]
	_ = x[SplitGroupByFilteredScanIntoUnionScans-280]
	_ = x[EliminateIndexJoinOrProjectInsideGroupBy-281]
	_ = x[GenerateLimitedGroupByScans-282]
	_ = x[ReorderJoins-283]
	_ = x[CommuteLeftJoin-284]
	_ = x[CommuteSemiJoin-285]
	_ = x[ConvertSemiToInnerJoin-286]
	_ = x[GenerateMergeJoins-287]
	_ = x[GenerateLookupJoins-288]
	_ = x[GenerateInvertedJoins-289]
	_ = x[GenerateInvertedJoinsFromSelect-290]
	_ = x[GenerateLookupJoinsWithFilter-291]
	_ = x[GenerateLookupJoinsWithVirtualCols-292]
	_ = x[GenerateLookupJoinsWithVirtualColsAndFilter-293]
	_ = x[PushJoinIntoIndexJoin-294]
	_ = x[HoistProjectFromInnerJoin-295]
	_ = x[HoistProjectFromLeftJoin-296]
	_ = x[GenerateLocalityOptimizedAntiJoin-297]
	_ = x[GenerateLocalityOptimizedLookupJoin-298]
	_ = x[GenerateLimitedScans-299]
	_ = x[PushLimitIntoFilteredScan-300]
	_ = x[PushLimitIntoIndexJoin-301]
	_ = x[SplitLimitedScanIntoUnionScans-302]
	_ = x[SplitLimitedSelectIntoUnionSelects-303]
	_ = x[GenerateTopK-304]
	_ = x[GenerateLimitedTopKScans-305]
	_ = x[GeneratePartialOrderTopK-306]
	_ = x[EliminateIndexJoinInsideProject-307]
	_ = x[GenerateIndexScans-308]
	_ = x[GenerateLocalityOptimizedScan-309]
	_ = x[GeneratePartialIndexScans-310]
	_ = x[GenerateConstrainedScans-311]
	_ = x[GenerateInvertedIndexScans-312]
	_ = x[GenerateZigzagJoins-313]
	_ = x[GenerateInvertedIndexZigzagJoins-314]
	_ = x[SplitDisjunction-315]
	_ = x[SplitDisjunctionAddKey-316]
	_ = x[GenerateStreamingSetOp-317]
	_ = x[NumRuleNames-318]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightFoldEqTrueFoldEqFalseFoldNeTrueFoldNeFalseNormCycleTestRelTrueToFalseNormCycleTestRelFalseToTrueDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateLimitTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldAssignmentCastFoldIndirectionFoldColumnAccessFoldFunctionWithNullArgFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsConvertRegressionCountToCountFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineSelectVirtualColumnsInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionProjectInnerJoinValuesEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowSimplifyPartialIndexProjectionsFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingSimplifyWithBindingOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesPushAssignmentCastsIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullSimplifyInSingleElementSimplifyNotInSingleElementUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifyNotDisjointSimplifySelectFiltersConsolidateSelectFiltersDeduplicateSelectFiltersEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateSetLeftEliminateSetRightEliminateDistinctSetLeftEliminateDistinctSetRightSimplifyExceptSimplifyIntersectLeftSimplifyIntersectRightConvertUnionToDistinctUnionAllEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleMemoCycleTestRelRuleReplaceScalarMinMaxWithScalarSubqueriesReplaceFilteredScalarMinMaxWithSubqueriesReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupBySplitGroupByScanIntoUnionScansSplitGroupByFilteredScanIntoUnionScansEliminateIndexJoinOrProjectInsideGroupByGenerateLimitedGroupByScansReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateLookupJoinsWithFilterGenerateLookupJoinsWithVirtualColsGenerateLookupJoinsWithVirtualColsAndFilterPushJoinIntoIndexJoinHoistProjectFromInnerJoinHoistProjectFromLeftJoinGenerateLocalityOptimizedAntiJoinGenerateLocalityOptimizedLookupJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitLimitedScanIntoUnionScansSplitLimitedSelectIntoUnionSelectsGenerateTopKGenerateLimitedTopKScansGeneratePartialOrderTopKEliminateIndexJoinInsideProjectGenerateIndexScansGenerateLocalityOptimizedScanGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsSplitDisjunctionSplitDisjunctionAddKeyGenerateStreamingSetOpNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 917, 928, 938, 949, 976, 1003, 1018, 1039, 1059, 1080, 1107, 1137, 1160, 1187, 1208, 1235, 1257, 1279, 1298, 1322, 1342, 1363, 1380, 1400, 1419, 1439, 1456, 1475, 1498, 1522, 1544, 1571, 1596, 1608, 1621, 1639, 1658, 1676, 1687, 1701, 1710, 1720, 1729, 1743, 1751, 1769, 1784, 1800, 1823, 1835, 1852, 1876, 1899, 1928, 1958, 1975, 1993, 2018, 2045, 2080, 2106, 2138, 2163, 2189, 2219, 2242, 2271, 2292, 2306, 2328, 2349, 2372, 2396, 2426, 2452, 2474, 2490, 2509, 2532, 2562, 2583, 2605, 2636, 2658, 2681, 2697, 2714, 2731, 2762, 2779, 2810, 2833, 2857, 2878, 2898, 2925, 2946, 2963, 2985, 3008, 3031, 3055, 3081, 3103, 3117, 3132, 3152, 3173, 3192, 3215, 3236, 3258, 3268, 3291, 3315, 3331, 3362, 3374, 3386, 3399, 3410, 3421, 3431, 3442, 3461, 3482, 3504, 3527, 3553, 3576, 3603, 3632, 3662, 3678, 3691, 3713, 3742, 3771, 3796, 3820, 3844, 3860, 3873, 3888, 3902, 3917, 3934, 3952, 3978, 3990, 4006, 4021, 4040, 4056, 4075, 4096, 4116, 4138, 4160, 4183, 4200, 4213, 4230, 4249, 4269, 4287, 4311, 4336, 4354, 4364, 4376, 4393, 4409, 4422, 4438, 4448, 4471, 4497, 4517, 4540, 4562, 4584, 4607, 4627, 4647, 4673, 4697, 4720, 4742, 4764, 4775, 4801, 4826, 4853, 4872, 4893, 4917, 4941, 4956, 4968, 4989, 5009, 5047, 5069, 5090, 5112, 5136, 5155, 5171, 5188, 5212, 5237, 5251, 5272, 5294, 5324, 5339, 5364, 5386, 5406, 5425, 5435, 5451, 5471, 5510, 5551, 5579, 5598, 5617, 5641, 5671, 5709, 5749, 5776, 5788, 5803, 5818, 5840, 5858, 5877, 5898, 5929, 5958, 5992, 6035, 6056, 6081, 6105, 6138, 6173, 6193, 6218, 6240, 6270, 6304, 6316, 6340, 6364, 6395, 6413, 6442, 6467, 6491, 6517, 6536, 6568, 6584, 6606, 6628, 6640}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}