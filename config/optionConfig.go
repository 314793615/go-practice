package config

type CommitFlags struct{

}

type ReviewFlags struct{
	CreatReview bool 
	AppendReview bool
	Judges []string
	PatchPath string
	LogMessage string
	SvnPath string
	ReviewID string
}

func (r *ReviewFlags) Validate() error{
	return nil
}