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
	_ = x[DecorrelateJoin-53]
	_ = x[DecorrelateProjectSet-54]
	_ = x[TryDecorrelateSelect-55]
	_ = x[TryDecorrelateProject-56]
	_ = x[TryDecorrelateProjectSelect-57]
	_ = x[TryDecorrelateProjectInnerJoin-58]
	_ = x[TryDecorrelateInnerJoin-59]
	_ = x[TryDecorrelateInnerLeftJoin-60]
	_ = x[TryDecorrelateGroupBy-61]
	_ = x[TryDecorrelateScalarGroupBy-62]
	_ = x[TryDecorrelateSemiJoin-63]
	_ = x[TryDecorrelateLimitOne-64]
	_ = x[TryDecorrelateProjectSet-65]
	_ = x[TryDecorrelateWindow-66]
	_ = x[TryDecorrelateMax1Row-67]
	_ = x[HoistSelectExists-68]
	_ = x[HoistSelectNotExists-69]
	_ = x[HoistSelectSubquery-70]
	_ = x[HoistProjectSubquery-71]
	_ = x[HoistJoinSubquery-72]
	_ = x[HoistValuesSubquery-73]
	_ = x[HoistProjectSetSubquery-74]
	_ = x[NormalizeSelectAnyFilter-75]
	_ = x[NormalizeJoinAnyFilter-76]
	_ = x[NormalizeSelectNotAnyFilter-77]
	_ = x[NormalizeJoinNotAnyFilter-78]
	_ = x[FoldNullCast-79]
	_ = x[FoldNullUnary-80]
	_ = x[FoldNullBinaryLeft-81]
	_ = x[FoldNullBinaryRight-82]
	_ = x[FoldNullInNonEmpty-83]
	_ = x[FoldInEmpty-84]
	_ = x[FoldNotInEmpty-85]
	_ = x[FoldArray-86]
	_ = x[FoldBinary-87]
	_ = x[FoldUnary-88]
	_ = x[FoldComparison-89]
	_ = x[FoldCast-90]
	_ = x[FoldIndirection-91]
	_ = x[FoldColumnAccess-92]
	_ = x[FoldFunctionWithNullArg-93]
	_ = x[FoldFunction-94]
	_ = x[FoldEqualsAnyNull-95]
	_ = x[ConvertGroupByToDistinct-96]
	_ = x[EliminateGroupByProject-97]
	_ = x[EliminateJoinUnderGroupByLeft-98]
	_ = x[EliminateJoinUnderGroupByRight-99]
	_ = x[EliminateDistinct-100]
	_ = x[ReduceGroupingCols-101]
	_ = x[ReduceNotNullGroupingCols-102]
	_ = x[EliminateAggDistinctForKeys-103]
	_ = x[EliminateAggFilteredDistinctForKeys-104]
	_ = x[EliminateDistinctNoColumns-105]
	_ = x[EliminateEnsureDistinctNoColumns-106]
	_ = x[EliminateDistinctOnValues-107]
	_ = x[PushAggDistinctIntoGroupBy-108]
	_ = x[PushAggFilterIntoScalarGroupBy-109]
	_ = x[ConvertCountToCountRows-110]
	_ = x[ConvertRegressionCountToCount-111]
	_ = x[FoldGroupingOperators-112]
	_ = x[InlineConstVar-113]
	_ = x[InlineProjectConstants-114]
	_ = x[InlineSelectConstants-115]
	_ = x[InlineJoinConstantsLeft-116]
	_ = x[InlineJoinConstantsRight-117]
	_ = x[PushSelectIntoInlinableProject-118]
	_ = x[InlineSelectVirtualColumns-119]
	_ = x[InlineProjectInProject-120]
	_ = x[CommuteRightJoin-121]
	_ = x[SimplifyJoinFilters-122]
	_ = x[DetectJoinContradiction-123]
	_ = x[PushFilterIntoJoinLeftAndRight-124]
	_ = x[MapFilterIntoJoinLeft-125]
	_ = x[MapFilterIntoJoinRight-126]
	_ = x[MapEqualityIntoJoinLeftAndRight-127]
	_ = x[PushFilterIntoJoinLeft-128]
	_ = x[PushFilterIntoJoinRight-129]
	_ = x[SimplifyLeftJoin-130]
	_ = x[SimplifyRightJoin-131]
	_ = x[EliminateSemiJoin-132]
	_ = x[SimplifyZeroCardinalitySemiJoin-133]
	_ = x[EliminateAntiJoin-134]
	_ = x[SimplifyZeroCardinalityAntiJoin-135]
	_ = x[EliminateJoinNoColsLeft-136]
	_ = x[EliminateJoinNoColsRight-137]
	_ = x[HoistJoinProjectRight-138]
	_ = x[HoistJoinProjectLeft-139]
	_ = x[SimplifyJoinNotNullEquality-140]
	_ = x[ExtractJoinEqualities-141]
	_ = x[SortFiltersInJoin-142]
	_ = x[LeftAssociateJoinsLeft-143]
	_ = x[LeftAssociateJoinsRight-144]
	_ = x[RightAssociateJoinsLeft-145]
	_ = x[RightAssociateJoinsRight-146]
	_ = x[RemoveJoinNotNullCondition-147]
	_ = x[ProjectInnerJoinValues-148]
	_ = x[EliminateLimit-149]
	_ = x[EliminateOffset-150]
	_ = x[PushLimitIntoProject-151]
	_ = x[PushOffsetIntoProject-152]
	_ = x[PushLimitIntoOffset-153]
	_ = x[PushLimitIntoOrdinality-154]
	_ = x[PushLimitIntoJoinLeft-155]
	_ = x[PushLimitIntoJoinRight-156]
	_ = x[FoldLimits-157]
	_ = x[AssociateLimitJoinsLeft-158]
	_ = x[AssociateLimitJoinsRight-159]
	_ = x[EliminateMax1Row-160]
	_ = x[SimplifyPartialIndexProjections-161]
	_ = x[FoldPlusZero-162]
	_ = x[FoldZeroPlus-163]
	_ = x[FoldMinusZero-164]
	_ = x[FoldMultOne-165]
	_ = x[FoldOneMult-166]
	_ = x[FoldDivOne-167]
	_ = x[InvertMinus-168]
	_ = x[EliminateUnaryMinus-169]
	_ = x[SimplifyLimitOrdering-170]
	_ = x[SimplifyOffsetOrdering-171]
	_ = x[SimplifyGroupByOrdering-172]
	_ = x[SimplifyOrdinalityOrdering-173]
	_ = x[SimplifyExplainOrdering-174]
	_ = x[EliminateJoinUnderProjectLeft-175]
	_ = x[EliminateJoinUnderProjectRight-176]
	_ = x[EliminateProject-177]
	_ = x[MergeProjects-178]
	_ = x[MergeProjectWithValues-179]
	_ = x[PushColumnRemappingIntoValues-180]
	_ = x[FoldTupleAccessIntoValues-181]
	_ = x[FoldJSONAccessIntoValues-182]
	_ = x[ConvertZipArraysToValues-183]
	_ = x[PruneProjectCols-184]
	_ = x[PruneScanCols-185]
	_ = x[PruneSelectCols-186]
	_ = x[PruneLimitCols-187]
	_ = x[PruneOffsetCols-188]
	_ = x[PruneJoinLeftCols-189]
	_ = x[PruneJoinRightCols-190]
	_ = x[PruneSemiAntiJoinRightCols-191]
	_ = x[PruneAggCols-192]
	_ = x[PruneGroupByCols-193]
	_ = x[PruneValuesCols-194]
	_ = x[PruneOrdinalityCols-195]
	_ = x[PruneExplainCols-196]
	_ = x[PruneProjectSetCols-197]
	_ = x[PruneWindowOutputCols-198]
	_ = x[PruneWindowInputCols-199]
	_ = x[PruneMutationFetchCols-200]
	_ = x[PruneMutationInputCols-201]
	_ = x[PruneMutationReturnCols-202]
	_ = x[PruneWithScanCols-203]
	_ = x[PruneWithCols-204]
	_ = x[PruneUnionAllCols-205]
	_ = x[RejectNullsLeftJoin-206]
	_ = x[RejectNullsRightJoin-207]
	_ = x[RejectNullsGroupBy-208]
	_ = x[RejectNullsUnderJoinLeft-209]
	_ = x[RejectNullsUnderJoinRight-210]
	_ = x[RejectNullsProject-211]
	_ = x[CommuteVar-212]
	_ = x[CommuteConst-213]
	_ = x[EliminateCoalesce-214]
	_ = x[SimplifyCoalesce-215]
	_ = x[EliminateCast-216]
	_ = x[NormalizeInConst-217]
	_ = x[FoldInNull-218]
	_ = x[SimplifyInSingleElement-219]
	_ = x[SimplifyNotInSingleElement-220]
	_ = x[UnifyComparisonTypes-221]
	_ = x[EliminateExistsZeroRows-222]
	_ = x[EliminateExistsProject-223]
	_ = x[EliminateExistsGroupBy-224]
	_ = x[InlineExistsSelectTuple-225]
	_ = x[IntroduceExistsLimit-226]
	_ = x[EliminateExistsLimit-227]
	_ = x[SimplifyCaseWhenConstValue-228]
	_ = x[InlineAnyValuesSingleCol-229]
	_ = x[InlineAnyValuesMultiCol-230]
	_ = x[SimplifyEqualsAnyTuple-231]
	_ = x[SimplifyAnyScalarArray-232]
	_ = x[FoldCollate-233]
	_ = x[NormalizeArrayFlattenToAgg-234]
	_ = x[SimplifySameVarEqualities-235]
	_ = x[SimplifySameVarInequalities-236]
	_ = x[SimplifyNotDisjoint-237]
	_ = x[SimplifySelectFilters-238]
	_ = x[ConsolidateSelectFilters-239]
	_ = x[EliminateSelect-240]
	_ = x[MergeSelects-241]
	_ = x[PushSelectIntoProject-242]
	_ = x[MergeSelectInnerJoin-243]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-244]
	_ = x[PushSelectIntoJoinLeft-245]
	_ = x[PushSelectIntoGroupBy-246]
	_ = x[RemoveNotNullCondition-247]
	_ = x[PushSelectIntoProjectSet-248]
	_ = x[PushFilterIntoSetOp-249]
	_ = x[EliminateUnionAllLeft-250]
	_ = x[EliminateUnionAllRight-251]
	_ = x[EliminateUnionLeft-252]
	_ = x[EliminateUnionRight-253]
	_ = x[EliminateWindow-254]
	_ = x[ReduceWindowPartitionCols-255]
	_ = x[SimplifyWindowOrdering-256]
	_ = x[PushSelectIntoWindow-257]
	_ = x[PushLimitIntoWindow-258]
	_ = x[InlineWith-259]
	_ = x[startExploreRule-260]
	_ = x[ReplaceScalarMinMaxWithLimit-261]
	_ = x[ReplaceMinWithLimit-262]
	_ = x[ReplaceMaxWithLimit-263]
	_ = x[GenerateStreamingGroupBy-264]
	_ = x[SplitGroupByScanIntoUnionScans-265]
	_ = x[SplitGroupByFilteredScanIntoUnionScans-266]
	_ = x[ReorderJoins-267]
	_ = x[CommuteLeftJoin-268]
	_ = x[CommuteSemiJoin-269]
	_ = x[ConvertSemiToInnerJoin-270]
	_ = x[GenerateMergeJoins-271]
	_ = x[GenerateLookupJoins-272]
	_ = x[GenerateInvertedJoins-273]
	_ = x[GenerateInvertedJoinsFromSelect-274]
	_ = x[GenerateLookupJoinsWithFilter-275]
	_ = x[PushJoinIntoIndexJoin-276]
	_ = x[HoistProjectFromInnerJoin-277]
	_ = x[HoistProjectFromLeftJoin-278]
	_ = x[GenerateLocalityOptimizedAntiJoin-279]
	_ = x[GenerateLocalityOptimizedLookupJoin-280]
	_ = x[GenerateLimitedScans-281]
	_ = x[PushLimitIntoFilteredScan-282]
	_ = x[PushLimitIntoIndexJoin-283]
	_ = x[SplitLimitedScanIntoUnionScans-284]
	_ = x[EliminateIndexJoinInsideProject-285]
	_ = x[GenerateIndexScans-286]
	_ = x[GenerateLocalityOptimizedScan-287]
	_ = x[GeneratePartialIndexScans-288]
	_ = x[GenerateConstrainedScans-289]
	_ = x[GenerateInvertedIndexScans-290]
	_ = x[GenerateZigzagJoins-291]
	_ = x[GenerateInvertedIndexZigzagJoins-292]
	_ = x[SplitDisjunction-293]
	_ = x[SplitDisjunctionAddKey-294]
	_ = x[NumRuleNames-295]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldNullTupleIsTupleNullFoldNonNullTupleIsTupleNullFoldIsNotNullFoldNonNullIsNotNullFoldNonNullTupleIsTupleNotNullFoldNullTupleIsTupleNotNullCommuteNullIsNormalizeCmpTimeZoneFunctionNormalizeCmpTimeZoneFunctionTZFoldEqZeroSTDistanceFoldCmpSTDistanceLeftFoldCmpSTDistanceRightFoldCmpSTMaxDistanceLeftFoldCmpSTMaxDistanceRightFoldEqTrueFoldEqFalseFoldNeTrueFoldNeFalseDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionWithNullArgFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateGroupByProjectEliminateJoinUnderGroupByLeftEliminateJoinUnderGroupByRightEliminateDistinctReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsConvertRegressionCountToCountFoldGroupingOperatorsInlineConstVarInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineSelectVirtualColumnsInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinLeftAssociateJoinsLeftLeftAssociateJoinsRightRightAssociateJoinsLeftRightAssociateJoinsRightRemoveJoinNotNullConditionProjectInnerJoinValuesEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoJoinLeftPushLimitIntoJoinRightFoldLimitsAssociateLimitJoinsLeftAssociateLimitJoinsRightEliminateMax1RowSimplifyPartialIndexProjectionsFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateJoinUnderProjectLeftEliminateJoinUnderProjectRightEliminateProjectMergeProjectsMergeProjectWithValuesPushColumnRemappingIntoValuesFoldTupleAccessIntoValuesFoldJSONAccessIntoValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByRejectNullsUnderJoinLeftRejectNullsUnderJoinRightRejectNullsProjectCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullSimplifyInSingleElementSimplifyNotInSingleElementUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByInlineExistsSelectTupleIntroduceExistsLimitEliminateExistsLimitSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifyNotDisjointSimplifySelectFiltersConsolidateSelectFiltersEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateUnionLeftEliminateUnionRightEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupBySplitGroupByScanIntoUnionScansSplitGroupByFilteredScanIntoUnionScansReorderJoinsCommuteLeftJoinCommuteSemiJoinConvertSemiToInnerJoinGenerateMergeJoinsGenerateLookupJoinsGenerateInvertedJoinsGenerateInvertedJoinsFromSelectGenerateLookupJoinsWithFilterPushJoinIntoIndexJoinHoistProjectFromInnerJoinHoistProjectFromLeftJoinGenerateLocalityOptimizedAntiJoinGenerateLocalityOptimizedLookupJoinGenerateLimitedScansPushLimitIntoFilteredScanPushLimitIntoIndexJoinSplitLimitedScanIntoUnionScansEliminateIndexJoinInsideProjectGenerateIndexScansGenerateLocalityOptimizedScanGeneratePartialIndexScansGenerateConstrainedScansGenerateInvertedIndexScansGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 607, 634, 647, 667, 697, 724, 737, 765, 795, 815, 836, 858, 882, 907, 917, 928, 938, 949, 964, 985, 1005, 1026, 1053, 1083, 1106, 1133, 1154, 1181, 1203, 1225, 1249, 1269, 1290, 1307, 1327, 1346, 1366, 1383, 1402, 1425, 1449, 1471, 1498, 1523, 1535, 1548, 1566, 1585, 1603, 1614, 1628, 1637, 1647, 1656, 1670, 1678, 1693, 1709, 1732, 1744, 1761, 1785, 1808, 1837, 1867, 1884, 1902, 1927, 1954, 1989, 2015, 2047, 2072, 2098, 2128, 2151, 2180, 2201, 2215, 2237, 2258, 2281, 2305, 2335, 2361, 2383, 2399, 2418, 2441, 2471, 2492, 2514, 2545, 2567, 2590, 2606, 2623, 2640, 2671, 2688, 2719, 2742, 2766, 2787, 2807, 2834, 2855, 2872, 2894, 2917, 2940, 2964, 2990, 3012, 3026, 3041, 3061, 3082, 3101, 3124, 3145, 3167, 3177, 3200, 3224, 3240, 3271, 3283, 3295, 3308, 3319, 3330, 3340, 3351, 3370, 3391, 3413, 3436, 3462, 3485, 3514, 3544, 3560, 3573, 3595, 3624, 3649, 3673, 3697, 3713, 3726, 3741, 3755, 3770, 3787, 3805, 3831, 3843, 3859, 3874, 3893, 3909, 3928, 3949, 3969, 3991, 4013, 4036, 4053, 4066, 4083, 4102, 4122, 4140, 4164, 4189, 4207, 4217, 4229, 4246, 4262, 4275, 4291, 4301, 4324, 4350, 4370, 4393, 4415, 4437, 4460, 4480, 4500, 4526, 4550, 4573, 4595, 4617, 4628, 4654, 4679, 4706, 4725, 4746, 4770, 4785, 4797, 4818, 4838, 4876, 4898, 4919, 4941, 4965, 4984, 5005, 5027, 5045, 5064, 5079, 5104, 5126, 5146, 5165, 5175, 5191, 5219, 5238, 5257, 5281, 5311, 5349, 5361, 5376, 5391, 5413, 5431, 5450, 5471, 5502, 5531, 5552, 5577, 5601, 5634, 5669, 5689, 5714, 5736, 5766, 5797, 5815, 5844, 5869, 5893, 5919, 5938, 5970, 5986, 6008, 6020}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
