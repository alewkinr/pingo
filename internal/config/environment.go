package config

// Environment — окружение приложения
type Environment string

const (
	// Development — окружение development
	Development Environment = "development"
	// Staging — окружение staging
	Staging = "staging"
	// Production – окружение production
	Production = "production"
)

// IsDevelopment — проверяем, что текущее окрежение — DEV среда
func (e Environment) IsDevelopment() bool {
	return e == Development
}

// IsStaging — проверяем, что текущее окрежение — STAGE среда
func (e Environment) IsStaging() bool {
	return e == Staging
}

// IsProduction — проверяем, что текущее окрежение — PROD среда
func (e Environment) IsProduction() bool {
	return e == Production
}
