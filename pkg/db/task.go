package db

import "fmt"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(t *Task) (int64, error) {
	var id int64

	query := "INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)"
	res, err := db.Exec(query, t.Date, t.Title, t.Comment, t.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}

	return id, err
}

func GetTask(id string) (*Task, error) {
	query := "SELECT * FROM scheduler WHERE id = ?"
	res := db.QueryRow(query, id)

	t := Task{}
	err := res.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func UpdateTask(task *Task) error {
	query := "UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?"
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func UpdateDate(next string, id string) error {
	query := "UPDATE scheduler SET date = ? WHERE id = ?"
	res, err := db.Exec(query, next, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func DeleteTask(id string) error {
	query := "DELETE FROM scheduler WHERE id = ?"
	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func TasksByText(search string, limit int) ([]*Task, error) {
	var res []*Task

	query := "SELECT * FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?"
	pattern := "%" + search + "%"
	rows, err := db.Query(query, pattern, pattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := Task{}
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	if res == nil {
		res = make([]*Task, 0)
	}

	return res, nil
}

func TasksByDate(date string, limit int) ([]*Task, error) {
	var res []*Task

	query := "SELECT * FROM scheduler WHERE date = ? LIMIT ?"
	rows, err := db.Query(query, date, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := Task{}
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	if res == nil {
		res = make([]*Task, 0)
	}

	return res, nil
}
