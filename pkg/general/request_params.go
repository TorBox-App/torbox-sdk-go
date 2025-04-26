package general

type GetSpeedtestFilesRequestParams struct {
	TestLength *string `explode:"true" serializationStyle:"form" queryParam:"test_length"`
	Region     *string `explode:"true" serializationStyle:"form" queryParam:"region"`
}

func (params *GetSpeedtestFilesRequestParams) SetTestLength(testLength string) {
	params.TestLength = &testLength
}
func (params *GetSpeedtestFilesRequestParams) SetRegion(region string) {
	params.Region = &region
}
