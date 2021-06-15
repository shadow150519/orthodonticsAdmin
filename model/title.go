package model

type Title int

const (
	PHYSICIAN Title = iota + 1
	ASSISTANT_DIRECTOR_PHYSICIAN
	DIRECTOR_PHYSICIAN
)
