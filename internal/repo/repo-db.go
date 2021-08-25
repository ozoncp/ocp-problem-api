package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type repoPgMetricHolder struct {
	addCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	describeCounter *prometheus.CounterVec
	listCounter *prometheus.CounterVec
	removeCounter *prometheus.CounterVec
}

type repoPg struct {
	RepoRemover
	connectUrl string
	pool *pgxpool.Pool
	metricHolder *repoPgMetricHolder
}

func (rp *repoPg) AddEntities(ctx context.Context, problems []utils.Problem) error {
	var errMethod error
	defer func() {
		if errMethod != nil {
			rp.metricHolder.addCounter.WithLabelValues("error").Inc()
		} else {
			rp.metricHolder.addCounter.WithLabelValues("success").Inc()
		}
	}()

	conn, errMethod := rp.getConnection()
	if errMethod != nil {
		return errMethod
	}

	batch := &pgx.Batch{}
	insertCount := 0

	findProblems, _ := rp.findProblems(ctx, problems)
	for _, problem := range problems {
		if _, ok := findProblems[problem.Id]; ok {
			errMethod = NewRepoError(fmt.Sprintf("duplicate problem #%d", problem.Id), &problem, errMethod)
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

	br := conn.SendBatch(ctx, batch)
	for i := 0; i < insertCount; i++ {
		if ct, err := br.Exec(); ct.Insert() {
			if err != nil {
				errMethod = NewRepoError(err.Error(), nil, err)
			} else if ct.RowsAffected() != 1 {
				errMethod = NewRepoError(fmt.Sprintf("rows in unaffected: %v", ct.String()), nil, err)
			}
		}
	}

	return errMethod
}

func (rp *repoPg) UpdateEntity(ctx context.Context, problem utils.Problem) error {
	var errMethod error
	defer func() {
		if errMethod != nil {
			rp.metricHolder.updateCounter.WithLabelValues("error").Inc()
		} else {
			rp.metricHolder.updateCounter.WithLabelValues("success").Inc()
		}
	}()

	conn, errMethod := rp.getConnection()
	if errMethod != nil {
		return errMethod
	}

	localProblem, errMethod := rp.DescribeEntity(ctx, problem.Id)
	if errMethod != nil {
		return errMethod
	}

	if localProblem == nil {
		errMethod = errors.New("problem not found")
		return errMethod
	}

	ct, errMethod := conn.Exec(
		ctx,
		"update problem set user_id=$1, message=$2 where id = $3;",
		problem.UserId,
		problem.Text,
		problem.Id,
		)

	if errMethod != nil {
		return errMethod
	}

	if ct.RowsAffected() != 1 {
		errMethod = errors.New("problem is not updated")
		return errMethod
	}

	return errMethod
}

func (rp *repoPg) getConnection() (*pgxpool.Pool, error) {
	if rp.pool != nil {
		return rp.pool, nil
	}

	pool, err := pgxpool.Connect(context.Background(), rp.connectUrl)
	if err != nil {
		return nil ,err
	}

	rp.pool = pool
	return rp.pool, nil
}

func (rp *repoPg) findProblems(ctx context.Context, problems []utils.Problem) (map[uint64]utils.Problem, error) {
	var errMethod error

	conn, errMethod := rp.getConnection()
	if errMethod != nil {
		return nil, errMethod
	}

	problemSize := len(problems)
	paramsList := make([]string, 0, problemSize)
	problemIds := make([]interface{}, 0, problemSize)
	for i, problem := range problems {
		paramsList = append(paramsList, fmt.Sprintf("$%v", i+1))
		problemIds = append(problemIds, problem.Id)
	}

	params := strings.Join(paramsList, ",")
	rows, errMethod := conn.Query(
		ctx,
		fmt.Sprintf("select id, user_id, message from problem where id in (%v)", params),
		problemIds...)
	if errMethod != nil {
		return nil, errMethod
	}

	result := map[uint64]utils.Problem{}
	for rows.Next() {
		problem := utils.Problem{}
		errMethod = rows.Scan(&problem.Id, &problem.UserId, &problem.Text)
		if errMethod != nil {
			return nil, errMethod
		}

		result[problem.Id] = problem
	}

	return result, errMethod
}

func (rp *repoPg) DescribeEntity(ctx context.Context, entityId uint64) (*utils.Problem, error) {
	var errMethod error
	defer func() {
		if errMethod != nil {
			rp.metricHolder.describeCounter.WithLabelValues("error").Inc()
		} else {
			rp.metricHolder.describeCounter.WithLabelValues("success").Inc()
		}
	}()

	conn, errMethod := rp.getConnection()
	if errMethod != nil {
		return nil, errMethod
	}

	rows, errMethod := conn.Query(ctx, "select id, user_id, message from problem where id = $1;", entityId)
	if errMethod != nil {
		return nil, errMethod
	}

	resultProblem := &utils.Problem{}
	if !rows.Next() {
		errMethod = errors.New("problem not found")
		return nil, errMethod
	}

	if err := rows.Scan(&resultProblem.Id, &resultProblem.UserId, &resultProblem.Text); err != nil {
		errMethod = err
		return nil, errMethod
	}

	return resultProblem, errMethod
}

func (rp *repoPg) ListEntities(ctx context.Context, limit, offset uint64) ([]utils.Problem, error) {
	var (
		rows pgx.Rows
		errMethod error
	)
	defer func() {
		if errMethod != nil {
			rp.metricHolder.listCounter.WithLabelValues("error").Inc()
		} else {
			rp.metricHolder.listCounter.WithLabelValues("success").Inc()
		}
	}()

	conn, errMethod := rp.getConnection()
	if errMethod != nil {
		return nil, errMethod
	}

	baseQuery := &strings.Builder{}
	baseQuery.WriteString("select id, user_id, message from problem")
	if limit > 0 {
		baseQuery.WriteString(" limit $1")
	}

	if offset > 0 && limit == 0 {
		baseQuery.WriteString(" offset $1")
	}

	if offset > 0 && limit > 0 {
		baseQuery.WriteString(" offset $2")
	}

	baseQuery.WriteByte(';')
	rows, errMethod = conn.Query(ctx, baseQuery.String(), limit, offset)

	if errMethod != nil {
		return nil, errMethod
	}

	result := make([]utils.Problem, 0, limit)
	for rows.Next() {
		problem := utils.Problem{}
		if err := rows.Scan(&problem.Id, &problem.UserId, &problem.Text); err != nil {
			continue
		}

		result = append(result, problem)
	}

	return result, errMethod
}

func (rp *repoPg) RemoveEntity(ctx context.Context, entityId uint64) error {
	var errMethod error
	defer func() {
		if errMethod != nil {
			rp.metricHolder.removeCounter.WithLabelValues("error").Inc()
		} else {
			rp.metricHolder.removeCounter.WithLabelValues("success").Inc()
		}
	}()

	conn, errMethod := rp.getConnection()
	if errMethod != nil {
		return errMethod
	}

	ct, errMethod := conn.Exec(ctx, "delete from problem where id = $1;", entityId)
	if errMethod != nil {
		return errMethod
	}

	if ct.RowsAffected() != 1 {
		errMethod = errors.New("problem is not removed")
		return errMethod
	}

	return errMethod
}

func makeCounterVecMetric(name, help string, labels []string) *prometheus.CounterVec  {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	)
}

func NewPgRepo(connectUrl string) RepoRemover {
	metricHolder := &repoPgMetricHolder{
		addCounter: makeCounterVecMetric(
			"pg_repo_add_counter",
			"Number of AddEntities method calls",
			[]string{"status"},
		),
		updateCounter: makeCounterVecMetric(
			"pg_repo_update_counter",
			"Number of UpdateEntity method calls",
			[]string{"status"},
		),
		describeCounter: makeCounterVecMetric(
			"pg_repo_describe_counter",
			"Number of DescribeEntity method calls",
			[]string{"status"},
		),
		listCounter: makeCounterVecMetric(
			"pg_repo_list_counter",
			"Number of ListEntities method calls",
			[]string{"status"},
		),
		removeCounter: makeCounterVecMetric(
			"pg_repo_remove_counter",
			"Number of RemoveEntity method calls",
			[]string{"status"},
		),
	}

	prometheus.MustRegister(
		metricHolder.addCounter,
		metricHolder.updateCounter,
		metricHolder.describeCounter,
		metricHolder.listCounter,
		metricHolder.removeCounter,
		)

	return &repoPg{
		connectUrl: connectUrl,
		metricHolder: metricHolder,
	}
}
