// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

const _RuleName_name = "InvalidRuleNameSimplifyProjectOrderingSimplifyRootOrderingPruneRootColsNumManualRuleNamesEliminateEmptyAndEliminateEmptyOrEliminateSingletonAndOrSimplifyAndSimplifyOrSimplifyFiltersFoldNullAndOrNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantClauseExtractRedundantSubclauseCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateInnerJoinTryDecorrelateScalarGroupByHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryNormalizeAnyFilterNormalizeNotAnyFilterEliminateDistinctEliminateGroupByProjectPushSelectIntoInlinableProjectInlineProjectInProjectEnsureJoinFiltersAndEnsureJoinFiltersPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinEliminateAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightEliminateLimitPushLimitIntoProjectPushOffsetIntoProjectEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusFoldUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyRowNumberOrderingSimplifyExplainOrderingEliminateProjectEliminateProjectProjectPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneRowNumberColsPruneExplainColsCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldNullInEmptyFoldNullNotInEmptyNormalizeInConstFoldInNullEliminateExistsProjectEliminateExistsGroupByEliminateSelectEnsureSelectFiltersAndEnsureSelectFiltersMergeSelectsPushSelectIntoProjectSimplifySelectLeftJoinSimplifySelectRightJoinMergeSelectInnerJoinPushSelectIntoJoinLeftPushSelectIntoJoinRightPushSelectIntoGroupByRemoveNotNullConditionstartExploreRuleReplaceMinWithLimitCommuteJoinCommuteLeftJoinCommuteRightJoinGenerateMergeJoinsPushLimitIntoScanPushLimitIntoIndexJoinGenerateIndexScansConstrainScanPushFilterIntoIndexJoinNoRemainderPushFilterIntoIndexJoinConstrainIndexJoinScanNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 38, 58, 71, 89, 106, 122, 145, 156, 166, 181, 194, 210, 222, 231, 239, 261, 286, 306, 328, 349, 371, 393, 415, 437, 460, 470, 487, 500, 520, 533, 548, 568, 589, 616, 639, 666, 683, 703, 722, 742, 759, 778, 796, 817, 834, 857, 887, 909, 929, 946, 976, 997, 1019, 1041, 1064, 1080, 1097, 1114, 1131, 1154, 1178, 1192, 1212, 1233, 1249, 1261, 1273, 1286, 1297, 1308, 1318, 1329, 1348, 1362, 1383, 1405, 1428, 1453, 1476, 1492, 1515, 1531, 1544, 1559, 1573, 1588, 1605, 1623, 1635, 1651, 1666, 1684, 1700, 1710, 1722, 1739, 1755, 1768, 1780, 1793, 1811, 1830, 1848, 1863, 1881, 1897, 1907, 1929, 1951, 1966, 1988, 2007, 2019, 2040, 2062, 2085, 2105, 2127, 2150, 2171, 2193, 2209, 2228, 2239, 2254, 2270, 2288, 2305, 2327, 2345, 2358, 2392, 2415, 2437, 2449}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
