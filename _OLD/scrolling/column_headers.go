package scrolling

import "sort"

type ColumnDescriptor struct {
	Name       string
	Header     string
	Field      string
	Visible    bool
	Filtered   bool
	Order      int
	WidthClass string
	ValueClass string
}

type ColumnDescriptorSlice []*ColumnDescriptor

type Columns struct {
	list     ColumnDescriptorSlice
	maxOrder int
	maxWidth int
}

func NewColumns() *Columns {
	return &Columns{
		list:     make(ColumnDescriptorSlice, 0),
		maxOrder: 0,
	}
}

// добавляет колонку, если уже есть не добавляет а возвращает существующую колонку
func (c *Columns) AddColumn(colName string, header string, field string, visible bool) *ColumnDescriptor {
	if col := c.Column(colName); col != nil {
		return col
	}
	desc := &ColumnDescriptor{
		Name:       colName,
		Header:     header,
		Field:      field,
		Visible:    visible,
		Order:      c.maxOrder,
		Filtered:   false,
		WidthClass: "w-10",
		ValueClass: "p-2 text-left",
	}
	c.maxOrder += 1
	c.maxWidth = 3000
	c.list = append(c.list, desc)
	return desc
}

// возвращает массив строк из всех колонок
func (c *Columns) ColumnNameAll() []string {
	keys := make([]string, 0, len(c.list))
	// keys := maps.Keys(c)
	for _, col := range c.list {
		keys = append(keys, col.Name)
	}
	return keys
}

// возвращает массив строк из видимых колонок
func (c *Columns) ColumnNameVisible() []string {
	keys := make([]string, 0, len(c.list))
	for _, val := range c.list {
		if val.Visible {
			keys = append(keys, val.Name)
		}
	}
	return keys
}

// возвращает массив строк из видимых колонок
func (c *Columns) VisibleColumns() ColumnDescriptorSlice {
	out := make(ColumnDescriptorSlice, 0)
	for i, val := range c.list {
		if val.Visible {
			out = append(out, c.list[i])
		}
	}
	return out
}

// возвращает массив строк из видимых колонок
func (c *Columns) AllColumns() ColumnDescriptorSlice {
	return c.list
}

// сортирует список по возростанию Order
func (c *Columns) SortByOrder() {
	sort.Slice(c.list, func(i, j int) bool {
		return c.list[i].Order < c.list[j].Order
	})
}

func (c *Columns) Column(name string) *ColumnDescriptor {
	for i, col := range c.list {
		if name == col.Name {
			return c.list[i]
		}
	}
	return nil
}

func (c *Columns) ColumnMap(name string) map[string]*ColumnDescriptor {
	out := make(map[string]*ColumnDescriptor)
	for i, col := range c.list {
		out[col.Name] = c.list[i]
	}
	return out
}

// количество видимых
func (c *Columns) CountVisible() (count int) {
	count = 0
	for _, col := range c.list {
		if col.Visible {
			count++
		}
	}
	return count
}

// количество видимых
func (c *Columns) CountAll() int {
	return len(c.list)
}

// ширина вся таблицы
func (c *Columns) SetMaxWidth(width int) {
	c.maxWidth = width
}

// ширина вся таблицы
func (c *Columns) MaxWidth() int {
	return c.maxWidth
}

func (cd *ColumnDescriptor) ClassWidth(s string) *ColumnDescriptor {
	cd.WidthClass = s
	return cd
}

func (cd *ColumnDescriptor) ClassValue(s string) *ColumnDescriptor {
	cd.ValueClass = s
	return cd
}

func (cd *ColumnDescriptor) SetVisible(b bool) *ColumnDescriptor {
	cd.Visible = b
	return cd
}

func (cd *ColumnDescriptor) SetFiltered(b bool) *ColumnDescriptor {
	cd.Filtered = b
	return cd
}
