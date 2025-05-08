// Package store abstracts the database layer behind store.Store methods.
package store

import (
	"context"
	"fmt"
	"github.com/reason4me/tasker/api"
	"log/slog"
	"time"

	"github.com/go-srvc/mods/sqlxmod"
	"github.com/go-srvc/srvc"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Store wraps the sqlxmod module and provides an interface to interact with the database.
type Store struct {
	srvc.Module
	db NamedDB
}

func New(opts ...sqlxmod.Opt) *Store {
	s := &Store{}
	s.Module = sqlxmod.New(append(opts, setDB(s))...)
	return s
}

func (s *Store) Healthy(ctx context.Context) error {
	t := time.Time{}
	if err := s.db.GetContext(ctx, &t, "SELECT NOW()"); err != nil {
		return err
	}
	slog.Info("DB healthy", slog.Time("time_from_db", t))
	return nil
}

func (s *Store) AddTasks(ctx context.Context, newTask api.NewTask) (*api.Task, error) {
	const query = `INSERT INTO tasks (name) VALUES (:name) RETURNING id, name`
	task := &api.Task{}
	err := s.db.NamedGetContext(ctx, task, query, newTask)
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %w", err)
	}

	return task, nil

	// INSERT INTO tasks (name, description, status) VALUES ($1, $2, $3) RETURNING id
}

func (s *Store) GetTasks(ctx context.Context) ([]api.Task, error) {
	const query = `select id, name from tasks`
	tasks := []api.Task{}
	err := s.db.SelectContext(ctx, &tasks, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return tasks, nil

	// INSERT INTO tasks (name, description, status) VALUES ($1, $2, $3) RETURNING id
}

func (s *Store) DeleteTasks(ctx context.Context, id int64) error {
	const query = `DELETE FROM tasks WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
