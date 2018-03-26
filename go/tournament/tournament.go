package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Tally writes the results of a tournament into a given buffer.
func Tally(reader io.Reader, writer io.Writer) (err error) {
	table := NewTable()

	if err = table.Populate(reader); err != nil {
		return err
	}

	if err = table.Display(writer); err != nil {
		return err
	}

	return
}

// Table represents the current standing of the different teams in tournament.
type Table struct {
	results map[string]*result
	ranking []string
}
type result struct {
	MP, W, D, L, P int
}

// NewTable initialises and returns a pointer to a Table struct.
func NewTable() (table *Table) {
	return &Table{
		results: make(map[string]*result),
		ranking: []string{},
	}
}

// Rank sorts the teams by the number of points they have.
func (table *Table) Rank() {
	sort.Strings(table.ranking)
	sort.Sort(table)
}
func (table *Table) Len() int {
	return len(table.ranking)
}
func (table *Table) Swap(i, j int) {
	table.ranking[i], table.ranking[j] = table.ranking[j], table.ranking[i]
}
func (table *Table) Less(i, j int) bool {
	return table.results[table.ranking[i]].P > table.results[table.ranking[j]].P
}

// Populate builds the internal table state from a give io.Reader.
func (table *Table) Populate(reader io.Reader) (err error) {
	scanner := bufio.NewScanner(reader)

	skip := func() bool {
		isNewLine := strings.TrimSpace(scanner.Text()) == ""
		isComment := strings.IndexRune(scanner.Text(), '#') == 0
		return isComment || isNewLine
	}

	validate := func(input []string) (err error) {
		if len(input) < 3 {
			return fmt.Errorf("invalid input")
		}
		switch input[2] {
		case "win", "loss", "draw":
		default:
			return fmt.Errorf("invalid outcome")
		}
		return
	}

	var input []string
	for scanner.Scan() {
		if skip() {
			continue
		}

		input = strings.Split(scanner.Text(), ";")
		if err = validate(input); err != nil {
			return err
		}

		table.Update(input[0], input[1], input[2])
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	return
}

// Update records the result of a match in the table.
func (table *Table) Update(home, away, outcome string) {
	if table.results[home] == nil {
		table.results[home] = &result{}
		table.ranking = append(table.ranking, home)
	}
	if table.results[away] == nil {
		table.results[away] = &result{}
		table.ranking = append(table.ranking, away)
	}

	table.results[home].MP++
	table.results[away].MP++

	switch outcome {
	case "win":
		table.results[home].W++
		table.results[away].L++

		table.results[home].P += 3
	case "loss":
		table.results[home].L++
		table.results[away].W++

		table.results[away].P += 3
	case "draw":
		table.results[home].D++
		table.results[away].D++

		table.results[home].P++
		table.results[away].P++
	}
}

// Display outputs the table to a given io.Writer.
func (table *Table) Display(writer io.Writer) (err error) {
	printer := bufio.NewWriter(writer)

	_, err = printer.WriteString(table.String())
	if err != nil {
		return err
	}

	err = printer.Flush()
	if err != nil {
		return err
	}

	return
}

// String provides Table with the Stringer interface.
func (table *Table) String() (output string) {
	row := func(name string, mp, w, d, l, p interface{}) string {
		const template = "%-30s | %2v | %2v | %2v | %2v | %2v\n"
		return fmt.Sprintf(template, name, mp, w, d, l, p)
	}

	table.Rank()

	output += row("Team", "MP", "W", "D", "L", "P")
	for _, team := range table.ranking {
		output += row(
			team,
			table.results[team].MP,
			table.results[team].W,
			table.results[team].D,
			table.results[team].L,
			table.results[team].P,
		)
	}

	return
}
