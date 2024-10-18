package cmd

import (
	"errors"
	"fmt"
	"go-practice/config"
	"go-practice/option"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const(
	createReviewLabel = "--create-review"
	appendReviewLabel = "--append-review"
	reviewLog = "--review-log"
	judges = "--judges"
	patchPath = "--patch-path"
	svnPath = "--svn-path"
	reviewId = "--review-id"
)

var reviewCmd = &cobra.Command{
	Use: "rv",
	Aliases: []string{"review"},
	Short: "create review or append review",
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Print("reviewCmd starts!")
		conf, err := getFlags(cmd.Flags())
		if err != nil{
			fmt.Printf("happend a error in parse review flags: %s", err.Error())
			os.Exit(1)
		}
		err = conf.Validate()
		if err != nil{
			fmt.Printf("happend a err in validate the review flags: %s", err.Error())
			os.Exit(1)
		}
		err = option.HanleReviewOpt(conf)
		if err != nil{
			fmt.Printf("handleReviewOpt failed: %s", err.Error())
			os.Exit(1)
		}
	},
}

func init(){
	rootCmd.AddCommand(reviewCmd)
	reviewCmd.Flags().BoolP(createReviewLabel, "c", false, "create review label")
	reviewCmd.Flags().BoolP(appendReviewLabel, "u", false, "append review label")
	reviewCmd.Flags().StringP(reviewLog, "m", "", "review log for create review")
	reviewCmd.Flags().StringP(judges, "j", "", "judges for review")
	reviewCmd.Flags().StringP(patchPath, "f", "", "patch for review")
	reviewCmd.Flags().StringP(svnPath, "a", "", "svn path for create review")
	reviewCmd.Flags().StringP(reviewId, "i", "", "review id for append review")
}


func getFlags(flags *pflag.FlagSet) (*config.ReviewFlags, error){
	conf := &config.ReviewFlags{}
	var nErr error
	var err error
	flags.Visit(func(f *pflag.Flag) {
		switch f.Name{
		case createReviewLabel: conf.CreatReview, nErr = strconv.ParseBool(f.Value.String())
		case appendReviewLabel: conf.AppendReview, nErr = strconv.ParseBool(f.Value.String())
		case reviewLog: conf.LogMessage = f.Value.String()
		case judges: conf.Judges = parseJudges(f.Value.String())
		case patchPath: conf.PatchPath = f.Value.String()
		case svnPath: conf.PatchPath = f.Value.String()
		case reviewId: conf.ReviewID = f.Value.String()
		}
		err = errors.Join(err, nErr)
	})
	fmt.Printf("review flags parsed finished")
	return conf, err
}

func parseJudges(judgeString string) []string{
	judgeString = strings.TrimSpace(judgeString)
	judgeList := strings.Split(judgeString, ",")
	judges := make([]string, 0)
	for _, j := range judgeList{
		judges = append(judges, strings.TrimSpace(j))
	}
	return judges
}

