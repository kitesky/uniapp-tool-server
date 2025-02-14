package cmd

import (
	"fmt"

	"app-api/boot"

	"github.com/spf13/cobra"
)

// gentool -dsn 'root:Wang86089793@tcp(rm-7xvh2t49ms2b9s3fljo.mysql.rds.aliyuncs.com:3306)/idcd_com?charset=utf8mb4&parseTime=True&loc=Local' -outPath './models' -modelPkgName 'models' -fieldWithTypeTag -onlyModel -tables 'users'
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate model",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := boot.Config.Database
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User,     // user
			cfg.Password, // password
			cfg.Host,     // host
			cfg.Port,     // port
			cfg.Database, // DB name
		)
		tables := "users"
		line := fmt.Sprintf("gentool -dsn '%s' -outPath './models' -modelPkgName 'models' -fieldWithTypeTag   -onlyModel -tables '%s'", dns, tables)
		fmt.Println(line)
		// command := exec.Command(line)
		// var out bytes.Buffer
		// command.Stdout = &out
		// err := command.Run()
		// if err != nil {
		// 	fmt.Println("命令执行出错:", err)
		// 	return
		// }
		// fmt.Println("命令输出:", out.String())
	},
}
