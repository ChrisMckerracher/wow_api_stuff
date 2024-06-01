package model

type Region string

const (
	US Region = "us"
	EU Region = "eu"
	TW Region = "tw"
)

type Namespace string

const (
	DYNAMIC_US Namespace = "dynamic-us"
)

type Local string

const (
	EN_US Local = "en_us"
)

type ConnectedRealm int32

const (
	CR_ONE ConnectedRealm = iota
	CR_TWO
	CR_THREE
	CR_FOUR
	CR_FIVE
	_
	_
	_
	_
	_
	_
	CR_ELEVEN
)
