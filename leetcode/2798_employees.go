package main

func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
	employees := 0
	for _,v := range hours{
		if v >= target{
			employees++
		}
	}

	return employees
}