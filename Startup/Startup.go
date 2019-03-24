package startup

import (
	"SimpleProject/Domin/data"
	"fmt"

	"github.com/kardianos/osext"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/kardianos/minwinsvc"
	"github.com/spf13/viper"
)

func init() {
	path, err := osext.ExecutableFolder()

	if err != nil {
		panic(fmt.Errorf("Fatal error ExecutableFolder: %s \n", err))

	}

	viper.SetConfigName("config")
	viper.AddConfigPath(path)  // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	data.GetDB()

}
