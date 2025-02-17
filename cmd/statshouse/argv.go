// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vkcom/statshouse/internal/vkgo/build"
	"github.com/vkcom/statshouse/internal/vkgo/rpc"

	"github.com/vkcom/statshouse/internal/agent"
	"github.com/vkcom/statshouse/internal/aggregator"
	"github.com/vkcom/statshouse/internal/receiver"
)

const (
	defaultPathToPwd = `/etc/engine/pass`
	defaultUser      = `kitten`
	defaultGroup     = `kitten`
)

var (
	argv struct {
		// common
		logFile         string
		logLevel        string
		userLogin       string // логин для setuid
		userGroup       string // логин для setguid
		maxOpenFiles    uint64
		pprofListenAddr string
		pprofHTTP       bool
		aesPwdFile      string
		cacheDir        string // different default
		customHostName  string // useful for testing and in some environments

		aggAddr string // common, different meaning

		cluster string // common for agent and ingress proxy

		configAgent                  agent.Config
		maxCores                     int
		listenAddr                   string
		coresUDP                     int
		bufferSizeUDP                int
		promRemoteMod                bool
		hardwareMetricScrapeInterval time.Duration
		hardwareMetricScrapeDisable  bool

		configAggregator aggregator.ConfigAggregator

		configIngress             aggregator.ConfigIngressProxy
		ingressExtAddr            string
		ingressPwdDir             string
		maxClientsPerShardReplica int

		// for old mode
		historicStorageDir string
		diskCacheFilename  string
	}
)

func readAESPwd() string {
	var aesPwd []byte
	var err error
	if argv.aesPwdFile == "" {
		aesPwd, _ = os.ReadFile(defaultPathToPwd)
	} else {
		aesPwd, err = os.ReadFile(argv.aesPwdFile)
		if err != nil {
			log.Fatalf("Could not read AES password file %s: %s", argv.aesPwdFile, err)
		}
	}
	return string(aesPwd)
}

func argvCreateClient() (*rpc.Client, string) {
	cryptoKey := readAESPwd()
	return rpc.NewClient(
		rpc.ClientWithLogf(logErr.Printf), rpc.ClientWithCryptoKey(cryptoKey), rpc.ClientWithTrustedSubnetGroups(build.TrustedSubnetGroups())), cryptoKey
}

func argvAddDeprecatedFlags() {
	// Deprecated args still used by statshouses in prod
	var (
		sampleFactor int
		maxMemLimit  uint64
	)
	flag.IntVar(&sampleFactor, "sample-factor", 1, "Deprecated - If 2, 50% of stats will be throw away, if 10, 90% of stats will be thrown away. If <= 1, keep all stats.")
	flag.Uint64Var(&maxMemLimit, "m", 0, "Deprecated - max memory usage limit")
	flag.StringVar(&argv.historicStorageDir, "historic-storage", "", "Data that cannot be immediately sent will be stored here together with metric cache.")
	flag.StringVar(&argv.diskCacheFilename, "disk-cache-filename", "", "disk cache file name")
}

func argvAddCommonFlags() {
	// common flags
	flag.StringVar(&argv.aesPwdFile, "aes-pwd-file", "", "path to AES password file, will try to read "+defaultPathToPwd+" if not set")

	flag.StringVar(&argv.logFile, "l", "/dev/stdout", "log file")
	flag.StringVar(&argv.logLevel, "log-level", "info", "log level. can be 'info' or 'trace' for now. 'trace' will print all incoming packets")

	flag.StringVar(&argv.userLogin, "u", defaultUser, "sets user name to make setuid")
	flag.StringVar(&argv.userGroup, "g", defaultGroup, "sets user group to make setguid")

	flag.StringVar(&argv.pprofListenAddr, "pprof", "", "HTTP pprof listen address (deprecated)")
	flag.BoolVar(&argv.pprofHTTP, "pprof-http", true, "Serve Go pprof HTTP on RPC port")

	flag.StringVar(&argv.cacheDir, "cache-dir", "", "Data that cannot be immediately sent will be stored here together with metric metadata cache.")

	flag.Uint64Var(&argv.maxOpenFiles, "max-open-files", 131072, "open files limit")

	flag.StringVar(&argv.aggAddr, "agg-addr", "", "Comma-separated list of 3 aggregator addresses (shard 1 is recommended). For aggregator, listen addr.")

	flag.StringVar(&argv.cluster, "cluster", aggregator.DefaultConfigAggregator().Cluster, "clickhouse cluster name to autodetect configuration, local shard and replica")

	flag.StringVar(&argv.customHostName, "hostname", "", "override auto detected hostname")
}

func argvAddAgentFlags(legacyVerb bool) {
	argv.configAgent.Bind(flag.CommandLine, agent.DefaultConfig(), legacyVerb)
	flag.StringVar(&argv.listenAddr, "p", ":13337", "RAW UDP & RPC TCP listen address")

	flag.IntVar(&argv.coresUDP, "cores-udp", 1, "CPU cores to use for udp receiving. 0 switches UDP off")
	flag.IntVar(&argv.bufferSizeUDP, "buffer-size-udp", receiver.DefaultConnBufSize, "UDP receiving buffer size")

	flag.IntVar(&argv.maxCores, "cores", -1, "CPU cores usage limit. 0 all available, <0 use (cores-udp*3/2 + 1)")

	flag.BoolVar(&argv.promRemoteMod, "prometheus-push-remote", false, "use remote pusher for prom metrics")

	flag.DurationVar(&argv.hardwareMetricScrapeInterval, "hardware-metric-scrape-interval", time.Second, "how often hardware metrics will be scraped")
	flag.BoolVar(&argv.hardwareMetricScrapeDisable, "hardware-metric-scrape-disable", false, "disable hardware metric scraping")
}

func argvAddAggregatorFlags(legacyVerb bool) {
	flag.IntVar(&argv.configAggregator.ShortWindow, "short-window", aggregator.DefaultConfigAggregator().ShortWindow, "Short admission window. Shorter window reduces latency, but also reduces recent stats quality as more agents come too late")
	flag.IntVar(&argv.configAggregator.RecentInserters, "recent-inserters", aggregator.DefaultConfigAggregator().RecentInserters, "How many parallel inserts to make for recent data")
	flag.IntVar(&argv.configAggregator.HistoricInserters, "historic-inserters", aggregator.DefaultConfigAggregator().HistoricInserters, "How many parallel inserts to make for historic data")
	flag.IntVar(&argv.configAggregator.InsertHistoricWhen, "insert-historic-when", aggregator.DefaultConfigAggregator().InsertHistoricWhen, "Aggregator will insert historic data when # of ongoing recent data inserts is this number or less")

	flag.IntVar(&argv.configAggregator.CardinalityWindow, "cardinality-window", aggregator.DefaultConfigAggregator().CardinalityWindow, "Aggregator will use this window (seconds) to estimate cardinality")
	flag.IntVar(&argv.configAggregator.MaxCardinality, "max-cardinality", aggregator.DefaultConfigAggregator().MaxCardinality, "Aggregator will sample metrics which cardinality estimates are higher")
	argv.configAggregator.Bind(flag.CommandLine, aggregator.DefaultConfigAggregator().ConfigAggregatorRemote)

	flag.Float64Var(&argv.configAggregator.SimulateRandomErrors, "simulate-errors-random", aggregator.DefaultConfigAggregator().SimulateRandomErrors, "Probability of errors for recent buckets from 0.0 (no errors) to 1.0 (all errors)")

	if legacyVerb { // TODO - remove
		var unused string
		var unused1 uint64
		flag.StringVar(&unused, "rpc-proxy-net", "", "rpc-proxy listen network")
		flag.StringVar(&unused, "rpc-proxy-addr", "", "rpc-proxy listen address")
		flag.StringVar(&unused, "dolphin-net", "", "dolphin listen network")
		flag.StringVar(&unused, "dolphin-addr", "", "dolphin listen address")
		flag.StringVar(&unused, "dolphin-table", "", "dolphin table with meta metrics")
		flag.Uint64Var(&unused1, "pmc-mapping-actor-id", 0, "actor ID of PMC mapping cluster")
	} else {
		flag.BoolVar(&argv.configAggregator.AutoCreate, "auto-create", aggregator.DefaultConfigAggregator().AutoCreate, "Enable metric auto-create.")
		flag.BoolVar(&argv.configAggregator.DisableRemoteConfig, "disable-remote-config", aggregator.DefaultConfigAggregator().DisableRemoteConfig, "disable remote configuration")
	}

	flag.StringVar(&argv.configAggregator.ExternalPort, "agg-external-port", aggregator.DefaultConfigAggregator().ExternalPort, "external port for aggregator autoconfiguration if different from port set in agg-addr")
	flag.IntVar(&argv.configAggregator.PreviousNumShards, "previous-shards", aggregator.DefaultConfigAggregator().PreviousNumShards, "Previous number of shard*replicas in cluster. During transition, clients with previous configuration are also allowed to send data.")

	flag.Uint64Var(&argv.configAggregator.MetadataActorID, "metadata-actor-id", aggregator.DefaultConfigAggregator().MetadataActorID, "")
	flag.StringVar(&argv.configAggregator.MetadataAddr, "metadata-addr", aggregator.DefaultConfigAggregator().MetadataAddr, "")
	flag.StringVar(&argv.configAggregator.MetadataNet, "metadata-net", aggregator.DefaultConfigAggregator().MetadataNet, "")

	flag.StringVar(&argv.configAggregator.KHAddr, "kh", "127.0.0.1:13338,127.0.0.1:13339", "clickhouse HTTP address:port")
}

func argvAddIngressProxyFlags() {
	flag.StringVar(&argv.configIngress.ListenAddr, "ingress-addr", "", "Listen address of ingress proxy")
	flag.StringVar(&argv.ingressExtAddr, "ingress-external-addr", "", "Comma-separate list of 3 external addresses of ingress proxies.")
	flag.StringVar(&argv.ingressPwdDir, "ingress-pwd-dir", "", "path to AES passwords dir for clients of ingress proxy.")

	flag.IntVar(&argv.maxClientsPerShardReplica, "ingress-max-conn-per-shard-replica", 3000, "")

}

func printVerbUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "Daemons usage:\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse agent <options>             daemon receiving data from clients and sending to aggregators\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse aggregator <options>        daemon receiving data from agents and inserting into clickhouse\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse ingress_proxy <options>     proxy between agents in unprotected and aggregators in protected environment\n")
	_, _ = fmt.Fprintf(os.Stderr, "Tools usage:\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse tlclient <options>          use as TL client to send JSON metrics to another statshouse\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse test_map <options>          test key mapping pipeline\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse test_parser <options>       parse and print packets received by UDP\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse test_longpoll <options>     test longpoll journal\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse simple_fsync <options>      simple SSD benchmark\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse tlclient.api <options>      test API\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse simulator <options>         simulate 10 agents sending data\n")
	_, _ = fmt.Fprintf(os.Stderr, "statshouse benchmark <options>         some brnchmark\n")
}
