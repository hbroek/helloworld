package main

import (
	"errors"
	"fmt"
)

const (
	genderBoy  = "boy"
	genderGirl = "girl"
)

var (
	defaultBoyNames = []string{
		"Liam", "Noah", "Oliver", "Elijah", "William", "James", "Benjamin", "Lucas", "Henry", "Theodore",
		"Alexander", "Michael", "Daniel", "Matthew", "Sebastian", "Jack", "Jayden", "John", "David", "Samuel",
	}

	defaultGirlNames = []string{
		"Olivia", "Emma", "Charlotte", "Amelia", "Sophia", "Isabella", "Ava", "Mia", "Ella", "Luna",
		"Camila", "Harper", "Evelyn", "Abigail", "Emily", "Elizabeth", "Sofia", "Mila", "Samantha", "Layla",
	}

	errNoNames = errors.New("no names configured")
)

type RandomFunc func(int) int

type NameService struct {
	boyNames  []string
	girlNames []string
	random    RandomFunc
}

type NameResponse struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Total   int    `json:"total"`
	Message string `json:"message"`
}

func NewNameService(boyNames, girlNames []string, random RandomFunc) *NameService {
	return &NameService{
		boyNames:  boyNames,
		girlNames: girlNames,
		random:    random,
	}
}

func (s *NameService) Generate(requestedGender string) (NameResponse, error) {
	switch requestedGender {
	case genderBoy:
		return s.generateFromList(requestedGender, genderBoy, s.boyNames)
	case genderGirl:
		return s.generateFromList(requestedGender, genderGirl, s.girlNames)
	default:
		return s.generateRandomGender(requestedGender)
	}
}

func (s *NameService) generateRandomGender(requestedGender string) (NameResponse, error) {
	totalNames := len(s.boyNames) + len(s.girlNames)
	if totalNames == 0 {
		return NameResponse{}, errNoNames
	}

	nameIndex := s.random(totalNames)
	if nameIndex < len(s.boyNames) {
		return s.response(requestedGender, genderBoy, s.boyNames[nameIndex], len(s.boyNames)), nil
	}

	girlIndex := nameIndex - len(s.boyNames)
	return s.response(requestedGender, genderGirl, s.girlNames[girlIndex], len(s.girlNames)), nil
}

func (s *NameService) generateFromList(requestedGender, selectedGender string, names []string) (NameResponse, error) {
	if len(names) == 0 {
		return NameResponse{}, errNoNames
	}

	name := names[s.random(len(names))]
	return s.response(requestedGender, selectedGender, name, len(names)), nil
}

func (s *NameService) response(requestedGender, selectedGender, name string, total int) NameResponse {
	return NameResponse{
		Name:    name,
		Gender:  selectedGender,
		Total:   total,
		Message: responseMessage(requestedGender, selectedGender),
	}
}

func responseMessage(requestedGender, actualGender string) string {
	if requestedGender == "" {
		return fmt.Sprintf("Random %s name", actualGender)
	}

	if requestedGender == actualGender {
		return fmt.Sprintf("Requested %s and received %s", requestedGender, requestedGender)
	}

	return fmt.Sprintf("Random %s name (requested: %s)", actualGender, requestedGender)
}
