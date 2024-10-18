package option

import (
	"fmt"
	"go-practice/config"
)

func HanleReviewOpt(f *config.ReviewFlags) error {
	fmt.Printf("flags: %#v", *f)
	return nil
}