package snowflake

import "sfw/pkg/utils/generator"

var (
	UserIDGenerator  *generator.Snowflake
	VideoIDGenerator *generator.Snowflake
)

func Init() {
	var err error
	UserIDGenerator, err = generator.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	VideoIDGenerator, err = generator.NewSnowflake(2)
	if err != nil {
		panic(err)
	}
}
