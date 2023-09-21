package query

type Query struct {
	AQ []*AndQuery
}

func newQuery(andQuery any, andQueries any) (*Query, error) {
	var q Query

	q.AQ = make([]*AndQuery, 1)
	q.AQ[0] = andQuery.(*AndQuery)

	aqs := andQueries.([]interface{})
	for _, aq := range aqs {
		q.AQ = append(q.AQ, aq.(*AndQuery))
	}

	return &q, nil
}

type AndQuery struct {
	FQ []*FieldQuery
}

func newAndQuery(fieldQuery interface{}, fieldQueries interface{}) (*AndQuery, error) {
	var aq AndQuery

	aq.FQ = make([]*FieldQuery, 1)
	aq.FQ[0] = fieldQuery.(*FieldQuery)

	fqs := fieldQueries.([]interface{})
	for _, fq := range fqs {
		aq.FQ = append(aq.FQ, fq.(*FieldQuery))
	}

	return &aq, nil
}

type FieldQuery struct {
	Query *Query
	Field *Field
}

func newFieldQueryFromQuery(query interface{}) (*FieldQuery, error) {
	return &FieldQuery{
		Query: query.(*Query),
		Field: nil,
	}, nil
}

func newFieldQueryFromField(field interface{}) (*FieldQuery, error) {
	return &FieldQuery{
		Query: nil,
		Field: field.(*Field),
	}, nil
}

type Field struct {
	Key   *Source
	Op    string
	Value interface{} // String / Int /Float / Measure
}

func newField(source interface{}, op interface{}, value interface{}) (*Field, error) {
	return &Field{
		Key:   source.(*Source),
		Op:    op.(string),
		Value: value,
	}, nil
}

type Source struct {
	Name string
	Path []string
}

func newSource(name interface{}, path interface{}) (*Source, error) {
	ps := path.([]interface{})

	paths := make([]string, 0)
	for _, p := range ps {
		pa := p.([]interface{})
		paths = append(paths, pa[1].(string))
	}

	return &Source{
		Name: name.(string),
		Path: paths,
	}, nil
}

type Measure struct {
	Number interface{} // int64 / float64
	Unit   string
}

func newMeasure(number interface{}, unit interface{}) (*Measure, error) {
	return &Measure{
		Number: number,
		Unit:   unit.(string),
	}, nil
}
