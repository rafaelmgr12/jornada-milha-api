// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import ()

type Destination struct {
	ID    string
	Name  string
	Price float64
	Photo string
}

type Testimonial struct {
	ID          string
	Name        string
	Testimonial string
}
