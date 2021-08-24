package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	"strings"
)

type repoPg struct {
	pool *pgxpool.Pool
}

func (rp *repoPg) AddEntities(ctx context.Context, problems []utils.Problem) error {
	var err error
	batch := &pgx.Batch{}
	insertCount := 0

	findProblems, _ := rp.findProblems(ctx, problems)
	for _, problem := range problems {
		if _, ok := findProblems[problem.Id]; ok {
		//if p, _ := rp.DescribeEntity(ctx, problem.Id); p != nil {
			err = NewRepoError(fmt.Sprintf("duplicate problem #%d", problem.Id), &problem, err)
		} else {
			batch.Queue(
				"insert into problem (id, user_id, message) values($1, $2, $3);",
				problem.Id,
				problem.UserId,
				problem.Text,
			)
			insertCount++
		}
	}

	br := rp.pool.SendBatch(ctx, batch)
	for i := 0; i < insertCount; i++ {
		if ct, err := br.Exec(); ct.Insert() {
			if err != nil {
				err = NewRepoError(err.Error(), nil, err)
			} else if ct.RowsAffected() != 1 {
				err = NewRepoError(fmt.Sprintf("rows in unaffected: %v", ct.String()), nil, err)
			}
		}
	}

	return err
}

func (rp *repoPg) findProblems(ctx context.Context, problems []utils.Problem) (map[uint64]utils.Problem, error) {
	problemSize := len(problems)
	paramsList := make([]string, 0, problemSize)
	problemIds := make([]interface{}, 0, problemSize)
	for i, problem := range problems {
		paramsList = append(paramsList, fmt.Sprintf("$%v", i+1))
		problemIds = append(problemIds, problem.Id)
	}

	params := strings.Join(paramsList, ",")
	rows, err := rp.pool.Query(
		ctx,
		fmt.Sprintf("select id, user_id, message from problem where id in (%v)", params),
		problemIds...)
	if err != nil {
		return nil, err
	}

	result := map[uint64]utils.Problem{}
	for rows.Next() {
		problem := utils.Problem{}
		err := rows.Scan(&problem.Id, &problem.UserId, &problem.Text)
		if err != nil {
			return nil, err
		}

		result[problem.Id] = problem
	}

	return result, nil
}

func (rp *repoPg) DescribeEntity(ctx context.Context, entityId uint64) (*utils.Problem, error) {
	rows, err := rp.pool.Query(ctx, "select id, user_id, message from problem where id = $1;", entityId)
	if err != nil {
		return nil, err
	}

	resultProblem := &utils.Problem{}
	if !rows.Next() {
		return nil, errors.New("problem not found")
	}

	for rows.Next() {
		err := rows.Scan(&resultProblem.Id, &resultProblem.UserId, &resultProblem.Text)
		if err != nil {
			return nil, err
		}
	}

	return resultProblem, nil
}

func (rp *repoPg) ListEntities(ctx context.Context, limit, offset uint64) ([]utils.Problem, error) {
	var (
		rows pgx.Rows
		err error
	)

	switch {
	case limit == 0 && offset == 0:
		rows, err = rp.pool.Query(ctx, "select id, user_id, message from problem;")
	case limit > 0 && offset == 0:
		rows, err = rp.pool.Query(ctx, "select id, user_id, message from problem limit $1;", limit)
	case limit == 0 && offset > 0:
		rows, err = rp.pool.Query(ctx, "select id, user_id, message from problem offset $1;", offset)
	default:
		rows, err = rp.pool.Query(ctx, "select id, user_id, message from problem limit $1 offset $2;", limit, offset)
	}

	if err != nil {
		return nil, err
	}

	result := make([]utils.Problem, 0, limit)
	for rows.Next() {
		problem := utils.Problem{}
		if err := rows.Scan(&problem.Id, &problem.UserId, &problem.Text); err != nil {
			continue
		}

		result = append(result, problem)
	}

	return result, nil
}

func (rp *repoPg) RemoveEntity(ctx context.Context, entityId uint64) error {
	ct, err := rp.pool.Exec(ctx, "delete from problem where id = $1;", entityId)
	if err != nil {
		return err
	}

	if ct.RowsAffected() != 1 {
		return errors.New("problem is not removed")
	}

	return nil
}

func NewPgRepo(connectUrl string) (RepoRemover, error) {
	pool, err := pgxpool.Connect(context.Background(), connectUrl)
	if err != nil {
		return nil ,err
	}

	return &repoPg{
		pool: pool,
	}, nil
}
