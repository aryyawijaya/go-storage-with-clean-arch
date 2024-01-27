package entity

type Access string

const (
	AccessPUBLIC  Access = "PUBLIC"
	AccessPRIVATE Access = "PRIVATE"
)

func (e Access) Valid() bool {
	switch e {
	case AccessPUBLIC,
		AccessPRIVATE:
		return true
	}
	return false
}

func AllAccessValues() []Access {
	return []Access{
		AccessPUBLIC,
		AccessPRIVATE,
	}
}
