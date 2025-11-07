package imagesaver

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

type Resume struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	ProfileImage string `json:"profile_image"`
}

func SaveImage(c echo.Context) error {
	var resume Resume 
	
	if err := c.Bind(&resume); err != nil {
		return err
	}
	
	resumeBytes, err := json.Marshal(resume)
	if err != nil {
		return err
	}
	
	err = ioutil.WriteFile("output.json", resumeBytes, 0644)
	if err != nil {
		return err
	}

	return nil  
}