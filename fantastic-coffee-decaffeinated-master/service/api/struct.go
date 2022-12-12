package api

type Photo struct {
	photoID   uint64
	profileId uint64
	image     []string
	time      string
	likes     []uint64
}

type User struct {
	ID        uint64
	username  string
	follows   []uint64
	followers []uint64
	banned    []uint64
	photos    []Photo
}
