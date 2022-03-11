package store

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/bytebase/bytebase/api"
	"github.com/bytebase/bytebase/common"
	"go.uber.org/zap"
)

var (
	_ api.ActivityService = (*ActivityService)(nil)
)

// ActivityService represents a service for managing activity.
type ActivityService struct {
	l  *zap.Logger
	db *DB
}

// NewActivityService returns a new instance of ActivityService.
func NewActivityService(logger *zap.Logger, db *DB) *ActivityService {
	return &ActivityService{l: logger, db: db}
}

// CreateActivity creates a new activity.
func (s *ActivityService) CreateActivity(ctx context.Context, create *api.ActivityCreate) (*api.ActivityRaw, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.PTx.Rollback()

	activity, err := createActivity(ctx, tx.PTx, create)
	if err != nil {
		return nil, err
	}

	if err := tx.PTx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return activity, nil
}

// FindActivityList retrieves a list of activities based on the find condition.
func (s *ActivityService) FindActivityList(ctx context.Context, find *api.ActivityFind) ([]*api.ActivityRaw, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.PTx.Rollback()

	list, err := findActivityList(ctx, tx.PTx, find)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// FindActivity retrieves a single activity based on find.
// Returns ECONFLICT if finding more than 1 matching records.
func (s *ActivityService) FindActivity(ctx context.Context, find *api.ActivityFind) (*api.ActivityRaw, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.PTx.Rollback()

	list, err := findActivityList(ctx, tx.PTx, find)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	} else if len(list) > 1 {
		return nil, &common.Error{Code: common.Conflict, Err: fmt.Errorf("found %d activities with filter %+v, expect 1. ", len(list), find)}
	}
	return list[0], nil
}

// PatchActivity updates an existing activity by ID.
// Returns ENOTFOUND if activity does not exist.
func (s *ActivityService) PatchActivity(ctx context.Context, patch *api.ActivityPatch) (*api.ActivityRaw, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.PTx.Rollback()

	activity, err := patchActivity(ctx, tx.PTx, patch)
	if err != nil {
		return nil, FormatError(err)
	}

	if err := tx.PTx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return activity, nil
}

// DeleteActivity deletes an existing activity by ID.
func (s *ActivityService) DeleteActivity(ctx context.Context, delete *api.ActivityDelete) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return FormatError(err)
	}
	defer tx.PTx.Rollback()

	if err := deleteActivity(ctx, tx.PTx, delete); err != nil {
		return FormatError(err)
	}

	if err := tx.PTx.Commit(); err != nil {
		return FormatError(err)
	}

	return nil
}

// createActivity creates a new activity.
func createActivity(ctx context.Context, tx *sql.Tx, create *api.ActivityCreate) (*api.ActivityRaw, error) {
	// Insert row into activity.
	if create.Payload == "" {
		create.Payload = "{}"
	}
	row, err := tx.QueryContext(ctx, `
		INSERT INTO activity (
			creator_id,
			updater_id,
			container_id,
			type,
			level,
			comment,
			payload
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, container_id, type, level, comment, payload
	`,
		create.CreatorID,
		create.CreatorID,
		create.ContainerID,
		create.Type,
		create.Level,
		create.Comment,
		create.Payload,
	)

	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	row.Next()
	var activityRaw api.ActivityRaw
	if err := row.Scan(
		&activityRaw.ID,
		&activityRaw.CreatorID,
		&activityRaw.CreatedTs,
		&activityRaw.UpdaterID,
		&activityRaw.UpdatedTs,
		&activityRaw.ContainerID,
		&activityRaw.Type,
		&activityRaw.Level,
		&activityRaw.Comment,
		&activityRaw.Payload,
	); err != nil {
		return nil, FormatError(err)
	}

	return &activityRaw, nil
}

func findActivityList(ctx context.Context, tx *sql.Tx, find *api.ActivityFind) ([]*api.ActivityRaw, error) {
	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, fmt.Sprintf("id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.ContainerID; v != nil {
		where, args = append(where, fmt.Sprintf("container_id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.CreatorID; v != nil {
		where, args = append(where, fmt.Sprintf("creator_id = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.Type; v != nil {
		where, args = append(where, fmt.Sprintf("type = $%d", len(args)+1)), append(args, *v)
	}
	if v := find.Level; v != nil {
		where, args = append(where, fmt.Sprintf("level = $%d", len(args)+1)), append(args, *v)
	}

	var query = `
		SELECT
			id,
			creator_id,
			created_ts,
			updater_id,
			updated_ts,
			container_id,
			type,
			level,
			comment,
			payload
		FROM activity
		WHERE ` + strings.Join(where, " AND ")
	if v := find.Limit; v != nil {
		query += fmt.Sprintf(" ORDER BY updated_ts DESC LIMIT %d", *v)
	}

	rows, err := tx.QueryContext(ctx, query,
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()

	// Iterate over result set and deserialize rows into activityRawList.
	var activityRawList []*api.ActivityRaw
	for rows.Next() {
		var activity api.ActivityRaw
		if err := rows.Scan(
			&activity.ID,
			&activity.CreatorID,
			&activity.CreatedTs,
			&activity.UpdaterID,
			&activity.UpdatedTs,
			&activity.ContainerID,
			&activity.Type,
			&activity.Level,
			&activity.Comment,
			&activity.Payload,
		); err != nil {
			return nil, FormatError(err)
		}

		activityRawList = append(activityRawList, &activity)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return activityRawList, nil
}

// patchActivity updates a activity by ID. Returns the new state of the activity after update.
func patchActivity(ctx context.Context, tx *sql.Tx, patch *api.ActivityPatch) (*api.ActivityRaw, error) {
	// Build UPDATE clause.
	set, args := []string{"updater_id = $1"}, []interface{}{patch.UpdaterID}
	if v := patch.Comment; v != nil {
		set, args = append(set, fmt.Sprintf("comment = $%d", len(args)+1)), append(args, api.Role(*v))
	}

	args = append(args, patch.ID)

	// Execute update query with RETURNING.
	row, err := tx.QueryContext(ctx, fmt.Sprintf(`
		UPDATE activity
		SET `+strings.Join(set, ", ")+`
		WHERE id = $%d
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, container_id, type, level, comment, payload
	`, len(args)),
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	if row.Next() {
		var activityRaw api.ActivityRaw
		if err := row.Scan(
			&activityRaw.ID,
			&activityRaw.CreatorID,
			&activityRaw.CreatedTs,
			&activityRaw.UpdaterID,
			&activityRaw.UpdatedTs,
			&activityRaw.ContainerID,
			&activityRaw.Type,
			&activityRaw.Level,
			&activityRaw.Comment,
			&activityRaw.Payload,
		); err != nil {
			return nil, FormatError(err)
		}

		return &activityRaw, nil
	}

	return nil, &common.Error{Code: common.NotFound, Err: fmt.Errorf("activity ID not found: %d", patch.ID)}
}

// deleteActivity permanently deletes a activity by ID.
func deleteActivity(ctx context.Context, tx *sql.Tx, delete *api.ActivityDelete) error {
	// Remove row from activity.
	if _, err := tx.ExecContext(ctx, `DELETE FROM activity WHERE id = $1`, delete.ID); err != nil {
		return FormatError(err)
	}
	return nil
}
