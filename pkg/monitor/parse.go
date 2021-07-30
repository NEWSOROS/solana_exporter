package monitor

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	ParsedResult struct {
		PublicKey               string  `json:"pubkey"`
		Status                  float64 `json:"status"`
		RootSlot                float64 `json:"rootSlot"`
		LastVote                float64 `json:"lastVote"`
		LeaderSlots             float64 `json:"leaderSlots"`
		SkippedSlots            float64 `json:"skippedSlots"`
		PercentSkipped          float64 `json:"pctSkipped"`
		PercentTotalSkipped     float64 `json:"pctTotSkipped"`
		PercentSkippedDelta     float64 `json:"pctSkippedDelta"`
		PercentTotalDelinquent  float64 `json:"pctTotDelinquent"`
		PercentNewerVersions    float64 `json:"pctNewerVersions"`
		PercentEpochElapsed     float64 `json:"pctEpochElapsed"`
		Commission              float64 `json:"commission"`
		ActivatedStake          float64 `json:"activatedStake"`
		Credits                 float64 `json:"credits"`
		SolanaPrice             float64 `json:"solanaPrice"`
		OpenFiles               float64 `json:"openFiles"`
		ValidatorBalance        float64 `json:"validatorBalance"`
		ValidatorVoteBalance    float64 `json:"validatorVoteBalance"`
		Nodes                   float64 `json:"nodes"`
		Epoch                   float64 `json:"epoch"`
		ValidatorCreditsCurrent float64 `json:"validatorCreditsCurrent"`
		Version                 string  `json:"version"`
	}
)

// Usage for NewParsedResult:
//	src, err := Exec("/bin/bash", "-e", "./run.sh")
//	if err != nil {
//		panic(err)
//	}
//	_ = NewParsedResult().Parse(src)
func NewParsedResult() *ParsedResult {
	return &ParsedResult{}
}

// Parse(`nodemonitor,pubkey=383df5a46392f60c187a5aeb3ad7c0281af81c00c5ec status=0,rootSlot=111,lastVote=111,leaderSlots=111,skippedSlots=11,pctSkipped=11.11,pctTotSkipped=11.11,pctSkippedDelta=-11.11,pctTotDelinquent=11.11,version="1.7.8",pctNewerVersions=.11,commission=100,activatedStake=11.11,credits=11,solanaPrice=11.11,openFiles=11,validatorBalance=11.11,validatorVoteBalance=11.11,nodes=11,epoch=11,pctEpochElapsed=11.11,validatorCreditsCurrent=11 1627641121663342912`)
func (result *ParsedResult) Parse(src string) *ParsedResult {
	spaces := strings.Split(src, " ")
	for _, line := range spaces {
		elements := strings.Split(line, ",")
		for _, el := range elements {
			s := strings.Split(el, "=")
			if len(s) != 2 {
				continue
			}
			switch s[0] {
			case "pubkey":
				result.PublicKey = s[1]
			case "status":
				i, err := strconv.Atoi(s[1])
				if err == nil {
					result.Status = float64(i)
				}
			case "version":
				result.Version = strings.Replace(s[1], "\"", "", -1)
			case "openFiles":
				i, err := strconv.Atoi(s[1])
				if err == nil {
					result.OpenFiles = float64(i)
				}
			case "epoch":
				i, err := strconv.Atoi(s[1])
				if err == nil {
					result.Epoch = float64(i)
				}
			case "nodes":
				i, err := strconv.Atoi(s[1])
				if err == nil {
					result.Nodes = float64(i)
				}
			case "credits":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.Credits = i
				}
			case "validatorCreditsCurrent":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.ValidatorCreditsCurrent = i
				}
			case "validatorBalance":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.ValidatorBalance = i
				}
			case "validatorVoteBalance":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.ValidatorVoteBalance = i
				}
			case "activatedStake":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.ActivatedStake = i
				}
			case "solanaPrice":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.SolanaPrice = i
				}
			case "commission":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.Commission = i
				}
			case "rootSlot":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.RootSlot = i
				}
			case "lastVote":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.LastVote = i
				}
			case "leaderSlots":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.LeaderSlots = i
				}
			case "skippedSlots":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.SkippedSlots = i
				}
			case "pctSkipped":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.PercentSkipped = i
				}
			case "pctEpochElapsed":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.PercentEpochElapsed = i
				}
			case "pctTotSkipped":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.PercentTotalSkipped = i
				}
			case "pctNewerVersions":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.PercentNewerVersions = i
				}
			case "pctSkippedDelta":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.PercentSkippedDelta = i
				}
			case "pctTotDelinquent":
				i, err := strconv.ParseFloat(s[1], 64)
				if err == nil {
					result.PercentTotalDelinquent = i
				}
			default:
				fmt.Println(s)
			}
		}
	}
	return result
}
