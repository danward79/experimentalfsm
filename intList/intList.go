package intList

import (
	"fmt"
	"strconv"
	"strings"
)

//List list of ints
type List []int64

//String prints our list
func (l List) String() string {
	s := ""
	for _, v := range l {
		s += fmt.Sprintf("%d", v)
	}
	return s
}

//Len length of l
func (l List) Len() int {
	return len(l)
}

//Swap i & j
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//Less returns bool for i < j
func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

//New from string columns specified
func New(columnList string) (List, error) {

	var out []int64

	if columnList == "" {
		return out, nil
	}

	byComma := strings.Split(columnList, ",")

	for _, v := range byComma {

		if strings.Contains(v, "-") {

			ssv := strings.Split(v, "-")
			start, err := strconv.ParseInt(ssv[0], 10, 64)
			if err != nil {
				return nil, err
			}
			end, err := strconv.ParseInt(ssv[1], 10, 64)
			if err != nil {
				return nil, err
			}

			for i := start; i <= end; i++ {
				out = append(out, i)
			}

		} else {

			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}

			out = append(out, i)
		}
	}

	return out, nil
}
