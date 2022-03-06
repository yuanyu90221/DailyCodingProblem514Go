# DailyCodingProblem514Go

## 題目描述

給定一個未排序的整數陣列，找出最長連續數列的長度。

舉例來說， 給定 [100, 4, 200, 1, 3, 2] ， 最長連續數列是 [1, 2, 3, 4] ， 所以最長長度是 4 。

實作的演算法需要符合時間複雜度為 O(n) 。

## 我的解法

要找出最長連續數列長度

先設定最長連續數列長度為 longestAccum = 0

利用 HashSet 先把所有數值先存入

接著把 HashSet 內部所有元素

找出不具有該元素值 - 1的元素(也就是最小值元素)

然後紀錄目前找到的值為 currentNum ， 目前累計長度為 currentAccum = 1

然後檢查 HashSet 中有沒有 currentNum + 1 ，如果有則把 currentNum 與 currentAccum 都各加 1

如果 HashSet 沒有 currentNum + 1 ， 則將 longestAccum 設定為 longestAccum 與 currentAccum 的最大值

當走訪完所有 HashSet 中元素

longestAccum 則為最長連續長度

## 程式碼
### 主要邏輯
```golang
var longesetAccum int = 0
	valueSet := NewSet(nums...)
	var values []int = valueSet.Values()
	for _, num := range values {
		if !valueSet.Contains(num - 1) {
			currentNum := num
			currentAccum := 1
			for valueSet.Contains(currentNum + 1) {
				currentNum += 1
				currentAccum += 1
			}
			if currentAccum > longesetAccum {
				longesetAccum = currentAccum
			}
		}
	}
	return longesetAccum
```
### HashSet 實作
```golang
var itemExists = struct{}{}

type HashSet struct {
	Items map[int]struct{}
}

func NewSet(values ...int) *HashSet {
	set := &HashSet{Items: make(map[int]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (set *HashSet) Add(items ...int) {
	for _, item := range items {
		set.Items[item] = itemExists
	}
}

func (set *HashSet) Remove(items ...int) {
	for _, item := range items {
		delete(set.Items, item)
	}
}

func (set *HashSet) Contains(items ...int) bool {
	for _, item := range items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

func (set *HashSet) Size() int {
	return len(set.Items)
}

func (set *HashSet) Clear() {
	set.Items = make(map[int]struct{})
}

func (set *HashSet) Values() []int {
	values := make([]int, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

```