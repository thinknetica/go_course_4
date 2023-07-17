/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution,
and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
*/
package main

import "reflect"

var a int

func main() {
	a = 10
	_ = a
	go double(a)

	res := twoSum([]int{2, 7, 11, 15}, 13)
	if !reflect.DeepEqual(res, [2]int{0, 2}) {
		println("неверный результат")
		return
	}
	println("Номера элементов:", res[0], res[1])

	res = twoSum([]int{2, 10, 6, 12, 14}, 20)
	if !reflect.DeepEqual(res, [2]int{2, 4}) {
		println("неверный результат")
		return
	}
	println("Номера элементов:", res[0], res[1])
}

func twoSum(nums []int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{-1, -1}
}

func double(a int) int {
	b := a * a
	return b
}
