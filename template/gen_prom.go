// username: vonhng
// create_time: 2019/12/15 - 22:29
// mail: qianyong.feng@woqutech.com
package template

type PromYaml struct {
	Global        *BaseConfig `json:"global, omitempty" yaml:"global, omitempty"`
	ScrapeConfigs []*Job      `json:"scrape_configs, omitempty" yaml:"scrape_configs, omitempty"`
}

type BaseConfig struct {
	ScrapeInterval string            `json:"scrape_interval, omitempty" yaml:"scrape_interval, omitempty"`
	ExternalLabels map[string]string `json:"external_labels, omitempty" yaml:"external_labels, omitempty"`
}

type Job struct {
	JobName        string    `json:"job_name, omitempty" yaml:"job_name, omitempty"`
	ScrapeInterval string    `json:"scrape_interval, omitempty" yaml:"scrape_interval, omitempty"`
	ScrapeConfigs  []*Target `json:"scrape_configs, omitempty" yaml:"scrape_configs, omitempty"`
}

type Target struct {
	Targets []string `json:"targets, omitempty" yaml:"targets, omitempty"`
}
