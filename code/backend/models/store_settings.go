type StoreSettings struct {
    ID              uint    `gorm:"primaryKey"`
    LogoPath        string  `json:"logo_path"`
    ServiceCharge   float64 `json:"service_charge" gorm:"default:0.0"`
    POSViewMode     string  `json:"pos_view_mode" gorm:"default:'detailed'"`
    Timezone        string  `json:"timezone" gorm:"default:'UTC'"`
    DefaultLanguage string  `json:"default_language" gorm:"default:'en'"`
    // ... existing fields
}