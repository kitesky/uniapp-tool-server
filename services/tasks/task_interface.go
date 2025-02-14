package tasks

type TaskHandler interface {
	TaskCheck(userID int64, code string) (ok bool, err error)
	TaskReward(userID int64, code string) (ok bool, err error)
}
