package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["China"] = "shenzhen"

	//使用key输出map值
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country,"is", countryCapitalMap[country])
	}

	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["China"]
	fmt.Println("captial:", capital, "ok:", ok)
	if(ok) {
		fmt.Println("Capital of China is", capital);
	}else {
		fmt.Println("Capital of China is not present")
	}

	tmpMap := map[string] string {"France":"Paris", "a":"b"}
	printMap(tmpMap)

	delete(tmpMap, "France")
	fmt.Println("删除后:")
	printMap(tmpMap)

}

func printMap(tmpMap map[string]string) {

	for key := range tmpMap {
		fmt.Println("key is", key, "value is", tmpMap[key])
	}
	
}
