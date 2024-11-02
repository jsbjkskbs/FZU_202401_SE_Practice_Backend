package generator

var (
	UserIDGenerator     *Snowflake
	VideoIDGenerator    *Snowflake
	ActivityIDGenerator *Snowflake
)

func Init() {
	var err error
	UserIDGenerator, err = NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	VideoIDGenerator, err = NewSnowflake(2)
	if err != nil {
		panic(err)
	}
	ActivityIDGenerator, err = NewSnowflake(3)
	if err != nil {
		panic(err)
	}
}
