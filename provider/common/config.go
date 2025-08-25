package common

type Config struct {
	OtoroshiAdminUrl          string `pulumi:"otoroshiAdminUrl"`
	OtoroshiAdminClientID     string `pulumi:"otoroshiAdminClientId" provider:"secret"`
	OtoroshiAdminClientSecret string `pulumi:"otoroshiAdminClientSecret" provider:"secret"`
}
