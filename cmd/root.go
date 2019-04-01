package cmd

import (
	"fmt"
	"github.com/GaruGaru/flaw/flaws"
	"github.com/GaruGaru/flaw/proxy"
	"github.com/spf13/cobra"
	"net/url"
	"os"
)

const (
	Name             = "flaw"
	DescriptionShort = `Flaw proxy and injects failures on api calls for local chaos engineering`
	DescriptionLong  = `Flaw proxy and injects failures on api calls for local chaos engineering`
)

var (
	Config proxy.Config
	Host   string
	Scheme string
)

var (
	RunType            string
	RunPercentageValue int
)

var (
	HttpStatusFlawEnabled    bool
	HttpStatusFlawStatusCode int

	LatencyFlawEnabled  bool
	LatencyFlawValue    int
	LatencyFlawMaxValue int
	LatencyFlawMinValue int
)

var rootCmd = &cobra.Command{
	Use:   Name,
	Short: DescriptionShort,
	Long:  DescriptionLong,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 && Host == "" {
			Host = args[0]
		}

		hostUrl, err := url.Parse(Host)

		if err != nil {
			panic(err)
		}

		if hostUrl.Host == "" {
			panic("Host not set")
		}

		if hostUrl.Scheme != "" {
			Config.Scheme = hostUrl.Scheme
		} else {
			Config.Scheme = Scheme
		}

		if hostUrl.Port() == "" {
			Config.Host = hostUrl.Host
		} else {
			Config.Host = hostUrl.Host + ":" + hostUrl.Port()
		}

		Config.EnabledFlaws = createFlawsMiddlewares()

		proxy := proxy.New(Config)

		err = proxy.Run()

		if err != nil {
			panic(err)
		}

	},
}

func createFlawsMiddlewares() []flaws.FlawMiddleware {
	var runOption flaws.RunOption

	if RunType == "always" {
		runOption = flaws.RunAlways{}
	} else if RunType == "perc" {
		runOption = flaws.RunPerc{Percentage: RunPercentageValue}
	} else {
		panic("unsupported run type " + RunType)
	}

	middlewares := make([]flaws.FlawMiddleware, 0)

	if HttpStatusFlawEnabled {
		middlewares = append(middlewares, flaws.Of(flaws.NewHttpStatusCode(HttpStatusFlawStatusCode), runOption))
	}

	if LatencyFlawEnabled || LatencyFlawValue > 0 {

		if LatencyFlawMinValue != LatencyFlawMaxValue {
			middlewares = append(middlewares, flaws.Of(flaws.NewRandomLatency(LatencyFlawMinValue, LatencyFlawMaxValue), runOption))
		} else {
			middlewares = append(middlewares, flaws.Of(flaws.NewFixedLatency(LatencyFlawValue), runOption))
		}

	}

	return middlewares
}

func init() {

	rootCmd.PersistentFlags().StringVar(&Config.Bind, "bind", "0.0.0.0:9999", "interface and port binding for the web server")
	rootCmd.PersistentFlags().StringVar(&Host, "host", "", "real web server Host:port")
	rootCmd.PersistentFlags().StringVar(&Scheme, "schema", "http", "real web server url schema")

	rootCmd.PersistentFlags().StringVar(&RunType, "run", "perc", "run configuration [always, perc]")
	rootCmd.PersistentFlags().IntVar(&RunPercentageValue, "percentage", 10, "percentage of the flawed requests")

	rootCmd.PersistentFlags().IntVar(&HttpStatusFlawStatusCode, "status-code", 500, "Status code used for http flaws")
	rootCmd.PersistentFlags().BoolVar(&HttpStatusFlawEnabled, "http-status-enabled", true, "enable or disable http status code failure injection")

	rootCmd.PersistentFlags().BoolVar(&LatencyFlawEnabled, "latency-enabled", false, "enable or disable latency injection")
	rootCmd.PersistentFlags().IntVar(&LatencyFlawMinValue, "latency-min", 0, "set minimum latency flaw value (ms)")
	rootCmd.PersistentFlags().IntVar(&LatencyFlawMaxValue, "latency-max", 0, "set maximum latency flaw value (ms)")
	rootCmd.PersistentFlags().IntVar(&LatencyFlawValue, "latency", 0, "set fixed latency flaw value (ms)")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
