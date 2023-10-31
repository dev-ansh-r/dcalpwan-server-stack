package commands

import (
	"strconv"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func paginationFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.Uint32("limit", 50, "maximum number of results to get")
	flagSet.Uint32("page", 1, "results page number")
	return flagSet
}

func withPagination(flagSet *pflag.FlagSet) (limit, page uint32, opt grpc.CallOption, getTotal func() uint64) {
	limit, _ = flagSet.GetUint32("limit")
	page, _ = flagSet.GetUint32("page")
	responseHeaders := metadata.MD{}
	opt = grpc.Header(&responseHeaders)
	getTotal = func() uint64 {
		totalHeader := responseHeaders.Get("x-total-count")
		if len(totalHeader) > 0 {
			total, _ := strconv.ParseUint(totalHeader[len(totalHeader)-1], 10, 64)
			if total != 0 && total > uint64(limit)*uint64(page) {
				logger.WithField("total", total).Infof("Use the flags \"--limit=%d --page=%d\" to get the next page of results", limit, page+1)
			} else {
				logger.Debugf("Total results: %d", total)
			}
			return total
		}
		return 0
	}
	return
}

func orderFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("order", "", "order by this field")
	return flagSet
}

func getOrder(flagSet *pflag.FlagSet) string {
	order, _ := flagSet.GetString("order")
	return order
}
