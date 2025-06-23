package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func fineClosestNumber(nums []int) int {
	closestZero := nums[0] // เก็บค่าแรกใว้เช็ค ในตัวแปร ค่าที่ใกล้กับ เลข 0 // cls = -10000
	minAbs := abs(nums[0]) // ตัวแปรเก็บค่าที่น้อยที่สุด และ แปลงค่าสัมบูรณ์ //  min = 10000

	for i := 1; i < len(nums); i++ { //เริ่มต้นที่ num[1]
		num := nums[i]         //เก็บค่า num[1] // num = -10000
		currentAbs := abs(num) //แปลงเป็นค่าสัมบูรณ์ไม่ให้ติดลบ // curr = 10000
		//หาระยะทางที่ใกล้ 0 ที่สุด
		if currentAbs < minAbs { // 10000 < 10000
			minAbs = currentAbs //เก็บค่าที่น้อยที่สุด //
			closestZero = num   //เก็บค่าที่น้อยที่สุด //
		} else if currentAbs == minAbs { //มีค่าเท่ากัน  // 10000 = 10000
			if num > closestZero { //ถ้าค่า num มีค่ามากกว่า // -10000 > -10000
				closestZero = num // ให้เก็บค่านั้น //
			}
		}
	}

	return closestZero // cls = -10000
}

func main() {
	num1 := []int{-4, -2, 1, 4, 8}
	num2 := []int{2, -1, 1}
	num3 := []int{-10000, -10000}
	f := fineClosestNumber(num1)
	f2 := fineClosestNumber(num2)
	f3 := fineClosestNumber(num3)
	fmt.Println("Result1: ", f)
	fmt.Println("Result2: ", f2)
	fmt.Println("Result3: ", f3)
}
