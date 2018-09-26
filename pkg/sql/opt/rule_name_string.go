// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

const _RuleName_name = "InvalidRuleNameSimplifyProjectOrderingSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctEliminateEmptyFiltersEliminateSingletonAndOrSimplifyAndSimplifyOrSimplifyFiltersFoldNullAndOrFoldNotTrueFoldNotFalseNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantClauseExtractRedundantSubclauseCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateZipHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistZipSubqueryNormalizeAnyFilterNormalizeNotAnyFilterFoldArrayFoldBinaryFoldUnaryFoldComparisonConvertGroupByToDistinctEliminateDistinctEliminateGroupByProjectReduceGroupingColsEliminateAggDistinctForKeysPushSelectIntoInlinableProjectInlineProjectInProjectPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinWithoutFiltersSimplifyRightJoinWithoutFiltersSimplifyLeftJoinWithFiltersSimplifyRightJoinWithFiltersEliminateSemiJoinEliminateAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectSimplifyJoinNotNullEqualityExtractJoinEqualitiesEliminateLimitPushLimitIntoProjectPushOffsetIntoProjectEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyRowNumberOrderingSimplifyExplainOrderingEliminateProjectEliminateProjectProjectPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneRowNumberColsPruneExplainColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldNullInEmptyFoldNullNotInEmptyNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsProjectEliminateExistsGroupByNormalizeJSONFieldAccessNormalizeJSONContainsSimplifyCaseWhenConstValueEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectCondRightIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoJoinRightPushSelectIntoGroupByRemoveNotNullConditionEliminateUnionAllLeftEliminateUnionAllRightstartExploreRuleReplaceMinWithLimitReplaceMaxWithLimitCommuteJoinCommuteLeftJoinCommuteRightJoinGenerateMergeJoinsGenerateLookupJoinsGenerateLookupJoinsWithFilterGenerateLimitedScansPushLimitIntoConstrainedScanPushLimitIntoIndexJoinGenerateIndexScansGenerateConstrainedScansGenerateInvertedIndexScansNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 38, 58, 71, 99, 117, 137, 158, 181, 192, 202, 217, 230, 241, 253, 269, 281, 290, 298, 320, 345, 365, 387, 408, 430, 452, 474, 496, 519, 529, 546, 559, 579, 592, 607, 627, 648, 675, 705, 728, 755, 776, 803, 825, 847, 864, 881, 901, 920, 940, 957, 976, 992, 1010, 1031, 1040, 1050, 1059, 1073, 1097, 1114, 1137, 1155, 1182, 1212, 1234, 1264, 1285, 1307, 1329, 1352, 1382, 1413, 1440, 1468, 1485, 1502, 1525, 1549, 1565, 1592, 1613, 1627, 1647, 1668, 1684, 1696, 1708, 1721, 1732, 1743, 1753, 1764, 1783, 1804, 1826, 1849, 1874, 1897, 1913, 1936, 1952, 1965, 1980, 1994, 2009, 2026, 2044, 2056, 2072, 2087, 2105, 2121, 2140, 2160, 2178, 2188, 2200, 2217, 2233, 2246, 2258, 2271, 2289, 2308, 2326, 2341, 2359, 2375, 2385, 2405, 2427, 2449, 2473, 2494, 2520, 2535, 2547, 2568, 2588, 2626, 2665, 2687, 2710, 2731, 2753, 2774, 2796, 2812, 2831, 2850, 2861, 2876, 2892, 2910, 2929, 2958, 2978, 3006, 3028, 3046, 3070, 3096, 3108}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
