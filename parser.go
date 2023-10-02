package cachegrind

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/k1LoW/octocov/report"
)

var numRe = regexp.MustCompile(`^\d`)

func Parse(in io.Reader) (*report.CustomMetricSet, error) {
	scanner := bufio.NewScanner(in)
	cset := &report.CustomMetricSet{}
	var (
		events []string
		values []string
	)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "fn=") || strings.HasPrefix(line, "fl="):
			continue
		case numRe.MatchString(line):
			continue
		case strings.HasPrefix(line, "desc:"):
			kv := strings.Split(strings.TrimPrefix(line, "desc:"), ":")
			cset.Metadata = append(cset.Metadata, &report.MetadataKV{
				Key:   strings.TrimSpace(kv[0]),
				Value: strings.TrimSpace(kv[1]),
			})
		case strings.HasPrefix(line, "cmd:"):
			kv := strings.Split(line, ":")
			cset.Metadata = append(cset.Metadata, &report.MetadataKV{
				Key:   strings.TrimSpace(kv[0]),
				Value: strings.TrimSpace(kv[1]),
			})
			cset.Key = fmt.Sprintf("Cachegrind (%s)", strings.TrimSpace(kv[1]))
		case strings.HasPrefix(line, "events:"):
			events = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "events:")), " ")
		case strings.HasPrefix(line, "summary:"):
			values = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "summary:")), " ")
		}
	}
	for i, e := range events {
		v, err := strconv.Atoi(values[i])
		if err != nil {
			return nil, err
		}
		cset.Metrics = append(cset.Metrics, &report.CustomMetric{
			Key:   e,
			Value: float64(v),
		})
	}
	return cset, nil
}
