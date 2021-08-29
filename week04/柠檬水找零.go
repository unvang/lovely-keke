package week04

//很简单，看官方题解就好
func leonCharge(bills []int) bool{
	five ,ten := 0,0
	for _,v:= range bills{
		if v==5{
			five++
		}else if v == 10{
			if five ==0{
				return false
			}
			five--
			ten++
		}else{
			if five>0 && ten >0{
				five --
				ten --
			}else if five >=3{
				five = five -3
			}else{
				return false
			}
		}
	}
	return true
}
