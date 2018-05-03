package checker

type Checker interface {
	Check(repo string,ignore ...string) (string,error)
}
