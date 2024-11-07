package common

import "time"

const (
	ActionTypeOff = 0
	ActionTypeOn  = 1

	CommentTypeVideo    = "video"
	CommentTypeActivity = "activity"

	VideoStatusSubmit = "submit"
	VideoStatusReview = "review"
	VideoStatusPassed = "passed"
	VideoStatusLocked = "locked"

	ReportLimit = 3

	ReportResolved   = "resolved"
	ReportUnresolved = "unresolved"
	ReportRejected   = "rejected"

	VideoVisitInterval = 20 * time.Minute

	SyncInterval = 3 * time.Minute

	GorseFeedbackRead    = "read"
	GorseFeedbackDislike = "dislike"
	GorseFeedbackLike    = "like"
	GorseFeedbackStar    = "star"
	GorseFeedbackVisit   = "visit"
)
