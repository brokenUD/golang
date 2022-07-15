package jz_sort

import (
	// 
)

// 描述
// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007

// 数据范围：  对于 50\%50% 的数据, size\leq 10^4size≤10 4
// 对于 100\%100% 的数据, size\leq 10^5size≤10 5
// 数组中所有数字的值满足 0 \le val \le 10000000≤val≤1000000

// 要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
// 输入描述：
// 题目保证输入的数组中没有的相同的数字

func InversePairs( data []int ) int {
    tmp := make([]int, len(data))
	count := 0
	mergeSort(data, tmp, 0, len(data)-1, &count)
	return count
}

func mergeSort(data, tmp []int, start, end int, count *int) {
    if start >= end {
        return  
    }
    
    mid := start + (end - start) >> 1 
    
    mergeSort(data, tmp, start, mid, count)
    mergeSort(data, tmp, mid+1, end, count)
    merge(data, tmp, start, mid, end, count)
    
    return 
    
}

func merge(data, tmp []int, start, mid, end int, count *int) {
    if start >= end {
        return 
    }
    
    i := start 
    j := mid + 1 
    k := start 
    for i <= mid && j <= end {
        if data[i] <= data[j] { 
            tmp[k] = data[i]
            i++
        } else {
            *count = *count + (mid - i + 1) 
            *count = *count % 1000000007
            tmp[k] = data[j]
            j++
        }
        k++ 
    }
    
    for i <= mid {
        tmp[k] = data[i]
        i++
        k++ 
    }
    
    for j <= end {
        tmp[k] = data[j]
        j++ 
        k++ 
    }
    
    for i = start; i <= end; i++ {
        data[i] = tmp[i]
    }
    
    return 
}
