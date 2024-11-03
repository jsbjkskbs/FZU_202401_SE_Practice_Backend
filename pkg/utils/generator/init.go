package generator

var (
	UserIDGenerator                  *Snowflake
	VideoIDGenerator                 *Snowflake
	ActivityIDGenerator              *Snowflake
	VideoCommentIDGenerator          *Snowflake
	ActvityCommentIDGenerator        *Snowflake
	ImageIDGenerator                 *Snowflake
	VideoReportIDGenerator           *Snowflake
	ActivityReportIDGenerator        *Snowflake
	VideoCommentReportIDGenerator    *Snowflake
	ActivityCommentReportIDGenerator *Snowflake
)

func Init() {
	var err error
	list := []**Snowflake{
		&UserIDGenerator,
		&VideoIDGenerator,
		&ActivityIDGenerator,
		&VideoCommentIDGenerator,
		&ActvityCommentIDGenerator,
		&ImageIDGenerator,
		&VideoReportIDGenerator,
		&ActivityReportIDGenerator,
		&VideoCommentReportIDGenerator,
		&ActivityCommentReportIDGenerator,
	}
	for node, v := range list {
		*v, err = NewSnowflake(int64(node))
		if err != nil {
			panic(err)
		}
	}
}
