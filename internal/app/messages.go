package app

type (
	loadIdleMsg struct {
	}
	loadSuccessedMsg struct{
				notifications []Model

	}
	loadFailed       struct {
		Err error
	}
)
