package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	str := ""
	for i:=0;i<len(s);i++{
		//fmt.Println(string(s[i]));
		if string(s[i])==" " {
			//fmt.Println(str)
			mp[str] = mp[str]+1
			str = ""
		} else {
			str= str+string(s[i]);
		}
		
	}
	
	if str != ""{
		mp[str] = mp[str]+1
	}
	return mp;
}

func main() {
	wc.Test(WordCount)
	
}
