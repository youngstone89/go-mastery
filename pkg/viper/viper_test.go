package viper

import (
	"flag"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func TestBindPFlag(t *testing.T) {

	// using standard library "flag" package
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // retrieve value from viper
	t.Log(i)

}
