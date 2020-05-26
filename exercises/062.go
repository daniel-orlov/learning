package main

//Build and use an anonymous func

import "fmt"

func main()  {
	scandinavia := map[string]string{
		"Sweden": "Stockholm",
		"Norway": "Oslo",
		"Denmark": "Copenhagen",
		"Iceland": "Rejkjavik",
	}
	func (){
		for k,c := range scandinavia {
			if len(k) > len(c) {
				fmt.Printf("%v: true\n", k)
			}else{
				fmt.Printf("%v: false\n", k)
			}
		}
	}()
}

