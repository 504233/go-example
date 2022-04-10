

func HumanBytesLoaded(bytesLength int64) string {
	var resp string

	// bytesLength
	fmt.Printf("需要转换的值: %dB", bytesLength)

	// 单位转换的逻辑 将 bytesLength ==> resp, 比如 1024 ==> 1KB
	// 1023 ==> 1023 Byte
	cosnt a = 10
	arr1 := []string{"B","KB","MB","GB","TB","EB"}
	var r int
	fmt.Scan(&r)
	var b float32=r/(1<<(a*i))
	for i := 1; i < 7; i++ {
		if b>1 {
			i +=1
			var b float32=r/(1<<(a*i))
		}else{
			resp = b+arr1[i-1]
		}
	}
	// 返回转换的结果
	return resp
}
