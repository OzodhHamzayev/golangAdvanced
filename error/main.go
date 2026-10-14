package main


//! errors.New -> unwrapped error
//! error.Is -> wrapped errorni ochib tekshiradi
//! error.As -> aynan shu xato turidanmisan yani, 
	// validationError structga xato yozildi va errors.As bilan tekshiriladi kelgan xato validationErrordanmi
//!! fmt.Errorf(%w, err) -> wrapped error

import (
	"errors"
	"fmt"
	"strings"
)


func loadFile(fileName string) (string, error) {
	if !strings.HasSuffix(fileName, ".text") {
		return fileName, errors.New("bu kutilgan file emas iltimos tekshiring")
	}
	return fileName, nil
}
//!----------------------------------------------------------------------------------

//! errors.IS
var ErrNotFound = errors.New("topilmadi")

func findUser (id int) error {
	if id != 1 {
		return fmt.Errorf("findUSer(%d): %w", id, ErrNotFound)
	}
	return nil
}

//!-----------------------------------------------------------------------------------

//! errrors.As

type ValidationError struct {
	Field string
	Msg string
}

func (v *ValidationError) Error() string {
	return v.Field + ": " + v.Msg
}

func validate(age int) error {
	if age < 0 {
		return fmt.Errorf("validate: %w", &ValidationError{Field: "age", Msg: "manfiy bolishi mumkin emas"})
	}
	return nil
}

func main () {
	resp, err := loadFile("mod.ext")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	err = findUser(5)
	fmt.Println(err == ErrNotFound)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("you are not a manager, try again")
	}


	var ve *ValidationError
	err = validate(-5)
	if errors.As(err, &ve) {
		fmt.Println("Maydon:", ve.Field)
		fmt.Println("Xabar:", ve.Msg)
	}
}