package id




//使用string 类型的id变为identifier type设计模式，为了保证aid不是因为都是普遍类型而导致传错
//将每种ID设计成固有的type  string 类型，这样可以起到保护作用

type AccountID string

func (a AccountID)String()  string {
	return  string(a)
}



type TripID string

func (t TripID)String() string  {
	return  string(t)
}


type CarID string

func (t CarID)String() string  {
	return  string(t)
}


type IdentityID string

func (t IdentityID)String() string  {
	return  string(t)
}