// Inspiration boiler plate code logic for LinkedIn to consider
// full work duration on anniversary updates
// NOT FINISHED :)

package main

import (
	"fmt"
	"time"
)

type account struct {
	userName string
	firstName string
	lastName string
	email string
  workExperience []workExperience
}

type workExperience struct {
	company string
	currentCompany bool
	title string
	startDate time.Time
	endDate time.Time
}

// TODO - func logic should already exist in main program | no need to recreate
// func isWorkExperienceAnniversary(acct account) {
// 	currentDate := time.Now()

// 	for _, currentRole := range acct.workExperience {
// 		if currentRole.currentCompany && currentRole.endDate.IsZero() {

// 			fmt.Println("After: ", currentDate.AddDate(0, -12, 0))
// 			fmt.Println("Before: ", currentDate.AddDate(0, 11, 0))

// 		}
// 	}
// }


func calcAllCurrentCompanyWorkDuration(acct *account) int {

	workExperienceJobDuration := []int{}
	totalMonths := 0

	for _, currentRole := range acct.workExperience {
		if currentRole.currentCompany {
			if currentRole.endDate.IsZero() {
				years := time.Now().Year() - currentRole.startDate.Year()
				months := time.Now().Month() - currentRole.startDate.Month()
				workExperienceJobDuration = append(workExperienceJobDuration, (years * 12) + int(months))
			} else {
				years := currentRole.endDate.Year() - currentRole.startDate.Year()
				months := currentRole.endDate.Month() - currentRole.startDate.Month()
				workExperienceJobDuration = append(workExperienceJobDuration, (years * 12) + int(months))
			}
		}
	}

	for _, val := range workExperienceJobDuration {
		totalMonths += val
	}

	return totalMonths
}

func newAccount(userName, firstName, lastName, email string, experiences *[]workExperience) (account, error) {

	return account{
		userName: userName,
		firstName: firstName,
		lastName: lastName,
		email: email,
		workExperience: *experiences,
	}, nil
}

func main() {


	workExperiences := []workExperience {
		{
			company: "Microsoft",
			currentCompany: true,
			title: "Senior Product Manager - Azure Compute",
			startDate: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			company: "Microsoft",
			currentCompany: true,
			title: "Senior Azure Customer Engineer",
			startDate: time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
			endDate: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			company: "Docker",
			currentCompany: false,
			title: "Technical Alliance Manager",
			startDate: time.Date(2017, time.April, 1, 0, 0, 0, 0, time.UTC),
			endDate: time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			company: "Microsoft",
			currentCompany: true,
			title: "Cloud Solution Architect",
			startDate: time.Date(2016, time.April, 1, 0, 0, 0, 0, time.UTC),
			endDate: time.Date(2017, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			company: "Microsoft",
			currentCompany: true,
			title: "Solution Architect",
			startDate: time.Date(2013, time.October, 1, 0, 0, 0, 0, time.UTC),
			endDate: time.Date(2016, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			company: "Microsoft",
			currentCompany: true,
			title: "Senior Consultant",
			startDate: time.Date(2008, time.January, 1, 0, 0, 0, 0, time.UTC),
			endDate: time.Date(2013, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	phillipgibson, err := newAccount("phillipgibson", "Phill", "Gibson", "phillipgibson@notreal.com", &workExperiences)

	if err != nil {
		fmt.Println("please check your account info for input errors")

	}


	// isWorkExperienceAnniversary(phillipgibson)
	totalWorkDurationMonths := calcAllCurrentCompanyWorkDuration(&phillipgibson)
	fmt.Printf(`Congratulate %v %v on his current {TODO} year work anniversary. They have been at %v for a total of %v years and %v months.`, 
	phillipgibson.firstName, phillipgibson.lastName, phillipgibson.workExperience[0].company,
  (totalWorkDurationMonths / 12), (totalWorkDurationMonths % 12))
  

}