package main

// functions relating to accessing the scratch api,
// and converting project code to structs

import (
	//"io"
	//"net/http"
	"os"
)

// function for getting information about a project
func GetProject() (data *Project, err error) {
	// todo: allow project IDs to be specified

	//res, err := http.Get("https://projects.scratch.mit.edu/600000129/")
	//if(err != nil) {
	//	return nil, err
	//}
	//body, err := io.ReadAll(res.Body)
	//if(err != nil) {
	//	return nil, err
	//}

	body, err := os.ReadFile("./test.json")
	if(err != nil) {
		return nil, err
	}
	data, err = JSONToMemory(body)
	if(err != nil) {
		return nil, err
	}
	return data, nil
}