package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sort"
)

func Pwd() {
	password := "123"
	//hashedPassword := "$2b$12$Y8S9341NSNzpr1捡1VZlO.9Ed3rKFWsVBZ8dy.7V4JLO7厮iUfKLOuC"
	// 对密码进行加盐并生成哈希值
	newhashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating password hash:", err)
		return
	}
	n, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(string(newhashedPassword), "--vs--", string(n))
	err = bcrypt.CompareHashAndPassword(n, []byte(password))
	if err == nil {
		fmt.Println("Password correct")
	} else {
		fmt.Println("Password incorrect")
	}
}
func Search(nums []int, target int) int {
	left, right := 0, len(nums)
	var isorder func(int, int, int) int
	isorder = func(l, mid, r int) int {
		var res int
		if nums[l] <= nums[mid-1] { //左有序
			if target <= nums[mid-1] && target >= nums[l] { //在左有序中0
				return res
			} else {
				res += 2
				if nums[mid] < nums[r-1] { //右也有序
					if target <= nums[r-1] && target >= nums[mid] { //在右有序中2
						return res
					} else { //左右都有序都没有5
						return 5
					}
				} else { //在右无序中3
					res++
					return res
				}
			}
		} else { //右有序
			if target <= nums[r-1] && target >= nums[mid] { //在右有序2
				return 2
			} else {
				return 1 //在左无序1
			}
		}
		return 0
	}
	var half func(int, int, bool) int
	half = func(l int, r int, flag bool) int {
		if r-l <= 1 {
			if nums[l] == target {
				return l
			} else {
				return -1
			}
		} else if r-l == 2 {
			if nums[l] == target {
				return l
			} else if nums[l+1] == target {
				return l + 1
			} else {
				return -1
			}
		} else {
			mid := (l + r) / 2
			if flag { //有序
				if nums[mid] > target {
					return half(l, mid, flag)
				} else if nums[mid] < target {
					return half(mid+1, r, flag)
				} else {
					return mid
				}
			} else { //无序
				temp := isorder(l, mid, r)
				if temp == 5 {
					return -1
				}
				if temp < 2 { //在左边
					return half(l, mid, temp%2 == 0)
				} else { //在右边
					return half(mid, r, temp%2 == 0)
				}
			}
		}
		return -1
	}
	res := half(left, right, false)
	return res
}
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(int, int)
	sort.Ints(candidates)
	dfs = func(target, start int) {
		if target == 0 {
			res = append(res, append([]int{}, tmp...))
			return
		}
		if start == len(candidates) {
			return
		}
		dfs(target, start+1)
		if target-candidates[start] >= 0 {
			tmp = append(tmp, candidates[start])
			dfs(target-candidates[start], start+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return res
}
