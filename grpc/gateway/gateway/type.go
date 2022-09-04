package main



type rig struct {
	Man   bool `json:"man"`
	Woman bool `json:"woman"`
	Nbr   string `json:"nbr"`
}

type rigs struct {
	Man   bool `json:"man"`
	Woman bool `json:"woman"`
	Nbr   int `json:"nbr"`
}
