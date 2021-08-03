package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
)

type Collector struct {
	pathToScript string

	descStatus                  *prometheus.Desc
	descOpenFiles               *prometheus.Desc
	descEpoch                   *prometheus.Desc
	descNodes                   *prometheus.Desc
	descCredits                 *prometheus.Desc
	descValidatorBalance        *prometheus.Desc
	descValidatorVoteBalance    *prometheus.Desc
	descValidatorCreditsCurrent *prometheus.Desc
	descActivatedStake          *prometheus.Desc
	descSolanaPrice             *prometheus.Desc
	descCommission              *prometheus.Desc
	descRootSlot                *prometheus.Desc
	descLastVote                *prometheus.Desc
	descLeaderSlots             *prometheus.Desc
	descSkippedSlots            *prometheus.Desc
	descPercentSkipped          *prometheus.Desc
	descPercentTotalSkipped     *prometheus.Desc
	descPercentSkippedDelta     *prometheus.Desc
	descPercentTotalDelinquent  *prometheus.Desc
	descPercentNewerVersions    *prometheus.Desc
	descPercentEpochElapsed     *prometheus.Desc
}

func (self *Collector) Describe(descs chan<- *prometheus.Desc) {
	// klog.Info("Collector.Describe")
	descs <- self.descStatus
	descs <- self.descOpenFiles
	descs <- self.descEpoch
	descs <- self.descNodes
	descs <- self.descCredits
	descs <- self.descValidatorBalance
	descs <- self.descValidatorVoteBalance
	descs <- self.descValidatorCreditsCurrent
	descs <- self.descActivatedStake
	descs <- self.descSolanaPrice
	descs <- self.descCommission
	descs <- self.descRootSlot
	descs <- self.descLastVote
	descs <- self.descLeaderSlots
	descs <- self.descSkippedSlots
	descs <- self.descPercentSkipped
	descs <- self.descPercentTotalSkipped
	descs <- self.descPercentSkippedDelta
	descs <- self.descPercentTotalDelinquent
	descs <- self.descPercentNewerVersions
	descs <- self.descPercentEpochElapsed
}

func NewCollector(pathToScript string) *Collector {
	// klog.Info("Collector.NewCollector")
	return &Collector{
		pathToScript: pathToScript,
		descStatus: prometheus.NewDesc(
			"solana_monitor_status",
			"status",
			[]string{
				"pubkey",
			},
			nil,
		),
		descOpenFiles: prometheus.NewDesc(
			"solana_monitor_open_files",
			"opened files",
			[]string{
				"pubkey",
			},
			nil,
		),
		descEpoch: prometheus.NewDesc(
			"solana_monitor_epoch",
			"epoch",
			[]string{
				"pubkey",
			},
			nil,
		),
		descNodes: prometheus.NewDesc(
			"solana_monitor_nodes",
			"nodes",
			[]string{
				"pubkey",
			},
			nil,
		),
		descCredits: prometheus.NewDesc(
			"solana_monitor_credits",
			"credits",
			[]string{
				"pubkey",
			},
			nil,
		),
		descValidatorBalance: prometheus.NewDesc(
			"solana_monitor_validator_balance",
			"validator balance",
			[]string{"pubkey"},
			nil,
		),
		descValidatorVoteBalance: prometheus.NewDesc(
			"solana_monitor_validator_vote_balance",
			"validator vote balance",
			[]string{"pubkey"},
			nil,
		),
		descValidatorCreditsCurrent: prometheus.NewDesc(
			"solana_monitor_validator_credits_current",
			"validator credits current",
			[]string{"pubkey"},
			nil,
		),
		descActivatedStake: prometheus.NewDesc(
			"solana_monitor_activated_stake",
			"activated stake",
			[]string{"pubkey"},
			nil,
		),
		descSolanaPrice: prometheus.NewDesc(
			"solana_monitor_solana_price",
			"solana price",
			[]string{"pubkey"},
			nil,
		),
		descCommission: prometheus.NewDesc(
			"solana_monitor_commission",
			"commission",
			[]string{"pubkey"},
			nil,
		),
		descRootSlot: prometheus.NewDesc(
			"solana_monitor_root_slot",
			"root slot",
			[]string{"pubkey"},
			nil,
		),
		descLastVote: prometheus.NewDesc(
			"solana_monitor_last_vote",
			"last vote",
			[]string{"pubkey"},
			nil,
		),
		descLeaderSlots: prometheus.NewDesc(
			"solana_monitor_leader_slots",
			"leader slots",
			[]string{"pubkey"},
			nil,
		),
		descSkippedSlots: prometheus.NewDesc(
			"solana_monitor_skipped_slots",
			"skipped slots",
			[]string{"pubkey"},
			nil,
		),
		descPercentSkipped: prometheus.NewDesc(
			"solana_monitor_percent_skipped",
			"",
			[]string{"pubkey"},
			nil,
		),
		descPercentTotalSkipped: prometheus.NewDesc(
			"solana_monitor_percent_total_skipped",
			"",
			[]string{"pubkey"},
			nil,
		),
		descPercentSkippedDelta: prometheus.NewDesc(
			"solana_monitor_percent_skipped_delta",
			"",
			[]string{"pubkey"},
			nil,
		),
		descPercentTotalDelinquent: prometheus.NewDesc(
			"solana_monitor_percent_total_delinquent",
			"",
			[]string{"pubkey"},
			nil,
		),
		descPercentNewerVersions: prometheus.NewDesc(
			"solana_monitor_percent_newer_versions",
			"",
			[]string{"pubkey"},
			nil,
		),
		descPercentEpochElapsed: prometheus.NewDesc(
			"solana_monitor_percent_epoch_elapsed",
			"",
			[]string{"pubkey"},
			nil,
		),
	}
}

func (self *Collector) Collect(ch chan<- prometheus.Metric) {
	// klog.Info("Collector.Collect")
	src, err := Exec("/bin/bash", "-e", self.pathToScript)
	result := NewParsedResult().Parse(src)
	if err != nil {
		klog.Error("/bin/bash -e " + self.pathToScript + ": ", err, " | ", src)
		ch <- prometheus.NewInvalidMetric(self.descStatus, err)
		ch <- prometheus.NewInvalidMetric(self.descOpenFiles, err)
		ch <- prometheus.NewInvalidMetric(self.descEpoch, err)
		ch <- prometheus.NewInvalidMetric(self.descNodes, err)
		ch <- prometheus.NewInvalidMetric(self.descCredits, err)
		ch <- prometheus.NewInvalidMetric(self.descValidatorBalance, err)
		ch <- prometheus.NewInvalidMetric(self.descValidatorVoteBalance, err)
		ch <- prometheus.NewInvalidMetric(self.descValidatorCreditsCurrent, err)
		ch <- prometheus.NewInvalidMetric(self.descActivatedStake, err)
		ch <- prometheus.NewInvalidMetric(self.descSolanaPrice, err)
		ch <- prometheus.NewInvalidMetric(self.descCommission, err)
		ch <- prometheus.NewInvalidMetric(self.descRootSlot, err)
		ch <- prometheus.NewInvalidMetric(self.descLastVote, err)
		ch <- prometheus.NewInvalidMetric(self.descLeaderSlots, err)
		ch <- prometheus.NewInvalidMetric(self.descSkippedSlots, err)
		ch <- prometheus.NewInvalidMetric(self.descPercentSkipped, err)
		ch <- prometheus.NewInvalidMetric(self.descPercentTotalSkipped, err)
		ch <- prometheus.NewInvalidMetric(self.descPercentSkippedDelta, err)
		ch <- prometheus.NewInvalidMetric(self.descPercentTotalDelinquent, err)
		ch <- prometheus.NewInvalidMetric(self.descPercentNewerVersions, err)
		ch <- prometheus.NewInvalidMetric(self.descPercentEpochElapsed, err)
	} else {
		self.mustEmitMetrics(
			ch,
			result,
		)
	}
}

func (self *Collector) mustEmitMetrics(ch chan<- prometheus.Metric, result *ParsedResult) {
	// klog.Info("Collector.mustEmitMetrics")
	ch <- prometheus.MustNewConstMetric(
		self.descStatus,
		prometheus.GaugeValue,
		result.Status,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descOpenFiles,
		prometheus.GaugeValue,
		result.OpenFiles,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descEpoch,
		prometheus.GaugeValue,
		result.Epoch,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descNodes,
		prometheus.GaugeValue,
		result.Nodes,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descCredits,
		prometheus.GaugeValue,
		result.Credits,
		result.PublicKey,
	)

	ch <- prometheus.MustNewConstMetric(
		self.descValidatorBalance,
		prometheus.GaugeValue,
		result.ValidatorBalance,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descValidatorVoteBalance,
		prometheus.GaugeValue,
		result.ValidatorVoteBalance,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descValidatorCreditsCurrent,
		prometheus.GaugeValue,
		result.ValidatorCreditsCurrent,
		result.PublicKey,
	)

	ch <- prometheus.MustNewConstMetric(
		self.descActivatedStake,
		prometheus.GaugeValue,
		result.ActivatedStake,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descSolanaPrice,
		prometheus.GaugeValue,
		result.SolanaPrice,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descCommission,
		prometheus.GaugeValue,
		result.Commission,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descRootSlot,
		prometheus.GaugeValue,
		result.RootSlot,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descLastVote,
		prometheus.GaugeValue,
		result.LastVote,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descLeaderSlots,
		prometheus.GaugeValue,
		result.LeaderSlots,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descSkippedSlots,
		prometheus.GaugeValue,
		result.SkippedSlots,
		result.PublicKey,
	)

	ch <- prometheus.MustNewConstMetric(
		self.descPercentSkipped,
		prometheus.GaugeValue,
		result.PercentSkipped,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descPercentTotalSkipped,
		prometheus.GaugeValue,
		result.PercentTotalSkipped,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descPercentSkippedDelta,
		prometheus.GaugeValue,
		result.PercentSkippedDelta,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descPercentTotalDelinquent,
		prometheus.GaugeValue,
		result.PercentTotalDelinquent,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descPercentNewerVersions,
		prometheus.GaugeValue,
		result.PercentNewerVersions,
		result.PublicKey,
	)
	ch <- prometheus.MustNewConstMetric(
		self.descPercentEpochElapsed,
		prometheus.GaugeValue,
		result.PercentEpochElapsed,
		result.PublicKey,
	)
}
