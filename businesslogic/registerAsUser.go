package businesslogic

import (
	"Ecommerce/dataaccess"
	"Ecommerce/utilities"
	"regexp"
)

func RegistrAsUser(email string, password string) (errors []DomainError) {

	errors = ValidateEmail(email)
	if len(errors) != 0 {
		return errors

	}

	errors = ValidatePassword(password)
	if len(errors) != 0 {
		return errors
	}

	found, err := dataaccess.CheckIfUserExsistByEmail(email)
	if !found {
		return []DomainError{
			NewInvalidFormatDomainError("This email has been found in the registry, try to sign in or sign up with another email "),
		}
	}
	if err != nil {
		return []DomainError{
			NewInvalidFormatDomainError("Internal error occured"),
		}
	}

	//hash password function call
	hashedPassword, err := utilities.HashPassword(password)
	if err != nil {
		//error hnadling
		return []DomainError{
			NewInvalidFormatDomainError("Internal error occured"),
		}
	}

	err = dataaccess.CreateUser(email, hashedPassword)
	if err != nil {
		//error handling
	}

	return nil
}

func ValidateEmail(email string) (errors []DomainError) {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	matchEmail := emailRegex.MatchString(email)
	if !matchEmail {
		errors = append(errors, NewInvalidFormatDomainError("Email must contain character followed by @ then domain name and . followed by type"))
	}
	return errors
}

func ValidatePassword(password string) (errors []DomainError) {
	// passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})`)

	containsSmallCharacters, _ := regexp.MatchString(`[a-z]`, password)
	if !containsSmallCharacters {
		errors = append(errors, NewInvalidFormatDomainError("Password must contain at least 1 small character"))
	}

	containsCapitalCharacters, _ := regexp.MatchString(`[A-Z]`, password)
	if !containsCapitalCharacters {
		errors = append(errors, NewInvalidFormatDomainError("Password must contain at least 1 capital character"))
	}

	containsNumbers, _ := regexp.MatchString(`[0-9]`, password)
	if !containsNumbers {
		errors = append(errors, NewInvalidFormatDomainError("Password must contain at least 1 number"))
	}

	containsSpecialCharacters, _ := regexp.MatchString(`[!@#\$%\^&\*]`, password)
	if !containsSpecialCharacters {
		errors = append(errors, NewInvalidFormatDomainError("Password must contain at least 1 special character (any of !@#$%^&*)"))
	}

	return errors
}
