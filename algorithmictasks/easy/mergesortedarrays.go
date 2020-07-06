package easy
/*
88. Merge Sorted Array

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]


Constraints:

-10^9 <= nums1[i], nums2[i] <= 10^9
nums1.length == m + n
nums2.length == n
Source: LeetCode.com
*/

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int)  {
	j:=0
	i:=0
	for j<n {
		if i<m {
			if nums1[i] > nums2[j] {
				ShiftArrayByOne(nums1,i)
				nums1[i]=nums2[j]
				j=j+1
				m=m+1
			}
			i = i+1

		} else {
			nums1[i]=nums2[j]
			j=j+1
			i=i+1
		}
	}
}


func ShiftArrayByOne(A[]int, pos1 int){
	current:= A[pos1]
	i:=pos1
	for i<=len(A)-2 {
		next := A[i+1]
		A[i+1]=current
		current = next
		i=i+1
	}
}