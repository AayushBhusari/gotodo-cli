package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

const (
	ColorDefault = "\x1b[39m"
	ColorRed = "\x1b[39m"
	ColorGreen = "\x1b[32m"
	ColorBlue = "\x1b[94m"
)

func red(s string) string{
	return fmt.Sprintf("%s%s%s", ColorRed,s,ColorDefault)
}
func green(s string) string{
	return fmt.Sprintf("%s%s%s", ColorGreen,s,ColorDefault)
}

func blue(s string) string{
	return fmt.Sprintf("%s%s%s", ColorBlue,s,ColorDefault)
}

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

// Add a new task
func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

// Mark a task as completed
func (t *Todos) Completed(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil
}

// Delete a task by index
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

// Load todos from a file
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

// Store todos to a file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// List all todos in a table
func (t *Todos) List() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Done{
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822Z)},
			{Text: item.CompletedAt.Format(time.RFC822Z)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func(t *Todos) CountPending() int {
	total:= 0
	for _, item := range *t{
		if !item.Done {
			total++
		}
	}
	return total
}

// Reset the Todos slice to an empty slice
func (t *Todos) Clear() {
	*t = Todos{}
}