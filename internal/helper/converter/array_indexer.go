/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

type ArrIndexer struct {
	IndicesInts []int
}

func (arrIndex *ArrIndexer) New() {
	arrIndex.AppendIndex(0)
}

func (arrIndex *ArrIndexer) AppendIndex(index int) {
	arrIndex.IndicesInts = append(arrIndex.IndicesInts, index)
}

func (arrIndex *ArrIndexer) GetAllIndexes() []int {
	return arrIndex.IndicesInts
}

func (arrIndex *ArrIndexer) RemoveLastIndex() {
	if arrIndex.IsActive() {
		// Remove the last index from IndicesInts
		arrIndex.IndicesInts = arrIndex.IndicesInts[:len(arrIndex.IndicesInts)-1]
	}
}

func (arrIndex *ArrIndexer) IncrementLastIndex() {
	if arrIndex.IsActive() {
		arrIndex.IndicesInts[len(arrIndex.IndicesInts)-1]++
	}
}

func (arrIndex *ArrIndexer) IsActive() bool {
	return len(arrIndex.IndicesInts) > 0
}
