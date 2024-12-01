package map

import "fmt"

func map() {

	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon web services": "https://www.aws.com",
	}
	fmt.Println(websites)

	//mutating map
	fmt.Println(websites["Amazon web services"])

	//adding new website
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	//delete key
	delete(websites, "Amazon web services")
	fmt.Println(websites)

}
