// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: leaderboards.sql

package database

import (
	"context"

	snowflake "github.com/disgoorg/snowflake/v2"
)

const averageSpeedLeaderboard = `-- name: AverageSpeedLeaderboard :many
SELECT user_id, cast(avg(speed) as float) as average_speed
FROM bongs
WHERE guild_id = ?
  and message_id > ?
  and won = true
GROUP BY user_id
ORDER BY average_speed, user_id DESC
LIMIT 10
`

type AverageSpeedLeaderboardParams struct {
	GuildID   snowflake.ID `json:"guild_id"`
	MessageID snowflake.ID `json:"message_id"`
}

type AverageSpeedLeaderboardRow struct {
	UserID       snowflake.ID `json:"user_id"`
	AverageSpeed float64      `json:"average_speed"`
}

func (q *Queries) AverageSpeedLeaderboard(ctx context.Context, arg AverageSpeedLeaderboardParams) ([]AverageSpeedLeaderboardRow, error) {
	rows, err := q.db.QueryContext(ctx, averageSpeedLeaderboard, arg.GuildID, arg.MessageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AverageSpeedLeaderboardRow
	for rows.Next() {
		var i AverageSpeedLeaderboardRow
		if err := rows.Scan(&i.UserID, &i.AverageSpeed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fastestSpeedLeaderboard = `-- name: FastestSpeedLeaderboard :many
SELECT user_id, cast(min(speed) as float) as min_speed
FROM bongs
WHERE guild_id = ?
  and message_id > ?
  and won = true
GROUP BY user_id
ORDER BY min_speed, user_id DESC
LIMIT 10
`

type FastestSpeedLeaderboardParams struct {
	GuildID   snowflake.ID `json:"guild_id"`
	MessageID snowflake.ID `json:"message_id"`
}

type FastestSpeedLeaderboardRow struct {
	UserID   snowflake.ID `json:"user_id"`
	MinSpeed float64      `json:"min_speed"`
}

func (q *Queries) FastestSpeedLeaderboard(ctx context.Context, arg FastestSpeedLeaderboardParams) ([]FastestSpeedLeaderboardRow, error) {
	rows, err := q.db.QueryContext(ctx, fastestSpeedLeaderboard, arg.GuildID, arg.MessageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FastestSpeedLeaderboardRow
	for rows.Next() {
		var i FastestSpeedLeaderboardRow
		if err := rows.Scan(&i.UserID, &i.MinSpeed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const slowestSpeedLeaderboard = `-- name: SlowestSpeedLeaderboard :many
SELECT user_id, cast(max(speed) as float) as max_speed
FROM bongs
WHERE guild_id = ?
  and message_id > ?
  and won = true
GROUP BY user_id
ORDER BY max_speed DESC, user_id DESC
LIMIT 10
`

type SlowestSpeedLeaderboardParams struct {
	GuildID   snowflake.ID `json:"guild_id"`
	MessageID snowflake.ID `json:"message_id"`
}

type SlowestSpeedLeaderboardRow struct {
	UserID   snowflake.ID `json:"user_id"`
	MaxSpeed float64      `json:"max_speed"`
}

func (q *Queries) SlowestSpeedLeaderboard(ctx context.Context, arg SlowestSpeedLeaderboardParams) ([]SlowestSpeedLeaderboardRow, error) {
	rows, err := q.db.QueryContext(ctx, slowestSpeedLeaderboard, arg.GuildID, arg.MessageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SlowestSpeedLeaderboardRow
	for rows.Next() {
		var i SlowestSpeedLeaderboardRow
		if err := rows.Scan(&i.UserID, &i.MaxSpeed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const totalClicksLeaderboard = `-- name: TotalClicksLeaderboard :many
SELECT user_id, count(user_id) as bong_count
FROM bongs
WHERE guild_id = ?
  and message_id > ?
  and won = true
GROUP BY user_id
ORDER BY bong_count DESC, user_id DESC
LIMIT 10
`

type TotalClicksLeaderboardParams struct {
	GuildID   snowflake.ID `json:"guild_id"`
	MessageID snowflake.ID `json:"message_id"`
}

type TotalClicksLeaderboardRow struct {
	UserID    snowflake.ID `json:"user_id"`
	BongCount int64        `json:"bong_count"`
}

func (q *Queries) TotalClicksLeaderboard(ctx context.Context, arg TotalClicksLeaderboardParams) ([]TotalClicksLeaderboardRow, error) {
	rows, err := q.db.QueryContext(ctx, totalClicksLeaderboard, arg.GuildID, arg.MessageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TotalClicksLeaderboardRow
	for rows.Next() {
		var i TotalClicksLeaderboardRow
		if err := rows.Scan(&i.UserID, &i.BongCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
