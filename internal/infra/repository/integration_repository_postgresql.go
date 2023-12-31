package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-integrations-service/internal/entity"
)

type IntegrationRepositoryPostgresql struct {
	db *sql.DB
}

func NewIntegrationRepositoryPostgresql(db *sql.DB) *IntegrationRepositoryPostgresql {
	return &IntegrationRepositoryPostgresql{
		db: db,
	}
}

func (ir *IntegrationRepositoryPostgresql) Create(integration *entity.Integration) error {
	sql := `
		INSERT INTO
			integrations (
				id,
				name,
				company_id,
				platform,
				platform_username,
				platform_token,
				active,
				created_at,
				updated_at
			) VALUES (
				$1, 
				$2, 
				$3, 
				$4, 
				$5, 
				$6, 
				$7,
				NOW(),
				NOW()
			)
	`

	_, err := ir.db.Exec(
		sql,
		integration.ID,
		integration.Name,
		integration.CompanyID,
		integration.Platform,
		integration.PlatformUsername,
		integration.PlatformToken,
		integration.Active,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ir *IntegrationRepositoryPostgresql) Update(integration *entity.Integration) error {
	sql := `
		UPDATE
			integrations
		SET
			name = $1,
			company_id = $2,
			platform = $3,
			platform_username = $4,
			platform_token = $5,
			active = $6,
			updated_at = NOW()
		WHERE
			id = $7
	`

	_, err := ir.db.Exec(
		sql,
		integration.Name,
		integration.CompanyID,
		integration.Platform,
		integration.PlatformUsername,
		integration.PlatformToken,
		integration.Active,
		integration.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ir *IntegrationRepositoryPostgresql) Delete(id string) error {
	sql := `
		DELETE FROM
			integrations
		WHERE
			id = $1
	`
	_, err := ir.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (ir *IntegrationRepositoryPostgresql) FindAllByCompanyIDAndActive(companyID string) ([]*entity.Integration, error) {
	sql := `
		SELECT
			id,
			name,
			company_id,
			platform,
			platform_username,
			platform_token,
			active
		FROM
			integrations
		WHERE
			company_id = $1
		AND
			active = true
	`

	rows, err := ir.db.Query(sql, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var integrations []*entity.Integration
	for rows.Next() {
		integration := &entity.Integration{}
		err := rows.Scan(
			&integration.ID,
			&integration.Name,
			&integration.CompanyID,
			&integration.Platform,
			&integration.PlatformUsername,
			&integration.PlatformToken,
			&integration.Active,
		)
		if err != nil {
			return nil, err
		}

		integrations = append(integrations, integration)
	}

	return integrations, nil
}
