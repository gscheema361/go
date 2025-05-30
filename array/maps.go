package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":    "https://www.google.com",
		"Facebook":  "https://www.facebook.com",
		"Instagram": "https://www.instagram.com",
		"Aws":       "https://aws.amazon.com",
		"Amazon":    "https://www.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	for key, value := range websites {
		fmt.Println(key, value)
	}
	websites["Linkedin"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
