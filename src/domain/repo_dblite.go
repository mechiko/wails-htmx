package domain

import "firstwails/repo/lite/liteboil"

// рефактор
type DbLite interface {
	Version() (string, error)
	// InsertCisRequest(in []string) (err error)
	InsertCisRequest(in []Cises) (err error)
	CisRequestAll() (out liteboil.CisRequestSlice, err error)
	CisRequestDeleteAll() (err error)
	InsertCisRequestPost(in []CisesPost) (err error)
}
