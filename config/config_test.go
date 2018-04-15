package config

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func helperLoadConfig(contents []byte) (*Config, error) {
	cfg := &Config{}
	err := yaml.Unmarshal(contents, cfg)
	return cfg, err
}

func TestLoadFile(t *testing.T) {
	testCases := []struct {
		file string
		cfg  Config
		ok   bool
	}{
		{
			file: "../config.example.yaml",
			cfg: Config{
				Notifier: Notifier{
					Slack: SlackNotifier{
						Token:    "xoxb-XXXX",
						Channel:  "test-channel",
						Template: "test format\n",
					},
				},
				Feed: []Feed{
					Feed{
						URL: "https://example.com/feed1.xml",
						Auth: BasicAuth{
							User: "hogefuga",
							Pass: "password",
						},
					},
					Feed{
						URL:  "https://example.com/feed2.xml",
						Auth: BasicAuth{},
					},
				},
				Interval: 5,
				path:     "../config.example.yaml",
			},
			ok: true,
		},
	}

	var cfg Config
	for _, testCase := range testCases {
		err := cfg.LoadFile(testCase.file)
		if !reflect.DeepEqual(cfg, testCase.cfg) {
			t.Errorf("got \n%q\n but want \n%q", cfg, testCase.cfg)
		}
		if (err == nil) != testCase.ok {
			t.Errorf("got error %q", err)
		}
	}
}

func TestValidation(t *testing.T) {
	testCases := []struct {
		contents []byte
		expected string
	}{
		{
			contents: []byte(""),
			expected: "notifier is missing",
		},
		{
			contents: []byte(`
notifier:
  slack:
    token: "xoxb-XXXX"
`),
			expected: "feed is missing",
		},
	}

	for _, testCase := range testCases {
		cfg, err := helperLoadConfig(testCase.contents)
		if err != nil {
			t.Fatal(err)
		}
		err = cfg.Validation()
		if err == nil {
			if testCase.expected != "" {
				t.Errorf("got no error but want %q", testCase.expected)
			}
		} else {
			if err.Error() != testCase.expected {
				t.Errorf("got %q but want %q", err.Error(), testCase.expected)
			}
		}
	}
}
