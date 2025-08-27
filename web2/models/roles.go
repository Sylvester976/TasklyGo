package models

import (
	"context"
	"log"
	"web2/db" // your DB pool
)

type Role struct {
	ID   int
	Role string // map `role` column here
}

func GetAllRoles(ctx context.Context) ([]Role, error) {
	rows, err := db.Pool.Query(ctx, "SELECT id, role FROM roles ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var r Role
		if err := rows.Scan(&r.ID, &r.Role); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		roles = append(roles, r)
	}

	return roles, nil
}
