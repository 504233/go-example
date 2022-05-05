package main

import "fmt"

func main() {
	var liftrun = makelift()
	runstart(liftrun)
}

// 电梯对象
type lift struct {
	floorlist    []int //停靠楼层
	rundirection int   //运行方向(上行1，下行2，停止0)
	floornow     int   //当前停靠楼层（默认1）
	floormun     int   //楼层总数（默认5）
}

// 实例化电梯对象
func makelift() lift {
	var floor lift
	floor.floormun = 5
	floor.rundirection = 0
	floor.floornow = 1
	floor.floorlist = append(floor.floorlist, 1)
	return floor
}

//电梯运行方向查询
func selectdirection(ststus lift) string {
	var direction = ststus.rundirection
	switch direction {
	case 1:
		return "上行"
	case 2:
		return "下行"
	case 0:
		return "静止"
	default:
		return "静止"
	}
}

//查询楼层是否重复或超过数值
func floorcheck(target int, liftrun lift) {
	for _, value := range liftrun.floorlist {
		if value == target {
			fmt.Println("已选择该楼层，请勿重复选择\n")
		} else if target > liftrun.floormun || target <= 0 {
			fmt.Println("仅可选择0-", liftrun.floormun, "之间的楼层，请重新选择\n")
			runtime(liftrun)
		} else {
			liftrun.floorlist[0] = target
			fmt.Println("添加成功\n")
		}
	}
}

//检查乘客中途上梯
func checkcuston(newcostum string, liftrun lift) {
	var target int
	if newcostum == "y" {
		fmt.Println("请输入需要的楼层：\n")
		fmt.Scan(&target)
		floorcheck(target, liftrun)
	} else if newcostum == "n" {
		return
	} else {
		fmt.Println("输入错误，请重新输入：\n")
		checkcuston(newcostum, liftrun)
	}

}

//电梯欢迎函数
func runstart(liftrun lift) {
	var target int
	fmt.Println("欢迎乘坐电梯，当前停靠楼层：", liftrun.floornow, "    电梯运行方向：", selectdirection(liftrun), "\n")
	fmt.Println("请输入需要的楼层：\n")
	fmt.Scan(&target)
	floorcheck(target, liftrun)
	runtime(liftrun)

}

//判断并切换电梯运行方向
func changedirection(liftrun lift) {
	if liftrun.floornow == 1 {
		liftrun.rundirection = 1
	} else if liftrun.floornow == liftrun.floormun {
		liftrun.rundirection = 2
	} else {
		liftrun.rundirection = 0
	}
}

//电梯运行函数
func runtime(liftrun lift) {
	var newcostum string
	changedirection(liftrun)

	//判断是否为初始化状态
	if liftrun.floorlist[0] == 1 {
		fmt.Println("请输入需要的楼层：\n")
	} else {
		switch liftrun.rundirection {
		case 1:
			liftrun.floornow += 1
		case 2:
			liftrun.floornow -= 1
		case 0:
			liftrun.floornow += 0
		default:
			liftrun.floornow += 0
		}

		//遍历楼层判断是否选中
		for index, value := range liftrun.floorlist {
			if value == liftrun.floornow {
				fmt.Println(liftrun.floornow, "层到了")
				liftrun.floorlist[index] = 0
			}
		}

		fmt.Println("电梯运行，当前运行方向：", selectdirection(liftrun), "当前楼层：", liftrun.floornow, "是否有新乘客上梯[y/n]\n")
		fmt.Scan(&newcostum)
		checkcuston(newcostum, liftrun)

	}
}
